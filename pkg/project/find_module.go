package project

import (
	"fmt"

	"github.com/hanymamdouh82/operatree/pkg/module"
)

// findModule recursively searches for a module by type within the project hierarchy.
func findModule(modules []module.Module, tmt module.ModuleType) (*module.Module, error) {
	for i, m := range modules {
		// Check if this module matches the target type
		if m.Type == tmt {
			return &modules[i], nil
		}

		// Recursively search submodules
		if len(m.Modules) > 0 {
			if found, err := findModule(m.Modules, tmt); err == nil && found != nil {
				return found, nil
			}
		}
	}

	return nil, fmt.Errorf("module type %s not found in project", string(tmt))
}
