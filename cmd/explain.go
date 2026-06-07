package cmd

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/charmbracelet/glamour"
	"github.com/hanymamdouh82/operatree/internal/help"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(explainCmd)
}

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Explain the directory structure",
	Long: `Print the full OperaTree directory philosophy guide.

Covers what each folder is for, what belongs in it, and how the layers
relate to each other — useful when onboarding new team members or setting
up a project from scratch.

Examples:
  operatree explain
  operatree explain | less`,
	Run: explain,
}

func explain(cmd *cobra.Command, args []string) {

	data, _ := help.FS.ReadFile("dir_struct.md")

	out, err := glamour.Render(string(data), "dark")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(out)
}
