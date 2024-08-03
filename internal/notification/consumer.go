package notification

import (
	"context"
	"encoding/json"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/dawitel/addispay-project-2/internal/models"
)

// ConsumeUpdates consumes the update messages from the wallet services 
func ConsumeUpdates(consumer pulsar.Consumer) {
	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			logger.Error("Could not receive update message: ", err)
			continue
		}

		var notification models.Notification
		
		if err := json.Unmarshal(msg.Payload(), &notification); err != nil {
			logger.Error("Could not unmarshal log message: ", err)
			consumer.Ack(msg)
			continue
		}
		
        
		// send the notification to the frontend
		sendNotificationsToAPIgateway(&notification)

		// send the email notification
		sendEmailNotification(&notification)
		
		// send the SMS message
		sendSMSNotifications(&notification)

		// Acknowledge the message
		consumer.Ack(msg)
	}
}

