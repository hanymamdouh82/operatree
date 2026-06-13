// internal/config/config.go
package config

import "fmt"

type Config struct {
	StandardDir string    `yaml:"standardDir"` // default project base dir
	Default     Project   `yaml:"default"`     // default project - used when current dir doen't include project yaml of -d is not used
	Editor      string    `yaml:"editor"`      // Default file editor - if not provided fallback to $EDITOR
	FileManager string    `yaml:"fileManager"` // Default file manager
	Projects    []Project `yaml:"projects"`    // tracked projects
	Daemon      Daemon    `yaml:"daemon"`      // future daemon config
}

type Project struct {
	Name     string `yaml:"name"`
	AbsPath  string `yaml:"absPath"`
	Template string `yaml:"template"` // e.g. "dev", "research"
}

type Daemon struct {
	Enabled  bool   `yaml:"enabled"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DBDriver string `yaml:"dbDriver"` // sqlite, mysql, etc
	DSN      string `yaml:"dsn"`      // connection string
}

func (c *Config) ListProjects() {
	pn := "├──"
	cn := "│   └──"

	fmt.Printf("Tracked Projects:\n")

	for i, p := range c.Projects {
		if i == len(c.Projects)-1 {
			pn = "└──"
			cn = "    └──"
		}

		fmt.Printf("%s %s\n", pn, p.Name)
		fmt.Printf("%s %s\n", cn, p.AbsPath)
	}
}
