package annotation

import "github.com/pkg/errors"

// SimpleSpec model represents an identifier of type.
type SimpleSpec struct {
	PackageName string
	TypeName    string
	IsPointer   bool
}

// Validates SimpleSpec model fields.
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

// Renders SimpleSpec model to string.
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

// Creates copy of SimpleSpec model.
func (m *SimpleSpec) Clone() interface{} {
	return &SimpleSpec{
		PackageName: m.PackageName,
		TypeName:    m.TypeName,
		IsPointer:   m.IsPointer,
	}
}

// Checks that value is equal to SimpleSpec model.
func (m *SimpleSpec) EqualSpec(value interface{}) bool {
	model, ok := value.(*SimpleSpec)

	return ok &&
		model.PackageName == m.PackageName &&
		model.TypeName == m.TypeName &&
		model.IsPointer == m.IsPointer
}

// Fetches list of Import models registered in file argument, which are used by PackageName field.
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

// Renames PackageName field if its same as oldAlias argument.
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
