package wallet

import (
	"context"
	"encoding/json"

	"github.com/dawitel/addispay-project-2/configs"
	"github.com/dawitel/addispay-project-2/internal/models"
	"github.com/dawitel/addispay-project-2/internal/utils"

	"github.com/apache/pulsar-client-go/pulsar"
)

var pulsarClient pulsar.Client
var logger = utils.GetLogger()

// ConsumeOrderResponseForWalletService consumes messages from the processed orders topic 
// to save them to the merchant wallet
func ConsumeOrderResponseForWalletService() {
	// Load configuration files to the environment
	config, err := configs.LoadConfig()
	if err != nil {
		logger.Error("Could not load configuration files")
	}
	// Create a consumer
	consumer, err := pulsarClient.Subscribe(pulsar.ConsumerOptions{
		Topic:            config.TransactionsTopic,
		SubscriptionName: config.OrderResponseSubscription,
		Type:             pulsar.Shared,
	})
	if err != nil {
		logger.Error("Could not create consumer for wallet service: ", err)
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

		// Process the wallet data
		ProcessWalletData(&txn)

		// Acknowledge the message
		consumer.Ack(msg)
	}

}

// PublishWalletUpdates publishes the wallet updates for a merchant to the wallet-update-topic
// then the notification service picks up the message and notifies the merchant about the order
func PublishWalletUpdates(walletData *models.Wallet) error {
	config, err := configs.LoadConfig()
    if err != nil {
        logger.Error("Could not load configuration files: ", err)
    }
    producer, err := pulsarClient.CreateProducer(pulsar.ProducerOptions{
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