package database

import (
	"database/sql"
	"fmt"
	"time"
)

// Change connection string here
const (
	hostName = "localhost"
	hostPort = 5432
	username = "postgres"
	password = "postgres"
	dbName   = "testdb"
)

var DBConn *sql.DB

// SetupDatabase :Setup database connection
func SetupDatabase() {
	connStr := fmt.Sprintf("port=%d host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", hostPort, hostName, username, password, dbName)
	var err error
	DBConn, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	DBConn.SetMaxOpenConns(4)
	DBConn.SetMaxIdleConns(4)
	DBConn.SetConnMaxLifetime(60 * time.Second)
	fmt.Println("Database successfully connected!")
}
