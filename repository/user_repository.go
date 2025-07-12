package repository

import (
	"belaja-golang-crud/db"
	"belaja-golang-crud/models"
	"errors"

	"gorm.io/gorm"
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

	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
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

func GetUserByEmailOrName(name, email string) (models.User, error) {
	var user models.User
	query := db.DB

	if name != "" && email != "" {
		query = query.Where("name = ? OR email = ?", name, email)
	} else if name != "" {
		query = query.Where("name = ?", name)
	} else if email != "" {
		query = query.Where("email = ?", email)
	} else {
		return user, errors.New("name atau email harus diisi")
	}

	err := query.First(&user).Error
	return user, err
}

func GetUserById(id int) (models.User, error) {
	var user models.User

	err := db.DB.First(&user, id).Error
	return user, err
}

func UpdateUserById(id int, name string, email string) (models.User, error) {
	var user models.User
	updates := make(map[string]interface{})

	if name != "" {
		updates["name"] = name
	}
	if email != "" {
		updates["email"] = email
	}

	if len(updates) == 0 {
		return user, errors.New("tidak ada field yang diupdate")
	}

	result := db.DB.Model(&models.User{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func DeleteUser(id int) (models.User, error) {
	var user models.User

	result := db.DB.Delete(&user, id)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
