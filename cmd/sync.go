package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

var confirmSync bool

func init() {
	syncCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	syncCmd.Flags().BoolVarP(&confirmSync, "yes", "y", false, "confirm and apply sync changes")
	syncCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(syncCmd)
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Syncs project",
	Long:  "Syncs project subjects with project metadata - use WatchExec to automate sync",
	Args:  cobra.NoArgs,
	Run:   syncProject,
}

func syncProject(cmd *cobra.Command, args []string) {
	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}

	result, err := project.Sync(&p, confirmSync)
	if err != nil {
		log.Fatal(err)
	}

	result.Print(confirmSync)
}
