package user

import (
	"backend/pkg/database"
	"backend/pkg/database/models"
	testutil "backend/test/util"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

func TestCreate(t *testing.T) {
	database.Use(testutil.TestDB(t))

	t.Run("success", func(t *testing.T) {
		id, err := Create("test_user", "123456", RoleUser)
		assert.NoError(t, err)

		u := models.User{}
		database.Instance().Where("id = ?", id).Find(&u)
		assert.Equal(t, "test_user", u.Username)
	})

	t.Run("duplicate", func(t *testing.T) {
		_, err := Create("test_user", "123456", RoleUser)
		assert.ErrorIs(t, err, ErrUserAlreadyExist)
	})
}

func TestVerify(t *testing.T) {
	database.Use(testutil.TestDB(t))
	expected := models.User{
		Id:       uuid.New().String(),
		Username: "test_user",
		Password: "123456",
		Role:     "user",
	}
	database.Instance().Create(&expected)

	t.Run("success", func(t *testing.T) {
		id, err := Verify("test_user", "123456")
		assert.NoError(t, err)
		assert.Equal(t, expected.Id, id)
	})

	t.Run("wrong username", func(t *testing.T) {
		_, err := Verify("test_user2", "123456")
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	})

	t.Run("wrong password", func(t *testing.T) {
		_, err := Verify("test_user", "1234567")
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	})
}

func TestGet(t *testing.T) {
	database.Use(testutil.TestDB(t))
	expected := models.User{
		Id:       uuid.New().String(),
		Username: "test_user",
		Role:     "user",
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
