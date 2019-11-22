package model

import (
	"strconv"

	"github.com/pkg/errors"
)

type ArraySpec struct {
	Value      Spec
	Length     int
	IsEllipsis bool
}

func (m *ArraySpec) Validate() {
	if m.Value == nil {
		panic(errors.New("Variable 'Value' must be not nil"))
	}

	if m.Length < 0 {
		panic(errors.Errorf("Variable 'Length' must be greater than or equal to 0, actual value: %d", m.Length))
	}

	if m.Length > 0 && m.IsEllipsis {
		panic(errors.Errorf("Array must have only length or only ellipsis"))
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

	if m.IsEllipsis {
		result = "[...]"
	} else if m.Length != 0 {
		result = "[" + strconv.Itoa(m.Length) + "]"
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
		Value:      m.Value.Clone().(Spec),
		Length:     m.Length,
		IsEllipsis: m.IsEllipsis,
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
