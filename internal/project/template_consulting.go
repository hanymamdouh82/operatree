package project

import (
	"github.com/hanymamdouh82/operatree/internal/module"
)

// returns Dev project template
func tmpltConsulting(name string) Project {

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

	return p
}
