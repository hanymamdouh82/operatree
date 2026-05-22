package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/internal/config"
	"github.com/spf13/cobra"
)

var showDefault bool

func init() {
	setDPCmd.Flags().BoolVar(&showDefault, "show", false, "show current default project")
	rootCmd.AddCommand(setDPCmd)
}

var setDPCmd = &cobra.Command{
	Use:   "default",
	Short: "Set or show default project",
	Long:  "Sets a default project from tracked projects, or shows the current default",
	Args:  cobra.NoArgs,
	Run:   setDefaultProject,
}

func setDefaultProject(cmd *cobra.Command, args []string) {

	if showDefault {
		config.ShowDefulatProject()
		return
	}

	if prjDir != "" {
		if err := config.SetDefaultProjectCLI(prjDir); err != nil {
			log.Fatal(err)
		}
		return
	}

	if err := config.SetDefaultProjectInteractive(); err != nil {
		log.Fatal(err)
	}
}
