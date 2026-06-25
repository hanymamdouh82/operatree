package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/pkg/project"
	"github.com/spf13/cobra"
)

var newStatus string

func init() {
	editStatusCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	editStatusCmd.Flags().StringVarP(&uuid, "uuid", "u", "", "task UUID")
	editStatusCmd.Flags().StringVarP(&newStatus, "new-status", "n", "", "New task status")
	editStatusCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(editStatusCmd)
}

var editStatusCmd = &cobra.Command{
	Use:   "status [type] [term]",
	Short: "Update Task status",
	Long: `Fuzzy-find a task and updates its status.

Optionally narrow the search by providing a search term
before launching the interactive finder. The project metadata index is updated
automatically once the editor is closed.

Flags:
  -d, --dest         Project directory to operate on
  -u, --uuid         Subject UUID for non-interactive mode
  -n, --new-status   New status to assign to the task (required with --uuid)

Examples:
  operatree status                    # browse all tasks interactively
  operatree status report             # filter to tasks matching "report"
  operatree status --uuid a1b2c3d4 --new-status "Planning"
  operatree status --uuid a1b2c3d4 --new-status "planning" -d /path/to/project`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(1)),
	Run:  editStatus,
}

func editStatus(cmd *cobra.Command, args []string) {

	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}

	var term string
	if len(args) != 0 {
		term = args[0]
	}

	if err := p.EditTaskStatus(uuid, term, newStatus); err != nil {
		log.Fatal(err)
	}

}
