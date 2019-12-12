package annotation

import (
	"go/parser"
	"strings"

	"github.com/pkg/errors"
)

// Func represents declaration of function.
type Func struct {
	Name        string
	Content     string
	Comment     string
	Annotations []interface{}
	Spec        *FuncSpec
	Related     *Field
}

// Validates Func model fields.
func (m *Func) Validate() {
	if m.Name == "" {
		panic(errors.New("Variable 'Name' must be not empty"))
	}

	if !identRegexp.MatchString(m.Name) {
		panic(errors.Errorf("Variable 'Name' must be valid identifier, actual value: '%s'", m.Name))
	}

	if m.Spec != nil {
		m.Spec.Validate()
	}

	if m.Related != nil {
		m.Related.Validate()

		if related, ok := m.Related.Spec.(*SimpleSpec); !ok {
			panic(errors.Errorf("Variable 'Related.Spec.(%T)' has invalid type for %T", m.Related.Spec, m))
		} else if related.PackageName != "" {
			panic(errors.Errorf("Variable 'Related.Spec.(%T).PackageName' must be empty for %T", related, m))
		}
	}

	if m.Content != "" {
		content := "func(){\n" + m.Content + "\n}"

		if _, err := parser.ParseExpr(content); err != nil {
			panic(err)
		}
	}
}

// Renders Func model to string.
func (m *Func) String() string {
	result := ""

	if m.Comment != "" {
		result += "// " + strings.Join(strings.Split(strings.TrimSpace(m.Comment), "\n"), "\n// ") + "\n"
	}

	result += "func "

	if m.Related != nil {
		result += "("

		if m.Related.Comment != "" {
			result += "\n// " + strings.Join(strings.Split(strings.TrimSpace(m.Related.Comment), "\n"), "\n// ") + "\n"
		}

		result += m.Related.Name + " " + m.Related.Spec.String() + ") "
	}

	result += m.Name

	if m.Spec == nil {
		result += "()"
	} else {
		result += m.Spec.String()
	}

	return result + " {\n" + m.Content + "\n}\n"
}

// Creates deep copy of Func model.
func (m *Func) Clone() interface{} {
	result := &Func{
		Name:        m.Name,
		Content:     m.Content,
		Comment:     m.Comment,
		Annotations: cloneAnnotations(m.Annotations),
	}

	if m.Spec != nil {
		result.Spec = m.Spec.Clone().(*FuncSpec)
	}

	if m.Related != nil {
		result.Related = m.Related.Clone().(*Field)
	}

	return result
}

// Checks that value is deeply equal to Func model.
// Ignores Comment and Annotations.
func (m *Func) EqualSpec(value interface{}) bool {
	model, ok := value.(*Func)

	if !ok ||
		model.Name != m.Name ||
		model.Content != m.Content ||
		((m.Spec == nil) != (model.Spec == nil)) ||
		((m.Related == nil) != (model.Related == nil)) {
		return false
	}

	if m.Spec != nil && !m.Spec.EqualSpec(model.Spec) {
		return false
	}

	if m.Related != nil && !m.Related.EqualSpec(model.Related) {
		return false
	}

	return true
}

// Fetches list of Import models registered in file argument, which are used by Spec, Content and Related fields.
func (m *Func) FetchImports(file *File) []*Import {
	result := []*Import{}

	if m.Spec != nil {
		result = append(result, m.Spec.FetchImports(file)...)
	}

	if m.Related != nil {
		result = append(result, m.Related.FetchImports(file)...)
	}

	result = append(result, fetchImportsFromContent(m.Content, file)...)

	return uniqImports(result)
}

// Renames import aliases, which are used by Spec, Content and Related fields.
func (m *Func) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	if m.Spec != nil {
		m.Spec.RenameImports(oldAlias, newAlias)
	}

	if m.Related != nil {
		m.Related.RenameImports(oldAlias, newAlias)
	}

	m.Content = renameImportsInContent(m.Content, oldAlias, newAlias)
}
