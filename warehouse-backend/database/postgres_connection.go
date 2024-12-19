package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"warehouse-backend/config"
)

var DB *gorm.DB

func ConnectPostgres() {
	cfg := config.GetConfig()

	dsn := cfg.GetPostgresDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	DB = db
	fmt.Println("Successfully connected to PostgreSQL!")
}
