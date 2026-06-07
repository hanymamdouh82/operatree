package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/hanymamdouh82/operatree/internal/template"
	"github.com/hanymamdouh82/operatree/pkg/config"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(showCmd)
}

var va = []cobra.Completion{"tracked", "config", "templates", "default"}

var showCmd = &cobra.Command{
	Use:   fmt.Sprintf("show [%s]", strings.Join(va, " | ")),
	Short: "Show OperaTree information",
	Long: `Display information about OperaTree configuration and state.

Arguments:
  tracked     List all tracked projects
  config      Display current configuration
  templates   List available project templates
  default     Show the currently set default project

Examples:
  operatree show tracked
  operatree show config
  operatree show templates
  operatree show default`,
	ValidArgs: va,
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run:       show,
}

func show(cmd *cobra.Command, args []string) {
	// Load config
	c, err := config.Load()
	if err != nil {
		log.Fatal(fmt.Errorf("cannot load config file. Use operatree init to initialize config"))
	}

	switch args[0] {
	case va[0]:
		c.ListProjects()
		return
	case va[1]:
		b, err := yaml.Marshal(c)
		if err != nil {
			log.Fatal()
		}
		fmt.Printf("%s\n", b)
	case va[2]:
		template.ListTemplates()
		return
	case va[3]:
		config.ShowDefulatProject()
		return
	default:
		log.Fatal("unknow command")
	}

}
