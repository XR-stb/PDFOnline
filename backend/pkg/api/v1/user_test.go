package v1

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"backend/pkg/api/apiutil"
	"backend/pkg/database"
	"backend/pkg/database/models"
	"backend/pkg/user/role"
	"backend/pkg/util"
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
		cookieHeader := rec.Header().Get("Set-Cookie")
		assert.NotNil(t, cookieHeader)
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
		assert.Equal(t, "username is too short, min 5 chars\npassword is required", payload["error"])
	})

	t.Run("duplicate username", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = testutil.NewRequest(t, http.MethodPost, "/users/register", UserRegisterOrLoginReq{
			Username: "testUser",
			Password: "123456",
		})
		UserAPI{}.Register(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		payload := map[string]any{}
		err := json.NewDecoder(rec.Body).Decode(&payload)
		assert.NoError(t, err)
		assert.Equal(t, "username already exist", payload["error"])
	})
}

func TestUserAPI_Login(t *testing.T) {
	database.Use(testutil.TestDB(t))
	database.Instance().Create(&models.User{
		Id:       uuid.New().String(),
		Username: "testUser",
		Password: util.MD5("123456"),
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
		cookieHeader := rec.Header().Get("Set-Cookie")
		assert.NotNil(t, cookieHeader)
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

	t.Run("username too short and password is nil", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = testutil.NewRequest(t, http.MethodPost, "/users/login", UserRegisterOrLoginReq{
			Username: "test",
		})
		UserAPI{}.Login(c)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		payload := map[string]any{}
		err := json.NewDecoder(rec.Body).Decode(&payload)
		assert.NoError(t, err)
		assert.Equal(t, "username is too short, min 5 chars\npassword is required", payload["error"])
	})
}

func TestUserAPI_Show(t *testing.T) {
	database.Use(testutil.TestDB(t))
	user := models.User{
		Id:       uuid.New().String(),
		Username: "testUser",
		Email:    "email@example.com",
		Password: "123456",
		Role:     role.RoleUser,
	}
	database.Instance().Create(&user)

	t.Run("success", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = testutil.NewRequest(t, http.MethodGet, "/users", nil)
		c.AddParam("user_id", user.Id)
		UserAPI{}.Show(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		payload := map[string]any{}
		err := json.NewDecoder(rec.Body).Decode(&payload)
		assert.NoError(t, err)
		expected := map[string]any{
			"user": map[string]any{
				"id":       user.Id,
				"username": user.Username,
				"email":    user.Email,
			},
		}
		assert.EqualValues(t, expected, payload)
	})

	t.Run("user not exist", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = testutil.NewRequest(t, http.MethodGet, "/users", nil)
		c.AddParam("user_id", uuid.New().String())
		UserAPI{}.Show(c)

		c.Writer.WriteHeaderNow()
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestUserAPI_Logout(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		UserAPI{}.Logout(c)

		assert.Equal(t, http.StatusOK, rec.Code)
		cookieHeader := rec.Header().Get("Set-Cookie")
		assert.Equal(t, `token=; Path=/; Max-Age=0; HttpOnly`, cookieHeader)
	})
}

func TestUserAPI_Update(t *testing.T) {
	user := models.User{
		Id:   uuid.New().String(),
		Role: role.RoleUser,
	}
	database.Use(testutil.TestDB(t))
	db := database.Instance()
	db.Create(&user)

	t.Run("success", func(t *testing.T) {
		t.Run("oneself", func(t *testing.T) {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			c.Request = testutil.NewRequest(t, http.MethodPatch, "/users", UpdateUserReq{
				Username: testutil.StringPtr("testUser"),
			})
			c.AddParam("user_id", user.Id)
			c.Set(apiutil.CtxUserId, user.Id)
			c.Set(apiutil.CtxRole, role.RoleUser)
			UserAPI{}.Update(c)

			c.Writer.WriteHeaderNow()
			assert.Equal(t, http.StatusNoContent, rec.Code)
			actual := models.User{}
			err := db.First(&actual, "id = ?", user.Id).Error
			assert.NoError(t, err)
			assert.Equal(t, "testUser", actual.Username)
		})

		t.Run("admin", func(t *testing.T) {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			c.Request = testutil.NewRequest(t, http.MethodPatch, "/users", UpdateUserReq{
				Username: testutil.StringPtr("testUser2"),
			})
			c.AddParam("user_id", user.Id)
			c.Set(apiutil.CtxRole, role.RoleAdmin)
			UserAPI{}.Update(c)

			c.Writer.WriteHeaderNow()
			assert.Equal(t, http.StatusNoContent, rec.Code)
			actual := models.User{}
			err := db.First(&actual, "id = ?", user.Id).Error
			assert.NoError(t, err)
			assert.Equal(t, "testUser2", actual.Username)
		})
	})

	t.Run("user not exist", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = testutil.NewRequest(t, http.MethodPatch, "/users", UpdateUserReq{
			Username: testutil.StringPtr("testUser"),
		})
		c.AddParam("user_id", uuid.New().String())
		c.Set(apiutil.CtxUserId, user.Id)
		c.Set(apiutil.CtxRole, role.RoleUser)
		UserAPI{}.Update(c)

		c.Writer.WriteHeaderNow()
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})
}

func TestUserAPI_UpdateRole(t *testing.T) {
	user := models.User{
		Id:   uuid.New().String(),
		Role: role.RoleUser,
	}
	database.Use(testutil.TestDB(t))
	db := database.Instance()
	db.Create(&user)

	t.Run("success", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = testutil.NewRequest(t, http.MethodPatch, "/users/role", UpdateUserRoleReq{
			Role: testutil.RolePtr(role.RoleAdmin),
		})
		c.AddParam("user_id", user.Id)
		UserAPI{}.UpdateRole(c)

		c.Writer.WriteHeaderNow()
		assert.Equal(t, http.StatusNoContent, rec.Code)
	})

	t.Run("user not exist", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = testutil.NewRequest(t, http.MethodPatch, "/users/role", UpdateUserRoleReq{
			Role: testutil.RolePtr(role.RoleAdmin),
		})
		c.AddParam("user_id", uuid.New().String())
		UserAPI{}.UpdateRole(c)

		c.Writer.WriteHeaderNow()
		assert.Equal(t, http.StatusNotFound, rec.Code)
	})

	t.Run("invalid role", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = testutil.NewRequest(t, http.MethodPatch, "/users/role", UpdateUserRoleReq{
			Role: testutil.RolePtr(role.Role(3)),
		})
		c.AddParam("user_id", user.Id)
		UserAPI{}.UpdateRole(c)

		c.Writer.WriteHeaderNow()
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		payload := map[string]any{}
		err := json.NewDecoder(rec.Body).Decode(&payload)
		assert.NoError(t, err)
		assert.Equal(t, "role is invalid", payload["error"])
	})
}
