package project

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hanymamdouh82/operatree/pkg/module"
	"github.com/hanymamdouh82/operatree/pkg/subject"
)

type SyncEntry struct {
	Name string
	Type string
}

type SyncResult struct {
	Updated []SyncEntry
	New     []SyncEntry
	Skipped []SyncEntry
}

func (r SyncResult) Print(confirmed bool) {
	for _, e := range r.New {
		fmt.Printf("→  new       %-12s %q\n", e.Type, e.Name)
	}
	for _, e := range r.Updated {
		fmt.Printf("→  updated   %-12s %q\n", e.Type, e.Name)
	}
	for _, e := range r.Skipped {
		fmt.Printf("→  skipped   %-12s %q (malformed or missing METADATA.yml)\n", e.Type, e.Name)
	}

	if confirmed {
		fmt.Printf("\nSync complete: %d updated, %d new, %d skipped\n",
			len(r.Updated), len(r.New), len(r.Skipped))
	} else {
		fmt.Printf("\nDry-run: %d to update, %d new, %d skipped — run with -y to apply\n",
			len(r.Updated), len(r.New), len(r.Skipped))
	}
}

func Sync(p *Project, confirm bool) (SyncResult, error) {
	var result SyncResult

	for i := range p.Modules {
		if err := syncModule(&p.Modules[i], confirm, &result); err != nil {
			return result, err
		}

		if err := discoverModule(&p.Modules[i], confirm, &result); err != nil {
			return result, err
		}
	}

	if confirm {
		if err := p.WriteMetadata(); err != nil {
			return result, err
		}
	}

	return result, nil
}

func syncModule(m *module.Module, confirm bool, result *SyncResult) error {
	for j := range m.Subjects {
		onDisk, err := m.Subjects[j].ReadMetadata()
		if err != nil {
			log.Printf("skipping subject %s: %v\n", m.Subjects[j].DirName, err)
			result.Skipped = append(result.Skipped, SyncEntry{Name: m.Subjects[j].Name, Type: string(m.Subjects[j].Type)})
			continue
		}

		result.Updated = append(result.Updated, SyncEntry{Name: onDisk.Name, Type: string(onDisk.Type)})
		if confirm {
			m.Subjects[j] = *onDisk
		}
	}

	for i := range m.Modules {
		if err := syncModule(&m.Modules[i], confirm, result); err != nil {
			return err
		}
	}

	return nil
}

func discoverModule(m *module.Module, confirm bool, result *SyncResult) error {
	entries, err := os.ReadDir(m.AbsPath)
	if err != nil {
		return err
	}

	indexed := make(map[string]bool, len(m.Subjects))
	for _, s := range m.Subjects {
		if s.UUID != "" {
			indexed[s.UUID] = true
		}
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		candidate := subject.Subject{DirName: filepath.Join(m.AbsPath, entry.Name())}

		onDisk, err := candidate.ReadMetadata()
		if err != nil {
			continue
		}
		if _, ok := SubjectModuleMap[onDisk.Type]; !ok {
			continue
		}

		if onDisk.UUID != "" && indexed[onDisk.UUID] {
			continue
		}

		result.New = append(result.New, SyncEntry{Name: onDisk.Name, Type: string(onDisk.Type)})
		if confirm {
			if onDisk.UUID == "" {
				if err := onDisk.SetID(); err != nil {
					return err
				}
				if err := onDisk.WriteMetadata(); err != nil {
					return err
				}
			}
			m.Subjects = append(m.Subjects, *onDisk)
		}
	}

	for i := range m.Modules {
		if err := discoverModule(&m.Modules[i], confirm, result); err != nil {
			return err
		}
	}

	return nil
}
