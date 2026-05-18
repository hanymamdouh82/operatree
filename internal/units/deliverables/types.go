package deliverables

import (
	"fmt"
	"path"

	"github.com/hanymamdouh82/operatree/internal/filesystem"
)

const (
	UNIT_NAME = "08_DELIVERABLES"
)

var (
	SubDirs = map[string]string{
		"client_documents": "client_documents description here",
		"presentations":    "presentations description here",
		"reports":          "reports description here",
		"sumbmissions":     "sumbmissions description here",
	}
)

type UnitDeliverables struct {
	Name       string `yaml:"name"`
	ParentPath string `yaml:"parentPath"`
}

func (u *UnitDeliverables) SetParentDir(pth string) {
	u.ParentPath = pth
}

func (u UnitDeliverables) UnitDir() string {
	return path.Join(u.ParentPath, UNIT_NAME)
}

// Cannot use *Event since it will not implement the interface
func (u UnitDeliverables) Bootstrap(ppth string) error {
	if err := filesystem.CreateDir(u.UnitDir()); err != nil {
		return err
	}

	// create sub dirs, ADMIN doesn't include units, instead it includes only flat sub dirs
	errs := make([]error, 0)
	for k, v := range SubDirs {
		if err := filesystem.CreateDir(path.Join(u.UnitDir(), k)); err != nil {
			errs = append(errs, err)
			continue
		}
		fmt.Printf("Created sub-dir: %s\n", v)
	}

	return nil
}
