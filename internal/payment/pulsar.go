package payment

import (
	"context"
	"encoding/json"

	"github.com/dawitel/addispay-project-2/configs"
	"github.com/dawitel/addispay-project-2/internal/models"

	"github.com/apache/pulsar-client-go/pulsar"
)

var pulsarClient pulsar.Client

// InitPulsar initializes the Pulsar client.
func InitPulsar(pulsarURL string) error {
    client, err := pulsar.NewClient(pulsar.ClientOptions{
        URL: pulsarURL,
    })
    if err != nil {
        return err
    }
    pulsarClient = client
    return nil
}

// ConsumeOrders consumes order messages from the specified Pulsar topic and processes them.
func ConsumeOrders(topic string) {
    consumer, err := pulsarClient.Subscribe(pulsar.ConsumerOptions{
        Topic:            topic,
        SubscriptionName: "order-subscription",
        Type:             pulsar.Shared,
    })
    if err != nil {
        logger.Error("Could not subscribe to Pulsar topic: ", err)
    }
    defer consumer.Close()

    for {
        msg, err := consumer.Receive(context.Background())
        if err != nil {
            logger.Error("Could not receive message: ", err)
            continue
        }

        var order models.Order
        if err := json.Unmarshal(msg.Payload(), &order); err != nil {
            logger.Error("Could not unmarshal order message: ", err)
            consumer.Ack(msg)
            continue
        }

        logger.Success("Received order:", order.OrderID)

        // Process the order
        transaction := ProcessPayment(&order)
        logMessage := &models.PaymentLogMessage{
            TransactionID: transaction.TransactionID,
            OrderID: transaction.OrderID,
            CustID: transaction.CustID,
            Status: transaction.Status,
            Amount: transaction.Amount,
            Timestamp: "", // replace to time.Time()
            LogLevel: transaction.Status,
            Message: transaction.Message,
        }
        // Save transaction to database
        if err := SaveTransaction(transaction); err != nil {
            logger.Error("Failed to save transaction:", err)
        }

        // Publish transaction result
        if err := PublishTransaction(transaction); err != nil {
            logger.Error("Failed to publish transaction:", err)
        }

        // Publish transaction logs 
        if err := PublishLogs(logMessage); err != nil{
            logger.Error("Failed to publish transaction logs: ", err)
        }
        // Acknowledge the message
        consumer.Ack(msg)
    }
}

// PublishTransaction publishes the transaction result to the specified Pulsar topic.
func PublishTransaction(transaction *models.Transaction) error {
    // Load configuration files to the environment
	config, err := configs.LoadConfig()
    if err != nil {
        logger.Error("Could not load configuration files")
    }
    producer, err := pulsarClient.CreateProducer(pulsar.ProducerOptions{
        Topic: config.TransactionsTopic,
    })
    if err != nil {
        return err
    }
    defer producer.Close()

    msg := pulsar.ProducerMessage{
        Payload: []byte(transaction.TxnToJSON()),
    }

    _, err = producer.Send(context.Background(), &msg)
    return err
}

// PublishLogs publishes the transaction logs to the payment-logs-topic.
func PublishLogs(logMessage *models.PaymentLogMessage) error {
    // Load configuration files to the environment
	config, err := configs.LoadConfig()
    if err != nil {
        logger.Error("Could not load configuration files")
    }
    producer, err := pulsarClient.CreateProducer(pulsar.ProducerOptions{
        Topic: config.PaymentsLogTopic,
    })
    if err != nil {
        return err
    }
    defer producer.Close()

    msg := pulsar.ProducerMessage{
        Payload: []byte(logMessage.PaymentLogToJSON()),
    }

    _, err = producer.Send(context.Background(), &msg)
    return err
}
