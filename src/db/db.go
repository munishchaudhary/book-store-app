package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// DB is the package-level variable for the database connection.
var DB *sql.DB

// Connect establishes a connection to the MySQL database.
// The function pings the database to ensure the connection is live
// and returns an error if the connection fails.
func Connect() {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Invalid DB_PORT: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		port,
		os.Getenv("DB_NAME"))
	var dbErr error

	// Open a database connection.
	DB, dbErr = sql.Open("mysql", dsn)
	if dbErr != nil {
		log.Fatalf("failed to open database connection: %v", dbErr)
	}

	// Ping the database to verify the connection is active.
	if dbErr = DB.Ping(); dbErr != nil {
		// If ping fails, close the database connection before returning the error.
		DB.Close()
		log.Fatalf("failed to ping database: %v", dbErr)
	}

	log.Println("Successfully connected to MySQL database!")
}

// GetBookDB returns the global BookDB instance
func GetBookDB() *BookDB {
	return &BookDB{db: DB}
}

// Close closes the database connection.
func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
