package controllers

import (
	"belaja-golang-crud/services"
	"belaja-golang-crud/utils"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 10

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}
	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}
	users, err := services.GetPaginationUsers(page, limit)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, "gagal mengambil data")
		return
	}
	utils.ResponseJSON(w, http.StatusOK, true, "Sukses get data", users)
}

func GetUserByEmailOrName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")

	user, err := services.GetUserByEmailOrName(name, email)
	if err != nil {
		if err.Error() == "nama atau email harus diisi" {
			utils.ResponseError(w, http.StatusBadRequest, err.Error())
			return
		}
		utils.ResponseError(w, http.StatusNotFound, "user tidak ditemukan")
	}
	utils.ResponseJSON(w, http.StatusOK, true, "data ditemukan", user)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "ID tidak valid")
		return
	}

	user, err := services.GetUserById(id)
	if err != nil {
		utils.ResponseError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.ResponseJSON(w, http.StatusOK, true, "Id ditemukan", user)
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Id tidak valid")
		return
	}

	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "Format tidak valid")
	}

	user, err := services.UpdateUserById(id, req.Name, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ResponseError(w, http.StatusNotFound, "user tidak ditemukan")
		} else {
			utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	utils.ResponseJSON(w, http.StatusOK, true, "User berhasil diupdate", user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Id tidak valid", http.StatusNotFound)
		return
	}

	user, err := services.DeleteUser(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			utils.ResponseError(w, http.StatusNotFound, "User tidak ditemukan")
		} else {
			utils.ResponseError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	utils.ResponseJSON(w, http.StatusOK, true, "User berhasil didelete", user)
}
