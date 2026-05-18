package cmd

import (
	"fmt"
	"log"

	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/hanymamdouh82/operatree/internal/units/event"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var entityName string

func init() {
	newCmd.Flags().StringVarP(&pDir, "dest", "d", "/mnt/extra/onfly/testprj", "project directory")
	newCmd.Flags().StringVarP(&entityName, "name", "n", "", "entity name")
	rootCmd.AddCommand(newCmd)
}

var newCmd = &cobra.Command{
	Use:       "new [event | task]",
	Short:     "creates new entity",
	Long:      "creates new entity within project",
	ValidArgs: []cobra.Completion{"event", "task"},
	Args:      cobra.MatchAll(cobra.OnlyValidArgs, cobra.ExactArgs(1)),
	Run:       newUnitEntity,
}

func newUnitEntity(cmd *cobra.Command, args []string) {
	a := args[0]
	p, err := project.Load(pDir)
	if err != nil {
		log.Fatal(err)
	}

	switch a {
	case "event":
		newEvent(&p)
	case "task":
		fmt.Println("will create task")
	default:
		fmt.Println("i shouldn't be here")
	}
}

func newEvent(p *project.Project) error {
	u, err := p.UnitEvents()
	if err != nil {
		return err
	}

	// if name flag is not provided, we open interactive CLI
	if entityName == "" {
		var e event.Event
		if e, err = u.NewInteractive(); err != nil {
			return err
		}
		y, _ := yaml.Marshal(e)
		fmt.Printf("%s\n", y)
		return nil
	}

	// create empty event using provided name flag only and default values
	u.New(entityName)

	return nil
}
