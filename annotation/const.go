package annotation

import (
	"go/parser"
	"strings"

	"github.com/pkg/errors"
)

// Const represents const declaration.
type Const struct {
	Name string
	// Expression with, which could be calculated in compilation time, or empty.
	Value       string
	Comment     string
	Annotations []interface{}
	Spec        *SimpleSpec
}

// Validates Const model fields.
func (m *Const) Validate() {
	if m.Name == "" {
		panic(errors.New("Variable 'Name' must be not empty"))
	}

	if !identRegexp.MatchString(m.Name) {
		panic(errors.Errorf("Variable 'Name' must be valid identifier, actual value: '%s'", m.Name))
	}

	if m.Spec != nil {
		if m.Spec.IsPointer {
			panic(errors.Errorf("Variable 'Spec.(%T).IsPointer' must be 'false' for %T", m.Spec, m))
		}

		m.Spec.Validate()
	}

	if m.Value != "" {
		if _, err := parser.ParseExpr(m.Value); err != nil {
			panic(err)
		}
	}
}

// Renders Const model to string.
// Const inside ConstGroup could be without Value, but single const requires Value field.
func (m *Const) String() string {
	if m.Value == "" {
		panic(errors.New("Variable 'Value' must be not empty"))
	}

	result := ""

	if m.Comment != "" {
		result += "// " + strings.Join(strings.Split(strings.TrimSpace(m.Comment), "\n"), "\n// ") + "\n"
	}

	result += "const " + m.Name

	if m.Spec != nil {
		result += " " + m.Spec.String()
	}

	return result + " = " + m.Value + "\n"
}

// Creates deep copy of Const model.
func (m *Const) Clone() interface{} {
	result := &Const{
		Name:        m.Name,
		Value:       m.Value,
		Comment:     m.Comment,
		Annotations: cloneAnnotations(m.Annotations),
	}

	if m.Spec != nil {
		result.Spec = m.Spec.Clone().(*SimpleSpec)
	}

	return result
}

// Checks that value is deeply equal to Const model.
// Ignores Comment and Annotations.
func (m *Const) EqualSpec(value interface{}) bool {
	model, ok := value.(*Const)

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
func (m *Const) FetchImports(file *File) []*Import {
	result := []*Import{}

	if m.Spec != nil {
		result = append(result, m.Spec.FetchImports(file)...)
	}

	result = append(result, fetchImportsFromContent(m.Value, file)...)

	return uniqImports(result)
}

// Renames import aliases, which are used by Spec and Value fields.
func (m *Const) RenameImports(oldAlias string, newAlias string) {
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
