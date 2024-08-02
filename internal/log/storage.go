package log

import (
	"database/sql"

	"github.com/dawitel/addispay-project-2/internal/models"
)

var db *sql.DB

func SaveOrderLogToDB(logMessage *models.OrderLogMessage) error {
	query := "INSERT INTO order_logs (time_stamp, order_id, customer_id, amount, status, message, log_level) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, logMessage.Timestamp, logMessage.OrderID, logMessage.CustID, logMessage.TotalAmount, logMessage.LogLevel)
	return err
}
func SavePaymentLogToDB(logMessage models.PaymentLogMessage) error {
	query := "INSERT INTO payment_logs (time_stamp, transaction_ID, order_ID, customer_id, amount, status, message, log_level) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, logMessage.Timestamp, logMessage.TransactionID, logMessage.OrderRequest.OrderID, logMessage.OrderRequest.CustID, logMessage.OrderRequest.TotalAmount, logMessage.Status, logMessage.Message, logMessage.LogLevel)
	return err
}