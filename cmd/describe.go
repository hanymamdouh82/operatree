package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/pkg/project"
	"github.com/spf13/cobra"
)

var plain bool

func init() {
	descCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	descCmd.Flags().BoolVarP(&plain, "plain", "p", false, "output raw YAML for piping")
	descCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(descCmd)
}

var descCmd = &cobra.Command{
	Use:   "describe",
	Short: "Describe a project",
	Long: `Print a structured, colored view of the project directory and its metadata.

Use --plain to output raw YAML instead — useful for piping into grep, sed, or other UNIX tools.

Flags:
  -p, --plain   Output raw YAML instead of styled view
  -d, --dest    Project directory to operate on

Examples:
  operatree describe
  operatree describe --plain
  operatree describe --plain | grep tags
  operatree describe -d /path/to/project`,
	Args: cobra.NoArgs,
	Run:  describe,
}

func describe(cmd *cobra.Command, args []string) {

	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}

	p.Describe(plain)
}
