package event

import "fmt"

// Creates new event in 01_EVENTS dir and creates sub dirs for event
func save(e Event, unit *UnitEvents) error {
	// create event dir
	// create subdirs

	pth := unit.UnitDir()
	fmt.Printf("Will create event: %s\n", e.Name)
	fmt.Printf("will create into: %s\n", pth)

	return nil
}
