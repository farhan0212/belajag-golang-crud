package handlers

import (
	"belaja-golang-crud/db"
	"belaja-golang-crud/models"
	"belaja-golang-crud/utils"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "gagal mengambil data")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "gagal mengenkripsi password")
		return
	}

	user.Password = string(hashedPassword)

	if user.Email == "" || user.Password == "" || user.Name == "" {
		utils.ResponseError(w, http.StatusForbidden, "semua field harus diisi yaa")
		return
	}

	if user.Role == "" {
		user.Role = "user"
	}

	if err := db.DB.Create(&user).Error; err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "gagal create user")
		return
	}

	utils.ResponseJSON(w, http.StatusAccepted, true, "Sukses input user", map[string]interface{}{
		"name":  user.Name,
		"email": user.Email,
		"role":  user.Role,
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	var req models.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "gagal membaca data login")
		return
	}

	var user models.User

	if err := db.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		utils.ResponseError(w, http.StatusNotFound, "user tidak ditemukan")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		utils.ResponseError(w, http.StatusUnauthorized, "Email atau password salah")
		return
	}

	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		utils.ResponseError(w, http.StatusInternalServerError, "Gagal membuat token")
		return
	}

	utils.ResponseJSON(w, http.StatusOK, true, "berhasil login", map[string]interface{}{
		"email": user.Email,
		"name":  user.Name,
		"token": token,
		"role":  user.Role,
	})
}
