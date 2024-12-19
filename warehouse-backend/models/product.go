package models

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null"`
	Price       float64 `gorm:"not null"`
	Quantity    int     `gorm:"not null"`
}

func (p *Product) TableName() string {
	return "products"
}
