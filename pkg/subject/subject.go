package subject

import (
	"fmt"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/hanymamdouh82/operatree/internal/filesystem"
	"github.com/hanymamdouh82/operatree/internal/runner"
)

// Generates a unique ID for subject using new Google's v7 UUID.
// We used v7 since it is sortable based on timestamp
func (s *Subject) SetID() error {
	// generate UUID
	id7, err := uuid.NewV7()
	if err != nil {
		return err
	}

	s.UUID = id7.String()
	return nil
}

// A method to create module directory
func (s *Subject) MkDir() error {

	if err := filesystem.CreateDir(s.DirName); err != nil {
		return err
	}

	return nil
}

// A method to create module sub directories
func (s *Subject) MkSubDirs() error {

	for _, v := range s.SubDirs {
		sdp := filepath.Join(s.DirName, v)

		if err := filesystem.CreateDir(sdp); err != nil {
			return err
		}
	}

	return nil
}

// Creates empty files on disk
func (s *Subject) WriteFiles() error {
	for _, f := range s.Files {
		sdp := filepath.Join(s.DirName, f)
		t := fmt.Sprintf("# %s\n", s.Name)
		if err := filesystem.TextToMDFile(t, sdp); err != nil {
			return err
		}
	}

	return nil
}

// Reads metadata.yml file for the subject from subject dir
func (s *Subject) ReadMetadata() (*Subject, error) {
	fn := filepath.Join(s.DirName, METADATA_FILE)

	var onDisk Subject
	if err := filesystem.FileToStruct(&onDisk, fn); err != nil {
		return nil, err
	}

	// restore DirName since it's excluded from YAML
	onDisk.DirName = s.DirName

	return &onDisk, nil
}

// Writes metadata.yml file for the subject at subject dir
func (s *Subject) WriteMetadata() error {

	fn := filepath.Join(s.DirName, METADATA_FILE)
	if err := filesystem.StructToFile(s, fn); err != nil {
		return err
	}

	return nil
}

// Writes subject to disk
func (s *Subject) WriteToDisk() error {

	// make dir
	if err := s.MkDir(); err != nil {
		return err
	}

	// make subdirs
	if err := s.MkSubDirs(); err != nil {
		return err
	}

	// Create emtpy subject files
	if err := s.WriteFiles(); err != nil {
		return err
	}

	// write metadata file
	if err := s.WriteMetadata(); err != nil {
		return err
	}

	return nil
}

// Prints subject on screen in a prettey style
// Not all properties are displayed, such as subDirs not included
// This function is intended to be used for subject briefing only
func (s *Subject) Describe() {
	describe(s)
}

func (s *Subject) EditMetadata() error {

	fn := filepath.Join(s.DirName, METADATA_FILE)
	if err := runner.EditFile(fn); err != nil {
		return err
	}

	return nil
}

func (s *Subject) EditTaskStatus(newStatus string) error {

	if err := EditStatus(s, newStatus); err != nil {
		return err
	}

	return s.WriteMetadata()
}

func (s *Subject) Rename(newName string) (string, error) {
	return renameSubject(s, newName)
}
