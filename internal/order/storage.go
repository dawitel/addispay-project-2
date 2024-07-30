package order

import (
    "database/sql"
    "github.com/dawitel/addispay-project-2/internal/models"
    "github.com/dawitel/addispay-project-2/internal/utils"

    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
// InitDB initializes the database connection.
func InitDB() error {
    dsn := utils.GoDotEnvVariable("mysql_dsn")
    var err error
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        return err
    }
    return db.Ping()
}

// SaveOrder stores the order details in the database.
func SaveOrder(order *models.Order) error {
    query := "INSERT INTO orders (order_id, customer_id, customer_bacc, amount, phone_number, callback_url) VALUES (?, ?, ?, ?, ?, ?)"
    _, err := db.Exec(query, order.OrderID, order.CustID, order.CustBankAcc, order.Amount, order.PhoneNumber, order.CallbackURL)
    return err
}
