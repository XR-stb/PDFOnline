package pdf

import (
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	"backend/pkg/api/apiutil"
	"backend/pkg/database"
	"backend/pkg/database/models"
	"backend/pkg/static"
)

var (
	ErrForbidden = errors.New("forbidden")
)

func Create(author, host, title, description string, f *multipart.FileHeader) error {
	pdf := models.PDF{
		Id:          uuid.New().String(),
		Author:      author,
		Title:       title,
		Description: description,
	}

	err := static.UploadPdf(pdfFileName(pdf.Id), f)
	if err != nil {
		return err
	}

	pdf.Url = pdfUrl(host, pdf.Id)

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

type UpdateOption struct {
	Id          string
	UserId      string
	IsAdmin     bool
	Title       *string
	Description *string
}

func Update(opt UpdateOption) error {
	pdf := models.PDF{}
	db := database.Instance()
	if err := db.First(&pdf, "id = ?", opt.Id).Error; err != nil {
		return err
	}

	if !opt.IsAdmin && pdf.Author != opt.UserId {
		return ErrForbidden
	}

	if opt.Title != nil {
		pdf.Title = *opt.Title
	}

	if opt.Description != nil {
		pdf.Description = *opt.Description
	}

	if err := db.Save(&pdf).Error; err != nil {
		return err
	}

	return nil
}

type DeleteOption struct {
	Id      string
	UserId  string
	IsAdmin bool
}

func Delete(opt DeleteOption) error {
	pdf := models.PDF{}
	db := database.Instance()
	if err := db.First(&pdf, "id = ?", opt.Id).Error; err != nil {
		return err
	}

	if !opt.IsAdmin && pdf.Author != opt.UserId {
		return ErrForbidden
	}

	if err := db.Delete(&pdf).Error; err != nil {
		return err
	}

	// async delete pdf file
	go func() {
		if err := static.DeletePdf(pdfFileName(opt.Id)); err != nil {
			logrus.Error(err)
		}
	}()

	return nil
}

func pdfFileName(id string) string {
	return fmt.Sprintf("%s.pdf", id)
}

func pdfUrl(host, id string) string {
	return fmt.Sprintf("http://%s%s/%s", host, apiutil.StaticRootPdf, pdfFileName(id))
}
