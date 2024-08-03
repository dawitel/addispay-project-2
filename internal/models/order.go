package models

import (
    "encoding/json"
)

// Order represents an order structure.
type Order struct {
    Merchant          Merchant `json:"merchant"`
    OrderID           string   `json:"order_id"`
    CustID            string   `json:"cust_id"`
    CustName          string   `json:"cust_name"`
    PhoneNumber       string   `json:"phone_number"`
    CustBankAcc       string   `json:"cust_bank_acc"` 
    ProductAmount     ProductAmount `json:"product_amount"`
    TotalAmount       float64  `json:"total_amount"`
}

type ProductAmount struct {
    Product        Product `json:"product"`
    ProductAmount  int64 `json:"product_amount"`
}

// ToJSON converts the Order to a JSON string.
func (o *Order) OrderToJSON() string {
    data, _ := json.Marshal(o)
    return string(data)
}
