package project

import (
	"path/filepath"

	"github.com/hanymamdouh82/operatree/pkg/module"
)

// hydratePath sets the runtime absolute paths on a project and everything under
// it. METADATA.yml stores only names, never paths, so these must be rebuilt on
// every load; that is what lets the same project resolve correctly when shared
// storage is mounted at a different location per user. projectBaseDir is the
// project root on the current machine. The result is in-memory only and is
// never written back to disk.
func hydratePath(projectBaseDir string, p *Project) {
	p.absDir = filepath.Clean(projectBaseDir)
	for i, m := range p.Modules {
		p.Modules[i].AbsPath = filepath.Join(p.absDir, m.Name)
		hydrateModule(&p.Modules[i])
	}
}

// hydrateModule fills in paths below a module whose AbsPath is already set: a
// DirName for each subject and an AbsPath for each nested module, recursing all
// the way down. A module is the same at any depth, so one routine handles every
// level.
func hydrateModule(m *module.Module) {

	for i, s := range m.Subjects {
		m.Subjects[i].DirName = filepath.Join(m.AbsPath, s.Name)
	}

	for i, sm := range m.Modules {
		m.Modules[i].AbsPath = filepath.Join(m.AbsPath, sm.Name)
		hydrateModule(&m.Modules[i])
	}
}
