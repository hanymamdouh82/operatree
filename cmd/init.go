package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/pkg/config"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize OperaTree",
	Long: `Create the OperaTree configuration file in the system config directory.

Interactively prompts for your standard projects directory, preferred editor,
and file manager. Run this once before using any other command.

Config file location (in order of priority):
  1. $XDG_CONFIG_HOME/operatree/          if XDG_CONFIG_HOME is set
  2. ~/.config/operatree/                 Linux default
  3. ~/Library/Application Support/operatree/   macOS default

Examples:
  operatree init`,
	Args: cobra.NoArgs,
	Run:  initConfig,
}

func initConfig(cmd *cobra.Command, args []string) {
	if err := config.InitializeConfig(); err != nil {
		log.Fatal(err)
	}
}
