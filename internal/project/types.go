package project

import "github.com/hanymamdouh82/operatree/internal/module"

const (
	METADATA_FILE = "METADATA.yml"
	ARCHIVED_DEST = "closed_tasks"
)

type Project struct {
	Name     string          `yaml:"name"`
	Template string          `yaml:"template"`
	BaseDir  string          `yaml:"baseDir"`
	Tags     []string        `yaml:"tags"`
	Modules  []module.Module `yaml:"modules"`
}

// project templates map
type tmpltMap map[string]func(name string, bpth string) Project

var templates tmpltMap = tmpltMap{
	"general": tmpltGeneral,
	"dev":     tmpltDev,
}
