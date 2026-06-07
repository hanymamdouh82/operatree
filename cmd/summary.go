package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/pkg/project"
	"github.com/spf13/cobra"
)

func init() {
	summaryCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	summaryCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(summaryCmd)
}

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show project summary",
	Long: `Print a high-level overview of the project at a glance.

Displays subject counts broken down by type and status — useful for a quick
pulse check on project activity without browsing individual subjects.

Examples:
  operatree summary
  operatree summary -d /path/to/project`,
	Args: cobra.NoArgs,
	Run:  summary,
}

func summary(cmd *cobra.Command, args []string) {
	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}
	p.Summary()
}
