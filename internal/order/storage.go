package order

import (
    "database/sql"
    
    "github.com/dawitel/addispay-project-2/internal/models"
    "github.com/dawitel/addispay-project-2/api/orders/pb"

    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// SaveOrder stores the order details in the database.
func SaveOrder(order *models.Order) error {
    query := "INSERT INTO orders (order_id, customer_id, customer_bacc, amount, phone_number, callback_url) VALUES (?, ?, ?, ?, ?, ?)"
    _, err := db.Exec(query, order.OrderID, order.CustID, order.CustBankAcc, order.Amount, order.PhoneNumber, order.CallbackURL)
    return err
}

// SaveOrder stores the order details in the database.
func SaveTransactionResponseForOrder(txnResp *pb.TransactionResponseForOrder) error {
    query := "INSERT INTO orders_response (order_id, customer_id,transaction_id, status, time_stamp, message) VALUES (?, ?, ?, ?, ?, ?)"
    _, err := db.Exec(query, txnResp.OrderId, txnResp.CustId, txnResp.TransactionId, txnResp.Status, txnResp.Timestamp, txnResp.Message)
    return err
}