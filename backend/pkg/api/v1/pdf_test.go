package v1

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"backend/pkg/database"
	"backend/pkg/database/models"
	"backend/pkg/static"
	"backend/test/testutil"
)

func TestPdfAPI_Upload(t *testing.T) {
	t.Skip("waiting for repair")

	tmpDir := t.TempDir()
	static.PdfDir = tmpDir

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	defer writer.Close()

	part, err := writer.CreateFormFile("file", "test.pdf")
	assert.NoError(t, err)
	_, err = io.Copy(part, bytes.NewReader([]byte("test")))
	assert.NoError(t, err)

	w, err := writer.CreateFormField("title")
	assert.NoError(t, err)
	w.Write([]byte("title"))

	w, err = writer.CreateFormField("description")
	assert.NoError(t, err)
	w.Write([]byte("description"))

	c.Request = httptest.NewRequest(http.MethodPost, "http://127.0.0.1:8080/pdfs", body)
	c.Request.Header.Set("Content-Type", writer.FormDataContentType())
	c.Request.Header.Add("S-COOKIE2", "a=2l=310260000000000&m=460&n=00")
	PdfAPI{}.Upload(c)

	assert.Equal(t, http.StatusCreated, rec.Code)
	t.Log(rec.Body.String())
}

func TestPdfAPI_List(t *testing.T) {
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

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = testutil.NewRequest(t, http.MethodGet, "/pdfs")
	PdfAPI{}.List(c)

	assert.Equal(t, 200, rec.Code)
	var payload struct {
		Pdfs []models.PDF `json:"pdfs"`
	}
	err := json.Unmarshal(rec.Body.Bytes(), &payload)
	assert.NoError(t, err)
	assert.Equal(t, len(pdfs), len(payload.Pdfs))
}

func TestPdfAPI_Update(t *testing.T) {
	pdf := models.PDF{
		Id: uuid.New().String(),
	}
	database.Use(testutil.TestDB(t))
	db := database.Instance()
	db.Create(&pdf)

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = testutil.NewRequest(t, http.MethodPut, "/pdfs", UpdatePdfReq{
		Title:       stringPtr("title"),
		Description: stringPtr("description"),
	})
	c.AddParam("id", pdf.Id)
	PdfAPI{}.Update(c)

	c.Writer.WriteHeaderNow()
	assert.Equal(t, rec.Code, http.StatusNoContent)
	var payload models.PDF
	db.Find(&payload, "id = ?", pdf.Id)
	assert.Equal(t, "title", payload.Title)
	assert.Equal(t, "description", payload.Description)
}

func TestPdfAPI_Delete(t *testing.T) {
	pdf := models.PDF{
		Id: uuid.New().String(),
	}
	database.Use(testutil.TestDB(t))
	db := database.Instance()
	db.Create(&pdf)

	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = testutil.NewRequest(t, http.MethodDelete, "/pdfs")
	c.AddParam("id", pdf.Id)
	PdfAPI{}.Delete(c)

	c.Writer.WriteHeaderNow()
	assert.Equal(t, http.StatusNoContent, rec.Code)
	var payload models.PDF
	err := db.First(&payload, "id = ?", pdf.Id).Error
	assert.Equal(t, err, gorm.ErrRecordNotFound)
}

func stringPtr(s string) *string {
	return &s
}
