package repository

import (
	"belaja-golang-crud/db"
	"belaja-golang-crud/models"
)

func CreateUser(user *models.User) error {
	return db.DB.Create(user).Error
}

func IsEmailExist(email string) bool {
	var user []models.User

	err := db.DB.Where("email = ?", email).First(&user).Error
	return err == nil
}

func LoginRequest(email string) (*models.User, error) {
	var user models.User

	if err := db.DB.Where("email = ?", user.Email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CountUsers() (int64, error) {
	var total int64

	err := db.DB.Model(&models.User{}).Count(&total).Error
	return total, err
}

func GetUsers(limit, offset int) ([]models.User, error) {
	var users []models.User
	err := db.DB.Limit(limit).Offset(offset).Order("created_at DESC").Find(&users).Error

	return users, err
}
