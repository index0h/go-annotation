package model

import (
	"strings"

	"github.com/pkg/errors"
)

type TypeGroup struct {
	Comment     string
	Annotations []interface{}
	Types       []*Type
}

func (m *TypeGroup) Validate() {
	for i, element := range m.Types {
		if element == nil {
			panic(errors.Errorf("Variable 'Types[%d]' must be not nil", i))
		}

		element.Validate()
	}
}

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

func (m *TypeGroup) FetchImports(file *File) []*Import {
	result := []*Import{}

	for _, field := range m.Types {
		result = append(result, field.FetchImports(file)...)
	}

	return uniqImports(result)
}

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
