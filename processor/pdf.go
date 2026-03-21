package processor

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/schollz/progressbar/v3"
	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func showSpinner(desc string, done chan bool) {
	bar := progressbar.NewOptions(-1,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetDescription(fmt.Sprintf("[cyan]%s[reset]", desc)),
		progressbar.OptionSpinnerType(14),
	)
	for {
		select {
		case <-done:
			bar.Finish()
			fmt.Println()
			return
		default:
			bar.Add(1)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func MergePDF(inputFiles []string, outputDir, outputFilename string) error {
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	outputPath := filepath.Join(outputDir, outputFilename)

	done := make(chan bool)
	go showSpinner("Merging PDFs...", done)

	err := api.MergeCreateFile(inputFiles, outputPath, false, nil)
	done <- true
	
	if err != nil {
		return fmt.Errorf("failed to merge pdfs: %w", err)
	}

	return nil
}

func SplitPDF(inputFile, outputDir string, span int) error {
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	done := make(chan bool)
	go showSpinner("Splitting PDF...", done)

	err := api.SplitFile(inputFile, outputDir, span, nil)
	done <- true
	
	if err != nil {
		return fmt.Errorf("failed to split pdf: %w", err)
	}

	return nil
}
