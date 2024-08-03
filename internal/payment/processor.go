package payment

import (
    "time"

    "github.com/dawitel/addispay-project-2/internal/models"
    "github.com/dawitel/addispay-project-2/internal/utils"
)

var logger = utils.GetLogger()

// ProcessPayment simulates payment processing for an order.
func ProcessPayment(order *models.Order) *models.Transaction {
    transaction := &models.Transaction{
        Merchant: order.Merchant,
        TransactionID: utils.GenerateID(),
        OrderRequest:       *order,
        Status:        "PENDING",
        Timestamp:     "" , // replace -> time.now()
    }

    maxRetries := 3
    retryInterval := 30 * time.Second // use exponential back off strategy for the retry itervals -> f(n) = 2^n
    startTime := time.Now()
    for i := 0; i < maxRetries; i++ {
        // Simulate payment processing
        time.Sleep(2 * time.Second) // Simulated processing time

        if utils.RandomSuccess() {
            transaction.Status = "SUCCESS"
            break
        } else {
            logger.Info("Retrying payment for order:", order.OrderID)
            time.Sleep(retryInterval)
        }

        if time.Since(startTime) > 2*time.Minute {
            transaction.Status = "EXPIRED"
            transaction.Message = "Transaction expired after 3 retries."
            break
        }
    }

    if transaction.Status != "SUCCESS" {
        transaction.Status = "FAILED"
        transaction.Message = "Transaction failed after 3 retries."
    }

    return transaction
}
