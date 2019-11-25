package model

import (
	"path/filepath"

	"github.com/pkg/errors"
)

type Namespace struct {
	Name      string
	Path      string
	IsIgnored bool
	Files     []*File
}

func (m *Namespace) Validate() {
	if m.Name == "" {
		panic(errors.New("Variable 'Name' must be not empty"))
	}

	if m.Path == "" {
		panic(errors.New("Variable 'Path' must be not empty"))
	}

	if !filepath.IsAbs(m.Path) {
		panic(errors.Errorf("Variable 'Path' must be absolute path, actual value: '%s'", m.Path))
	}

	if m.IsIgnored && len(m.Files) > 0 {
		panic(errors.Errorf("Ignored namespace with name: '%s' must have no files", m.Name))
	}

	fileNames := map[string]bool{}
	packageName := ""

	for i, element := range m.Files {
		if element == nil {
			panic(errors.Errorf("Variable 'Files[%d]' must be not nil", i))
		}

		element.Validate()

		if _, ok := fileNames[element.Name]; ok {
			panic(errors.Errorf("Namespace has duplicate file name: %s", element.Name))
		} else {
			fileNames[element.Name] = true
		}

		if i == 0 {
			packageName = element.PackageName
		} else if element.PackageName != packageName {
			panic(errors.New("Namespace has different packages"))
		}
	}
}

func (m *Namespace) Clone() interface{} {
	result := &Namespace{
		Name:      m.Name,
		Path:      m.Path,
		IsIgnored: m.IsIgnored,
	}

	if m.Files != nil {
		result.Files = make([]*File, len(m.Files))
	}

	for i, element := range m.Files {
		result.Files[i] = element.Clone().(*File)
	}

	return result
}

func (m *Namespace) PackageName() string {
	if len(m.Files) > 0 {
		return m.Files[0].PackageName
	}

	return filepath.Base(m.Name)
}

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
