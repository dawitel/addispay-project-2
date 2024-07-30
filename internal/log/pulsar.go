package log

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/dawitel/addispay-project-2/configs"
	"github.com/dawitel/addispay-project-2/internal/utils"
)

var logger = utils.GetLogger()

func InitPulsar() error {
	// Load configuration files to the environment
	config, err := configs.LoadConfig("configs/configs.yml")
    if err != nil {
        logger.Error("Could not load configuration files")
    }

	client, err := pulsar.NewClient(pulsar.ClientOptions{
        URL: config.ProductionPulsarURL,
    })
    if err != nil {
        logger.Error("Could not initialize Pulsar client: ", err)
    }
    defer client.Close()

    // payment log consumer
    paymentConsumer, err := client.Subscribe(pulsar.ConsumerOptions{
        Topic:            "payment-logs-topic",
        SubscriptionName: "payment-log-subscription",
        Type:             pulsar.Shared,
    })
    if err != nil {
        logger.Error("Could not subscribe to logs topic: ", err)
    }
    defer paymentConsumer.Close()
    
    // order log consumer
    orderConsumer, err := client.Subscribe(pulsar.ConsumerOptions{
        Topic:            "order-logs-topic",
        SubscriptionName: "order-log-subscription",
        Type:             pulsar.Shared,
    })
    if err != nil {
        logger.Error("Could not subscribe to logs topic: ", err)
    }
    defer orderConsumer.Close()

    logger.Success("Logger service is running and waiting for messages...")
	ConsumePaymentLogs(paymentConsumer)
	ConsumeOrderLogs(orderConsumer)
	
	return nil
}

