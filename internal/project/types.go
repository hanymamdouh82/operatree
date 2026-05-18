package project

import (
	"fmt"
	"path"

	"github.com/hanymamdouh82/operatree/internal/types"
	"github.com/hanymamdouh82/operatree/internal/units/event"
	"gopkg.in/yaml.v3"
)

const (
	METADATA_FILE = "metadata.yml"
)

type Unit interface {
	Bootstrap(pth string) error
	UnitDir() string
	SetParentDir(pth string)
	SetUnitName()
	SetUnitDir()
	SetUnitType(t types.UnitType)
	UnitType() types.UnitType
}

type Project struct {
	Name    string   `yaml:"name"`
	BaseDir string   `yaml:"baseDir"`
	Tags    []string `yaml:"tags"`
	Units   []Unit   `yaml:"units"`
}

func (p *Project) ProjectName() string {
	return p.Name
}

func (p *Project) MetadataFile() string {
	return path.Join(p.ProjectDir(), METADATA_FILE)
}

// Returns base dir of the project. It is the dir where project resides
func (p *Project) ProjectBaseDir() string {
	return p.BaseDir
}

// Returns full project path including project name.
// Ex: /mnt/repos/porjects/my_project
// never use baseDir property, always use reciever function whenever project path is required
func (p *Project) ProjectDir() string {
	return path.Join(p.BaseDir, p.Name)
}

// Prints project contents on stdout
// It is useful since users of CLI always needs a way to display all project details
// the displayed information should follow UNIX/Linux style so it can be piped to
// or chained with GNU tools such as `sed`, `cut`, `grep`, etc
func (p *Project) Describe() error {
	k, err := yaml.Marshal(p)
	if err != nil {
		return err
	}

	// To-Do: format output instead of plain yaml format
	fmt.Printf("%s\n", k)
	return nil
}

// Use to add unit to project, it is reponsible to set required project properties into unit
// t: unit type, it is very important since it defines unit types during loading
func (p *Project) AddUnit(u Unit, t types.UnitType) {
	u.SetUnitType(t)
	u.SetUnitName()
	u.SetParentDir(p.ProjectDir())
	u.SetUnitDir()
	p.Units = append(p.Units, u)
}

func (p *Project) UnitEvents() (*event.UnitEvents, error) {
	return GetUnit[*event.UnitEvents](p, types.UnitEvents)
}
