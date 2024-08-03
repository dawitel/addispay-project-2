package payment

import (
    "database/sql"
    "github.com/dawitel/addispay-project-2/internal/models"

    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB


// SaveTransaction stores the transaction details in the database.
func SaveTransaction(transaction *models.Transaction) error {
    query := "INSERT INTO transactions (transaction_id, order_id, customer_id, amount, status, time_stamp, message) VALUES (?, ?, ?, ?, ?)"
    _, err := db.Exec(query, transaction.TransactionID, transaction.OrderRequest.OrderID, transaction.OrderRequest.CustID, transaction.OrderRequest.TotalAmount, transaction.Status, transaction.Timestamp, transaction.Message)
    return err
}
