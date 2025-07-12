package services

import (
	"belaja-golang-crud/models"
	"belaja-golang-crud/repository"
	"belaja-golang-crud/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func CreateUserService(user models.User) (*models.User, error) {
	if user.Email == "" || user.Password == "" || user.Name == "" {
		return nil, errors.New("Semua field harus diisi ya")
	}

	if user.Role == "" {
		user.Role = "user"
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("gagal mengenkripsi password")
	}

	user.Password = string(hashedPassword)
	if err := repository.CreateUser(&user); err != nil {
		return nil, errors.New("gagal create user")
	}

	return &user, nil
}

func LoginUser(req models.LoginRequest) (map[string]interface{}, error) {
	user, err := repository.LoginRequest(req.Email)
	if err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("Email atau password salah")
	}
	token, err := utils.GenerateJWT(user.ID, user.Email, user.Role)
	if err != nil {
		return nil, errors.New("gagal membuat token")
	}
	return map[string]interface{}{
		"email": user.Email,
		"name":  user.Name,
		"role":  user.Role,
		"token": token,
	}, nil
}
