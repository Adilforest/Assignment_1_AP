package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse-backend/database"
	"warehouse-backend/models"
)

// CreateProduct — обработчик POST-запроса для создания нового продукта
func CreateProduct(c *gin.Context) {
	var product models.Product

	// Привязка JSON-данных в структуру product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Сохраняем продукт в базе данных
	if result := database.DB.Create(&product); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Возвращаем успешный ответ
	c.JSON(http.StatusOK, gin.H{
		"message": "Product created successfully",
		"product": product,
	})
}

// SetupProductRoutes — функция для настройки маршрутов, связанных с продуктами
func SetupProductRoutes(r *gin.Engine) {
	r.POST("/products", CreateProduct)
}
