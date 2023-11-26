package jwt

import (
	"backend/pkg/user/role"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestGenToken(t *testing.T) {
	claims := Claims{
		UserId:   uuid.New().String(),
		Username: "testUser",
		Role:     role.RoleAdmin,
	}

	token := GenerateToken(&claims)
	assert.NotNil(t, token)
}

func TestVerifyToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		claims := Claims{
			UserId:   uuid.New().String(),
			Username: "testUser",
			Role:     role.RoleAdmin,
		}

		token := GenerateToken(&claims)
		assert.NotNil(t, token)

		claims2, err := VerifyToken(token)
		assert.NoError(t, err)
		assert.Equal(t, "testUser", claims2.Username)
		assert.Equal(t, role.RoleAdmin, claims2.Role)
	})

	t.Run("empty token", func(t *testing.T) {
		_, err := VerifyToken("")
		assert.Equal(t, "empty token", err.Error())
	})

	t.Run("invalid token", func(t *testing.T) {
		_, err := VerifyToken("INVALID_TOKEN")
		assert.Equal(t, "token is invalid", err.Error())
	})

	t.Run("expired token", func(t *testing.T) {
		tmp := expireTime
		expireTime = -time.Minute
		t.Cleanup(func() {
			expireTime = tmp
		})

		claims := Claims{
			UserId: uuid.New().String(),
			Role:   role.RoleAdmin,
		}
		token := GenerateToken(&claims)

		_, err := VerifyToken(token)
		assert.Equal(t, "token is expired", err.Error())
	})
}
