package routes

import (
	"belaja-golang-crud/controllers"

	"github.com/go-chi/chi/v5"
)

func UserRoutes(r chi.Router) {

	r.Post("/login", controllers.LoginUser)
	r.Post("/", controllers.CreateUser)

	r.Group(func(r chi.Router) {
		// r.Use(middleware.AuthMiddleware)
		// r.Use(middleware.IsAdmin)
		r.Get("/{id}", controllers.GetUserById)
		r.Get("/", controllers.GetUserByEmailOrName)
		r.Put("/{id}", controllers.UpdateUserById)
		r.Delete("/{id}", controllers.DeleteUser)
	})
}

func UserListGo(r chi.Router) {
	// r.Use(middleware.AuthMiddleware)
	r.Get("/", controllers.GetUsers)
}
