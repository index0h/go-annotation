package model

import (
	"strings"

	"github.com/pkg/errors"
)

type FuncSpec struct {
	Params     []*Field
	Results    []*Field
	IsVariadic bool
}

func (m *FuncSpec) Validate() {
	for i, param := range m.Params {
		if param == nil {
			panic(errors.Errorf("Variable 'Params[%d]' must be not nil", i))
		}

		param.Validate()
	}

	if m.IsVariadic {
		if len(m.Params) == 0 {
			panic(errors.Errorf("Variable 'Params' must be not empty for variadic %T", m))
		}

		lastParamArraySpec, ok := m.Params[len(m.Params)-1].Spec.(*ArraySpec)

		if !ok {
			panic(
				errors.Errorf(
					"Variable 'Params[%d].Spec' has invalid type for variadic %T",
					len(m.Params)-1,
					m,
				),
			)
		}

		if lastParamArraySpec.IsEllipsis {
			panic(
				errors.Errorf(
					"Variable 'Params[%d].Spec.(%T).IsEllipsis' must be 'false' for variadic %T",
					len(m.Params)-1,
					lastParamArraySpec,
					m,
				),
			)
		}

		if lastParamArraySpec.Length > 0 {
			panic(
				errors.Errorf(
					"Variable 'Params[%d].Spec.(%T).Length' must be '0' for variadic %T",
					len(m.Params)-1,
					lastParamArraySpec,
					m,
				),
			)
		}
	}

	hasName := false

	for i, result := range m.Results {
		if result == nil {
			panic(errors.Errorf("Variable 'Results[%d]' must be not nil", i))
		}

		if i == 0 {
			hasName = result.Name != ""
		} else if hasName != (result.Name != "") {
			panic(errors.New("Variable 'Results' must have all fields with names or all without names"))
		}

		result.Validate()
	}
}

func (m *FuncSpec) String() string {
	result := "("

	for i, field := range m.Params {
		if i > 0 {
			result += ", "
		}

		if field.Comment != "" {
			result += "\n// " + strings.Join(strings.Split(strings.TrimSpace(field.Comment), "\n"), "\n// ") + "\n"
		}

		if field.Name != "" {
			result += field.Name + " "
		}

		if _, ok := field.Spec.(*FuncSpec); ok {
			result += "func "
		}

		if m.IsVariadic && i == len(m.Params)-1 {
			result += "..." + field.Spec.(*ArraySpec).Value.String()
		} else {
			result += field.Spec.String()
		}
	}

	result += ")"

	if len(m.Results) == 0 {
		return result
	}

	result += " "

	if len(m.Results) > 0 || m.Results[0].Name != "" {
		result += "("
	}

	for i, field := range m.Results {
		if i > 0 {
			result += ", "
		}

		if field.Comment != "" {
			result += "\n// " + strings.Join(strings.Split(strings.TrimSpace(field.Comment), "\n"), "\n// ") + "\n"
		}

		if field.Name != "" {
			result += field.Name + " "
		}

		if _, ok := field.Spec.(*FuncSpec); ok {
			result += "func "
		}

		result += field.Spec.String()
	}

	if len(m.Results) > 0 || m.Results[0].Name != "" {
		result += ")"
	}

	return result
}

func (m *FuncSpec) Clone() interface{} {
	if m.Params == nil && m.Results == nil {
		return &FuncSpec{}
	}

	result := &FuncSpec{
		IsVariadic: m.IsVariadic,
	}

	if m.Params != nil {
		result.Params = make([]*Field, len(m.Params))
	}

	if m.Results != nil {
		result.Results = make([]*Field, len(m.Results))
	}

	for i, field := range m.Params {
		result.Params[i] = field.Clone().(*Field)
	}

	for i, field := range m.Results {
		result.Results[i] = field.Clone().(*Field)
	}

	return result
}

func (m *FuncSpec) FetchImports(file *File) []*Import {
	result := []*Import{}

	for _, field := range m.Params {
		result = append(result, field.FetchImports(file)...)
	}

	for _, field := range m.Results {
		result = append(result, field.FetchImports(file)...)
	}

	return uniqImports(result)
}

func (m *FuncSpec) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	for _, field := range m.Params {
		field.RenameImports(oldAlias, newAlias)
	}

	for _, field := range m.Results {
		field.RenameImports(oldAlias, newAlias)
	}
}
