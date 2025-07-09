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

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Incoming %s %s", r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})

	r.Route("/user", routes.UserRoutes)
	r.Route("/users", routes.UserListGo)

	fmt.Println("server berjalan di port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}
