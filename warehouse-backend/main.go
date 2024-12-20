package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"warehouse-backend/database"
	"warehouse-backend/models"
)

const (
	serverPort = ":8080"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	router := setupRoutes()

	log.Println("Server is running on port " + serverPort)
	if err := router.Run(serverPort); err != nil {
		log.Fatal(err)
	}
}

func setupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:63342"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	router.GET("/", handleHome)
	router.GET("/get", handleGetRequest)
	router.POST("/post", handlePostRequest)

	productRoutes := router.Group("/products")
	{
		productRoutes.POST("/create", createProductHandler)
		productRoutes.GET("/:id", getProductHandler)
		productRoutes.GET("/", getAllProductsHandler)
		productRoutes.PUT("/:id", updateProductHandler)
		productRoutes.DELETE("/:id", deleteProductHandler)
	}

	return router
}

//  General Handlers

func handleHome(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to the Warehouse Backend!")
}

func handleGetRequest(c *gin.Context) {
	response := createResponse("success", "GET request processed successfully")
	c.JSON(http.StatusOK, response)
}

func handlePostRequest(c *gin.Context) {
	var message Message
	if err := c.ShouldBindJSON(&message); err != nil || message.Message == "" {
		c.JSON(http.StatusBadRequest, createResponse("fail", "Invalid or empty JSON message"))
		return
	}

	c.JSON(http.StatusOK, createResponse("success", "Data received successfully"))
}

func createResponse(status, message string) map[string]string {
	return map[string]string{
		"status":  status,
		"message": message,
	}
}

// CRUD Handlers for Product

func createProductHandler(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, createResponse("fail", "Invalid JSON payload"))
		return
	}

	if err := database.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, createResponse("fail", "Failed to create product"))
		return
	}

	c.JSON(http.StatusCreated, product)
}

func getProductHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, createResponse("fail", "Invalid product ID"))
		return
	}

	product, err := database.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, createResponse("fail", "Product not found"))
		return
	}

	c.JSON(http.StatusOK, product)
}

func getAllProductsHandler(c *gin.Context) {
	products, err := database.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, createResponse("fail", "Failed to fetch products"))
		return
	}

	c.JSON(http.StatusOK, products)
}

func updateProductHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, createResponse("fail", "Invalid product ID"))
		return
	}

	var updatedProduct models.Product
	if err := c.ShouldBindJSON(&updatedProduct); err != nil {
		c.JSON(http.StatusBadRequest, createResponse("fail", "Invalid JSON payload"))
		return
	}

	err = database.UpdateProduct(uint(id), &updatedProduct)
	if err != nil {
		c.JSON(http.StatusNotFound, createResponse("fail", err.Error()))
		return
	}

	c.JSON(http.StatusOK, createResponse("success", "Product updated successfully"))
}

func deleteProductHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, createResponse("fail", "Invalid product ID"))
		return
	}

	err = database.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, createResponse("fail", err.Error()))
		return
	}

	c.JSON(http.StatusOK, createResponse("success", "Product deleted successfully"))
}
