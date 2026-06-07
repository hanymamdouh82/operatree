package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version   string
	commit    string
	buildDate string
)

// Called from main.go before Execute()
func SetVersion(v, c, d string) {
	version = v
	commit = c
	buildDate = d
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long: `Print the current OperaTree version, commit hash, and build date.

Examples:
  operatree version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("OperaTree %s\n", version)
		fmt.Printf("  Commit:     %s\n", commit)
		fmt.Printf("  Built:      %s\n", buildDate)
	},
}
