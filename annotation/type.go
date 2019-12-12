package annotation

import (
	"strings"

	"github.com/pkg/errors"
)

// Type represents type declaration.
type Type struct {
	Name        string
	Comment     string
	Annotations []interface{}
	// Allowed types: *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec.
	Spec Spec
}

// Validates Type model fields.
func (m *Type) Validate() {
	if m.Name == "" {
		panic(errors.New("Variable 'Name' must be not empty"))
	}

	if !identRegexp.MatchString(m.Name) {
		panic(errors.Errorf("Variable 'Name' must be valid identifier, actual value: '%s'", m.Name))
	}

	if m.Spec == nil {
		panic(errors.New("Variable 'Spec' must be not nil"))
	}

	switch m.Spec.(type) {
	case *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec:
		m.Spec.Validate()
	default:
		panic(errors.Errorf("Variable 'Spec' has invalid type: %T", m.Spec))
	}
}

// Renders Type model to string.
func (m *Type) String() string {
	result := ""

	if m.Comment != "" {
		result += "// " + strings.Join(strings.Split(strings.TrimSpace(m.Comment), "\n"), "\n// ") + "\n"
	}

	return result + "type " + m.Name + " " + m.Spec.String() + "\n"
}

// Creates deep copy of Type model.
func (m *Type) Clone() interface{} {
	return &Type{
		Name:        m.Name,
		Comment:     m.Comment,
		Annotations: cloneAnnotations(m.Annotations),
		Spec:        m.Spec.Clone().(Spec),
	}
}

// Checks that value is deeply equal to Type model.
// Ignores Comment and Annotations.
func (m *Type) EqualSpec(value interface{}) bool {
	model, ok := value.(*Type)

	if !ok || model.Name != m.Name {
		return false
	}

	return m.Spec.EqualSpec(model.Spec)
}

// Fetches list of Import models registered in file argument, which are used by Spec field.
func (m *Type) FetchImports(file *File) []*Import {
	return m.Spec.FetchImports(file)
}

// Renames import aliases, which are used by Spec field.
func (m *Type) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	m.Spec.RenameImports(oldAlias, newAlias)
}
