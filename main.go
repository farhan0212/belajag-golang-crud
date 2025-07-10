package main

import (
	"belaja-golang-crud/db"
	"belaja-golang-crud/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	db.InitDB()

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/user", routes.UserRoutes)
	r.Route("/users", routes.UserListGo)

	fmt.Println("server berjalan di port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}

}
