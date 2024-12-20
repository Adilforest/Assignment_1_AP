package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"warehouse-backend/database"
	"warehouse-backend/models"
)

const productCreatedMessage = "Product created successfully" // Success message constant

// HandleCreateProduct handles the POST request for creating a new product
func HandleCreateProduct(c *gin.Context) {
	var product models.Product

	// Bind JSON data into the Product struct and check for errors
	if err := c.ShouldBindJSON(&product); err != nil {
		handleError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Save the product in the database and check for errors
	if result := database.DB.Create(&product); result.Error != nil {
		handleError(c, http.StatusInternalServerError, result.Error.Error())
		return
	}

	// Return successful response
	c.JSON(http.StatusOK, gin.H{
		"message": productCreatedMessage,
		"product": product,
	})
}

// SetupProductRoutes configures routes related to products
func SetupProductRoutes(router *gin.Engine) {
	router.POST("/products", HandleCreateProduct)
}

// handleError is a helper function for sending an error response
func handleError(c *gin.Context, statusCode int, errorMessage string) {
	c.JSON(statusCode, gin.H{"error": errorMessage})
}
