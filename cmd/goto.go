package cmd

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/hanymamdouh82/operatree/internal/runner"
	"github.com/hanymamdouh82/operatree/pkg/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(gotoCmd)
}

var gotoCmd = &cobra.Command{
	Use:   "goto",
	Short: "Jump to a tracked project",
	Long: `Interactively select from your tracked projects and open the chosen
project directory in your configured file manager.

The file manager is set during 'operatree init' and stored in config.
To change it, update the 'fileManager' field in your config file.

Examples:
  operatree goto`,
	Run: gotoProject,
}

func gotoProject(cmd *cobra.Command, args []string) {

	c, err := config.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("cannot load config file. Use operatree init to initialize config"))
	}

	cp, err := config.Find(c)
	if err != nil {
		log.Fatal(err)
	}

	if err := runner.OpenFileManager(cp.AbsPath); err != nil {
		log.Fatal(err)
	}
}
