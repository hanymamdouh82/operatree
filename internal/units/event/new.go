package event

func (u *UnitEvents) New(name string) (Event, error) {
	e := Event{
		Type: "event",
		Name: name,
	}

	err := save(e, u)
	if err != nil {
		return e, err
	}

	return e, nil
}
