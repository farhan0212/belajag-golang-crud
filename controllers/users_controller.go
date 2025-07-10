package controllers

import (
	"belaja-golang-crud/services"
	"belaja-golang-crud/utils"
	"net/http"
	"strconv"
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
