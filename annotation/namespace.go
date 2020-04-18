package annotation

import (
	"path/filepath"

	"github.com/pkg/errors"
)

type Namespace struct {
	Name string
	Path string
	// Ignored namespace must have no files.
	// It could be useful to search declarations, imported with "." alias.
	IsIgnored bool
	Files     []*File
}

// Returns package name from first file, ok base of Name field.
func (m *Namespace) PackageName() string {
	if len(m.Files) > 0 {
		return m.Files[0].PackageName
	}

	return filepath.Base(m.Name)
}

// Returns file be its name.
func (m *Namespace) FindFileByName(name string) *File {
	if name == "" {
		panic(errors.New("Variable 'name' must be not empty"))
	}

	for _, element := range m.Files {
		if name == element.Name {
			return element
		}
	}

	return nil
}
