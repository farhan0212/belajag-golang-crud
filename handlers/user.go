package handlers

import (
	"belaja-golang-crud/db"
	"belaja-golang-crud/models"
	"belaja-golang-crud/utils"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

// CreateUser handles HTTP POST requests to create a new user.
// It decodes the JSON request body into a User model, attempts to save it to the database,
// and returns the created user as a JSON response with HTTP status 201 (Created).
// If decoding or database operations fail, it responds with the appropriate HTTP error status.

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "gagal mengambil data")
		return
	}

	if strings.TrimSpace(user.Name) == "" {
		utils.ResponseError(w, http.StatusForbidden, "Username harus diisi")
		return
	}

	if err := db.DB.Create(&user).Error; err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "gagal create user")
		return
	}

	utils.ResponseJSON(w, http.StatusAccepted, true, "Sukses input user", nil)
}

// GetUsers handles HTTP requests to retrieve all users from the database.
// It queries the database for all user records, encodes the result as JSON,
// and writes it to the response. If an error occurs during the database query,
// it responds with a 400 Bad Request status and an error message.
//
// Parameters:
//   - w: http.ResponseWriter to write the HTTP response.
//   - r: *http.Request representing the incoming HTTP request.

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

	offset := (page - 1) * limit

	var users []models.User
	var total int64

	db.DB.Model(&models.User{}).Count(&total)

	result := db.DB.Limit(limit).Offset(offset).Find(&users)

	if result.Error != nil {
		utils.ResponseError(w, http.StatusNotFound, "gagal mengambil data")
		return
	}

	totalPage := int(math.Ceil(float64(total) / float64(limit)))

	response := map[string]interface{}{
		"page":        page,
		"limit":       limit,
		"total":       total,
		"total_pages": totalPage,
		"data":        users,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetUserByEmailOrName(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	email := r.URL.Query().Get("email")

	var user = models.User{}
	query := db.DB

	if name != "" {
		query = query.Where("name = ?", name)
	}

	if email != "" {
		query = query.Where("email = ?", email)
	}

	err := query.First(&user).Error

	if err != nil {
		utils.ResponseError(w, http.StatusNotFound, "user tidak ditemukan")
		return
	}
	utils.ResponseJSON(w, http.StatusOK, true, "user ditemukan", user)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Id tidak valid", http.StatusNotFound)
		return
	}
	var user models.User

	if err := db.DB.First(&user, id).Error; err != nil {
		http.Error(w, "User tidak ditemukkan", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

	// userId := r.URL.Query().Get("id")

	// var user = models.User{}

	// err := db.DB.First(&user.ID)

	// if err != nil {
	// 	http.Error(w, "User tidak ditemukan", http.StatusNotFound)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(userId)
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		http.Error(w, "Id tidak valid", http.StatusNotFound)
		return
	}

	var input models.User

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "gagal decode json", http.StatusBadRequest)
		return
	}

	result := db.DB.Model(&models.User{}).Where("id = ?", id).Updates(models.User{
		Name:  input.Name,
		Email: input.Email,
	})

	if result.RowsAffected == 0 {
		http.Error(w, "user tidak ditemukan", http.StatusNotFound)
		return
	}
	if result.Error != nil {
		http.Error(w, "Gagal update user", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User dengan ID %d berhasil di update", id)

}

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
