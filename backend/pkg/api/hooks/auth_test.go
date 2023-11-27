package hooks

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"backend/pkg/api/apiutil"
	"backend/pkg/api/apiutil/jwt"
	"backend/pkg/user/role"
	"backend/test/testutil"
)

func TestAuth(t *testing.T) {
	t.Run("admin role", func(t *testing.T) {
		auth := Auth(role.RoleAdmin)

		t.Run("admin", func(t *testing.T) {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			c.Request = testutil.NewRequest(t, http.MethodPost, "/auth")
			c.Request.Header.Set("Cookie", apiutil.CookieToken+"="+jwt.GenerateToken(&jwt.Claims{
				UserId:   "admin",
				Username: "admin",
				Role:     role.RoleAdmin,
			}))
			auth(c)
			assert.Equal(t, http.StatusOK, rec.Code)
		})

		t.Run("user", func(t *testing.T) {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			c.Request = testutil.NewRequest(t, http.MethodPost, "/auth")
			c.Request.Header.Set("Cookie", apiutil.CookieToken+"="+jwt.GenerateToken(&jwt.Claims{
				Role: role.RoleUser,
			}))
			auth(c)
			assert.Equal(t, http.StatusForbidden, rec.Code)
		})

		t.Run("guest", func(t *testing.T) {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			c.Request = testutil.NewRequest(t, http.MethodPost, "/auth")
			auth(c)
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		})

		t.Run("invalid token", func(t *testing.T) {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			c.Request = testutil.NewRequest(t, http.MethodPost, "/auth")
			c.Request.Header.Set("Cookie", apiutil.CookieToken+"=invalid")
			auth(c)
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		})
	})

	t.Run("user role", func(t *testing.T) {
		auth := Auth(role.RoleUser)

		t.Run("admin", func(t *testing.T) {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			c.Request = testutil.NewRequest(t, http.MethodPost, "/auth")
			c.Request.Header.Set("Cookie", apiutil.CookieToken+"="+jwt.GenerateToken(&jwt.Claims{
				UserId: "admin",
				Role:   role.RoleAdmin,
			}))
			auth(c)
			assert.Equal(t, "admin", c.GetString(apiutil.CtxUserId))
			userRole, _ := c.Get(apiutil.CtxRole)
			assert.Equal(t, role.RoleAdmin, userRole)
			assert.Equal(t, http.StatusOK, rec.Code)
		})

		t.Run("user", func(t *testing.T) {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			c.Request = testutil.NewRequest(t, http.MethodPost, "/auth")
			c.Request.Header.Set("Cookie", apiutil.CookieToken+"="+jwt.GenerateToken(&jwt.Claims{
				Role: role.RoleUser,
			}))
			auth(c)
			assert.Equal(t, http.StatusOK, rec.Code)
		})

		t.Run("guest", func(t *testing.T) {
			rec := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(rec)
			c.Request = testutil.NewRequest(t, http.MethodPost, "/auth")
			auth(c)
			assert.Equal(t, http.StatusUnauthorized, rec.Code)
		})
	})
}

func TestUserAuth(t *testing.T) {
	auth := UserAuth()

	t.Run("admin", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = testutil.NewRequest(t, http.MethodPost, "/users")
		c.AddParam("user_id", "testUser")
		c.Set(apiutil.CtxRole, role.RoleAdmin)
		auth(c)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("oneself", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = testutil.NewRequest(t, http.MethodPost, "/users")
		c.AddParam("user_id", "testUser")
		c.Set(apiutil.CtxUserId, "testUser")
		auth(c)
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("other user", func(t *testing.T) {
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = testutil.NewRequest(t, http.MethodPost, "/users")
		c.AddParam("user_id", "testUser")
		c.Set(apiutil.CtxUserId, "otherUser")
		auth(c)
		assert.Equal(t, http.StatusForbidden, rec.Code)
	})
}
