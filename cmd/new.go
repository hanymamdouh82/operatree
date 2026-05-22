package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

var silent bool
var subjectName string
var subjectDate string

func init() {
	newCmd.Flags().StringVar(&subjectName, "name", "", "subject name")
	newCmd.Flags().StringVar(&subjectDate, "date", "", "subject date")
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:       "new [event | task | topic | objective]",
	Short:     "Creates new subject",
	Long:      "Creates new subject within project",
	ValidArgs: SubjectValidArgs,
	Args:      cobra.MatchAll(cobra.OnlyValidArgs, cobra.ExactArgs(1)),
	Run:       newSubject,
}

func newSubject(cmd *cobra.Command, args []string) {
	a := args[0]
	p, err := project.Load(prjDir)
	if err != nil {
		log.Fatal(err)
	}

	switch a {
	case "event":
		if err := project.NewEvent(&p, subjectName, subjectDate); err != nil {
			log.Fatal(err)
		}

	case "task":
		if err := project.NewTask(&p, subjectName, subjectDate); err != nil {
			log.Fatal(err)
		}

	case "topic":
		if err := project.NewTopic(&p, subjectName, subjectDate); err != nil {
			log.Fatal(err)
		}

	case "objective":
		if err := project.NewObjective(&p, subjectName, subjectDate); err != nil {
			log.Fatal(err)
		}

	default:
		return
	}
}
