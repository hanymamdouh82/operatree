package cmd

import (
	"fmt"
	"log"

	"github.com/hanymamdouh82/operatree/pkg/config"
	"github.com/spf13/cobra"
)

func init() {
	untrackCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)

	rootCmd.AddCommand(untrackCmd)
}

var untrackCmd = &cobra.Command{
	Use:   "untrack [project_name]",
	Short: "Untrack a project",
	Long: `Remove a project from your OperaTree tracked projects list.

The project can be identified either by name or by directory path.
The project directory and its contents are not affected.

Resolution order:
  1. project_name argument — untrack by name as registered in config
  2. -d flag              — untrack by directory path

Flags:
  -d, --dest   Project directory to untrack

Examples:
  operatree untrack myproject              # untrack by name
  operatree untrack -d /path/to/project   # untrack by path
  operatree untrack -d .                  # untrack current directory`,
	Args: cobra.MaximumNArgs(1),
	Run:  untrack,
}

func untrack(cmd *cobra.Command, args []string) {

	// untrack using project name
	if len(args) > 0 {
		pn := args[0]
		if err := config.RemoveProjectByName(pn); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Project untracked: %s\n", pn)
		return
	}

	// Untrack using -d flag
	resolveProjectDir(cmd, args)
	if err := config.RemoveProject(actDir); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Project untracked: %s\n", actDir)
}
