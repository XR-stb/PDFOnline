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
)

type User struct {
	user *models.User
	db   *gorm.DB
}

func Create(username, password, email string, role role.Role) (*User, error) {
	db := database.Instance()
	user := models.User{
		Id:       uuid.New().String(),
		Username: username,
		Password: util.MD5(password),
		Email:    email,
		Role:     role,
	}

	if err := db.Create(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) || strings.Contains(err.Error(), "UNIQUE") {
			return nil, ErrUserAlreadyExist
		}
		return nil, err
	}

	return &User{
		user: &user,
		db:   db,
	}, nil
}

func GetById(id string) (*User, error) {
	db := database.Instance()
	user := models.User{}
	return &User{
		user: &user,
		db:   db,
	}, db.First(&user, "id = ?", id).Error
}

func GetByUsername(username string) (*User, error) {
	db := database.Instance()
	user := models.User{}
	return &User{
		user: &user,
		db:   db,
	}, db.First(&user, "username = ?", username).Error
}

func GetByEmail(email string) (*User, error) {
	db := database.Instance()
	user := models.User{}
	return &User{
		user: &user,
		db:   db,
	}, db.First(&user, "email = ?", email).Error
}

func CreateInternalUser() error {
	if err := createAdminUser(); err != nil {
		return err
	}

	return createGuestUser()
}

func createAdminUser() error {
	var u *User
	var err error
	if u, err = GetByUsername(config.AdminUsername()); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if u, err = Create(config.AdminUsername(), config.AdminPassword(), "", role.RoleAdmin); err != nil {
				return fmt.Errorf("failed to create admin user: %v", err)
			}
			logrus.Infof("\u001B[31mAdmin user created. username: %s, password: %s. It will not be shown again, please keep it\u001B[0m", config.AdminUsername(), config.AdminPassword())
		} else {
			return fmt.Errorf("failed to query admin user: %v", err)
		}
	}

	if u.user.Role != role.RoleAdmin {
		return fmt.Errorf("admin user exist, but its role is not admin")
	}

	return nil
}

func createGuestUser() error {
	var u *User
	var err error
	if u, err = GetByUsername("guest"); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			if u, err = Create("guest", "guest123", "guest@mail.com", role.RoleGuest); err != nil {
				return fmt.Errorf("failed to create guest user: %v", err)
			}
		} else {
			return fmt.Errorf("failed to query guest user: %v", err)
		}
	}

	if u.user.Role != role.RoleGuest {
		return fmt.Errorf("guest user exist, but its role is not guest")
	}

	return nil
}

func (u *User) Show() models.User {
	return *u.user
}

func (u *User) Id() string {
	return u.user.Id
}

func (u *User) Username() string {
	return u.user.Username
}

func (u *User) Role() role.Role {
	return u.user.Role
}

func (u *User) Email() string {
	return u.user.Email
}

func (u *User) VerifyPassword(password string) bool {
	return u.user.Password == util.MD5(password)
}

type UpdateOption struct {
	Username *string
	Password *string
	Role     *role.Role
}

func (u *User) Update(opt *UpdateOption) error {
	if opt.Username != nil {
		u.user.Username = *opt.Username
	}

	if opt.Password != nil {
		u.user.Password = util.MD5(*opt.Password)
	}

	if opt.Role != nil {
		u.user.Role = *opt.Role
	}

	return u.save()
}

func (u *User) save() error {
	return u.db.Save(u.user).Error
}
