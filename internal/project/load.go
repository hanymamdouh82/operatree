package project

import (
	"os"
	"path"

	"github.com/hanymamdouh82/operatree/internal/types"
	"github.com/hanymamdouh82/operatree/internal/units/admin"
	"github.com/hanymamdouh82/operatree/internal/units/archive"
	"github.com/hanymamdouh82/operatree/internal/units/deliverables"
	"github.com/hanymamdouh82/operatree/internal/units/engineering"
	"github.com/hanymamdouh82/operatree/internal/units/event"
	"github.com/hanymamdouh82/operatree/internal/units/legal"
	"github.com/hanymamdouh82/operatree/internal/units/medialib"
	"gopkg.in/yaml.v3"
)

type unitLoader struct {
	Type       string `yaml:"type"`
	Name       string `yaml:"name"`
	ParentPath string `yaml:"parentPath"`
	UnitPath   string `yaml:"unitPath"`
}

type projectLoader struct {
	Name    string       `yaml:"name"`
	BaseDir string       `yaml:"baseDir"`
	Tags    []string     `yaml:"tags"`
	Units   []unitLoader `yaml:"units"`
}

// Loads a project by reading project metadata file and sets project structrue
// Path represents project root path
func Load(pth string) (Project, error) {

	b, err := os.ReadFile(path.Join(pth, METADATA_FILE))
	if err != nil {
		return Project{}, err
	}

	// unmarsahl into loader struct, this is because Unit is an interface not a struct
	// then we can convert units as per build logic
	var pl projectLoader
	if err := yaml.Unmarshal(b, &pl); err != nil {
		return Project{}, err
	}

	p, err := build(pl)
	if err != nil {
		return Project{}, err
	}

	return p, err
}

func build(pl projectLoader) (Project, error) {

	p := Project{
		Name:    pl.Name,
		BaseDir: pl.BaseDir,
		Tags:    pl.Tags,
	}

	// load units here
	// To-Do: implement LoadUnit instead of AddUnit
	for _, v := range pl.Units {
		switch types.UnitType(v.Type) {
		case types.UnitAdmin:
			p.AddUnit(&admin.UnitAdmin{}, types.UnitAdmin)
		case types.UnitArchive:
			p.AddUnit(&archive.UnitArchive{}, types.UnitArchive)
		case types.UnitDeliverables:
			p.AddUnit(&deliverables.UnitDeliverables{}, types.UnitDeliverables)
		case types.UnitEngineering:
			p.AddUnit(&engineering.UnitEngineering{}, types.UnitEngineering)
		case types.UnitEvents:
			p.AddUnit(&event.UnitEvents{}, types.UnitEvents)
		case types.UnitLegal:
			p.AddUnit(&legal.UnitLegal{}, types.UnitLegal)
		case types.UnitMediaLib:
			p.AddUnit(&medialib.UnitMediaLib{}, types.UnitMediaLib)
		}
	}

	return p, nil
}
