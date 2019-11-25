package annotation

import "github.com/pkg/errors"

type SimpleSpec struct {
	PackageName string
	TypeName    string
	IsPointer   bool
}

func (m *SimpleSpec) Validate() {
	if m.TypeName == "" {
		panic(errors.New("Variable 'TypeName' must be not empty"))
	}

	if !identRegexp.MatchString(m.TypeName) {
		panic(errors.Errorf("Variable 'TypeName' must be valid identifier, actual value: '%s'", m.TypeName))
	}

	if m.PackageName != "" && !identRegexp.MatchString(m.PackageName) {
		panic(errors.Errorf("Variable 'PackageName' must be valid identifier, actual value: '%s'", m.PackageName))
	}
}

func (m *SimpleSpec) String() string {
	result := ""

	if m.IsPointer {
		result += "*"
	}

	if m.PackageName != "" {
		result += m.PackageName + "."
	}

	return result + m.TypeName
}

func (m *SimpleSpec) Clone() interface{} {
	return &SimpleSpec{
		PackageName: m.PackageName,
		TypeName:    m.TypeName,
		IsPointer:   m.IsPointer,
	}
}

func (m *SimpleSpec) FetchImports(file *File) []*Import {
	if m.PackageName == "" {
		return []*Import{}
	}

	for _, group := range file.ImportGroups {
		for _, element := range group.Imports {
			if element.RealAlias() == m.PackageName {
				return []*Import{element}
			}
		}
	}

	return []*Import{}
}

func (m *SimpleSpec) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	if m.PackageName == oldAlias {
		m.PackageName = newAlias
	}
}
