package order

import (
	"context"
	"encoding/json"

	// "github.com/dawitel/addispay-project-2/api/orders/pb"
	"github.com/dawitel/addispay-project-2/configs"
	"github.com/dawitel/addispay-project-2/internal/models"

	"github.com/apache/pulsar-client-go/pulsar"
)

var pulsarClient pulsar.Client

// InitPulsar initializes a Pulsar client for the order service.
func InitPulsar() error {
    config, err := configs.LoadConfig()
    if err != nil {
        logger.Error("Could not load configuration files: ", err)
    }
    client, err := pulsar.NewClient(pulsar.ClientOptions{
        URL: config.ProductionPulsarURL,
    })
    if err != nil {
        return err
    }
    pulsarClient = client
    return nil
}

// PublishOrder publishes an order message to a Pulsar topic.
func PublishOrder(order *models.Order) error {
     config, err := configs.LoadConfig()
    if err != nil {
        logger.Error("Could not load configuration files: ", err)
    }
    producer, err := pulsarClient.CreateProducer(pulsar.ProducerOptions{
        Topic: config.OrdersTopic,
    })
    if err != nil {
        return err
    }
    defer producer.Close()

    logMessage := &models.OrderLogMessage{
        Merchant: order.Merchant,
        OrderID: order.OrderID,
        CustID: order.CustID,
        TotalAmount: order.TotalAmount,
        Timestamp: "",
        LogLevel: "INFO",
        Message: "Order created",
    }

    if err = PublishLogs(logMessage); err != nil {
        logger.Error("Could not publish order logs: ", err)
    }
    msg := pulsar.ProducerMessage{
        Payload: []byte(order.OrderToJSON()),
    }

    _, err = producer.Send(context.Background(), &msg)
    return err
}

// PublishLogs publishes order logs to the logs-topic.
func PublishLogs(logMessage *models.OrderLogMessage) error {
     config, err := configs.LoadConfig()
    if err != nil {
        logger.Error("Could not load configuration files: ", err)
    }
    producer, err := pulsarClient.CreateProducer(pulsar.ProducerOptions{
        Topic: config.OrdersLogTopic,
    })
    if err != nil {
        return err
    }
    defer producer.Close()

    msg := pulsar.ProducerMessage{
        Payload: []byte(logMessage.OrderLogToJSON()),
    }

    _, err = producer.Send(context.Background(), &msg)
    return err
}


func ConsumeOrderResponse() {
    // Load configuration files to the environment
	config, err := configs.LoadConfig()
    if err != nil {
        logger.Error("Could not load configuration files")
    }
    // Create a consumer
    consumer, err := pulsarClient.Subscribe(pulsar.ConsumerOptions{
        Topic: config.TransactionsTopic,
        SubscriptionName: config.OrderResponseSubscription,
        Type: pulsar.Shared,
    })
    if err!= nil {
        logger.Error("Could not create consumer: ", err)
    }
    defer consumer.Close()

    
    for {
        msg, err := consumer.Receive(context.Background())
        if err != nil {
            logger.Error("Could not receive processed orders messages: ", err)
            continue
        }

        var txn models.Transaction
        if err := json.Unmarshal(msg.Payload(), &txn); err != nil {
            logger.Error("Could not unmarshal processed transation message: ", err)
            consumer.Ack(msg)
            continue
        }

        logger.Success("Received order:", txn.OrderRequest.OrderID)
       
        // Process the order response
        ProcessOrderResults(&txn)

        // Acknowledge the message
        consumer.Ack(msg)
    }

}