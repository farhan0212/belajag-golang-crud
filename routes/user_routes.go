package routes

import (
	"belaja-golang-crud/controllers"
	"belaja-golang-crud/handlers"
	"belaja-golang-crud/middleware"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router) {

	r.Post("/login", controllers.LoginUser)
	r.Post("/", controllers.CreateUser)

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Use(middleware.IsAdmin)
		r.Get("/{id}", handlers.GetUserById)
		r.Get("/", handlers.GetUserByEmailOrName)
		r.Put("/{id}", handlers.UpdateUserById)
		r.Delete("/{id}", handlers.DeleteUser)
	})
}

func UserListGo(r chi.Router) {
	r.Use(middleware.AuthMiddleware)
	r.Get("/", handlers.GetUsers)
}
