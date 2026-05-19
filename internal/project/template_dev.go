package project

import (
	"path"

	"github.com/hanymamdouh82/operatree/internal/module"
)

// returns Dev project template
// `bpth` is the abs base dir for the project, without project name included into it
func tmpltDev(name string, bpth string) Project {

	ppth := path.Join(bpth, name)

	p := Project{
		Name:    name,
		BaseDir: bpth,
		Modules: []module.Module{
			module.FactoryAdmin(ppth),
			module.FactoryEvents(ppth),
			module.FactoryProjectManagement(ppth),
			module.FactoryResearch(ppth),
			module.FactoryEngineering(ppth),
			module.FactoryData(ppth),
			module.FactoryMediaLib(ppth),
			module.FactoryDeliverables(ppth),
			module.FactoryArchive(ppth),
		},
	}

	return p
}
