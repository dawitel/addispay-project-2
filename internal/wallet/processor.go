package wallet

import "github.com/dawitel/addispay-project-2/internal/models"

// update the merchant wallet with the data from the successful tranactions
func ProcessWalletData(txn *models.Transaction) {
	walletData := &models.Wallet {
		Merchant: txn.Merchant,
		Transaction: models.Transaction{
			Merchant: txn.Merchant,
			OrderRequest: models.Order{
				Merchant: txn.Merchant,
				OrderID: txn.OrderRequest.OrderID,
				CustID: txn.OrderRequest.CustID,
				CustName: txn.OrderRequest.CustName,
				CustBankAcc: txn.OrderRequest.CustBankAcc,
				PhoneNumber: txn.OrderRequest.PhoneNumber,
				ProductAmount: txn.OrderRequest.ProductAmount,
				TotalAmount: txn.OrderRequest.TotalAmount,
			},
			TransactionID: txn.TransactionID,
			Status: txn.Status,
			Timestamp: txn.Timestamp,
			Message: txn.Message,
		},
	}
	// save wallet data to database
	SaveWalletData(walletData)

	// publish the data to the wallet-update topic to notify merchant about the new transaction
	PublishWalletUpdates(walletData)
}
