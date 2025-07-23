package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/denisenkom/go-mssqldb" // SQL Server driver'ı
)

// NewSQLServerDB, SQL Server veritabanına bir bağlantı açar ve doğrular.
// Bağlantı havuzunu döndürür (*sql.DB).
func NewSQLServerDB(connString string) (*sql.DB, error) {
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Bağlantı havuzu ayarları (isteğe bağlı, performansı optimize etmek için)
	db.SetMaxOpenConns(25)                 // Aynı anda açılacak maksimum bağlantı sayısı
	db.SetMaxIdleConns(10)                 // Boşta duran bağlantı havuzundaki maksimum bağlantı sayısı
	db.SetConnMaxLifetime(5 * time.Minute) // Bir bağlantının yeniden kullanılmadan önce ne kadar süre açık kalacağı

	// Bağlantıyı doğrula (ping at)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // Bağlantı doğrulama için timeout
	defer cancel()

	err = db.PingContext(ctx) // PingContext ile timeout ekleyebiliriz
	if err != nil {
		db.Close() // Ping başarısız olursa bağlantıyı kapat
		return nil, fmt.Errorf("failed to connect to the database (ping failed): %w", err)
	}

	log.Println("Database connection established successfully.")
	return db, nil
}
