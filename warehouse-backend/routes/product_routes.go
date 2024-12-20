package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse-backend/database"
	"warehouse-backend/models"
)

const productCreatedMessage = "Product created successfully"

func HandleCreateProduct(c *gin.Context) {
	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		handleError(c, http.StatusBadRequest, err.Error())
		return
	}

	if result := database.DB.Create(&product); result.Error != nil {
		handleError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": productCreatedMessage,
		"product": product,
	})
}

func SetupProductRoutes(router *gin.Engine) {
	router.POST("/products", HandleCreateProduct)
}

func handleError(c *gin.Context, statusCode int, errorMessage string) {
	c.JSON(statusCode, gin.H{"error": errorMessage})
}
