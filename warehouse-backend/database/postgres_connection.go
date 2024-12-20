package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"warehouse-backend/config"
	"warehouse-backend/models"
)

var Database *gorm.DB

const dbConnectionSuccessMsg = "Successfully connected to PostgreSQL!"

// ConnectPostgres establishes a connection to the PostgreSQL database and performs automatic migrations.
func ConnectPostgres() {
	db := initializeDatabaseConnection()
	Database = db
	fmt.Println(dbConnectionSuccessMsg)

	// Perform migration for tables
	if err := Database.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("Error migrating the database: %v", err)
	}
	fmt.Println("Database migrated successfully!")
}

// initializeDatabaseConnection handles the database connection setup logic.
func initializeDatabaseConnection() *gorm.DB {
	cfg := config.GetConfig()
	dsn := cfg.GetPostgresDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database. DSN: %v, Error: %v", dsn, err)
	}
	return db
}
