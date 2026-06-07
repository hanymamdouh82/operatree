package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/pkg/config"
	"github.com/spf13/cobra"
)

func init() {
	setDPCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	setDPCmd.PreRun = resolveProjectDirSkippingConfig
	rootCmd.AddCommand(setDPCmd)
}

var setDPCmd = &cobra.Command{
	Use:   "use",
	Short: "Set the default project",
	Long: `Interactively select a default project from your tracked projects.

Once set, all commands use it automatically without requiring the -d flag.
To view the current default, run 'operatree show default'.

Examples:
  operatree use                  # pick default project interactively
  operatree show default         # show current default project`,
	Args: cobra.NoArgs,
	Run:  setDefaultProject,
}

func setDefaultProject(cmd *cobra.Command, args []string) {

	if destDir != "" {
		if err := config.SetDefaultProjectCLI(actDir); err != nil {
			log.Fatal(err)
		}
		return
	}

	if err := config.SetDefaultProjectInteractive(); err != nil {
		log.Fatal(err)
	}
}
