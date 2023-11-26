package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"backend/pkg/api/apiutil"
	"backend/pkg/user"
)

type UserAPI struct{}

func (u UserAPI) Routes() []apiutil.Route {
	return []apiutil.Route{
		{
			Method:  http.MethodPost,
			Pattern: "/v1/users/register",
			Handler: u.Register,
		},
		{
			Method:  http.MethodPost,
			Pattern: "/v1/users/login",
			Handler: u.Login,
		},
		{
			Method:  http.MethodGet,
			Pattern: "/v1/users/:user_id",
			Handler: u.Get,
		},
	}
}

type UserRegisterOrLoginReq struct {
	Username string `json:"username" binding:"required,min=6,max=36" required:"username is required" min:"username is too short, min 6 chars" max:"username is too long, max 36 chars"`
	Password string `json:"password" binding:"required,min=6,max=36" required:"password is required" min:"password is too short, min 6 chars" max:"password is too long, max 36 chars"`
}

func (u UserAPI) Register(c *gin.Context) {
	var req UserRegisterOrLoginReq
	if err := apiutil.ShouldBind(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId, err := user.Create(req.Username, req.Password, user.RoleUser)
	if err != nil {
		if errors.Is(err, user.ErrUserAlreadyExist) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			logrus.Warn(err)
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	// TODO: generate token
	token := ""

	c.JSON(http.StatusOK, gin.H{
		"user_id": userId,
		"token":   token,
	})
}

func (u UserAPI) Login(c *gin.Context) {
	var req UserRegisterOrLoginReq
	if err := apiutil.ShouldBind(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userId, err := user.Verify(req.Username, req.Password)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "username or password is incorrect",
			})
		} else {
			logrus.Warn(err)
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	// TODO: generate token
	token := ""

	c.JSON(http.StatusOK, gin.H{
		"user_id": userId,
		"token":   token,
	})
}

func (u UserAPI) Get(c *gin.Context) {
	user, err := user.Get(c.Param("user_id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
		} else {
			logrus.Warn(err)
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
