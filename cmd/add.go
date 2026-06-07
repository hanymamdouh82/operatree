package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/hanymamdouh82/operatree/pkg/project"
	"github.com/hanymamdouh82/operatree/pkg/subject"
	"github.com/spf13/cobra"
)

var subjectName string
var subjectDate string
var validSubjects []cobra.Completion
var addCmd = &cobra.Command{}

func init() {
	// build completion slice from available subjects dynamically
	for k := range project.SubjectModuleMap {
		sn := strings.ToLower(string(k))
		validSubjects = append(validSubjects, sn)
	}

	// define command
	addCmd = &cobra.Command{
		Use:   fmt.Sprintf("add [%s]", strings.Join(validSubjects, " | ")),
		Short: "Add a new subject",
		Long: `Launch an interactive form to add a new subject to the project.

The subject type is required and determines which fields are collected.
Use --name and --date to pre-fill those fields and skip their interactive prompts.
Every creation is appended to activity.log at the project root.

Flags:
  -d, --dest   Project directory to operate on
  --name       Subject name (skips interactive prompt)
  --date       Subject date (skips interactive prompt)

Examples:
  operatree add event                                        # fully interactive
  operatree add task --name "Prepare Report"                 # skip interactive prompt
  operatree add event --name "Site Visit" --date 2026-06-01  # skip interactive prompt`,
		ValidArgs: validSubjects,
		Args:      cobra.MatchAll(cobra.OnlyValidArgs, cobra.ExactArgs(1)),
		Run:       newSubject,
	}

	addCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_project)
	addCmd.Flags().StringVar(&subjectName, "name", "", "subject name")
	addCmd.Flags().StringVar(&subjectDate, "date", "", "subject date")
	addCmd.PreRun = resolveProjectDir
	rootCmd.AddCommand(addCmd)
}

func newSubject(cmd *cobra.Command, args []string) {
	a := args[0]

	// safety check -> dynamic loading prevents reach this error
	if a == "" {
		log.Fatal("unsupprted subject")
	}

	st := strings.ToUpper(a)

	p, err := project.Load(actDir)
	if err != nil {
		log.Fatal(err)
	}

	if err := project.NewSubject(&p, subjectName, subjectDate, subject.SubjectType(st)); err != nil {
		log.Fatal(err)
	}
}
