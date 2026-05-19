package event

import (
	"fmt"
	"log"
	"time"

	"github.com/hanymamdouh82/operatree/internal/metadata"
	"github.com/manifoldco/promptui"
)

func (u *UnitEvents) NewInteractive() (Event, error) {
	e := Event{
		Type: "event",
	}

	// Name prompt
	prompt := promptui.Prompt{
		Label: "Event name",
	}
	name, err := prompt.Run()
	if err != nil {
		return e, err
	}

	// Date prompt
	prompt = promptui.Prompt{
		Label:   "Date",
		Default: time.Now().Format("2006-01-02"),
	}
	date, err := prompt.Run()
	if err != nil {
		return e, err
	}

	// Location prompt
	prompt = promptui.Prompt{
		Label: "Location",
	}
	location, err := prompt.Run()
	if err != nil {
		return e, err
	}

	// Participants prompt
	prompt = promptui.Prompt{
		Label: "Participants (comma-separated)",
	}
	participants, err := prompt.Run()
	if err != nil {
		return e, err
	}

	// Tags prompt
	prompt = promptui.Prompt{
		Label: "Tags (comma-separated)",
	}
	tags, err := prompt.Run()
	if err != nil {
		return e, err
	}

	// Notes prompt
	prompt = promptui.Prompt{
		Label: "Notes",
	}
	notes, err := prompt.Run()
	if err != nil {
		return e, err
	}

	// Set prompt values to struct before calling save
	e.Name = fmt.Sprintf("%s-%s", date, metadata.FormatName(name))
	e.Date = date
	e.Location = location
	e.Participants = metadata.ParseParticipants(participants)
	e.Tags = metadata.ParseTags(tags)
	e.Notes = notes

	if err := save(e, u); err != nil {
		log.Fatal(err)
		return e, err
	}

	return e, nil
}
