package main

import (
    "github.com/dawitel/addispay-project-2/internal/payment"
    "github.com/dawitel/addispay-project-2/internal/utils"
    "github.com/dawitel/addispay-project-2/internal/db"
)

var logger = utils.GetLogger()

func main() {
   
    // Initialize the database
    if err := db.InitDB(); err != nil {
        logger.Error("Failed to initialize database: ", err)
    }

    // Initialize Pulsar client 
    if err := payment.InitPulsar(); err != nil {
        logger.Error("Failed to initialize Pulsar client: ", err)
    }

    // Start consuming order messages and processing payments
    logger.Success("Payment service is running...")
    payment.ConsumeOrders("test/mock/order-topic")
}
