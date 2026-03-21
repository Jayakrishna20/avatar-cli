package cmd

import (
	"fmt"
	"runtime"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"u"},
	Short:   "Instructions to update Avatar to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan("🚀 Updating Avatar CLI to the latest version...\n")

		fmt.Println("Depending on how you installed Avatar, run one of the following commands:")

		color.Green("▶ If installed via Go:")
		fmt.Println("  go install github.com/Jayakrishna20/file-converter@latest")
		fmt.Println("")

		if runtime.GOOS == "windows" {
			color.Green("▶ If installed via Winget:")
			fmt.Println("  winget upgrade avtr")
			fmt.Println("")
			
			color.Green("▶ If installed via Chocolatey:")
			fmt.Println("  choco upgrade avtr")
			fmt.Println("")
		} else if runtime.GOOS == "darwin" {
			color.Green("▶ If installed via Homebrew:")
			fmt.Println("  brew upgrade avtr")
			fmt.Println("")

			color.Green("▶ If installed via the Install Script:")
			fmt.Println("  curl -sSL https://raw.githubusercontent.com/Jayakrishna20/file-converter/main/install.sh | bash")
			fmt.Println("")
		} else if runtime.GOOS == "linux" {
			color.Green("▶ If installed via APT (Debian/Ubuntu):")
			fmt.Println("  sudo apt update && sudo apt install --only-upgrade avtr")
			fmt.Println("")

			color.Green("▶ If installed via DNF (Fedora/RHEL):")
			fmt.Println("  sudo dnf upgrade avtr")
			fmt.Println("")

			color.Green("▶ If installed via the Install Script:")
			fmt.Println("  curl -sSL https://raw.githubusercontent.com/Jayakrishna20/file-converter/main/install.sh | bash")
			fmt.Println("")
		}

		color.Yellow("💡 Note: You can always verify your version by running: avtr --version")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
