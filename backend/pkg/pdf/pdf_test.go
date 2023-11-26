package pdf

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"backend/pkg/database"
	"backend/pkg/database/models"
	"backend/test/testutil"
)

func Test_url(t *testing.T) {
	url := pdfUrl("localhost:8080", "123")
	assert.Equal(t, "http://localhost:8080/static/pdf/123.pdf", url)
}

func TestList(t *testing.T) {
	pdfs := []models.PDF{
		{
			Id: uuid.New().String(),
		},
		{
			Id: uuid.New().String(),
		},
		{
			Id: uuid.New().String(),
		},
	}
	database.Use(testutil.TestDB(t))
	db := database.Instance()
	db.Create(&pdfs)

	list, err := List()
	assert.NoError(t, err)
	assert.Equal(t, len(pdfs), len(list))
}

func TestUpdate(t *testing.T) {
	testPdf := models.PDF{
		Id:     uuid.New().String(),
		Author: uuid.New().String(),
	}
	database.Use(testutil.TestDB(t))
	db := database.Instance()
	db.Create(&testPdf)

	t.Run("admin", func(t *testing.T) {
		opt := UpdateOption{
			Id:      testPdf.Id,
			IsAdmin: true,
			Title:   testutil.StringPtr("title"),
		}
		err := Update(opt)
		assert.NoError(t, err)

		var pdf models.PDF
		err = db.First(&pdf, "id = ?", testPdf.Id).Error
		assert.NoError(t, err)
		assert.Equal(t, *opt.Title, pdf.Title)
	})

	t.Run("author", func(t *testing.T) {
		opt := UpdateOption{
			Id:          testPdf.Id,
			UserId:      testPdf.Author,
			IsAdmin:     false,
			Description: testutil.StringPtr("description"),
		}
		err := Update(opt)
		assert.NoError(t, err)

		var pdf models.PDF
		err = db.First(&pdf, "id = ?", testPdf.Id).Error
		assert.NoError(t, err)
		assert.Equal(t, *opt.Description, pdf.Description)
	})

	t.Run("forbidden", func(t *testing.T) {
		opt := UpdateOption{
			Id:          testPdf.Id,
			UserId:      uuid.New().String(),
			IsAdmin:     false,
			Description: testutil.StringPtr("description"),
		}
		err := Update(opt)
		assert.ErrorIs(t, err, ErrForbidden)
	})

	t.Run("not found", func(t *testing.T) {
		opt := UpdateOption{
			Id:          uuid.New().String(),
			UserId:      testPdf.Author,
			IsAdmin:     false,
			Description: testutil.StringPtr("description"),
		}
		err := Update(opt)
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	})
}

func TestDelete(t *testing.T) {
	testPdf := models.PDF{
		Id:     uuid.New().String(),
		Author: uuid.New().String(),
	}
	database.Use(testutil.TestDB(t))
	db := database.Instance()
	db.Create(&testPdf)

	t.Run("forbidden", func(t *testing.T) {
		opt := DeleteOption{
			Id:      testPdf.Id,
			UserId:  uuid.New().String(),
			IsAdmin: false,
		}
		err := Delete(opt)
		assert.ErrorIs(t, err, ErrForbidden)
	})

	t.Run("admin", func(t *testing.T) {
		opt := DeleteOption{
			Id:      testPdf.Id,
			IsAdmin: true,
		}
		err := Delete(opt)
		assert.NoError(t, err)

		var pdf models.PDF
		err = db.First(&pdf, "id = ?", testPdf.Id).Error
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	})

	t.Run("not found", func(t *testing.T) {
		opt := DeleteOption{
			Id:      uuid.New().String(),
			IsAdmin: true,
		}
		err := Delete(opt)
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	})
}
