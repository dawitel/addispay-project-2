package models

type Product struct {
	ProductId     string `json:"product_id"`
	ProductName   string `json:"product_name"`
	Price         float64 `json:"product_price"`
}