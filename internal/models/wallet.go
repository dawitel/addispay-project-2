package models

import "encoding/json"

type Merchant struct{
	Products 	 []Product
	MerchantId   string `json:"merchant_id"`
	MerchantName string `json:"merchant_name"`
}

type Wallet struct {
	Merchant Merchant
	Transaction  Transaction 
}

// WDToJSON converts wallet data to JSON
func WDToJSON(w *Wallet) (jsonData string) {
	data, _ := json.Marshal(w)
	return string(data)
}