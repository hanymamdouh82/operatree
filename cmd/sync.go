package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/pkg/project"
	"github.com/spf13/cobra"
)

func init() {
	syncCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	syncCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(syncCmd)
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync project metadata",
	Long: `Walk the full project tree, re-read every META.yaml from disk, and update
the project metadata index.

Run this after editing subject files manually outside of OperaTree — for example
after bulk edits, git pulls, or file syncs that modify META.yaml files directly.
Note: the 'edit' command runs sync automatically after the editor is closed.

Examples:
  operatree sync
  operatree sync -d /path/to/project`,
	Args: cobra.NoArgs,
	Run:  sync,
}

func sync(cmd *cobra.Command, args []string) {
	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}

	// Sync project metadata with subject metadata
	if err := project.Sync(&p); err != nil {
		log.Fatal(err)
	}
}
