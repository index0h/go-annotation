package annotation

import (
	"github.com/pkg/errors"
)

type ArraySpec struct {
	Value  Spec
	Length string
}

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

func (m *ArraySpec) Clone() interface{} {
	return &ArraySpec{
		Value:  m.Value.Clone().(Spec),
		Length: m.Length,
	}
}

func (m *ArraySpec) FetchImports(file *File) []*Import {
	return m.Value.FetchImports(file)
}

func (m *ArraySpec) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	m.Value.RenameImports(oldAlias, newAlias)
}
