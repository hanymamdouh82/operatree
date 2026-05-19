package event

import "path"

type Event struct {
	Type         string   `yaml:"type"`
	Name         string   `yaml:"name"`
	Date         string   `yaml:"time"`
	Location     string   `yaml:"location"`
	Participants []string `yaml:"participants"`
	Tags         []string `yaml:"tags"`
	Notes        string   `yaml:"notes"`
}

// returns event dir as abs path, `upth` is the unit abs path
func (e *Event) EventDir(upth string) string {
	return path.Join(upth, e.Name)
}

// returns sub-dirs of an event as an absolute path
// `upth` is unit path not event path
func (e *Event) SubDirs(upth string) []string {

	sds := []string{
		path.Join(e.EventDir(upth), "01_AGENDA"),
		path.Join(e.EventDir(upth), "02_MEDIA"),
		path.Join(e.EventDir(upth), "03_NOTES"),
		path.Join(e.EventDir(upth), "04_DOCUMENTS"),
		path.Join(e.EventDir(upth), "05_OUTCOMES"),
	}

	return sds
}

// returns abs path for event metadata file.
// `upath` is unit dir not event dir
func (e *Event) MetadataDir(upth string) string {
	return path.Join(e.EventDir(upth), "metadata.yml")
}
