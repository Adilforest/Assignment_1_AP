package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"warehouse-backend/config"
	"warehouse-backend/models"
)

var DB *gorm.DB

// ConnectPostgres подключается к базе данных PostgreSQL
func ConnectPostgres() {
	cfg := config.GetConfig()

	dsn := cfg.GetPostgresDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	DB = db
	fmt.Println("Successfully connected to PostgreSQL!")

	// Миграция для создания таблицы продуктов
	Migrate()
}

// Migrate выполняет миграцию базы данных, создавая таблицы, если они не существуют
func Migrate() {
	err := DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalf("Error migrating the database: %v", err)
	} else {
		fmt.Println("Database migrated successfully!")
	}
}
