package model

import "github.com/pkg/errors"

type Field struct {
	Name        string
	Tag         string
	Comment     string
	Annotations []interface{}
	Spec        Spec
}

func (m *Field) Validate() {
	if m.Name != "" && !identRegexp.MatchString(m.Name) {
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

func (m *Field) Clone() interface{} {
	return &Field{
		Name:        m.Name,
		Tag:         m.Tag,
		Comment:     m.Comment,
		Annotations: cloneAnnotations(m.Annotations),
		Spec:        m.Spec.Clone().(Spec),
	}
}

func (m *Field) FetchImports(file *File) []*Import {
	return m.Spec.FetchImports(file)
}

func (m *Field) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	m.Spec.RenameImports(oldAlias, newAlias)
}
