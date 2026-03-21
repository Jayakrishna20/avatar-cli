package cmd

import (
	"avtr/processor"
	"strings"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

var img2pdfCmd = &cobra.Command{
	Use:     "img2pdf",
	Aliases: []string{"i"},
	Short:   "Convert a directory of images to a single PDF",
	Example: `  avtr img2pdf --indir images/ --outdir output/ --output combined.pdf --pagesize A4
  avtr i -d images/ -o output/ -f combined.pdf`,
	Run: func(cmd *cobra.Command, args []string) {
		indir, _ := cmd.Flags().GetString("indir")
		outdir, _ := cmd.Flags().GetString("outdir")
		output, _ := cmd.Flags().GetString("output")
		pageSize, _ := cmd.Flags().GetString("pagesize")

		if !strings.HasSuffix(strings.ToLower(output), ".pdf") {
			output += ".pdf"
		}

		if indir == "" || outdir == "" || output == "" {
			color.Red("Error: --indir, --outdir, and --output are required.")
			cmd.Usage()
			os.Exit(1)
		}

		color.Cyan("🚀 Converting images in %s to %s...", indir, output)

		if err := processor.ImageToPDF(indir, outdir, output, pageSize); err != nil {
			color.Red("Task failed: %v", err)
			os.Exit(1)
		}
		color.Green("✨ Processing complete!")
	},
}

func init() {
	rootCmd.AddCommand(img2pdfCmd)
	img2pdfCmd.Flags().StringP("indir", "d", "", "Directory containing images to convert (required)")
	img2pdfCmd.Flags().StringP("outdir", "o", "", "Directory to save the generated PDF (required)")
	img2pdfCmd.Flags().StringP("output", "f", "", "Output filename for the PDF (required)")
	img2pdfCmd.Flags().StringP("pagesize", "p", "A4", "Page size for the PDF (A3, A4, A5)")
}
