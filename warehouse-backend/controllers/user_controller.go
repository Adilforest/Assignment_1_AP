package controllers

import (
	"errors"
	"warehouse-backend/database"
	"warehouse-backend/models"
)

// CreateUser creates a new user in the database
func CreateUser(name, email, password string) (models.User, error) {
	user := models.User{Name: name, Email: email, Password: password}
	result := database.DB.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

// GetUserByID retrieves a user by their ID
func GetUserByID(id uint) (models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New("user not found")
	}
	return user, result.Error
}

// UpdateUserByID updates a user's details by their ID
func UpdateUserByID(id uint, name, email, password string) (models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.User{}, errors.New("user not found")
	}
	// Update fields
	user.Name = name
	user.Email = email
	user.Password = password
	saveResult := database.DB.Save(&user)
	return user, saveResult.Error
}

// DeleteUserByID deletes a user by their ID
func DeleteUserByID(id uint) error {
	result := database.DB.Delete(&models.User{}, id)
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return result.Error
}

// GetAllUsers returns a list of all users
func GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result.Error
}
