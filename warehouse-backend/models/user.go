package models

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(100);not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}
