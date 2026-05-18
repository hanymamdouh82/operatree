package project

import "path"

type Unit interface {
	Bootstrap(pth string) error
	UnitDir() string
	SetParentDir(pth string)
}

type Project struct {
	name    string   `yaml:"name"`
	baseDir string   `yaml:"baseDir"`
	Tags    []string `yaml:"tags"`
	Units   []Unit   `yaml:"units"`
}

func (p *Project) ProjectName() string {
	return p.name
}

// Returns base dir of the project. It is the dir where project resides
func (p *Project) ProjectBaseDir() string {
	return p.baseDir
}

// Returns full project path including project name.
// Ex: /mnt/repos/porjects/my_project
// never use baseDir property, always use reciever function whenever project path is required
func (p *Project) ProjectDir() string {
	return path.Join(p.baseDir, p.name)
}

// Use to add unit to project, it is reponsible to set required project properties into unit
func (p *Project) AddUnit(u Unit) {

	u.SetParentDir(p.ProjectDir())
	p.Units = append(p.Units, u)
}
