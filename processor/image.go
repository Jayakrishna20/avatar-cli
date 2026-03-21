package processor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/schollz/progressbar/v3"
	"github.com/signintech/gopdf"
)

func ImageToPDF(inputDir, outputDir, outputFilename, pageSizeStr string) error {
	files, err := os.ReadDir(inputDir)
	if err != nil {
		return fmt.Errorf("failed to read input directory: %w", err)
	}

	pdf := gopdf.GoPdf{}
	
	pageSize := gopdf.PageSizeA4
	switch strings.ToUpper(pageSizeStr) {
	case "A3":
		pageSize = gopdf.PageSizeA3
	case "A5":
		pageSize = gopdf.PageSizeA5
	}
	
	pdf.Start(gopdf.Config{PageSize: *pageSize})

	count := 0
	bar := progressbar.Default(int64(len(files)), "Converting Images")
	for _, f := range files {
		_ = bar.Add(1)
		if f.IsDir() {
			continue
		}
		imgPath := filepath.Join(inputDir, f.Name())
		ext := strings.ToLower(filepath.Ext(imgPath))
		if ext == ".png" || ext == ".jpg" || ext == ".jpeg" {
			pdf.AddPage()
			err := pdf.Image(imgPath, 0, 0, &gopdf.Rect{W: pageSize.W, H: pageSize.H})
			if err != nil {
				return fmt.Errorf("failed to add image %s: %w", imgPath, err)
			}
			count++
		}
	}
	
	if count == 0 {
		return fmt.Errorf("no valid images found in input_dir")
	}

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	outputPath := filepath.Join(outputDir, outputFilename)
	err = pdf.WritePdf(outputPath)
	if err != nil {
		return fmt.Errorf("failed to write pdf: %w", err)
	}
	
	bar.Finish()
	fmt.Println()

	return nil
}
