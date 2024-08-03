package notification

import (
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/dawitel/addispay-project-2/configs"
)


// InitPulsar initializes a new pulsarclient for the notification service and satrts consuming update messages
// produced by the wallet service
func InitPulsar() error {
	// Load configuration files to the environment
	config, err := configs.LoadConfig()
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
    
    // order log consumer
    consumer, err := client.Subscribe(pulsar.ConsumerOptions{
        Topic:            config.WalletUpdatesTopic,
        SubscriptionName: config.WalletUpdatesSubscription,
        Type:             pulsar.Shared,
    })
    if err != nil {
        logger.Error("Could not subscribe to wallet-updates topic: ", err)
    }
    defer consumer.Close()

    logger.Success("Notification service is initialized successfully and waiting for messages...")
    
    ConsumeUpdates(consumer)
	
	return nil
}

