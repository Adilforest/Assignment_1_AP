package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"warehouse-backend/controllers"
)

func UserRoutes(router chi.Router) {
	router.Post("/users", func(w http.ResponseWriter, r *http.Request) {
		// Получение данных из тела запроса
		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := controllers.CreateUser(name, email, password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write([]byte("User created: " + user.Name))
	})

	router.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(chi.URLParam(r, "id"))
		user, err := controllers.GetUserByID(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.Write([]byte("User: " + user.Name))
	})
}
