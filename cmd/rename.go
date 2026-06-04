package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/pkg/project"
	"github.com/spf13/cobra"
)

func init() {
	renameSubjectCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	renameSubjectCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(renameSubjectCmd)
}

var renameSubjectCmd = &cobra.Command{
	Use:   "rename [type] [term]",
	Short: "Finds a subject in project and rename it",
	Long:  "Fuzzy-Find a subject in a project and rename it",
	Args:  cobra.MatchAll(cobra.MaximumNArgs(2)),
	Run:   renameSubject,
}

func renameSubject(cmd *cobra.Command, args []string) {

	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}

	var t, term string

	if len(args) == 2 {
		t = args[0]
		term = args[1]
	} else if len(args) == 1 {
		term = args[0]
	} else {
		t = ""
		term = ""
	}

	if err := p.RenameSubject(t, term, ""); err != nil {
		log.Fatal(err)
	}
}
