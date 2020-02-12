package database

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	// autoload env
	_ "github.com/joho/godotenv/autoload"
)

// DB struct
type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

// ConnectSQL function
func ConnectSQL() (*DB, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_DATABASE")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	conn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		panic(err)
	}

	dbConn.SQL = conn
	return dbConn, err
}
