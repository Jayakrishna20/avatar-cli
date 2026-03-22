package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Version string

var rootCmd = &cobra.Command{
	Use:     "avtr",
	Short:   "Avatar: A powerful CLI tool for PDF and Image manipulations",
	Long:    `Avatar (avtr) provides lightning-fast operations to split, merge, and convert PDFs and Images.`,
}

func Execute() {
	rootCmd.Version = Version
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Root flags can be added here
}
