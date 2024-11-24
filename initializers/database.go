package initializers

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error

	// Fetch the database connection string from the environment
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		log.Fatal("Environment variable DB_URL is not set")
	}

	// Open the database connection with improved settings
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // Avoids prepared statements
	}), &gorm.Config{
		PrepareStmt: false, // Disable prepared statement caching
	})

	// Handle connection errors
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Set database connection pooling settings
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get the generic database object: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)                 // Maximum number of idle connections
	sqlDB.SetMaxOpenConns(20)                 // Maximum number of open connections (Supabase free tier allows 20)
	sqlDB.SetConnMaxLifetime(time.Minute * 5) // Maximum lifetime of a connection

	log.Println("Database connection successfully established")
}
