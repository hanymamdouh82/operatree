package event

import (
	"os"
	"path"
	"time"

	"github.com/hanymamdouh82/operatree/internal/filesystem"
	"gopkg.in/yaml.v3"
)

const (
	UNIT_NAME = "01_EVENTS"
)

type UnitEvents struct {
	Name       string `yaml:"name"`
	ParentPath string `yaml:"parentPath"`
}

func (u *UnitEvents) SetParentDir(pth string) {
	u.ParentPath = pth
}

func (u UnitEvents) UnitDir() string {
	return path.Join(u.ParentPath, UNIT_NAME)
}

type Event struct {
	Type         string   `yaml:"type"`
	Name         string   `yaml:"name"`
	Date         string   `yaml:"time"`
	Location     string   `yaml:"location"`
	Participants []string `yaml:"participants"`
	Tags         []string `yaml:"tags"`
	Notes        string   `yaml:"notes"`
}

// Cannot use *Event since it will not implement the interface
func (e UnitEvents) Bootstrap(ppth string) error {
	if err := filesystem.CreateDir(e.UnitDir()); err != nil {
		return err
	}

	// create metadata template
	md := Event{
		Type:         "Event",
		Name:         "sample_event",
		Date:         time.Now().Format("2006-01-02"),
		Location:     "TBD",
		Participants: []string{},
		Tags:         []string{},
		Notes:        "",
	}

	b, err := yaml.Marshal(md)
	if err != nil {
		return err
	}

	fn := path.Join(e.UnitDir(), "sample_META.yml")
	if err := os.WriteFile(fn, b, 0775); err != nil {
		return err
	}

	return nil
}
