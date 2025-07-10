package middleware

import (
	"belaja-golang-crud/utils"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.ResponseError(w, http.StatusUnauthorized, "Authorization header tidak valid")
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		})

		if err != nil || !token.Valid {
			utils.ResponseError(w, http.StatusUnauthorized, "token tidak valid")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func IsAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Context().Value("user").(jwt.MapClaims)

		if token["role"] != "admin" {
			utils.ResponseError(w, http.StatusForbidden, "Hanya admin")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func IsUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Context().Value("user").(jwt.MapClaims)

		if token["role"] != "user" {
			utils.ResponseError(w, http.StatusForbidden, "Hanya User")
			return
		}
		next.ServeHTTP(w, r)
	})
}
