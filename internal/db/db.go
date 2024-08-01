package db

import (
	"database/sql"
	"github.com/dawitel/addispay-project-2/internal/utils"
)

var db *sql.DB
var logger = utils.GetLogger()

// InitDB initializes a database connection for the services.
func InitDB() error {
    dsn := utils.GoDotEnvVariable("mysql_dsn")
    var err error
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        return err
    }
    logger.Success("Initialized connection to database successfully")
    return db.Ping()
}
