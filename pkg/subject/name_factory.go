package subject

import (
	"fmt"

	"github.com/hanymamdouh82/operatree/internal/metadata"
)

// Name factory return name based on user input name after sanitizing it and joins with another properties
// based on type. For example, event name should follow pattern `yyyy-MM-dd-user-sanitized-name`
func nameFactory(s Subject) string {

	switch s.Type {
	case SubjectEvent:
		sn := metadata.FormatName(s.Name)
		if s.Date != "" {
			sn = fmt.Sprintf("%s-%s", s.Date, sn)
		}
		return sn
	case SubjectTask:
		sn := metadata.FormatName(s.Name)
		if s.Date != "" {
			sn = fmt.Sprintf("%s-%s", s.Date, sn)
		}
		return sn
	case SubjectTopic:
		sn := metadata.FormatName(s.Name)
		return sn
	case SubjectObjective:
		sn := metadata.FormatName(s.Name)
		if s.Date != "" {
			sn = fmt.Sprintf("%s-%s", s.Date, sn)
		}
		return sn
	case SubjectDataSource:
		sn := metadata.FormatName(s.Name)
		return sn
	default:
		return s.Name
	}
}
