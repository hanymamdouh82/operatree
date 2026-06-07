package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/pkg/project"
	"github.com/spf13/cobra"
)

func init() {
	findCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	findCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(findCmd)
}

var findCmd = &cobra.Command{
	Use:   "find [type] [term]",
	Short: "Find a subject",
	Long: `Fuzzy-find subjects across all metadata fields — name, tags, participants, notes, date, and location.

Optionally narrow the search by providing a subject type, a search term, or both
before launching the interactive finder. The finder includes a live preview panel
for the selected subject.

Flags:
  -d, --dest   Project directory to operate on

Examples:
  operatree find                        # browse all subjects interactively
  operatree find event                  # filter to events, then pick one
  operatree find event cairo            # filter to events matching "cairo"
  operatree find cairo                  # search "cairo" across all subject types`,
	Args: cobra.MatchAll(cobra.MaximumNArgs(2)),
	Run:  find,
}

func find(cmd *cobra.Command, args []string) {

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

	s, err := project.FindSubjects(&p, t, term)
	if err != nil {
		log.Fatal(err)
	}

	if s.Type != "" {
		s.Describe()
	}
}
