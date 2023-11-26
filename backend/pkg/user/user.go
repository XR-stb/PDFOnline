package user

import (
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"backend/pkg/config"
	"backend/pkg/database"
	"backend/pkg/database/models"
	"backend/pkg/user/role"
	"backend/pkg/util"
)

var (
	ErrUserAlreadyExist = errors.New("username already exist")
	ErrForbidden        = errors.New("forbidden")
)

var (
	GuestUser = models.User{
		Id:       "guest",
		Username: "guestUser",
		Password: "guestUser",
		Role:     role.RoleGuest,
	}
	AdminUser models.User
)

func Create(username, password string, role role.Role) (*models.User, error) {
	u := models.User{
		Id:       uuid.New().String(),
		Username: username,
		Password: util.MD5(password),
		Role:     role,
	}

	if err := database.Instance().Create(&u).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "UNIQUE") {
			return nil, ErrUserAlreadyExist
		}
		return nil, err
	}

	return &u, nil
}

func Get(id string) (*models.User, error) {
	u := &models.User{}
	err := database.Instance().First(u, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func Verify(username, password string) (*models.User, error) {
	if username == GuestUser.Username && password == GuestUser.Password {
		return &GuestUser, nil
	}

	u := &models.User{}
	err := database.Instance().First(u, "username = ? AND password = ?", username, util.MD5(password)).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func CreateAdminUser() error {
	db := database.Instance()

	var user models.User
	if err := db.First(&user, "username = ?", config.AdminUser()).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			user = models.User{
				Id:       uuid.New().String(),
				Username: config.AdminUser(),
				Password: util.MD5(config.AdminPassword()),
				Role:     role.RoleAdmin,
			}
			if err := db.Create(&user).Error; err != nil {
				return fmt.Errorf("failed to create admin user: %v", err)
			}
			logrus.Infof("\u001B[31madmin user: %s, password: %s, it will not be shown again, please keep it\u001B[0m", config.AdminUser(), config.AdminPassword())
		} else {
			return fmt.Errorf("failed to query admin user: %v", err)
		}
	}

	if user.Role != role.RoleAdmin {
		return fmt.Errorf("admin user exist, but its role is not admin")
	}

	AdminUser = user

	return nil
}

type UpdateOption struct {
	Id       string
	UserId   string
	IsAdmin  bool
	Username *string
	Password *string
	Role     *role.Role
}

func Update(opt *UpdateOption) error {
	db := database.Instance()

	var user models.User
	if err := db.First(&user, "id = ?", opt.Id).Error; err != nil {
		return err
	}

	if !opt.IsAdmin && user.Id != opt.UserId {
		return ErrForbidden
	}

	if opt.Username != nil {
		user.Username = *opt.Username
	}
	if opt.Password != nil {
		user.Password = util.MD5(*opt.Password)
	}

	if opt.Role != nil {
		user.Role = *opt.Role
	}

	if err := db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
