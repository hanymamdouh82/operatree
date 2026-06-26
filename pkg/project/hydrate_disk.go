package project

import (
	"os"
	"path/filepath"

	"github.com/hanymamdouh82/operatree/pkg/module"
	"github.com/hanymamdouh82/operatree/pkg/subject"
)

func hydrateFromDisk(p *Project) {
	for i := range p.Modules {
		hydrateModuleFromDisk(&p.Modules[i])
	}
}

func hydrateModuleFromDisk(m *module.Module) {
	byUUID := make(map[string]int, len(m.Subjects))
	byDir := make(map[string]int, len(m.Subjects))
	for i := range m.Subjects {
		if m.Subjects[i].UUID != "" {
			byUUID[m.Subjects[i].UUID] = i
		}
		byDir[filepath.Clean(m.Subjects[i].DirName)] = i
	}

	nested := make(map[string]bool, len(m.Modules))
	for _, sm := range m.Modules {
		nested[sm.Name] = true
	}

	if entries, err := os.ReadDir(m.AbsPath); err == nil {
		for _, entry := range entries {
			if !entry.IsDir() || nested[entry.Name()] {
				continue
			}

			dir := filepath.Join(m.AbsPath, entry.Name())
			candidate := subject.Subject{DirName: dir}
			onDisk, err := candidate.ReadMetadata()
			if err != nil {
				continue
			}
			if _, ok := SubjectModuleMap[onDisk.Type]; !ok {
				continue
			}
			onDisk.Name = entry.Name()

			if onDisk.UUID != "" {
				if idx, ok := byUUID[onDisk.UUID]; ok {
					m.Subjects[idx] = *onDisk
					continue
				}
			}
			if idx, ok := byDir[filepath.Clean(dir)]; ok {
				m.Subjects[idx] = *onDisk
				continue
			}
			m.Subjects = append(m.Subjects, *onDisk)
		}
	}

	for i := range m.Modules {
		hydrateModuleFromDisk(&m.Modules[i])
	}
}
