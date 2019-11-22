package model

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type StructSpec struct {
	Fields []*Field
}

func (m *StructSpec) Validate() {
	for i, field := range m.Fields {
		if field == nil {
			panic(errors.Errorf("Variable 'Fields[%d]' must be not nil", i))
		}

		field.Validate()

		if _, ok := field.Spec.(*SimpleSpec); field.Name == "" && !ok {
			panic(errors.Errorf("Variable 'Fields[%d]' with empty 'Name' has invalid type: %T", i, field.Spec))
		}
	}
}

func (m *StructSpec) String() string {
	if len(m.Fields) == 0 {
		return "struct{}"
	}

	result := "struct{\n"

	for _, field := range m.Fields {
		if field.Comment != "" {
			result += "// " + strings.Join(strings.Split(strings.TrimSpace(field.Comment), "\n"), "\n// ") + "\n"
		}

		result += field.Name + " "

		if _, ok := field.Spec.(*FuncSpec); ok {
			result += "func "
		}

		result += field.Spec.String()

		if field.Tag != "" {
			result += " " + strconv.Quote(field.Tag)
		}

		result += "\n"
	}

	return result + "}"
}

func (m *StructSpec) Clone() interface{} {
	if m.Fields == nil {
		return &StructSpec{}
	}

	result := &StructSpec{}

	if m.Fields != nil {
		result.Fields = make([]*Field, len(m.Fields))
	}

	for i, field := range m.Fields {
		result.Fields[i] = field.Clone().(*Field)
	}

	return result
}

func (m *StructSpec) FetchImports(file *File) []*Import {
	result := []*Import{}

	for _, field := range m.Fields {
		result = append(result, field.FetchImports(file)...)
	}

	return uniqImports(result)
}

func (m *StructSpec) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	for _, field := range m.Fields {
		field.RenameImports(oldAlias, newAlias)
	}
}