package model

import (
	"strings"

	"github.com/pkg/errors"
)

type VarGroup struct {
	Comment     string
	Annotations []interface{}
	Vars        []*Var
}

func (m *VarGroup) Validate() {
	for i, element := range m.Vars {
		if element == nil {
			panic(errors.Errorf("Variable 'Vars[%d]' must be not nil", i))
		}

		element.Validate()
	}
}

func (m *VarGroup) String() string {
	if len(m.Vars) == 1 && m.Vars[0].Comment == "" {
		result := ""

		if m.Comment != "" {
			result += "// " + strings.Join(strings.Split(strings.TrimSpace(m.Comment), "\n"), "\n// ") + "\n"
		}

		return result + m.Vars[0].String()
	}

	result := ""

	if m.Comment != "" {
		result += "// " + strings.Join(strings.Split(strings.TrimSpace(m.Comment), "\n"), "\n// ") + "\n"
	}

	result += "var (\n"

	for _, element := range m.Vars {
		if element.Comment != "" {
			result += "// " + strings.Join(strings.Split(strings.TrimSpace(element.Comment), "\n"), "\n// ") + "\n"
		}

		result += element.Name

		if element.Spec != nil {
			result += " " + element.Spec.(Stringer).String()
		}

		if element.Value != "" {
			result += " = " + element.Value
		}

		result += "\n"
	}

	return result + ")\n"
}

func (m *VarGroup) Clone() interface{} {
	result := &VarGroup{
		Comment:     m.Comment,
		Annotations: cloneAnnotations(m.Annotations),
	}

	if m.Vars != nil {
		result.Vars = make([]*Var, len(m.Vars))
	}

	for i, element := range m.Vars {
		result.Vars[i] = element.Clone().(*Var)
	}

	return result
}

func (m *VarGroup) FetchImports(file *File) []*Import {
	result := []*Import{}

	for _, field := range m.Vars {
		result = append(result, field.FetchImports(file)...)
	}

	return uniqImports(result)
}

func (m *VarGroup) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	for _, element := range m.Vars {
		element.RenameImports(oldAlias, newAlias)
	}
}
