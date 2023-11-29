package static

import (
	"backend/pkg/util"
	"fmt"
	"os"
	"path/filepath"
)

// SaveCoverFile saves the cover image of the pdf file.
// Notice: The filename is the name of the file without extension.
func SaveCoverFile(id string) (string, error) {
	f, err := os.Create(coverFilePath(id))
	if err != nil {
		return "", fmt.Errorf("create file error: %v", err)
	}
	defer f.Close()

	err = util.ExtractPageAsImage(pdfFilePath(id), f, 1)
	if err != nil {
		return "", fmt.Errorf("extract page as image error: %v", err)
	}

	return fmt.Sprintf("%s.jpg", id), nil
}

func RemoveCoverFile(id string) error {
	return os.Remove(coverFilePath(id))
}

func coverFilePath(id string) string {
	return filepath.Join(CoverDir, fmt.Sprintf("%s.jpg", id))
}
