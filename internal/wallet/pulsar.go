package wallet

import (
	"context"
	"encoding/json"

	"github.com/dawitel/addispay-project-2/configs"
	"github.com/dawitel/addispay-project-2/internal/models"
	"github.com/dawitel/addispay-project-2/internal/utils"

	"github.com/apache/pulsar-client-go/pulsar"
)

var logger = utils.GetLogger()

// ConsumeOrderResponseForWalletService consumes messages from the processed orders topic 
// to save them to the merchant wallet
func ConsumeOrderResponseForWalletService() {
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
	logger.Success("Pulsar client for wallet service successfully initialized")
	
	// Create a consumer
	consumer, err := client.Subscribe(pulsar.ConsumerOptions{
		Topic:            config.TransactionsTopic,
		SubscriptionName: config.OrderResponseSubscription,
		Type:             pulsar.Shared,
	})
	if err != nil {
		logger.Error("Could not create consumer for wallet service: ", err)
	}
	defer consumer.Close()
	logger.Success("wallet service is successfully initialized and consuming messages...")

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

		// Process the wallet data
		ProcessWalletData(&txn)

		// Acknowledge the message
		consumer.Ack(msg)
	}

}

// PublishWalletUpdates publishes the wallet updates for a merchant to the wallet-update-topic
// then the notification service picks up the message and notifies the merchant about the order
func PublishWalletUpdates(walletData *models.Wallet) error {
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
    producer, err := client.CreateProducer(pulsar.ProducerOptions{
        Topic: config.WalletUpdatesTopic,
    })
    if err != nil {
        return err
    }
    defer producer.Close()


    msg := pulsar.ProducerMessage{
        Payload: []byte(walletData.WDToJSON()),
    }

    _, err = producer.Send(context.Background(), &msg)
    return err
}