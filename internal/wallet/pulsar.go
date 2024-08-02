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

		// notify merchant about the new transaction
		

		// Acknowledge the message
		consumer.Ack(msg)
	}

}