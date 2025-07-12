package controllers

import (
	"belaja-golang-crud/models"
	"belaja-golang-crud/repository"
	"belaja-golang-crud/services"
	"belaja-golang-crud/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "gagal mengambil data")
		return
	}
	if repository.IsEmailExist(user.Email) {
		utils.ResponseError(w, http.StatusBadRequest, "Email sudah digunakan")
		return
	}

	createdUser, err := services.CreateUserService(user)
	if err != nil {
		utils.ResponseError(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.ResponseJSON(w, http.StatusCreated, true, "sukses input user", map[string]interface{}{
		"name":  createdUser.Name,
		"email": createdUser.Email,
		"role":  createdUser.Role,
	})

}
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.ResponseError(w, http.StatusBadRequest, "gagal membaca data login")
		return
	}

	fmt.Println("received : %s, pass : %s\n", req.Email, req.Password)

	if req.Email == "" || req.Password == "" {
		utils.ResponseError(w, http.StatusBadRequest, "Email dan password harus diisi")
		return
	}

	result, err := services.LoginUser(req)
	if err != nil {
		utils.ResponseError(w, http.StatusUnauthorized, err.Error())
		return
	}
	utils.ResponseJSON(w, http.StatusOK, true, "berhasil login", result)
}
