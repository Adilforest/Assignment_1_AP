package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"warehouse-backend/database"
)

const serverPort = ":8080"

// Message represents a structure to decode received JSON payloads
type Message struct {
	Message string `json:"message"`
}

func main() {
	setupDatabase()
	router := setupRoutes()

	log.Println("Server is running on port" + serverPort + "...")
	if err := router.Run(serverPort); err != nil {
		log.Fatal(err)
	}
}

// setupDatabase connects to the database and applies migrations
func setupDatabase() {
	database.ConnectPostgres()
}

// setupRoutes configures and returns a new gin router with routes and middleware
func setupRoutes() *gin.Engine {
	router := gin.Default()

	// Enable CORS for development
	router.Use(cors.Default())

	// Define Routes
	router.GET("/", handleHome)
	router.GET("/get", handleGetRequest)
	router.POST("/post", handlePostRequest)

	return router
}

// handleHome handles the home route ("/")
func handleHome(c *gin.Context) {
	c.String(200, "Welcome to the Warehouse Backend!")
}

// handleGetRequest processes GET requests to "/get"
func handleGetRequest(c *gin.Context) {
	jsonResponse := createResponse("success", "GET запрос успешен!")
	c.JSON(200, jsonResponse)
}

// handlePostRequest processes POST requests to "/post"
func handlePostRequest(c *gin.Context) {
	var message Message
	if err := c.ShouldBindJSON(&message); err != nil || message.Message == "" {
		// Return failure response if body is invalid or empty
		c.JSON(400, createResponse("fail", "Некорректное JSON-сообщение"))
		return
	}

	// Successfully received data
	c.JSON(200, createResponse("success", "Данные успешно приняты"))
}

// createResponse returns a standardized JSON response
func createResponse(status, message string) map[string]string {
	return map[string]string{
		"status":  status,
		"message": message,
	}
}
