package project

import (
	"path"

	"github.com/hanymamdouh82/operatree/internal/module"
)

// returns Dev project template
// `bpth` is the abs base dir for the project, without project name included into it
func tmpltConsulting(name string, bpth string) Project {

	ppth := path.Join(bpth, name)

	p := Project{
		Name: name,
		Modules: []module.Module{
			module.FactoryAdmin("00"),
			module.FactoryEvents("01"),
			module.FactoryLegal("02"),
			module.FactoryResearch("03"),
			module.FactoryDeliverables("04"),
			module.FactoryArchive("99"),
		},
	}

	hydratePath(ppth, &p)

	return p
}
