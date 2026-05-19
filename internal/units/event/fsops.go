package event

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Creates new event in 01_EVENTS dir and creates sub dirs for event
// create event dir, this should uses normalized event name based on e.EventDir() that returns
// abs path of event within the project.
// create subdirs. This also should used receiver functions (methods) to get abs path for each dir
func save(e Event, unit *UnitEvents) error {

	// create dir using event name
	dpth := e.EventDir(unit.UnitDir())
	if err := os.Mkdir(dpth, 0775); err != nil {
		return err
	}

	// create associated dirs
	for _, v := range e.SubDirs(unit.UnitDir()) {
		if err := os.Mkdir(v, 0775); err != nil {
			// we just skip the failed sub dir
			continue
		}
	}

	// Write event metadata file
	b, err := yaml.Marshal(e)
	if err != nil {
		return err
	}

	if err := os.WriteFile(e.MetadataDir(unit.UnitDir()), b, 0775); err != nil {
		return err
	}

	return nil
}
