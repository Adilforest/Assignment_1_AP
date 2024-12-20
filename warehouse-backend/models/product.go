package models

import "fmt"

const ProductTableName = "products"

type PriceType float64

type Product struct {
	ProductID   uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Price       PriceType `gorm:"not null"`
	Quantity    int       `gorm:"not null"`
}

func (p *Product) TableName() string {
	return ProductTableName
}

func (p *Product) Debug() string {
	return fmt.Sprintf("Product[ID=%d, Name='%s', Price=%.2f, Quantity=%d]", p.ProductID, p.Name, p.Price, p.Quantity)
}
