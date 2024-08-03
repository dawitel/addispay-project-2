package models

import (
	"encoding/json"
)

// import "time"

type OrderLogMessage struct {
    Merchant          Merchant `json:"merchant"`
    OrderID           string   `json:"order_id"`
    CustID            string   `json:"cust_id"`
    CustName          string   `json:"cust_name"`
    PhoneNumber       string   `json:"phone_number"`
    CustBankAcc       string   `json:"cust_bank_acc"` 
    ProductAmount     ProductAmount `json:"product_amount"`
    TotalAmount       float64  `json:"total_amount"`
    LogLevel          string   `json:"log_level"`
    Timestamp         string   `json:"time_stamp"`
    Message           string   `json:"message,omitempty"`
}

func (o *OrderLogMessage) OrderLogToJSON() string {
    data, _ := json.Marshal(o)
    return string(data)
}

type PaymentLogMessage struct {
    Merchant      Merchant `json:"merchant"`
    TransactionID string   `json:"transaction_id"`
    OrderRequest  Order    `json:"order_request"`
    Status        string   `json:"status"`
    Timestamp     string   `json:"timestamp"`
    Message       string   `json:"message,omitempty"`
    LogLevel      string   `json:"log_level"`
}

func (p *PaymentLogMessage) PaymentLogToJSON() string {
    data, _ := json.Marshal(p)
    return string(data)
}

