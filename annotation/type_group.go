package annotation

import (
	"strings"

	"github.com/pkg/errors"
)

// TypeGroup represents list of type declaration.
type TypeGroup struct {
	Comment     string
	Annotations []interface{}
	Types       []*Type
}

// Validates TypeGroup model fields.
func (m *TypeGroup) Validate() {
	for i, element := range m.Types {
		if element == nil {
			panic(errors.Errorf("Variable 'Types[%d]' must be not nil", i))
		}

		element.Validate()
	}
}

// Renders TypeGroup model to string.
// If TypeGroup contain only one type, without own comment, it will be rendered as single type, without braces.
func (m *TypeGroup) String() string {
	if len(m.Types) == 1 && m.Types[0].Comment == "" {
		result := ""

		if m.Comment != "" {
			result += "// " + strings.Join(strings.Split(strings.TrimSpace(m.Comment), "\n"), "\n// ") + "\n"
		}

		return result + m.Types[0].String()
	}

	result := ""

	if m.Comment != "" {
		result += "// " + strings.Join(strings.Split(strings.TrimSpace(m.Comment), "\n"), "\n// ") + "\n"
	}

	result += "type (\n"

	for _, element := range m.Types {
		if element.Comment != "" {
			result += "// " + strings.Join(strings.Split(strings.TrimSpace(element.Comment), "\n"), "\n// ") + "\n"
		}

		result += element.Name + " " + element.Spec.String() + "\n"
	}

	return result + ")\n"
}

// Creates deep copy of TypeGroup model.
func (m *TypeGroup) Clone() interface{} {
	result := &TypeGroup{
		Comment:     m.Comment,
		Annotations: cloneAnnotations(m.Annotations),
	}

	if m.Types != nil {
		result.Types = make([]*Type, len(m.Types))
	}

	for i, element := range m.Types {
		result.Types[i] = element.Clone().(*Type)
	}

	return result
}

// Fetches list of Import models registered in file argument, which are used by Types field.
func (m *TypeGroup) FetchImports(file *File) []*Import {
	result := []*Import{}

	for _, field := range m.Types {
		result = append(result, field.FetchImports(file)...)
	}

	return uniqImports(result)
}

// Renames import aliases, which are used by Types field.
func (m *TypeGroup) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	for _, element := range m.Types {
		element.RenameImports(oldAlias, newAlias)
	}
}
