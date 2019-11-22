package model

import (
	"strings"

	"github.com/pkg/errors"
)

type InterfaceSpec struct {
	Methods []*Field
}

func (m *InterfaceSpec) Validate() {
	for i, method := range m.Methods {
		if method == nil {
			panic(errors.Errorf("Variable 'Methods[%d]' must be not nil", i))
		}

		switch method.Spec.(type) {
		case *SimpleSpec:
			if method.Name != "" {
				panic(
					errors.Errorf(
						"Variable 'Methods[%d].Name' must be empty for 'Methods[%d].Spec' type *SimpleSpec",
						i,
						i,
					),
				)
			}

			if method.Spec.(*SimpleSpec).IsPointer {
				panic(errors.Errorf("Variable 'Methods[%d].Spec.(%T).IsPointer' must be 'false'", i, method.Spec))
			}
		case *FuncSpec:
			if method.Name == "" {
				panic(
					errors.Errorf(
						"Variable 'Methods[%d].Name' must be not empty for 'Methods[%d].Spec' type *FuncSpec",
						i,
						i,
					),
				)
			}
		default:
			panic(errors.Errorf("Variable 'Methods[%d]' has invalid type %T", i, method.Spec))
		}

		method.Validate()
	}
}

func (m *InterfaceSpec) String() string {
	if len(m.Methods) == 0 {
		return "interface{}"
	}

	result := "interface{\n"

	for _, method := range m.Methods {
		if method.Comment != "" {
			result += "// " + strings.Join(strings.Split(strings.TrimSpace(method.Comment), "\n"), "\n// ") + "\n"
		}

		result += method.Name + method.Spec.String() + "\n"
	}

	return result + "}"
}

func (m *InterfaceSpec) Clone() interface{} {
	if m.Methods == nil {
		return &InterfaceSpec{}
	}

	result := &InterfaceSpec{}

	if m.Methods != nil {
		result.Methods = make([]*Field, len(m.Methods))
	}

	for i, method := range m.Methods {
		result.Methods[i] = method.Clone().(*Field)
	}

	return result
}

func (m *InterfaceSpec) FetchImports(file *File) []*Import {
	result := []*Import{}

	for _, method := range m.Methods {
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

	for _, method := range m.Methods {
		method.RenameImports(oldAlias, newAlias)
	}
}
