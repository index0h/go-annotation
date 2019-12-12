package annotation

import "github.com/pkg/errors"

// Field represents field of struct, interface, func param, func result or func related field.
type Field struct {
	Name string
	// Used for render *StructSpec
	Tag         string
	Comment     string
	Annotations []interface{}
	// Allowed types: *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec.
	Spec Spec
}

// Validates Field model fields.
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

// Creates deep copy of Field model.
func (m *Field) Clone() interface{} {
	return &Field{
		Name:        m.Name,
		Tag:         m.Tag,
		Comment:     m.Comment,
		Annotations: cloneAnnotations(m.Annotations),
		Spec:        m.Spec.Clone().(Spec),
	}
}

// Checks that value is deeply equal to Field model.
// Ignores Comment and Annotations.
func (m *Field) EqualSpec(value interface{}) bool {
	model, ok := value.(*Field)

	return ok &&
		m.Name == model.Name &&
		m.Tag == model.Tag &&
		m.Spec.EqualSpec(model.Spec)
}

// Fetches list of Import models registered in file argument, which are used by Spec field.
func (m *Field) FetchImports(file *File) []*Import {
	return m.Spec.FetchImports(file)
}

// Renames import aliases, which are used by Spec field.
func (m *Field) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	m.Spec.RenameImports(oldAlias, newAlias)
}
