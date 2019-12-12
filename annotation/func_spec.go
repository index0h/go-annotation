package annotation

import (
	"strings"

	"github.com/pkg/errors"
)

// FuncSpec represents specification of func.
type FuncSpec struct {
	Params     []*Field
	Results    []*Field
	IsVariadic bool
}

// Validates FuncSpec model fields.
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

		if _, ok := m.Params[len(m.Params)-1].Spec.(*ArraySpec); !ok {
			panic(errors.Errorf("Variable 'Params[%d].Spec' has invalid type for variadic %T", len(m.Params)-1, m))
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

// Renders FuncSpec model to string.
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

// Creates deep copy of FuncSpec model.
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

// Checks that value is deeply equal to FuncSpec model.
// Ignores Comment and Annotations.
func (m *FuncSpec) EqualSpec(value interface{}) bool {
	model, ok := value.(*FuncSpec)

	if !ok ||
		m.IsVariadic != model.IsVariadic ||
		len(m.Params) != len(model.Params) ||
		len(m.Results) != len(model.Results) {
		return false
	}

	for i, field := range m.Params {
		if !field.EqualSpec(model.Params[i]) {
			return false
		}
	}

	for i, field := range m.Results {
		if !field.EqualSpec(model.Results[i]) {
			return false
		}
	}

	return true
}

// Fetches list of Import models registered in file argument, which are used by Params and Results fields.
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

// Renames import aliases, which are used by Params and Results fields.
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
