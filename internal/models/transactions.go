package models

import (
    "encoding/json"
)

// Transaction represents a transaction structure.
type Transaction struct {
    TransactionID string  `json:"transaction_id"`
    CustID        string  `json:"cust_id"`
    OrderID       string  `json:"order_id"`
    Status        string  `json:"status"`
    Amount        float64 `json:"amount"`
    Timestamp     string  `json:"timestamp"`
    Message       string  `json:"message,omitempty"`
}

// ToJSON converts the Transaction to a JSON string.
func (t *Transaction) TxnToJSON() string {
    data, _ := json.Marshal(t)
    return string(data)
}
