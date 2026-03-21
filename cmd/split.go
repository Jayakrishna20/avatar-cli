package cmd

import (
	"avtr/processor"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

var splitCmd = &cobra.Command{
	Use:     "split",
	Aliases: []string{"s"},
	Short:   "Split a PDF into multiple parts",
	Example: `  avtr split --input document.pdf --outdir output/ --span 1
  avtr s -i doc.pdf -o out/ -s 2`,
	Run: func(cmd *cobra.Command, args []string) {
		input, _ := cmd.Flags().GetString("input")
		outdir, _ := cmd.Flags().GetString("outdir")
		span, _ := cmd.Flags().GetInt("span")

		if input == "" || outdir == "" {
			color.Red("Error: --input and --outdir are required.")
			cmd.Usage()
			os.Exit(1)
		}

		color.Cyan("🚀 Splitting %s into span of %d pages...", input, span)

		if err := processor.SplitPDF(input, outdir, span); err != nil {
			color.Red("Task failed: %v", err)
			os.Exit(1)
		}
		color.Green("✨ Processing complete!")
	},
}

func init() {
	rootCmd.AddCommand(splitCmd)
	splitCmd.Flags().StringP("input", "i", "", "Input PDF file to split (required)")
	splitCmd.Flags().StringP("outdir", "o", "", "Directory to save the split PDFs (required)")
	splitCmd.Flags().IntP("span", "s", 1, "Number of pages per split file")
}
