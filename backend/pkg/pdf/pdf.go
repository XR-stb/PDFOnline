package pdf

import (
	"fmt"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"backend/pkg/api/apiutil"
	"backend/pkg/database"
	"backend/pkg/database/models"
	"backend/pkg/static"
)

func Create(c *gin.Context, title, description string, f *multipart.FileHeader) error {
	pdf := models.PDF{
		Id:          uuid.New().String(),
		Title:       title,
		Description: description,
	}

	err := static.UploadPdf(fmt.Sprintf("%s.pdf", pdf.Id), f)
	if err != nil {
		return err
	}

	pdf.Url = fmt.Sprintf("http://%s:%s%s/%s.pdf", c.Request.Host, c.Request.URL.Port(), apiutil.StaticRootPdf, pdf.Id)

	// TODO: generate cover and save

	if err := database.Instance().Create(&pdf).Error; err != nil {
		return err
	}

	return nil
}

func List() ([]*models.PDF, error) {
	var pdfs []*models.PDF

	if err := database.Instance().Find(&pdfs).Error; err != nil {
		return nil, err
	}

	return pdfs, nil
}

func Update(id string, title, description *string) error {
	pdf := models.PDF{}
	db := database.Instance()
	if err := db.First(&pdf, "id = ?", id).Error; err != nil {
		return err
	}

	if title != nil {
		pdf.Title = *title
	}

	if description != nil {
		pdf.Description = *description
	}

	if err := db.Save(&pdf).Error; err != nil {
		return err
	}

	return nil
}

func Delete(id string) error {
	pdf := models.PDF{}
	db := database.Instance()
	if err := db.First(&pdf, "id = ?", id).Error; err != nil {
		return err
	}

	if err := db.Delete(&pdf).Error; err != nil {
		return err
	}

	// async delete pdf file
	go func() {
		if err := static.DeletePdf(fmt.Sprintf("%s.pdf", pdf.Id)); err != nil {
			logrus.Error(err)
		}
	}()

	return nil
}
