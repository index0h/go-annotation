package annotation

import (
	"github.com/pkg/errors"
)

// ArraySpec represents specification of array or slice type.
type ArraySpec struct {
	// Allowed types: *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec.
	Value Spec
	// Expression with int result, which could be calculated at compilation time, or "...".
	Length string
}

// Validates ArraySpec model fields.
func (m *ArraySpec) Validate() {
	if m.Value == nil {
		panic(errors.New("Variable 'Value' must be not nil"))
	}

	switch m.Value.(type) {
	case *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec:
		m.Value.Validate()
	default:
		panic(errors.Errorf("Variable 'Value' has invalid type: %T", m.Value))
	}
}

// Renders ArraySpec model to string.
func (m *ArraySpec) String() string {
	result := ""

	if m.Length != "" {
		result = "[" + m.Length + "]"
	} else {
		result = "[]"
	}

	if _, ok := m.Value.(*FuncSpec); ok {
		result += "func "
	}

	return result + m.Value.String()
}

// Creates deep copy of ArraySpec model.
func (m *ArraySpec) Clone() interface{} {
	return &ArraySpec{
		Value:  m.Value.Clone().(Spec),
		Length: m.Length,
	}
}

// Checks that value is deeply equal to ArraySpec model.
func (m *ArraySpec) EqualSpec(value interface{}) bool {
	model, ok := value.(*ArraySpec)

	return ok && model.Length == m.Length && m.Value.EqualSpec(model.Value)
}

// Fetches list of Import models registered in file argument, which are used by Value and Length fields.
func (m *ArraySpec) FetchImports(file *File) []*Import {
	result := m.Value.FetchImports(file)
	result = append(result, fetchImportsFromContent(m.Length, file)...)

	return uniqImports(result)
}

// Renames import aliases, which are used by Value and Length fields.
func (m *ArraySpec) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	m.Value.RenameImports(oldAlias, newAlias)
	m.Length = renameImportsInContent(m.Length, oldAlias, newAlias)
}
