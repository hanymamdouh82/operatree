package project

import (
	"github.com/hanymamdouh82/operatree/internal/module"
)

// returns Dev project template
func tmpltGeneral(name string) Project {

	p := Project{
		Name: name,
		Modules: []module.Module{
			module.FactoryAdmin("00"),
			module.FactoryEvents("01"),
			module.FactoryProjectManagement("02"),
			module.FactoryMediaLib("03"),
			module.FactoryDeliverables("04"),
			module.FactoryArchive("99"),
		},
	}

	return p
}
