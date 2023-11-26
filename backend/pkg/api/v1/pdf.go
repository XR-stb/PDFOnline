package v1

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"

	"backend/pkg/api/apiutil"
	"backend/pkg/pdf"
)

type PdfAPI struct{}

func (p PdfAPI) Routes() []apiutil.Route {
	return []apiutil.Route{
		{
			Method:  http.MethodPost,
			Pattern: "/v1/pdfs",
			Handler: p.Upload,
		},
		{
			Method:  http.MethodGet,
			Pattern: "/v1/pdfs",
			Handler: p.List,
		},
		{
			Method:  http.MethodPut,
			Pattern: "/v1/pdfs/:id",
			Handler: p.Update,
		},
		{
			Method:  http.MethodDelete,
			Pattern: "/v1/pdfs/:id",
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

	if err := pdf.Create(c, params.Title, params.Description, f); err != nil {
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

	if err := pdf.Update(c.Param("id"), params.Title, params.Description); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "pdf not found",
			})
		} else {
			logrus.Warn(err)
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.Status(http.StatusNoContent)
}

func (p PdfAPI) Delete(c *gin.Context) {
	if err := pdf.Delete(c.Param("id")); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "pdf not found",
			})
		} else {
			logrus.Warn(err)
			c.Status(http.StatusInternalServerError)
		}
		return
	}

	c.Status(http.StatusNoContent)
}
