package project

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/hanymamdouh82/operatree/internal/filesystem"
	"github.com/hanymamdouh82/operatree/internal/module"
	"github.com/hanymamdouh82/operatree/internal/subject"
	"gopkg.in/yaml.v3"
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

var validSubjectTypes = map[subject.SubjectType]bool{
	subject.SubjectEvent:     true,
	subject.SubjectTask:      true,
	subject.SubjectTopic:     true,
	subject.SubjectObjective: true,
}

// Sync combines discovery and sync into a single operation.
// If confirm is false it is a dry-run — nothing is written.
// If confirm is true changes are applied and metadata is saved.
func Sync(p *Project, confirm bool) (SyncResult, error) {
	var result SyncResult

	for i := range p.Modules {
		// sync existing subjects from disk
		if err := syncModule(&p.Modules[i], confirm, &result); err != nil {
			return result, err
		}

		// discover new subjects on disk not yet in the index
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

// syncModule updates existing indexed subjects from their METADATA.yml on disk.
func syncModule(m *module.Module, confirm bool, result *SyncResult) error {
	for j, s := range m.Subjects {
		b, err := filesystem.ReadFile(path.Join(s.DirName, subject.METADATA_FILE))
		if err != nil {
			log.Printf("missing yml for subject %s\n", s.DirName)
			result.Skipped = append(result.Skipped, SyncEntry{Name: s.Name, Type: string(s.Type)})
			continue
		}

		var diskMeta subject.Subject
		if err := yaml.Unmarshal(b, &diskMeta); err != nil {
			log.Printf("malformed yml for subject %s\n", s.DirName)
			result.Skipped = append(result.Skipped, SyncEntry{Name: s.Name, Type: string(s.Type)})
			continue
		}

		result.Updated = append(result.Updated, SyncEntry{Name: diskMeta.Name, Type: string(diskMeta.Type)})
		if confirm {
			m.Subjects[j] = diskMeta
		}
	}

	for i := range m.Modules {
		if err := syncModule(&m.Modules[i], confirm, result); err != nil {
			return err
		}
	}

	return nil
}

// discoverModule walks the module directory on disk and registers any folder
// containing a METADATA.yml with a valid type that is not already indexed.
func discoverModule(m *module.Module, confirm bool, result *SyncResult) error {
	entries, err := os.ReadDir(m.AbsPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		dirPath := path.Join(m.AbsPath, entry.Name())

		// skip if already indexed — compare by Name since DirName is
		// never stored in YAML and may be empty after syncModule runs
		alreadyIndexed := false
		for _, s := range m.Subjects {
			if s.Name == entry.Name() {
				alreadyIndexed = true
				break
			}
		}
		if alreadyIndexed {
			continue
		}

		// criteria: must contain METADATA.yml with a valid type field
		b, err := filesystem.ReadFile(path.Join(dirPath, subject.METADATA_FILE))
		if err != nil {
			continue
		}

		var newSubject subject.Subject
		if err := yaml.Unmarshal(b, &newSubject); err != nil {
			continue
		}

		if !validSubjectTypes[newSubject.Type] {
			continue
		}

		result.New = append(result.New, SyncEntry{Name: newSubject.Name, Type: string(newSubject.Type)})
		if confirm {
			// DirName intentionally not set — hydratePath recomputes it on next load
			m.Subjects = append(m.Subjects, newSubject)
		}
	}

	for i := range m.Modules {
		if err := discoverModule(&m.Modules[i], confirm, result); err != nil {
			return err
		}
	}

	return nil
}
