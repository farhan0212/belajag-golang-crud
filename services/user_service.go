package services

import (
	"belaja-golang-crud/models"
	"belaja-golang-crud/repository"
	"errors"
	"math"
)

func GetPaginationUsers(page, limit int) (*models.PaginationResponse, error) {
	offset := (page - 1) * limit

	total, err := repository.CountUsers()
	if err != nil {
		return nil, err
	}

	users, err := repository.GetUsers(limit, offset)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &models.PaginationResponse{
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
		Data:       users,
	}, nil
}

func GetUserByEmailOrName(name, email string) (models.User, error) {
	if name == "" && email == "" {
		return models.User{}, errors.New("Nama atau email harus diisi")
	}
	user, err := repository.GetUserByEmailOrName(name, email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetUserById(id int) (models.User, error) {
	if id <= 0 {
		return models.User{}, errors.New("ID tidak valid")
	}

	user, err := repository.GetUserById(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func UpdateUserById(id int, name string, email string) (models.User, error) {
	user, err := repository.UpdateUserById(id, name, email)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func DeleteUser(id int) (models.User, error) {
	user, err := repository.DeleteUser(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
