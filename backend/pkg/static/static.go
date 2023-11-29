package static

import (
	"fmt"
	"os"
	"path/filepath"

	"backend/pkg/config"
)

var (
	PdfDir   string
	CoverDir string
)

func Init() error {
	staticDir := config.StaticDir()

	if staticDir == "" {
		return fmt.Errorf("static dir is empty")
	}

	PdfDir = filepath.Join(staticDir, "pdf")
	CoverDir = filepath.Join(staticDir, "cover")

	if err := createDirIfNotExist(PdfDir); err != nil {
		return err
	}

	if err := createDirIfNotExist(CoverDir); err != nil {
		return err
	}

	return nil
}

func createDirIfNotExist(dir string) error {
	_, err := os.Stat(dir)
	if err != nil {
		switch {
		case os.IsNotExist(err):
			err = os.MkdirAll(dir, os.ModePerm)
			if err != nil {
				return fmt.Errorf("create directory error: %v", err)
			}
		default:
			return fmt.Errorf("get directory error: %v", err)
		}
	}

	return nil
}
