package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
)

var database *sql.DB

// Init Initalize database
func Init() error {

	var err error = nil

	log.Println("Initializing database")
	dbIP := os.Getenv("dbIP")
	dbUser := os.Getenv("dbUSR")
	dbPassword := os.Getenv("dbPW")
	dbName := os.Getenv("dbName")

	if dbIP == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		return errors.New("(one or many of) your {$dbIP, $dbUSR, $dbPW, $dbName} environment variables are not set")
	}

	connectstr := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbIP, dbName)

	database, err = sql.Open("mysql", connectstr)

	return err
}
