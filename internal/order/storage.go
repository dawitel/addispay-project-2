package order

import (
    "database/sql"
    
    "github.com/dawitel/addispay-project-2/internal/models"

    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// SaveOrder stores the order details in the database.
func SaveOrder(order *models.Order) error {
    query := "INSERT INTO orders (order_id, customer_id, customer_bacc, amount, phone_number, callback_url) VALUES (?, ?, ?, ?, ?, ?)"
    _, err := db.Exec(query, order.OrderID, order.CustID, order.CustBankAcc, order.Amount, order.PhoneNumber, order.CallbackURL)
    return err
}
