package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"warehouse-backend/database"
	"warehouse-backend/models"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	// Подключение к базе данных
	database.ConnectPostgres()
	database.Migrate()

	// Автоматическая миграция
	err := database.DB.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatalf("Error migrating the database: %v", err)
	}

	log.Println("Database migrated successfully!")

	// Настройка маршрутов
	r := gin.Default()

	// Добавляем поддержку CORS
	r.Use(cors.Default()) // Разрешает все источники (для разработки)

	// Главная страница
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to the Warehouse Backend!")
	})

	// Обработчик для GET запроса
	r.GET("/get", func(c *gin.Context) {
		response := map[string]string{
			"status":  "success",
			"message": "GET запрос успешен!",
		}
		c.JSON(200, response)
	})

	// Обработчик для POST запроса
	r.POST("/post", func(c *gin.Context) {
		var msg Message
		// Декодируем JSON из тела запроса
		if err := c.ShouldBindJSON(&msg); err != nil || msg.Message == "" {
			// Возвращаем ошибку, если данные не валидны
			c.JSON(400, gin.H{
				"status":  "fail",
				"message": "Некорректное JSON-сообщение",
			})
			return
		}

		// Формируем ответ
		response := map[string]string{
			"status":  "success",
			"message": "Данные успешно приняты",
		}

		// Отправляем ответ
		c.JSON(200, response)
	})

	// Запуск сервера
	log.Println("Server is running on port 8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
