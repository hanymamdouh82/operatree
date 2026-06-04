package subject

import (
	"path/filepath"
)

func SubjectFactory(s Subject, ppth string, pss []Subject) (Subject, error) {

	// requires interactive
	if s.Name == "" {
		ns, err := interactive(s.Type, ppth, pss)
		if err != nil {
			return s, err
		}

		ns.SetID()
		return ns, nil
	}

	// Regular path is silent
	ns, err := silent(s, ppth)
	if err != nil {
		return s, err
	}

	ns.SetID()
	return ns, nil
}

// creates new Subject of type event in silent mode add it to module subjects and updates metadata file.
// s is the subject, the caller is responsible to fill the least required fields .
// this function is commonly used with `--silent` flag.
// `ppth` is the parent dir abs path.
// `pss` all project subjects. Use in interactive CLI to attach related.
func silent(s Subject, ppth string) (Subject, error) {

	s.SubDirs = SubDirs[s.Type]
	s.Files = Files[s.Type]

	// reset name and path since path depends on name
	s.Name = nameFactory(s)
	s.DirName = filepath.Join(ppth, s.Name)

	return s, nil
}

// creates new Subject of type event, add it to module subjects and updates metadata file
// `ppth` is the parent dir abs path
// `pss` all project subjects. Use in interactive CLI to attach related
func interactive(st SubjectType, ppth string, pss []Subject) (Subject, error) {

	s := Subject{
		Type:    st,
		SubDirs: SubDirs[st],
		Files:   Files[st],
	}

	// call interactive to collect subject properties and fill subject object
	// CLI uses switch to decide propmpt fields
	// This should run before any other function since rest of function depends on
	// captured inputs from user
	if err := interactiveCLI(st, &s, pss); err != nil {
		return Subject{}, err
	}

	// reset name and path since path depends on name
	s.Name = nameFactory(s)
	s.DirName = filepath.Join(ppth, s.Name)

	return s, nil
}
