package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dawitel/addispay-project-2/api/orders/pb"
	"github.com/dawitel/addispay-project-2/configs"
	"github.com/dawitel/addispay-project-2/internal/models"
)

func ProcessOrderResults(txn *models.Transaction) {
	config, err := configs.LoadConfig()
	if err != nil {
		logger.Error("Failed to load configuration file: ", err)
	}
	// // donot process the order if the status is "PENDING"
	// if txn.Status == "PENDING" {
	// 	return
	// }
	// construct the response for the API client
	orderResponse := &pb.TransactionResponseForOrder{
		TransactionId: txn.TransactionID,
		OrderId: txn.OrderID,
		CustId: txn.CustID,
		Status: txn.Status,
		Amount: txn.Amount,
		Timestamp: txn.Timestamp,
		Message: txn.Message,
	}
	// write the order response to the DB
	if err = SaveTransactionResponseForOrder(orderResponse); err != nil {
		logger.Error("Failed to write the order response to the Database: ", err)
	}

	// marshal the data to JSON
	orderResponseJSON, err := json.Marshal(orderResponse)
	if err!= nil {
		logger.Error("Failed to marshal the transactions response: ", err)
	}

	URL := fmt.Sprintf("%s/processed-transactions", config.APIGatewayAddr)
	
	// send the JSON data to the API gateway
	response, err := http.Post(URL, "application/json", bytes.NewBuffer(orderResponseJSON))
	if err != nil {
		logger.Error("Failed to send the order Response to the API gateway: ", err)
	}

	defer response.Body.Close()
	logger.Success("Order Response data sent to the API gateway: ", response.StatusCode)
}