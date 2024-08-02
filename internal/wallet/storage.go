package wallet

import (
	// "database/sql"
	// "fmt"

	"github.com/dawitel/addispay-project-2/internal/models"
)

// var db *sql.DB

func SaveWalletData(walletData *models.Wallet) error {
	// merchant_id :=  walletData.Merchant.MerchantId
	// merchant_name := walletData.Merchant.MerchantName
	// merchant_phone := walletData.Merchant.MerchantPhone

	// query_sales := fmt.Sprintf("SELECT FROM merchants WHERE %s", merchant_id)
	// numberOfSales, err := db.Exec(query_sales)

	// query_revenue := ""
	// total_revenue, err := db.Exec(query_revenue)
	
	// query_profit := ""
	// total_profit, err := db.Exec(query_profit)

	// add the amount to the corresponding data points form the wallet data
	// total_revenue += walletData.Transaction.OrderRequest.TotalAmount
	// total_profit = total_revenue - 10 // simulate transaction fees 
	// numberOfSales += 1

	// query := "INSERT INTO merchants (merchant_id, merchant_name, merchant_phone, total_profit, total_revenue, number_of_sales) VALUES (?, ?, ?, ?, ?, ?)"
    // _, err = db.Exec(query, merchant_id, merchant_name, merchant_phone, total_profit, total_revenue, numberOfSales)
    // return err
	return nil
}