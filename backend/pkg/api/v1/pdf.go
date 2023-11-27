package v1

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"backend/pkg/api/apiutil"
	"backend/pkg/api/hooks"
	"backend/pkg/pdf"
	"backend/pkg/user/role"
)

type PdfAPI struct{}

func (p PdfAPI) Routes() []apiutil.Route {
	return []apiutil.Route{
		{
			Method:  http.MethodPost,
			Pattern: "/v1/pdfs",
			Hooks:   gin.HandlersChain{hooks.Auth(role.RoleGuest)},
			Handler: p.Upload,
		},
		{
			Method:  http.MethodGet,
			Pattern: "/v1/pdfs",
			Handler: p.List,
		},
		{
			Method:  http.MethodPatch,
			Pattern: "/v1/pdfs/:id",
			Hooks:   gin.HandlersChain{hooks.Auth(role.RoleUser)},
			Handler: p.Update,
		},
		{
			Method:  http.MethodDelete,
			Pattern: "/v1/pdfs/:id",
			Hooks:   gin.HandlersChain{hooks.Auth(role.RoleUser)},
			Handler: p.Delete,
		},
	}
}

type UploadPdfReq struct {
	Title       string `form:"title" binding:"required" required:"title is required"`
	Description string `form:"description" binding:"required" required:"description is required"`
}

func (p PdfAPI) Upload(c *gin.Context) {
	var params UploadPdfReq

	if err := apiutil.ShouldBind(c, &params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	f, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "pdf file is required",
		})
		return
	}

	if !strings.HasSuffix(f.Filename, ".pdf") {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "file must be pdf",
		})
		return
	}

	if err := pdf.Create(c.GetString(apiutil.CtxUserId), c.Request.Host, params.Title, params.Description, f); err != nil {
		logrus.Warn(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func (p PdfAPI) List(c *gin.Context) {
	pdfs, err := pdf.List()
	if err != nil {
		logrus.Warn(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"pdfs": pdfs,
	})
}

type UpdatePdfReq struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

func (p PdfAPI) Update(c *gin.Context) {
	var params UpdatePdfReq

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userRole, _ := c.Get(apiutil.CtxRole)
	opt := pdf.UpdateOption{
		Id:          c.Param("id"),
		UserId:      c.GetString(apiutil.CtxUserId),
		IsAdmin:     userRole == role.RoleAdmin,
		Title:       params.Title,
		Description: params.Description,
	}

	if err := pdf.Update(opt); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
		} else if errors.Is(err, pdf.ErrForbidden) {
			c.Status(http.StatusForbidden)
		} else {
			logrus.Warn(err)
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.Status(http.StatusNoContent)
}

func (p PdfAPI) Delete(c *gin.Context) {
	userRole, _ := c.Get(apiutil.CtxRole)
	opt := pdf.DeleteOption{
		Id:      c.Param("id"),
		UserId:  c.GetString(apiutil.CtxUserId),
		IsAdmin: userRole == role.RoleAdmin,
	}
	if err := pdf.Delete(opt); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.Status(http.StatusNotFound)
		} else if errors.Is(err, pdf.ErrForbidden) {
			c.Status(http.StatusForbidden)
		} else {
			logrus.Warn(err)
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.Status(http.StatusNoContent)
}
