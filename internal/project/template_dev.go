package project

import (
	"github.com/hanymamdouh82/operatree/internal/module"
)

// returns Dev project template
func tmpltDev(name string) Project {

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

	return p
}
