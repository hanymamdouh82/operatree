package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/hanymamdouh82/operatree/internal/template"
	"github.com/hanymamdouh82/operatree/pkg/project"
	"github.com/spf13/cobra"
)

var templateName string

func init() {
	ts := make([]string, 0, len(template.Templates))
	for k := range template.Templates {
		ts = append(ts, k)
	}
	avts := strings.Join(ts, "|")
	fth := fmt.Sprintf("project template: %s", avts)

	createCmd.Flags().StringVarP(&destDir, "dest", "d", actDir, dFlagHelp_baseDir)
	createCmd.Flags().StringVarP(&templateName, "template", "t", "", fth)
	createCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "show operation output")

	if err := createCmd.MarkFlagRequired("template"); err != nil {
		log.Fatal(err)
	}

	rootCmd.AddCommand(createCmd)
}

var createCmd = &cobra.Command{
	Use:   "create [project_name]",
	Short: "Create a new project",
	Long: `Scaffold a new project with the full OperaTree directory structure and register it in config.

The project folder is created under the base directory (-d) using the provided name.
A template (-t) is required and determines which modules are included.
Use 'operatree show templates' to list available templates.

Flags:
  -t, --template   Project template to use (required)
  -d, --dest       Base directory for the new project
  -v, --verbose    Print the created directory structure after creation

Examples:
  operatree create myproject -t dev
  operatree create research-2026 -t research -v
  operatree create myproject -t dev -d /home/user/projects`,
	Args: cobra.ExactArgs(1),
	Run:  bootstrap,
}

func bootstrap(cmd *cobra.Command, args []string) {
	// -d flag here is used to define base dir not project dir
	resolveBaseDir(cmd, args)

	pn := args[0]
	p, err := project.Bootstrap(pn, actDir, templateName)
	if err != nil {
		log.Fatal(err)
	}

	if verbose {
		if err := p.Describe(false); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Project: %s (%s)\n", p.Name, p.ProjectDir())
}
