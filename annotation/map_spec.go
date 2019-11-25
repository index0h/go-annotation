package annotation

import "github.com/pkg/errors"

type MapSpec struct {
	Key   Spec
	Value Spec
}

func (m *MapSpec) Validate() {
	if m.Key == nil {
		panic(errors.New("Variable 'Key' must be not nil"))
	}

	if m.Value == nil {
		panic(errors.New("Variable 'Value' must be not nil"))
	}

	switch m.Key.(type) {
	case *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec:
		m.Key.Validate()
	default:
		panic(errors.Errorf("Variable 'Key' has invalid type: %T", m.Key))
	}

	switch m.Value.(type) {
	case *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec:
		m.Value.Validate()
	default:
		panic(errors.Errorf("Variable 'Value' has invalid type: %T", m.Value))
	}
}

func (m *MapSpec) String() string {
	result := "map["

	if _, ok := m.Key.(*FuncSpec); ok {
		result += "func "
	}

	result += m.Key.String() + "]"

	if _, ok := m.Value.(*FuncSpec); ok {
		result += "func "
	}

	return result + m.Value.String()
}

func (m *MapSpec) Clone() interface{} {
	return &MapSpec{
		Value: m.Value.Clone().(Spec),
		Key:   m.Key.Clone().(Spec),
	}
}

func (m *MapSpec) FetchImports(file *File) []*Import {
	result := []*Import{}
	result = append(result, m.Key.FetchImports(file)...)
	result = append(result, m.Value.FetchImports(file)...)

	return uniqImports(result)
}

func (m *MapSpec) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	m.Key.RenameImports(oldAlias, newAlias)
	m.Value.RenameImports(oldAlias, newAlias)
}
