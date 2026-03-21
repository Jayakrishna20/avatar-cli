package cmd

import (
	"avtr/processor"
	"strings"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"os"
)

var mergeCmd = &cobra.Command{
	Use:     "merge",
	Aliases: []string{"m"},
	Short:   "Merge multiple PDFs into a single file",
	Example: `  avtr merge --input part1.pdf --input part2.pdf --outdir output/ --output merged.pdf
  avtr m -i a.pdf -i b.pdf -d output/ -o merged.pdf`,
	Run: func(cmd *cobra.Command, args []string) {
		inputs, _ := cmd.Flags().GetStringSlice("input")
		outdir, _ := cmd.Flags().GetString("outdir")
		output, _ := cmd.Flags().GetString("output")

		if !strings.HasSuffix(strings.ToLower(output), ".pdf") {
			output += ".pdf"
		}

		if len(inputs) < 2 || outdir == "" || output == "" {
			color.Red("Error: at least two --input flags, --outdir, and --output are required.")
			cmd.Usage()
			os.Exit(1)
		}

		color.Cyan("🚀 Merging %d files into %s...", len(inputs), output)

		if err := processor.MergePDF(inputs, outdir, output); err != nil {
			color.Red("Task failed: %v", err)
			os.Exit(1)
		}
		color.Green("✨ Processing complete!")
	},
}

func init() {
	rootCmd.AddCommand(mergeCmd)
	mergeCmd.Flags().StringSliceP("input", "i", []string{}, "Input PDF files to merge (required, specify multiple times)")
	mergeCmd.Flags().StringP("outdir", "d", "", "Directory to save the merged PDF (required)")
	mergeCmd.Flags().StringP("output", "o", "", "Output filename for the merged PDF (required)")
}
