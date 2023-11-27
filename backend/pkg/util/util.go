package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

// 获取pdf的某一页并保存为图片
// input_22_Image0.jpg
// 其中input为pdf的文件名，22为页码，后面固定
func extractPageAsImage(inputPath, outputDir string, pageNum int) error {
	// Create a context.
	ctx, err := api.ReadContextFile(inputPath)
	if err != nil {
		fmt.Printf("TestExtractImagesLowLevel readContext: %v\n", err)
		return err
	}

	// Optimize resource usage of this context.
	if err := api.OptimizeContext(ctx); err != nil {
		fmt.Printf("TestExtractImagesLowLevel optimizeContext: %v\n", err)
		return err
	}

	// Extract images for page 1.
	i := pageNum
	ii, err := pdfcpu.ExtractPageImages(ctx, i, false)
	if err != nil {
		fmt.Printf("TestExtractImagesLowLeveloptimizeContext extractPageFonts(%d): %v\n", i, err)
	}

	baseFileName := strings.TrimSuffix(filepath.Base(inputPath), ".pdf")

	// Process extracted images.
	for _, img := range ii {
		fn := filepath.Join(outputDir, fmt.Sprintf("%s_%d_%s.%s", baseFileName, i, img.Name, img.FileType))
		if err := pdfcpu.WriteReader(fn, img); err != nil {
			fmt.Printf("write: %s", fn)
			return err
		}
	}
	return nil
}