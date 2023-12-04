package v1

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"backend/pkg/api/apiutil"
	"backend/pkg/api/apiutil/jwt"
	"backend/pkg/api/hooks"
	"backend/pkg/captcha"
	"backend/pkg/provider"
	"backend/pkg/user"
	"backend/pkg/user/role"
)

type UserAPI struct{}

func (u UserAPI) Routes() []apiutil.Route {
	return []apiutil.Route{
		{
			Method:  http.MethodPost,
			Pattern: "/v1/users",
			Handler: u.Register,
		},
		{
			Method:  http.MethodPost,
			Pattern: "/v1/users/captcha",
			Handler: u.SendCaptcha,
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

type RegisterUserReq struct {
	Username string `json:"username" binding:"required,min=2,max=32" required:"username is required" min:"username is too short, min 2 chars" max:"username is too long, max 32 chars"`
	Password string `json:"password" binding:"required,min=6,max=32" required:"password is required" min:"password is too short, min 6 chars" max:"password is too long, max 32 chars"`
	Email    string `json:"email" binding:"required,email" required:"email is required" email:"email is invalid"`
	Captcha  string `json:"captcha" binding:"required,len=6" required:"captcha is required" len:"length of captcha should be 6"`
}

func (UserAPI) Register(c *gin.Context) {
	var req RegisterUserReq
	if err := apiutil.ShouldBindJSON(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := captcha.Verify(req.Email, req.Captcha)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	u, err := user.Create(req.Username, req.Password, req.Email, role.RoleUser)
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

type SendCaptchaReq struct {
	Email string `json:"email" binding:"required,email" required:"email is required" email:"email is invalid"`
}

var (
	bodyTemplate    = `Your captcha is: %s, please use it within 10 minutes.`
	subjectTemplate = `Verification Code`
)

func (UserAPI) SendCaptcha(c *gin.Context) {
	var req SendCaptchaReq
	if err := apiutil.ShouldBindJSON(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := user.GetByEmail(req.Email)
	switch {
	case err == nil:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email is already registered",
		})
		return
	case errors.Is(err, gorm.ErrRecordNotFound):
	default:
		logrus.Warn(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	err = provider.Send(req.Email, subjectTemplate, fmt.Sprintf(bodyTemplate, captcha.Generate(req.Email)))
	if err != nil {
		logrus.Warn(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

type LoginUserReq struct {
	Username string `json:"username" binding:"required,min=5,max=32" required:"username is required" min:"username is too short, min 5 chars" max:"username is too long, max 32 chars"`
	Password string `json:"password" binding:"required,min=6,max=32" required:"password is required" min:"password is too short, min 6 chars" max:"password is too long, max 32 chars"`
}

func (UserAPI) Login(c *gin.Context) {
	var req LoginUserReq
	if err := apiutil.ShouldBindJSON(c, &req); err != nil {
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
	if err := apiutil.ShouldBindJSON(c, &req); err != nil {
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
	if err := apiutil.ShouldBindJSON(c, &req); err != nil {
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
	Role *role.Role `json:"role" binding:"required,oneof=1 2" required:"role is required" oneof:"role is invalid"`
}

func (UserAPI) UpdateRole(c *gin.Context) {
	var req UpdateUserRoleReq
	if err := apiutil.ShouldBindJSON(c, &req); err != nil {
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
