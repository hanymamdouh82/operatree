package project

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Loads a project by reading project metadata file and sets project structrue
// Path represents project root path
func Load(pth string) (Project, error) {

	b, err := os.ReadFile(filepath.Join(pth, METADATA_FILE))
	if err != nil {
		return Project{}, err
	}

	// unmarshal into loader struct
	// then we can convert units as per build logic
	var p Project
	if err := yaml.Unmarshal(b, &p); err != nil {
		return Project{}, err
	}

	// hydrate all internal absolute paths based on root from config loaded on user system
	hydratePath(pth, &p)

	// rebuild UUIDs and other important migrations
	dirty, err := p.backfillUUIDs()
	if err != nil {
		return p, err
	}

	// only writes to disk if something changed
	if dirty {
		if err := p.WriteMetadata(); err != nil {
			return p, err
		}
	}

	return p, err
}
