package pdf

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
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

type PDF struct {
	pdf *models.PDF
	db  *gorm.DB
}

func Create(host, uploader, author, title, description string, f *multipart.FileHeader) (*PDF, error) {
	pdf := models.PDF{
		Id:          uuid.New().String(),
		Uploader:    uploader,
		Author:      author,
		Title:       title,
		Description: description,
	}

	pdfFilename, err := static.SavePdfFile(pdf.Id, f)
	if err != nil {
		return nil, err
	}

	pdf.Url = url(host, apiutil.StaticRootPdf, pdfFilename)

	coverFilename, err := static.SaveCoverFile(pdf.Id)
	if err != nil {
		_ = static.RemovePdfFile(pdf.Id)
		return nil, err
	}

	pdf.CoverUrl = url(host, apiutil.StaticRootCover, coverFilename)

	db := database.Instance()

	if err := db.Create(&pdf).Error; err != nil {
		return nil, err
	}

	return &PDF{
		pdf: &pdf,
		db:  db,
	}, nil
}

func List() ([]*models.PDF, error) {
	var pdfs []*models.PDF

	if err := database.Instance().Find(&pdfs).Error; err != nil {
		return nil, err
	}

	return pdfs, nil
}

func GetById(id string) (*PDF, error) {
	db := database.Instance()
	pdf := models.PDF{}

	return &PDF{
		pdf: &pdf,
		db:  db,
	}, db.First(&pdf, "id = ?", id).Error
}

func (p *PDF) Id() string {
	return p.pdf.Id
}

func (p *PDF) Author() string {
	return p.pdf.Uploader
}

func (p *PDF) Save() error {
	return p.db.Save(p.pdf).Error
}

type UpdateOption struct {
	Title       *string
	Description *string
}

func (p *PDF) Update(opt *UpdateOption) error {
	if opt.Title != nil {
		p.pdf.Title = *opt.Title
	}

	if opt.Description != nil {
		p.pdf.Description = *opt.Description
	}

	return p.Save()
}

func (p *PDF) Delete() error {
	if err := p.db.Delete(&p.pdf).Error; err != nil {
		return err
	}

	// async delete pdf file
	go func() {
		if err := static.RemovePdfFile(p.Id()); err != nil {
			logrus.Error(err)
		}
	}()

	// async delete cover file
	go func() {
		if err := static.RemoveCoverFile(p.Id()); err != nil {
			logrus.Error(err)
		}
	}()

	return nil
}

func url(host, root, filename string) string {
	return fmt.Sprintf("http://%s%s/%s", host, root, filename)
}
