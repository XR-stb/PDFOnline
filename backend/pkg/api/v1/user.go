package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"backend/pkg/api/apiutil"
	"backend/pkg/api/apiutil/jwt"
	"backend/pkg/api/hooks"
	"backend/pkg/user"
	"backend/pkg/user/role"
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
			Method:  http.MethodPost,
			Pattern: "/v1/users/logout",
			Handler: u.Logout,
		},
		{
			Method:  http.MethodGet,
			Pattern: "/v1/users/:user_id",
			Handler: u.Show,
		},
		{
			Method:  http.MethodPatch,
			Pattern: "/v1/users/:user_id",
			Hooks:   gin.HandlersChain{hooks.Auth(role.RoleUser), hooks.UserAuth()},
			Handler: u.Update,
		},
		{
			Method:  http.MethodPut,
			Pattern: "/v1/users/:user_id/password",
			Hooks:   gin.HandlersChain{hooks.Auth(role.RoleUser), hooks.UserAuth()},
			Handler: u.UpdatePassword,
		},
		{
			Method:  http.MethodPut,
			Pattern: "/v1/users/:user_id/role",
			Hooks:   gin.HandlersChain{hooks.Auth(role.RoleAdmin)},
			Handler: u.UpdateRole,
		},
	}
}

type UserRegisterOrLoginReq struct {
	Username string `json:"username" binding:"required,min=5,max=32" required:"username is required" min:"username is too short, min 5 chars" max:"username is too long, max 32 chars"`
	Password string `json:"password" binding:"required,min=6,max=32" required:"password is required" min:"password is too short, min 6 chars" max:"password is too long, max 32 chars"`
}

func (UserAPI) Register(c *gin.Context) {
	var req UserRegisterOrLoginReq
	if err := apiutil.ShouldBind(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := user.Create(req.Username, req.Password, "", role.RoleUser)
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

	token := jwt.GenerateToken(&jwt.Claims{
		UserId:   u.Id(),
		Username: u.Username(),
		Role:     u.Role(),
	})
	c.SetCookie("token", token, 86400, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"user_id": u.Id(),
	})
}

func (UserAPI) Login(c *gin.Context) {
	var req UserRegisterOrLoginReq
	if err := apiutil.ShouldBind(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := user.GetByUsername(req.Username)
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

	if !u.VerifyPassword(req.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "username or password is incorrect",
		})
		return
	}

	token := jwt.GenerateToken(&jwt.Claims{
		UserId:   u.Id(),
		Username: u.Username(),
		Role:     u.Role(),
	})
	c.SetCookie("token", token, 86400, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"user_id": u.Id(),
	})
}

func (UserAPI) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.Status(http.StatusOK)
}

func (UserAPI) Show(c *gin.Context) {
	u, err := user.GetById(c.Param("user_id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
		} else {
			logrus.Warn(err)
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u.Show(),
	})
}

type UpdateUserReq struct {
	Username *string `json:"username" binding:"omitempty,min=5,max=32" min:"username is too short, min 5 chars" max:"username is too long, max 32 chars"`
}

func (UserAPI) Update(c *gin.Context) {
	var req UpdateUserReq
	if err := apiutil.ShouldBind(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := user.GetById(c.Param("user_id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
		} else {
			logrus.Warn(err)
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	err = u.Update(&user.UpdateOption{Username: req.Username})
	if err != nil {
		logrus.Warn(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

type UpdateUserPasswordReq struct {
	Password    string `json:"password"`
	NewPassword string `json:"new_password" binding:"required,min=6,max=32" required:"new_password is required" min:"new_password is too short, min 6 chars" max:"new_password is too long, max 32 chars"`
}

func (UserAPI) UpdatePassword(c *gin.Context) {
	var req UpdateUserPasswordReq
	if err := apiutil.ShouldBind(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := user.GetById(c.Param("user_id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
		} else {
			logrus.Warn(err)
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	if !u.VerifyPassword(req.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "password is incorrect",
		})
		return
	}

	err = u.Update(&user.UpdateOption{Password: &req.NewPassword})
	if err != nil {
		logrus.Warn(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}

type UpdateUserRoleReq struct {
	Role *role.Role `json:"role" binding:"required,oneof=0 1 2" required:"role is required" oneof:"role is invalid"`
}

func (UserAPI) UpdateRole(c *gin.Context) {
	var req UpdateUserRoleReq
	if err := apiutil.ShouldBind(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := user.GetById(c.Param("user_id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
		} else {
			logrus.Warn(err)
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	err = u.Update(&user.UpdateOption{Role: req.Role})
	if err != nil {
		logrus.Warn(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusNoContent)
}
