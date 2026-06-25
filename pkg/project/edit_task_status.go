package project

import (
	"log"

	"github.com/hanymamdouh82/operatree/pkg/subject"
)

func editTaskStatus(p *Project, uuid, term, newStatus string) error {

	var s subject.Subject
	var err error

	if uuid == "" {
		// find the required subject using interactive CLI for user to select the required one
		s, err = FindSubject(p, string(subject.SubjectTask), term)
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

	if err := s.EditTaskStatus(newStatus); err != nil {
		log.Fatal(err)
	}

	return nil
}
