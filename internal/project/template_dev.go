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
		Name: name,
		Modules: []module.Module{
			module.FactoryAdmin("00"),
			module.FactoryEvents("01"),
			module.FactoryProjectManagement("02"),
			module.FactoryLegal("03"),
			module.FactoryResearch("04"),
			module.FactoryEngineering("05"),
			module.FactoryData("06"),
			module.FactoryMediaLib("07"),
			module.FactoryDeliverables("08"),
			module.FactoryArchive("99"),
		},
	}

	hydratePath(ppth, &p)
	return p
}
