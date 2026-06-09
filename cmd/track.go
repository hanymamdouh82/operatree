package cmd

import (
	"fmt"
	"log"

	"github.com/hanymamdouh82/operatree/pkg/config"
	"github.com/hanymamdouh82/operatree/pkg/project"
	"github.com/spf13/cobra"
)

func init() {
	trackCmd.Flags().StringVarP(&destDir, "dest", "d", "", dFlagHelp_baseDir)

	if err := trackCmd.MarkFlagRequired("dest"); err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(trackCmd)
}

var trackCmd = &cobra.Command{
	Use:   "track",
	Short: "Track a project",
	Long: `Register a project directory in your OperaTree configuration.

Tracked projects are available for default selection, the 'goto' command,
and all commands that resolve the project directory automatically.
The project directory must be provided via -d and is required.

Flags:
  -d, --dest   Project directory to register (required)

Examples:
  operatree track -d /path/to/project
  operatree track -d .`,
	Args: cobra.NoArgs,
	Run:  track,
}

func track(cmd *cobra.Command, args []string) {
	resolveProjectDir(cmd, args)

	// load project to confirm its state
	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}

	if err := config.AddProject(p.Name, p.ProjectDir(), p.Template); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Project tracked: %s\n", actDir)
}
