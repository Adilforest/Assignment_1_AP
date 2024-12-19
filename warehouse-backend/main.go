package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"warehouse-backend/database"
)

type Message struct {
	Message string `json:"message"`
}

func main() {
	// Подключение к базе данных
	database.ConnectPostgres()

	// Обработчик для GET запроса
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]string{
			"status":  "success",
			"message": "GET запрос успешен!",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Обработчик для POST запроса
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		// Проверяем метод запроса
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		// Декодируем JSON из тела запроса
		var msg Message
		err := json.NewDecoder(r.Body).Decode(&msg)
		if err != nil || msg.Message == "" {
			// Возвращаем ошибку, если данные не валидны
			http.Error(w, `{"status": "fail", "message": "Некорректное JSON-сообщение"}`, http.StatusBadRequest)
			return
		}

		// Формируем ответ
		response := map[string]string{
			"status":  "success",
			"message": "Данные успешно приняты",
		}

		// Отправляем ответ
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	// Пример простого домашнего маршрута для теста
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the Warehouse Backend!")
	})

	// Запуск сервера
	fmt.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
