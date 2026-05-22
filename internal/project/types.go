package project

import (
	"fmt"
	"path"
	"slices"

	"github.com/hanymamdouh82/operatree/internal/filesystem"
	"github.com/hanymamdouh82/operatree/internal/module"
	"github.com/hanymamdouh82/operatree/internal/subject"
	"gopkg.in/yaml.v3"
)

const (
	METADATA_FILE = "METADATA.yml"
)

type Project struct {
	Name     string          `yaml:"name"`
	Template string          `yaml:"template"`
	BaseDir  string          `yaml:"baseDir"`
	Tags     []string        `yaml:"tags"`
	Modules  []module.Module `yaml:"modules"`
}

type tmpltMap map[string]func(name string, bpth string) Project

var tmplts tmpltMap = tmpltMap{
	"general": tmpltGeneral,
	"dev":     tmpltDev,
}

func (p *Project) ProjectName() string {
	return p.Name
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
func (p *Project) Describe(plain bool) error {
	if plain {
		y, err := yaml.Marshal(p)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", y)
		return nil
	}

	describeProject(p)
	return nil
}

func (p *Project) WriteMetadata() error {

	fn := path.Join(p.ProjectDir(), METADATA_FILE)
	if err := filesystem.StructToFile(p, fn); err != nil {
		return err
	}

	return nil
}

// Confirms module exists, if exists it returns the module.
func (p *Project) ModuleExists(name string) (module.Module, error) {

	midx := slices.IndexFunc(p.Modules, func(m module.Module) bool {
		return m.Name == name
	})

	if midx == -1 {
		return module.Module{}, fmt.Errorf("project doesn't contain module %s", name)
	}

	return p.Modules[midx], nil
}

// Archives a subject by moving to project Archive module
func (p *Project) Archive(s subject.Subject) error {

	if err := Archive(p, s); err != nil {
		return err
	}

	return nil
}
