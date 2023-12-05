package apiutil

import (
	"backend/pkg/api/apiutil/jwt"
	"backend/pkg/user/role"
	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	StaticRootPdf   = "/static/pdf"
	StaticRootCover = "/static/cover"
)

// ShouldBindJSON is a wrapper of gin.Context.ShouldBindJSON
func ShouldBindJSON(c *gin.Context, obj interface{}) error {
	err := c.ShouldBindJSON(obj)
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

func ShouldBindForm(c *gin.Context, obj interface{}) error {
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

func SetTokenCookie(c *gin.Context, userId, username string, role role.Role, maxAge int) {
	token := jwt.GenerateToken(&jwt.Claims{
		UserId:   userId,
		Username: username,
		Role:     role,
	})
	c.SetCookie(CookieToken, token, maxAge, "/", "", false, false)
}
