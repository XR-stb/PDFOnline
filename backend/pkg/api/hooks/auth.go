package hooks

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/pkg/api/apiutil"
	"backend/pkg/api/apiutil/jwt"
	"backend/pkg/user/role"
)

func Auth(minRole role.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie(apiutil.CookieToken)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, err := jwt.VerifyToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		if claims.Role < minRole {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Set(apiutil.CtxUserId, claims.UserId)
		c.Set(apiutil.CtxRole, claims.Role)
	}
}
