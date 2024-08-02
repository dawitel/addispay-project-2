package models

import (
    "encoding/json"
)

// Transaction represents a transaction structure.
type Transaction struct {
    Merchant      Merchant `json:"merchant"`
    TransactionID string  `json:"transaction_id"`
    OrderRequest  Order   `json:"order_request"`
    Status        string  `json:"status"`
    Timestamp     string  `json:"timestamp"`
    Message       string  `json:"message,omitempty"`
}

// ToJSON converts the Transaction to a JSON string.
func (t *Transaction) TxnToJSON() string {
    data, _ := json.Marshal(t)
    return string(data)
}
