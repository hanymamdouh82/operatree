package project

import (
	"github.com/hanymamdouh82/operatree/internal/module"
)

// returns Dev project template
func tmpltResearch(name string) Project {

	p := Project{
		Name: name,
		Modules: []module.Module{
			module.FactoryAdmin("00"),
			module.FactoryEvents("01"),
			module.FactoryMediaLib("02"),
			module.FactoryResearch("03"),
			module.FactoryPublications("04"),
			module.FactoryArchive("99"),
		},
	}

	return p
}
