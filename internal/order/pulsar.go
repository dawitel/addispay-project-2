package order

import (
	"context"

	"github.com/dawitel/addispay-project-2/configs"
	"github.com/dawitel/addispay-project-2/internal/models"

	"github.com/apache/pulsar-client-go/pulsar"
)

var pulsarClient pulsar.Client

// InitPulsar initializes the Pulsar client.
func InitPulsar() error {
    config, err := configs.LoadConfig("configs/configs.yml")
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
     config, err := configs.LoadConfig("configs/configs.yml")
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
        OrderID: order.OrderID,
        CustID: order.CustID,
        Status: "CREATED",
        Amount: order.Amount,
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
     config, err := configs.LoadConfig("configs/configs.yml")
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
