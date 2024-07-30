package models

import (
    "encoding/json"
)

// Order represents an order structure.
type Order struct {
    OrderID     string  `json:"order_id"`
    CustID      string  `json:"cust_id"`
    CustBankAcc string  `json:"cust_bank_acc"` 
    Amount      float64 `json:"amount"`
    PhoneNumber string  `json:"phone_number"`
    CallbackURL string  `json:"callback_url"`
}

// ToJSON converts the Order to a JSON string.
func (o *Order) OrderToJSON() string {
    data, _ := json.Marshal(o)
    return string(data)
}
