package handlers

import (
	"belaja-golang-crud/db"
	"belaja-golang-crud/models"

	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Id tidak valid", http.StatusNotFound)
		return
	}

	result := db.DB.Delete(&models.User{}, id)

	if result.RowsAffected == 0 {
		http.Error(w, "User tidak ditemukan", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "data %d berhasil didelete", id)
}
