package log

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/dawitel/addispay-project-2/internal/models"
)

func ConsumeOrderLogs(consumer pulsar.Consumer) {
	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			logger.Error("Could not receive message: ", err)
			continue
		}

		var logMessage models.OrderLogMessage
		
		if err := json.Unmarshal(msg.Payload(), &logMessage); err != nil {
			logger.Error("Could not unmarshal log message: ", err)
			consumer.Ack(msg)
			continue
		}

		WriteOrderLogToLogFile(&logMessage)

		// Save log message to database
		if err = SaveOrderLogToDB(&logMessage); err != nil {
			logger.Error("Could not save to DB: ", err)
		}

		// Acknowledge the message
		consumer.Ack(msg)
	}
}

func ConsumePaymentLogs(consumer pulsar.Consumer) {
	for {
		msg, err := consumer.Receive(context.Background())
		if err != nil {
			logger.Error("Could not receive message: ", err)
			continue
		}

		var logMessage models.PaymentLogMessage
		
		if err := json.Unmarshal(msg.Payload(), &logMessage); err != nil {
			logger.Error("Could not unmarshal log message: ", err)
			consumer.Ack(msg)
			continue
		}
		writePaymentLogToFile(logMessage)

		// Save log message to database
		if err = SavePaymentLogToDB(logMessage); err != nil {
			logger.Error("Could not save to DB: ", err)
		}

		// Acknowledge the message
		consumer.Ack(msg)
	}
}

func WriteOrderLogToLogFile(logMessage *models.OrderLogMessage){
	
		// Open the log file
		logFile, err := os.OpenFile("logs/order_service.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logger.Error("Failed to open order service log file: ", err)
		}
		defer logFile.Close()

		// Set the output of the default logger to the log file
		log.SetOutput(logFile)

		log.Printf("[LOG]: %s %s %s %f %s", logMessage.Merchant.MerchantId, logMessage.OrderID, logMessage.CustID, logMessage.TotalAmount, logMessage.ProductAmount.Product.ProductName)
	
}
func writePaymentLogToFile(logMessage models.PaymentLogMessage){
		// Open the log file
		logFile, err := os.OpenFile("logs/payment_service.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			logger.Error("Failed to open payment service log file: ", err)
		}
		defer logFile.Close()

		// Set the output of the default logger to the log file
		log.SetOutput(logFile)
		
		log.Printf("[LOG]: %s %s %s %s %s %f %s %s", logMessage.Timestamp, logMessage.TransactionID, logMessage.OrderRequest.OrderID, logMessage.OrderRequest.CustID, logMessage.Status, logMessage.OrderRequest.TotalAmount, logMessage.Message, logMessage.LogLevel)
}