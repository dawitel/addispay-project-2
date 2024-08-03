package main

import (
	"github.com/dawitel/addispay-project-2/internal/utils"
	"github.com/dawitel/addispay-project-2/internal/log"
	"github.com/dawitel/addispay-project-2/internal/db"
)


var logger = utils.GetLogger()

func main() {
    // initialize the database connection
    if err := db.InitDB(); err != nil {
        logger.Error("Failed to initialize database connection: ", err)
    }

    // Initialize Pulsar
    if err := log.InitPulsar(); err != nil {
        logger.Error("Could not initialize the pulsar client: ", err)
    }
}


