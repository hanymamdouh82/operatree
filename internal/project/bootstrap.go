package project

import (
	"fmt"
	"os"

	"github.com/hanymamdouh82/operatree/internal/filesystem"
	"github.com/hanymamdouh82/operatree/internal/units/admin"
	"github.com/hanymamdouh82/operatree/internal/units/archive"
	"github.com/hanymamdouh82/operatree/internal/units/deliverables"
	"github.com/hanymamdouh82/operatree/internal/units/engineering"
	"github.com/hanymamdouh82/operatree/internal/units/event"
	"github.com/hanymamdouh82/operatree/internal/units/legal"
	"github.com/hanymamdouh82/operatree/internal/units/medialib"
)

// bootstraps a project by creating directory structure for primary entities
// pth is the parent path where project will be created
// In future features, pth will be stored into DB
func Bootstrap(pth string, name string) (Project, error) {

	p := Project{
		name:    name,
		baseDir: pth,
	}

	if err := filesystem.CheckDirExists(p.ProjectDir()); err != nil {
		return p, err
	}

	if err := os.Mkdir(p.ProjectDir(), 0775); err != nil {
		return p, err
	}

	// Bootstrap units
	// To-Do: units can be parsed from CLI or template based on user type or project type
	p.AddUnit(&admin.UnitAdmin{})
	p.AddUnit(&event.UnitEvents{})
	p.AddUnit(&legal.UnitLegal{})
	p.AddUnit(&engineering.UnitEngineering{})
	p.AddUnit(&medialib.UnitMediaLib{})
	p.AddUnit(&deliverables.UnitDeliverables{})
	p.AddUnit(&archive.UnitArchive{})

	// iterate and invoke every bootstrap function
	// we collect bootstrapping results and we don't interrupt the process
	bes := []error{}
	for _, v := range p.Units {
		if err := v.Bootstrap(p.ProjectDir()); err != nil {
			bes = append(bes, err)
		}
	}

	// dump errors
	for _, v := range bes {
		fmt.Printf("%s\n", v.Error())
	}

	return p, nil
}
