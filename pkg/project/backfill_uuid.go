package project

import (
	"errors"
	"fmt"
	"io/fs"

	"github.com/hanymamdouh82/operatree/pkg/module"
)

// returns true if any subject was missing a UUID
func (p *Project) backfillUUIDs() (bool, error) {
	dirty := false

	for i := range p.Modules {
		changed, err := backfillModuleUUIDs(&p.Modules[i])
		if err != nil {
			return false, err
		}
		if changed {
			dirty = true
		}
	}
	return dirty, nil
}

func backfillModuleUUIDs(m *module.Module) (bool, error) {
	dirty := false
	valid := m.Subjects[:0] // same backing array, no allocation

	for i := range m.Subjects {
		onDisk, err := m.Subjects[i].ReadMetadata()
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				// orphaned subject — log and remove from project metadata
				fmt.Printf("warning: subject '%s' not found on disk, removing from project metadata\n", m.Subjects[i].Name)
				dirty = true
				continue // skip adding to valid slice
			}
			return false, err // real error, propagate
		}

		if onDisk.UUID == "" {
			if err := m.Subjects[i].SetID(); err != nil {
				return false, err
			}
			if err := m.Subjects[i].WriteMetadata(); err != nil {
				return false, err
			}
			dirty = true
		} else if m.Subjects[i].UUID != onDisk.UUID {
			m.Subjects[i].UUID = onDisk.UUID
			dirty = true
		}

		valid = append(valid, m.Subjects[i])
	}

	m.Subjects = valid // replace with cleaned slice

	for i := range m.Modules {
		changed, err := backfillModuleUUIDs(&m.Modules[i])
		if err != nil {
			return false, err
		}
		if changed {
			dirty = true
		}
	}
	return dirty, nil
}
