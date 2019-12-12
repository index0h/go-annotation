package annotation

import (
	"strings"

	"github.com/pkg/errors"
)

// InterfaceSpec represents specification of an interface type.
type InterfaceSpec struct {
	Fields []*Field
}

// Validates InterfaceSpec model fields.
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

// Renders InterfaceSpec model to string.
func (m *InterfaceSpec) String() string {
	if len(m.Fields) == 0 {
		return "interface{}"
	}

	result := "interface{\n"

	for _, field := range m.Fields {
		if field.Comment != "" {
			result += "// " + strings.Join(strings.Split(strings.TrimSpace(field.Comment), "\n"), "\n// ") + "\n"
		}

		result += field.Name + field.Spec.String() + "\n"
	}

	return result + "}"
}

// Creates deep copy of InterfaceSpec model.
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

// Checks that value is deeply equal to InterfaceSpec model.
// Ignores Comment and Annotations.
func (m *InterfaceSpec) EqualSpec(value interface{}) bool {
	model, ok := value.(*InterfaceSpec)

	if !ok || len(m.Fields) != len(model.Fields) {
		return false
	}

	checkedModelFields := make([]bool, len(m.Fields))

	for _, field := range m.Fields {
		fieldEqual := false

		for j, modelField := range model.Fields {
			if checkedModelFields[j] {
				continue
			}

			if field.EqualSpec(modelField) {
				fieldEqual = true
				checkedModelFields[j] = true

				break
			}
		}

		if !fieldEqual {
			return false
		}
	}

	return true
}

// Fetches list of Import models registered in file argument, which are used by Fields field.
func (m *InterfaceSpec) FetchImports(file *File) []*Import {
	result := []*Import{}

	for _, method := range m.Fields {
		result = append(result, method.FetchImports(file)...)
	}

	return uniqImports(result)
}

// Renames import aliases, which are used by Fields field.
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
