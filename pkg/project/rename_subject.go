package project

import (
	"fmt"
	"log"
	"os"
	"slices"

	"github.com/hanymamdouh82/operatree/internal/activitylog"
	"github.com/hanymamdouh82/operatree/pkg/module"
	"github.com/hanymamdouh82/operatree/pkg/subject"
)

func RenameSubject(p *Project, st, term, newName string, uuid string) error {

	var s subject.Subject
	var err error
	var oldName string

	if uuid == "" {
		// find the required subject using interactive CLI for user to select the required one
		s, err = FindSubject(p, st, term)
		if err != nil {
			return err
		}
	} else {
		// find the required subject using uuid -> for scripting by provided uuid flag
		sp, err := FindSubjectByID(p, uuid)
		if err != nil {
			return err
		}

		s = *sp
	}

	// save old name before mutating the object
	oldName = s.Name

	if s.Type == "" {
		return fmt.Errorf("couldn't identify subject type")
	}

	// rename subject. internally it updates subject METADATA also
	// we shoud grab the new name from Rename function, since new name can be provided from interactive prompt
	if newName, err = s.Rename(newName); err != nil {
		return err
	}

	// find subject within the project to update the project metadata
	ps, err := FindSubjectByID(p, s.UUID)
	if err != nil {
		return err
	}

	// deference edited subject with found in project
	*ps = s

	// update referenced subjects
	// we need to walk  the project, identify any subject that references the old name
	// once identified, we update the subject with new name, write its metadata to disk
	// then update the node in project
	if err := updateRerences(p, oldName, newName); err != nil {
		return err
	}

	// update project metadata for the subject
	if err := p.WriteMetadata(); err != nil {
		return err
	}

	if err := activitylog.Log(
		p.ProjectDir(),
		activitylog.ActionRename,
		string(s.Type),
		s.Name,
	); err != nil {
		fmt.Fprintf(os.Stderr, "Warning :could not write activity log :%v\n", err)
	}

	return nil
}

// Walks the project, identify any subject that references the old name
// once identified, we update the subject with new name, write its metadata to disk
// then update the node in project
func updateRerences(p *Project, oldName, newName string) error {

	uuidsToUpdate, err := identifyUUIDsToBeUpdated(p.Modules, oldName)
	if err != nil {
		return err
	}

	// we get each subject using UUID
	// we update oldName with newName
	// we write subject metadata to disk
	// we update project subject
	// we don't write updated project metadata to disk, caller is responsible to write

	for _, uuid := range uuidsToUpdate {
		sp, err := FindSubjectByID(p, uuid)
		if err != nil {
			// skip and log error only
			log.Printf("failed to find subject of UUID: %s\n", uuid)
			continue
		}

		// 1. update metadata object for related events slice
		idx := slices.Index(sp.RelatedEvents, oldName)
		if idx != -1 {
			sp.RelatedEvents[idx] = newName
		}

		// 2. update metadata object for related objective
		if sp.RelatedObjective == oldName {
			sp.RelatedObjective = newName
		}

		// 3. write metadata to disk
		if err := sp.WriteMetadata(); err != nil {
			// skip and log error only
			log.Printf("failed to write subject metadata disk of UUID: %s\n", uuid)
			continue
		}
	}

	return nil
}

func identifyUUIDsToBeUpdated(ms []module.Module, oldName string) ([]string, error) {

	var uuidsToBeUpdated []string

	for _, m := range ms {
		// we walk subjects in current module
		for _, ps := range m.Subjects {
			// check related events
			for _, re := range ps.RelatedEvents {
				if re == oldName {
					// we found match that needs to be updated
					uuidsToBeUpdated = append(uuidsToBeUpdated, ps.UUID)
				}
			}

			// check related objective
			if ps.RelatedObjective == oldName {
				// we found match that needs to be updated
				uuidsToBeUpdated = append(uuidsToBeUpdated, ps.UUID)
			}
		}

		// we then run recursion for nested modules
		nestedUUIDs, err := identifyUUIDsToBeUpdated(m.Modules, oldName)
		if err != nil {
			return nestedUUIDs, err
		}
		uuidsToBeUpdated = append(uuidsToBeUpdated, nestedUUIDs...)

	}
	return uuidsToBeUpdated, nil
}
