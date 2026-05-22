package project

import (
	"fmt"
	"os"
	"slices"

	"github.com/hanymamdouh82/operatree/internal/activitylog"
	"github.com/hanymamdouh82/operatree/internal/module"
	"github.com/hanymamdouh82/operatree/internal/subject"
)

// Adds new event to Events module
func NewEvent(p *Project, subjectName, subjectDate string) error {
	ss := ListSubjects(p, "")

	i := slices.IndexFunc(p.Modules, func(m module.Module) bool {
		return m.Type == module.ModuleEvents
	})
	m := p.Modules[i]

	// Build initial subject that captures passed flags
	is := subject.Subject{
		Type: subject.SubjectEvent,
		Name: subjectName,
		Date: subjectDate,
	}

	// module abs path defines where subject will reside
	s, err := subject.SubjectFactory(is, m.AbsPath, ss)
	if err != nil {
		return err
	}

	if err := s.WriteToDisk(); err != nil {
		return err
	}

	// update project metadata and write to disk
	p.Modules[i].Subjects = append(m.Subjects, s)
	if err := p.WriteMetadata(); err != nil {
		return err
	}

	fmt.Printf("Event created: %s\n", s.Name)

	if err := activitylog.Log(
		p.ProjectDir(),
		activitylog.ActionCreate,
		string(subject.SubjectEvent),
		s.Name,
	); err != nil {
		// non-fatal — log failure should never block subject creation
		fmt.Fprintf(os.Stderr, "warning: could not write activity log: %v\n", err)
	}

	return nil
}

// Adds new event to Project Management / Tasks module
func NewTask(p *Project, subjectName, subjectDate string) error {
	ss := ListSubjects(p, "")

	i := slices.IndexFunc(p.Modules, func(m module.Module) bool {
		return m.Type == module.ModuleProjectManagement
	})
	pmm := p.Modules[i]

	j := slices.IndexFunc(pmm.Modules, func(m module.Module) bool {
		return m.Type == module.ModuleTasks
	})

	m := pmm.Modules[j]

	// Build initial subject that captures passed flags
	is := subject.Subject{
		Type: subject.SubjectTask,
		Name: subjectName,
		Date: subjectDate,
	}

	// module abs path defines where subject will reside
	s, err := subject.SubjectFactory(is, m.AbsPath, ss)
	if err != nil {
		return err
	}

	if err := s.WriteToDisk(); err != nil {
		return err
	}

	// update project metadata and write to disk
	p.Modules[i].Modules[j].Subjects = append(m.Subjects, s)
	if err := p.WriteMetadata(); err != nil {
		return err
	}

	fmt.Printf("Task created: %s\n", s.Name)

	if err := activitylog.Log(
		p.ProjectDir(),
		activitylog.ActionCreate,
		string(subject.SubjectEvent),
		s.Name,
	); err != nil {
		// non-fatal — log failure should never block subject creation
		fmt.Fprintf(os.Stderr, "warning: could not write activity log: %v\n", err)
	}

	return nil
}

// Adds new topic to Research / Topics module
func NewTopic(p *Project, subjectName, subjectDate string) error {
	ss := ListSubjects(p, "")

	i := slices.IndexFunc(p.Modules, func(m module.Module) bool {
		return m.Type == module.ModuleResearch
	})
	pmm := p.Modules[i]

	j := slices.IndexFunc(pmm.Modules, func(m module.Module) bool {
		return m.Type == module.ModuleTopics
	})

	m := pmm.Modules[j]

	// Build initial subject that captures passed flags
	is := subject.Subject{
		Type: subject.SubjectTopic,
		Name: subjectName,
		Date: subjectDate,
	}

	// module abs path defines where subject will reside
	s, err := subject.SubjectFactory(is, m.AbsPath, ss)
	if err != nil {
		return err
	}

	if err := s.WriteToDisk(); err != nil {
		return err
	}

	// update project metadata and write to disk
	p.Modules[i].Modules[j].Subjects = append(m.Subjects, s)
	if err := p.WriteMetadata(); err != nil {
		return err
	}

	fmt.Printf("Topic created: %s\n", s.Name)

	if err := activitylog.Log(
		p.ProjectDir(),
		activitylog.ActionCreate,
		string(subject.SubjectEvent),
		s.Name,
	); err != nil {
		// non-fatal — log failure should never block subject creation
		fmt.Fprintf(os.Stderr, "warning: could not write activity log: %v\n", err)
	}

	return nil
}

// Adds new objective to Research / Objectives module
func NewObjective(p *Project, subjectName, subjectDate string) error {
	ss := ListSubjects(p, "")

	i := slices.IndexFunc(p.Modules, func(m module.Module) bool {
		return m.Type == module.ModuleResearch
	})
	pmm := p.Modules[i]

	j := slices.IndexFunc(pmm.Modules, func(m module.Module) bool {
		return m.Type == module.ModuleObjectives
	})

	m := pmm.Modules[j]

	// Build initial subject that captures passed flags
	is := subject.Subject{
		Type: subject.SubjectObjective,
		Name: subjectName,
		Date: subjectDate,
	}

	// module abs path defines where subject will reside
	s, err := subject.SubjectFactory(is, m.AbsPath, ss)
	if err != nil {
		return err
	}

	if err := s.WriteToDisk(); err != nil {
		return err
	}

	// update project metadata and write to disk
	p.Modules[i].Modules[j].Subjects = append(m.Subjects, s)
	if err := p.WriteMetadata(); err != nil {
		return err
	}

	fmt.Printf("Objective created: %s\n", s.Name)

	if err := activitylog.Log(
		p.ProjectDir(),
		activitylog.ActionCreate,
		string(subject.SubjectEvent),
		s.Name,
	); err != nil {
		// non-fatal — log failure should never block subject creation
		fmt.Fprintf(os.Stderr, "warning: could not write activity log: %v\n", err)
	}

	return nil
}
