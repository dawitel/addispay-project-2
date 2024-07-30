package db

import (
	"database/sql"
	"github.com/dawitel/addispay-project-2/internal/utils"
)

var db *sql.DB
// InitDB initializes the database connection.
func InitDB() error {
    dsn := utils.GoDotEnvVariable("mysql_dsn")
    var err error
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        return err
    }
    return db.Ping()
}
