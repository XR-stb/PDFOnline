package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"backend/pkg/config"
	"backend/pkg/database"
	"backend/pkg/database/models"
	"backend/pkg/user/role"
	"backend/pkg/util"
	"backend/test/testutil"
)

func TestCreate(t *testing.T) {
	database.Use(testutil.TestDB(t))

	t.Run("success", func(t *testing.T) {
		user, err := Create("test_user", "123456", role.RoleUser)
		assert.NoError(t, err)

		u := models.User{}
		database.Instance().Where("id = ?", user.Id).Find(&u)
		assert.Equal(t, "test_user", u.Username)
	})

	t.Run("duplicate", func(t *testing.T) {
		_, err := Create("test_user", "123456", role.RoleUser)
		assert.ErrorIs(t, err, ErrUserAlreadyExist)
	})
}

func TestVerify(t *testing.T) {
	database.Use(testutil.TestDB(t))
	expected := models.User{
		Id:       uuid.New().String(),
		Username: "test_user",
		Password: util.MD5("123456"),
		Role:     role.RoleUser,
	}
	database.Instance().Create(&expected)

	t.Run("success", func(t *testing.T) {
		user, err := Verify("test_user", "123456")
		assert.NoError(t, err)
		assert.Equal(t, expected.Id, user.Id)
	})

	t.Run("wrong username", func(t *testing.T) {
		_, err := Verify("test_user2", "123456")
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	})

	t.Run("wrong password", func(t *testing.T) {
		_, err := Verify("test_user", "1234567")
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	})

	t.Run("guest", func(t *testing.T) {
		user, err := Verify("guestUser", "guestUser")
		assert.NoError(t, err)
		assert.Equal(t, &GuestUser, user)
	})
}

func TestGet(t *testing.T) {
	database.Use(testutil.TestDB(t))
	expected := models.User{
		Id:       uuid.New().String(),
		Username: "test_user",
		Role:     role.RoleUser,
	}
	database.Instance().Create(&expected)

	t.Run("success", func(t *testing.T) {
		u, err := Get(expected.Id)
		assert.NoError(t, err)
		assert.EqualValues(t, expected, *u)
	})

	t.Run("not found", func(t *testing.T) {
		_, err := Get("NOT_EXIST")
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	})
}

func TestCreateAdminUser(t *testing.T) {
	database.Use(testutil.TestDB(t))
	config.Cfg = &config.Config{
		AdminUser:     "admin",
		AdminPassword: "123456",
	}

	t.Run("success", func(t *testing.T) {
		err := CreateAdminUser()
		assert.NoError(t, err)

		var user models.User
		database.Instance().First(&user, "username = ?", "admin")
		assert.Equal(t, "admin", user.Username)
		assert.Equal(t, role.RoleAdmin, user.Role)
	})

	t.Run("already exist", func(t *testing.T) {
		err := CreateAdminUser()
		assert.NoError(t, err)

		err = CreateAdminUser()
		assert.NoError(t, err)
	})

	t.Run("exist but not admin", func(t *testing.T) {
		config.Cfg.AdminUser = "adminUser"
		database.Instance().Create(&models.User{
			Id:       uuid.New().String(),
			Username: "adminUser",
			Password: "123456",
			Role:     role.RoleUser,
		})

		err := CreateAdminUser()
		assert.Equal(t, "admin user exist, but its role is not admin", err.Error())
	})
}

func TestUpdate(t *testing.T) {
	user := models.User{
		Id:   uuid.New().String(),
		Role: role.RoleUser,
	}
	database.Use(testutil.TestDB(t))
	db := database.Instance()
	db.Create(&user)

	t.Run("forbidden", func(t *testing.T) {
		err := Update(&UpdateOption{
			Id:       user.Id,
			UserId:   "FORBIDDEN",
			Username: testutil.StringPtr("test_user"),
		})
		assert.Equal(t, ErrForbidden, err)
	})

	t.Run("success", func(t *testing.T) {
		t.Run("oneself", func(t *testing.T) {
			err := Update(&UpdateOption{
				Id:       user.Id,
				UserId:   user.Id,
				Username: testutil.StringPtr("testUser"),
			})
			assert.NoError(t, err)

			var u models.User
			db.First(&u, "id = ?", user.Id)
			assert.Equal(t, "testUser", u.Username)
		})

		t.Run("admin", func(t *testing.T) {
			err := Update(&UpdateOption{
				Id:       user.Id,
				UserId:   user.Id,
				IsAdmin:  true,
				Username: testutil.StringPtr("testUser2"),
			})
			assert.NoError(t, err)

			var u models.User
			db.First(&u, "id = ?", user.Id)
			assert.Equal(t, "testUser2", u.Username)
		})
	})

	t.Run("not found", func(t *testing.T) {
		err := Update(&UpdateOption{
			Id:       "NOT_EXIST",
			UserId:   "NOT_EXIST",
			Username: testutil.StringPtr("testUser"),
		})
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	})
}
