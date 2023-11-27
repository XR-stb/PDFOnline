package apiutil

import (
	"errors"
	"reflect"

	"backend/pkg/api/apiutil/jwt"
	"backend/pkg/user/role"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	StaticRootPdf   = "/static/pdf"
	StaticRootCover = "/static/cover"
)

// ShouldBind is a wrapper of gin.Context.ShouldBind
func ShouldBind(c *gin.Context, obj interface{}) error {
	err := c.ShouldBind(obj)
	if err == nil {
		return nil
	}

	var errs validator.ValidationErrors
	if !errors.As(err, &errs) {
		return err
	}

	var ret error
	for _, err := range errs {
		field, ok := reflect.TypeOf(obj).Elem().FieldByName(err.Field())
		if !ok {
			panic("field not found")
		}

		errMsg := field.Tag.Get(err.Tag())
		if errMsg == "" {
			errMsg = err.Error()
		}

		if ret == nil {
			ret = errors.New(errMsg)
		} else {
			ret = errors.Join(ret, errors.New(errMsg))
		}
	}

	return ret
}

func SetTokenCookie(c *gin.Context, UserId, Username string, Role role.Role) {
	token := jwt.GenerateToken(&jwt.Claims{
		UserId:   UserId,
		Username: Username,
		Role:     Role,
	})
	c.SetCookie("token", token, 0, "/", "", false, true)
}
