package pdf

import (
	"gorm.io/gorm"
	"testing"

	"backend/pkg/database"
	"backend/pkg/database/models"
	"backend/test/testutil"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	t.Skip("TODO")
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

func TestGetById(t *testing.T) {
	testPdf := models.PDF{
		Id:     uuid.New().String(),
		Author: uuid.New().String(),
	}
	database.Use(testutil.TestDB(t))
	db := database.Instance()
	db.Create(&testPdf)

	t.Run("success", func(t *testing.T) {
		p, err := GetById(testPdf.Id)
		assert.NoError(t, err)
		assert.Equal(t, testPdf.Id, p.pdf.Id)
	})

	t.Run("not found", func(t *testing.T) {
		_, err := GetById(uuid.New().String())
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	})
}

func TestPDF_Update(t *testing.T) {
	testPdf := models.PDF{
		Id:     uuid.New().String(),
		Author: uuid.New().String(),
	}
	database.Use(testutil.TestDB(t))
	db := database.Instance()
	db.Create(&testPdf)

	p := &PDF{
		pdf: &testPdf,
		db:  db,
	}

	err := p.Update(&UpdateOption{
		Title:       testutil.StringPtr("title"),
		Description: testutil.StringPtr("description"),
	})
	assert.NoError(t, err)

	var pdf models.PDF
	err = db.First(&pdf, "id = ?", testPdf.Id).Error
	assert.NoError(t, err)
	assert.Equal(t, "title", pdf.Title)
	assert.Equal(t, "description", pdf.Description)
}

func TestPDF_Delete(t *testing.T) {
	testPdf := models.PDF{
		Id:     uuid.New().String(),
		Author: uuid.New().String(),
	}
	database.Use(testutil.TestDB(t))
	db := database.Instance()
	db.Create(&testPdf)

	p := &PDF{
		pdf: &testPdf,
		db:  db,
	}

	err := p.Delete()
	assert.NoError(t, err)
}

func Test_url(t *testing.T) {
	url := url("localhost:8080", "/static/pdf", "123.pdf")
	assert.Equal(t, "http://localhost:8080/static/pdf/123.pdf", url)
}
