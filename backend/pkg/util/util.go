package util

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

// ExtractPageAsImage get a page of the pdf and save it as a picture
func ExtractPageAsImage(inputPath string, w io.Writer, pageNum int) error {
	// Create a context.
	ctx, err := api.ReadContextFile(inputPath)
	if err != nil {
		return err
	}

	// Optimize resource usage of this context.
	if err := api.OptimizeContext(ctx); err != nil {
		return err
	}

	// Extract images for page 1.
	images, err := pdfcpu.ExtractPageImages(ctx, pageNum, false)
	if err != nil {
		return err
	}

	if len(images) != 1 {
		return errors.New("unexpected number of images")
	}

	// Write images to disk.
	for _, img := range images {
		if _, err = io.Copy(w, img); err != nil {
			return err
		}
	}

	return nil
}
