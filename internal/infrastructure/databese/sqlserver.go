package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb" // SQL Server driver
)

// NewSQLServerDB opens and validates a connection to a SQL Server database.
// It returns the connection pool (*sql.DB).
func NewSQLServerDB(connString string) (*sql.DB, error) {
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Connection pool settings (optional, for performance optimization)
	db.SetMaxOpenConns(25)                 // Maximum number of open connections at any given time
	db.SetMaxIdleConns(10)                 // Maximum number of idle connections in the pool
	db.SetConnMaxLifetime(5 * time.Minute) // How long a connection can be open before being reused

	// Validate the connection (ping)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Timeout for connection validation
	defer cancel()

	err = db.PingContext(ctx) // We can add a timeout with PingContext
	if err != nil {
		db.Close() // Close the connection if ping fails
		return nil, fmt.Errorf("failed to connect to the database (ping failed): %w", err)
	}

	log.Println("Database connection established successfully.")
	return db, nil
}
