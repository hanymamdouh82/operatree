package cmd

import (
	"log"

	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

var (
	bDir string
)

func init() {
	bootstrapCmd.Flags().StringVarP(&bDir, "dest", "d", "/mnt/extra/onfly", "project root directory")
	rootCmd.AddCommand(bootstrapCmd)
}

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap [project_name]",
	Short: "Bootstraps new project",
	Long:  `Bootstraps new porject in current working directory`,
	Args:  cobra.ExactArgs(1),
	Run:   bootstrap,
}

func bootstrap(cmd *cobra.Command, args []string) {
	pn := args[0]
	p, err := project.Bootstrap(bDir, pn)
	if err != nil {
		log.Fatal(err)
	}

	p.Describe()
}
