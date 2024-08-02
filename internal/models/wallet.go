package models

import "encoding/json"

type Merchant struct{
	Products 	  []Product `json:"products"`
	MerchantId    string `json:"merchant_id"`
	MerchantName  string `json:"merchant_name"`
	MerchantEmail string `json:"merchant_emial"`
	MerchantPhone string `json:"merchant_phone"`
	TotalRevenue  float64 `json:"total_revenue"`
	TotalProfit   float64 `json:"total_profit"`
	NumberOfSales   float64 `json:"number_of_sales"`
}

type Wallet struct {
	Merchant Merchant
	Transaction  Transaction 
}

// WDToJSON converts wallet data to JSON
func (w *Wallet) WDToJSON() (jsonData string) {
	data, _ := json.Marshal(w)
	return string(data)
}