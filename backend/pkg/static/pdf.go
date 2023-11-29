package static

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// SavePdfFile saves the uploaded pdf file to the pdf directory.
// Notice: The filename is the name of the file without extension.
func SavePdfFile(id string, file *multipart.FileHeader) (string, error) {
	f, err := file.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	targetFile, err := os.Create(pdfFilePath(id))
	if err != nil {
		return "", fmt.Errorf("create file error: %v", err)
	}
	defer f.Close()

	_, err = io.Copy(targetFile, f)

	return fmt.Sprintf("%s.pdf", id), err
}

func RemovePdfFile(id string) error {
	return os.Remove(pdfFilePath(id))
}

func pdfFilePath(id string) string {
	return filepath.Join(PdfDir, fmt.Sprintf("%s.pdf", id))
}
