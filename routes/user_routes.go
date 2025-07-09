package routes

import (
	"belaja-golang-crud/handlers"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router) {
	r.Post("/", handlers.CreateUser)
	r.Get("/{id}", handlers.GetUserById)
	r.Get("/", handlers.GetUserByEmailOrName)
	r.Put("/{id}", handlers.UpdateUserById)
	r.Delete("/{id}", handlers.DeleteUser)
}

func UserListGo(r chi.Router) {
	r.Get("/", handlers.GetUsers)
}
