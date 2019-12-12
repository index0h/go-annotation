package annotation

import (
	"go/parser"
	"strings"

	"github.com/pkg/errors"
)

// Var represents var declaration.
type Var struct {
	Name        string
	Value       string
	Comment     string
	Annotations []interface{}
	// Allowed types: *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec.
	Spec Spec
}

// Validates Var model fields.
func (m *Var) Validate() {
	if m.Name == "" {
		panic(errors.New("Variable 'Name' must be not empty"))
	}

	if !identRegexp.MatchString(m.Name) {
		panic(errors.Errorf("Variable 'Name' must be valid identifier, actual value: '%s'", m.Name))
	}

	if m.Spec == nil && m.Value == "" {
		panic(errors.Errorf("%T must have not nil 'Spec' or not empty 'Value'", m))
	}

	if m.Spec != nil {
		switch m.Spec.(type) {
		case *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec:
			m.Spec.(Validator).Validate()
		default:
			panic(errors.Errorf("Variable 'Spec' has invalid type: %T", m.Spec))
		}
	}

	if m.Value != "" {
		if _, err := parser.ParseExpr(m.Value); err != nil {
			panic(err)
		}
	}
}

// Renders Var model to string.
func (m *Var) String() string {
	result := ""

	if m.Comment != "" {
		result += "// " + strings.Join(strings.Split(strings.TrimSpace(m.Comment), "\n"), "\n// ") + "\n"
	}

	result += "var " + m.Name

	if m.Spec != nil {
		result += " " + m.Spec.String()
	}

	if m.Value != "" {
		result += " = " + m.Value
	}

	return result + "\n"
}

// Creates deep copy of Var model.
func (m *Var) Clone() interface{} {
	return &Var{
		Name:        m.Name,
		Value:       m.Value,
		Comment:     m.Comment,
		Annotations: cloneAnnotations(m.Annotations),
		Spec:        m.Spec.Clone().(Spec),
	}
}

// Checks that value is deeply equal to Var model.
// Ignores Comment and Annotations.
func (m *Var) EqualSpec(value interface{}) bool {
	model, ok := value.(*Var)

	if !ok ||
		model.Name != m.Name ||
		model.Value != m.Value ||
		((m.Spec == nil) != (model.Spec == nil)) {
		return false
	}

	if m.Spec != nil && !m.Spec.EqualSpec(model.Spec) {
		return false
	}

	return true
}

// Fetches list of Import models registered in file argument, which are used by Spec and Value fields.
func (m *Var) FetchImports(file *File) []*Import {
	result := []*Import{}

	if m.Spec != nil {
		result = append(result, m.Spec.FetchImports(file)...)
	}

	result = append(result, fetchImportsFromContent(m.Value, file)...)

	return uniqImports(result)
}

// Renames import aliases, which are used by Spec and Value fields.
func (m *Var) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	if m.Spec != nil {
		m.Spec.RenameImports(oldAlias, newAlias)
	}

	m.Value = renameImportsInContent(m.Value, oldAlias, newAlias)
}
