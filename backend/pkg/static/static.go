package static

import (
	"fmt"
	"io"
	"mime/multipart"
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

func UploadPdf(filename string, file *multipart.FileHeader) error {
	path := filepath.Join(PdfDir, filename)

	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	targetFile, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file error: %v", err)
	}
	defer f.Close()

	_, err = io.Copy(targetFile, f)

	return err
}

func UploadCover(filename string, file []byte) (string, error) {
	path := filepath.Join(CoverDir, filename)
	f, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("create file error: %v", err)
	}
	defer f.Close()

	_, err = f.Write(file)

	return path, err
}

func DeletePdf(filename string) error {
	path := filepath.Join(PdfDir, filename)
	return os.Remove(path)
}

func DeleteCover(filename string) error {
	path := filepath.Join(CoverDir, filename)
	return os.Remove(path)
}
