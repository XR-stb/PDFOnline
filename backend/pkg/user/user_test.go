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
		user, err := Create("testUser", "123456", "email@example.com", role.RoleUser)
		assert.NoError(t, err)

		u := models.User{}
		database.Instance().Where("id = ?", user.Id()).Find(&u)
		assert.Equal(t, "testUser", u.Username)
	})

	t.Run("duplicate", func(t *testing.T) {
		_, err := Create("testUser", "123456", "email@example.com", role.RoleUser)
		assert.ErrorIs(t, err, ErrUserAlreadyExist)
	})
}

func TestGetById(t *testing.T) {
	database.Use(testutil.TestDB(t))
	testUser := models.User{
		Id:       uuid.New().String(),
		Username: "testUser",
		Role:     role.RoleUser,
	}
	database.Instance().Create(&testUser)

	t.Run("success", func(t *testing.T) {
		u, err := GetById(testUser.Id)
		assert.NoError(t, err)
		assert.EqualValues(t, testUser, u.Show())
	})

	t.Run("not found", func(t *testing.T) {
		_, err := GetById("NOT_EXIST")
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	})
}

func TestGetByUsername(t *testing.T) {
	database.Use(testutil.TestDB(t))
	testUser := models.User{
		Id:       uuid.New().String(),
		Username: "testUser",
		Role:     role.RoleUser,
	}
	database.Instance().Create(&testUser)

	t.Run("success", func(t *testing.T) {
		u, err := GetByUsername(testUser.Username)
		assert.NoError(t, err)
		assert.EqualValues(t, testUser, u.Show())
	})

	t.Run("not found", func(t *testing.T) {
		_, err := GetByUsername("NOT_EXIST")
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	})
}

func TestCreateInternalUser(t *testing.T) {
	database.Use(testutil.TestDB(t))
	config.Cfg = &config.Config{
		Default: config.Default{
			AdminUsername: "admin",
			AdminPassword: "123456",
		},
	}
	err := CreateInternalUser()
	assert.NoError(t, err)

	var adminUser models.User
	err = database.Instance().First(&adminUser, "username = ?", "admin").Error
	assert.NoError(t, err)
	assert.Equal(t, "admin", adminUser.Username)
	assert.Equal(t, role.RoleAdmin, adminUser.Role)

	var guestUser models.User
	err = database.Instance().First(&guestUser, "username = ?", "guest").Error
	assert.NoError(t, err)
	assert.Equal(t, "guest", guestUser.Username)
	assert.Equal(t, role.RoleGuest, guestUser.Role)
}

func TestUser_VerifyPassword(t *testing.T) {
	testUser := models.User{
		Id:       uuid.New().String(),
		Username: "testUser",
		Password: util.MD5("123456"),
		Role:     role.RoleUser,
	}

	u := User{
		user: &testUser,
	}

	t.Run("success", func(t *testing.T) {
		ok := u.VerifyPassword("123456")
		assert.True(t, ok)
	})

	t.Run("wrong password", func(t *testing.T) {
		ok := u.VerifyPassword("1234567")
		assert.False(t, ok)
	})
}

func TestUser_Update(t *testing.T) {
	database.Use(testutil.TestDB(t))
	db := database.Instance()
	testUser := models.User{
		Id:       uuid.New().String(),
		Username: "testUser",
		Password: util.MD5("123456"),
		Role:     role.RoleUser,
	}
	db.Create(&testUser)

	u := User{
		user: &testUser,
		db:   db,
	}

	t.Run("success", func(t *testing.T) {
		err := u.Update(&UpdateOption{
			Username: testutil.StringPtr("testUser"),
		})
		assert.NoError(t, err)

		var u models.User
		db.First(&u, "id = ?", testUser.Id)
		assert.Equal(t, "testUser", u.Username)
	})
}
