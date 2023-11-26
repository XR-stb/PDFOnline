package user

import (
	"backend/pkg/database"
	"backend/pkg/database/models"
	"errors"
	"github.com/google/uuid"
	"strings"
)

var (
	ErrUserAlreadyExist = errors.New("user already exist")
)

func Create(username, password string, role UserRole) (string, error) {
	u := models.User{
		Id:       uuid.New().String(),
		Username: username,
		Password: password,
		Role:     string(role),
	}

	if err := database.Instance().Create(&u).Error; err != nil {
		if strings.Contains(err.Error(), "UNIQUE") {
			return "", ErrUserAlreadyExist
		}
		return "", err
	}

	return u.Id, nil
}

func Get(id string) (*models.User, error) {
	u := &models.User{}
	err := database.Instance().First(u, "id = ?", id).Error
	return u, err
}

func Verify(username, password string) (string, error) {
	u := &models.User{}
	err := database.Instance().First(u, "username = ? AND password = ?", username, password).Error
	return u.Id, err
}
