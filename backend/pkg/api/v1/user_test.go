package v1

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"backend/pkg/database"
	"backend/pkg/database/models"
	"backend/test/testutil"
)

func TestUserAPI_Register(t *testing.T) {
	database.Use(testutil.TestDB(t))

	t.Run("success", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = testutil.NewRequest(t, http.MethodPost, "/users/register", UserRegisterOrLoginReq{
			Username: "testUser",
			Password: "123456",
		})
		UserAPI{}.Register(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		payload := map[string]any{}
		err := json.NewDecoder(rec.Body).Decode(&payload)
		assert.NoError(t, err)
		_, err = uuid.Parse(payload["user_id"].(string))
		assert.NoError(t, err)
	})

	t.Run("username too short and password is nil", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = testutil.NewRequest(t, http.MethodPost, "/users/register", UserRegisterOrLoginReq{
			Username: "test",
		})
		UserAPI{}.Register(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		payload := map[string]any{}
		err := json.NewDecoder(rec.Body).Decode(&payload)
		assert.NoError(t, err)
		assert.Equal(t, "username is too short, min 6 chars\npassword is required", payload["error"])
	})
}

func TestUserAPI_Login(t *testing.T) {
	database.Use(testutil.TestDB(t))
	database.Instance().Create(&models.User{
		Id:       uuid.New().String(),
		Username: "testUser",
		Password: "123456",
	})

	t.Run("success", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = testutil.NewRequest(t, http.MethodPost, "/users/login", UserRegisterOrLoginReq{
			Username: "testUser",
			Password: "123456",
		})
		UserAPI{}.Login(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		payload := map[string]any{}
		err := json.NewDecoder(rec.Body).Decode(&payload)
		assert.NoError(t, err)
		_, err = uuid.Parse(payload["user_id"].(string))
		assert.NoError(t, err)
	})

	t.Run("username not exist", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = testutil.NewRequest(t, http.MethodPost, "/users/login", UserRegisterOrLoginReq{
			Username: "testUser2",
			Password: "123456",
		})
		UserAPI{}.Login(c)

		assert.Equal(t, http.StatusUnauthorized, rec.Code)
		payload := map[string]any{}
		err := json.NewDecoder(rec.Body).Decode(&payload)
		assert.NoError(t, err)
		assert.Equal(t, "username or password is incorrect", payload["error"])
	})
}
