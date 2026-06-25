package subject

import "github.com/charmbracelet/huh"

func EditStatus(s *Subject, newStatus string) error {

	if newStatus == "" {
		// interactive
		var err error
		newStatus, err = newStatusInteractivePrompt()
		if err != nil {
			return err
		}
	}

	s.Status = newStatus

	return nil
}

func newStatusInteractivePrompt() (string, error) {
	var newStatus string

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("New Status").
				Options(
					huh.NewOption("Planned", "planned"),
					huh.NewOption("In Progress", "in-progress"),
					huh.NewOption("Postponed", "postponed"),
					huh.NewOption("Done", "done"),
				).
				Value(&newStatus),
		),
	).Run()
	if err != nil {
		return newStatus, err
	}

	return newStatus, nil
}
