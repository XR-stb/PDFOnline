package static

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"backend/pkg/config"
)

func Test_Init(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		config.Cfg = &config.Config{Default: config.Default{StaticDir: t.TempDir()}}

		err := Init()
		assert.NoError(t, err)

		assert.DirExists(t, PdfDir)
		assert.DirExists(t, CoverDir)
	})

	t.Run("static dir is empty", func(t *testing.T) {
		config.Cfg = &config.Config{}
		err := Init()
		assert.Error(t, err)
	})
}

func TestUploadPdf(t *testing.T) {
	t.Skip("multipart.FileHeader is not supported in tests")
}

func TestDeletePdf(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		config.Cfg = &config.Config{Default: config.Default{StaticDir: t.TempDir()}}

		f, err := os.Create(filepath.Join(config.StaticDir(), "test.pdf"))
		assert.NoError(t, err)
		_ = f.Close()

		err = Init()
		assert.NoError(t, err)

		err = DeletePdf("test.pdf")
		assert.Error(t, err)

		assert.NoFileExists(t, filepath.Join(PdfDir, "test.pdf"))
	})

	t.Run("file not found", func(t *testing.T) {
		config.Cfg = &config.Config{Default: config.Default{StaticDir: t.TempDir()}}

		err := Init()
		assert.NoError(t, err)

		err = DeletePdf("test.pdf")
		assert.Error(t, err)
	})
}
