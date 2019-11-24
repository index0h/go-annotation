package model

import (
	"strings"

	"github.com/pkg/errors"
)

type InterfaceSpec struct {
	Fields []*Field
}

func (m *InterfaceSpec) Validate() {
	for i, method := range m.Fields {
		if method == nil {
			panic(errors.Errorf("Variable 'Fields[%d]' must be not nil", i))
		}

		switch method.Spec.(type) {
		case *SimpleSpec:
			if method.Name != "" {
				panic(
					errors.Errorf(
						"Variable 'Fields[%d].Name' must be empty for 'Fields[%d].Spec' type *SimpleSpec",
						i,
						i,
					),
				)
			}

			if method.Spec.(*SimpleSpec).IsPointer {
				panic(errors.Errorf("Variable 'Fields[%d].Spec.(%T).IsPointer' must be 'false'", i, method.Spec))
			}
		case *FuncSpec:
			if method.Name == "" {
				panic(
					errors.Errorf(
						"Variable 'Fields[%d].Name' must be not empty for 'Fields[%d].Spec' type *FuncSpec",
						i,
						i,
					),
				)
			}
		default:
			panic(errors.Errorf("Variable 'Fields[%d]' has invalid type %T", i, method.Spec))
		}

		method.Validate()
	}
}

func (m *InterfaceSpec) String() string {
	if len(m.Fields) == 0 {
		return "interface{}"
	}

	result := "interface{\n"

	for _, method := range m.Fields {
		if method.Comment != "" {
			result += "// " + strings.Join(strings.Split(strings.TrimSpace(method.Comment), "\n"), "\n// ") + "\n"
		}

		result += method.Name + method.Spec.String() + "\n"
	}

	return result + "}"
}

func (m *InterfaceSpec) Clone() interface{} {
	if m.Fields == nil {
		return &InterfaceSpec{}
	}

	result := &InterfaceSpec{}

	if m.Fields != nil {
		result.Fields = make([]*Field, len(m.Fields))
	}

	for i, method := range m.Fields {
		result.Fields[i] = method.Clone().(*Field)
	}

	return result
}

func (m *InterfaceSpec) FetchImports(file *File) []*Import {
	result := []*Import{}

	for _, method := range m.Fields {
		result = append(result, method.FetchImports(file)...)
	}

	return uniqImports(result)
}

func (m *InterfaceSpec) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	for _, method := range m.Fields {
		method.RenameImports(oldAlias, newAlias)
	}
}
