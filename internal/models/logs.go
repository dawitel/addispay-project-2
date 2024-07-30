package models

import "encoding/json"

// import "time"


type OrderLogMessage struct {
    OrderID     string `json:"order_id"`
    CustID      string `json:"cust_id"`
    Status      string `json:"status"`
    Amount      float64 `json:"amount"`
    Timestamp   string `json:"timestamp"` // TODO: move from string -> time.Time
    LogLevel    string `json:"logLevel"`
    Message     string `json:"message"`
}

func (o *OrderLogMessage) OrderLogToJSON() string {
    data, _ := json.Marshal(o)
    return string(data)
}

type PaymentLogMessage struct {
    TransactionID string `json:"transaction_id"`
    OrderID     string   `json:"order_id"`
    CustID      string   `json:"cust_id"`
    Status      string   `json:"status"` // -> EXPIRED, FAILED, SUCCESS, PENDING
    Amount      float64  `json:"amount"`
    Timestamp   string   `json:"timestamp"` // TODO: move from string -> time.Time
    LogLevel    string   `json:"logLevel"`
    Message     string   `json:"message,omitempty"`
}

func (p *PaymentLogMessage) PaymentLogToJSON() string {
    data, _ := json.Marshal(p)
    return string(data)
}

