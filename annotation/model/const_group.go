package model

import (
	"strings"

	"github.com/pkg/errors"
)

type ConstGroup struct {
	Comment     string
	Annotations []interface{}
	Consts      []*Const
}

func (m *ConstGroup) Validate() {
	for i, element := range m.Consts {
		if element == nil {
			panic(errors.Errorf("Variable 'Consts[%d]' must be not nil", i))
		}

		element.Validate()
	}
}

func (m *ConstGroup) String() string {
	if len(m.Consts) == 1 && m.Consts[0].Comment == "" {
		result := ""

		if m.Comment != "" {
			result += "// " + strings.Join(strings.Split(strings.TrimSpace(m.Comment), "\n"), "\n// ") + "\n"
		}

		return result + m.Consts[0].String()
	}

	result := ""

	if m.Comment != "" {
		result += "// " + strings.Join(strings.Split(strings.TrimSpace(m.Comment), "\n"), "\n// ") + "\n"
	}

	result += "const (\n"

	for _, element := range m.Consts {
		if element.Comment != "" {
			result += "// " + strings.Join(strings.Split(strings.TrimSpace(element.Comment), "\n"), "\n// ") + "\n"
		}

		result += element.Name

		if element.Spec != nil {
			result += " " + element.Spec.String()
		}

		if element.Value != "" {
			result += " = " + element.Value
		}

		result += "\n"
	}

	return result + ")\n"
}

func (m *ConstGroup) Clone() interface{} {
	result := &ConstGroup{
		Comment:     m.Comment,
		Annotations: cloneAnnotations(m.Annotations),
	}

	if m.Consts != nil {
		result.Consts = make([]*Const, len(m.Consts))
	}

	for i, element := range m.Consts {
		result.Consts[i] = element.Clone().(*Const)
	}

	return result
}

func (m *ConstGroup) FetchImports(file *File) []*Import {
	result := []*Import{}

	for _, field := range m.Consts {
		result = append(result, field.FetchImports(file)...)
	}

	return uniqImports(result)
}

func (m *ConstGroup) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	for _, element := range m.Consts {
		element.RenameImports(oldAlias, newAlias)
	}
}
