package models

import "fmt"

// Constants
const ProductTableName = "products"

// Type Aliases
type PriceType float64

// Product represents a product in the database.
type Product struct {
	ProductID   uint      `gorm:"primaryKey"` // Renamed for clarity
	Name        string    `gorm:"not null"`   // Product name
	Description string    `gorm:"not null"`   // Detailed description
	Price       PriceType `gorm:"not null"`   // Product price
	Quantity    int       `gorm:"not null"`   // Available quantity
}

// TableName returns the database table name for the Product model.
func (p *Product) TableName() string {
	return ProductTableName
}

// Debugging (Optional Example for Developers)
func (p *Product) Debug() string {
	return fmt.Sprintf("Product[ID=%d, Name='%s', Price=%.2f, Quantity=%d]", p.ProductID, p.Name, p.Price, p.Quantity)
}
