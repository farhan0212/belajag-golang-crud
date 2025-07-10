package main

import (
	"belaja-golang-crud/db"
	"belaja-golang-crud/routes"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {

	db.InitDB()

	r := chi.NewRouter()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("request",
				"method", r.Method,
				"path", r.URL.Path,
				"ip", r.RemoteAddr,
			)
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
