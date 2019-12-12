package annotation

import (
	"strings"

	"github.com/pkg/errors"
)

// VarGroup represents list of var declaration.
type VarGroup struct {
	Comment     string
	Annotations []interface{}
	Vars        []*Var
}

// Validates VarGroup model fields.
func (m *VarGroup) Validate() {
	for i, element := range m.Vars {
		if element == nil {
			panic(errors.Errorf("Variable 'Vars[%d]' must be not nil", i))
		}

		element.Validate()
	}
}

// Renders VarGroup model to string.
// If VarGroup contain only one var, without own comment, it will be rendered as single var, without braces.
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

// Creates deep copy of VarGroup model.
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

// Checks that value is deeply equal to VarGroup model.
// Ignores Comment and Annotations.
func (m *VarGroup) EqualSpec(value interface{}) bool {
	model, ok := value.(*VarGroup)

	if !ok || len(m.Vars) != len(model.Vars) {
		return false
	}

	checkedModelElements := make([]bool, len(m.Vars))

	for _, element := range m.Vars {
		elementEqual := false

		for j, modelElement := range model.Vars {
			if checkedModelElements[j] {
				continue
			}

			if element.EqualSpec(modelElement) {
				elementEqual = true
				checkedModelElements[j] = true

				break
			}
		}

		if !elementEqual {
			return false
		}
	}

	return true
}

// Checks that VarGroup contain deeply equal value.
// Ignores Comment and Annotations.
func (m *VarGroup) ContainsSpec(value *Var) bool {
	for _, element := range m.Vars {
		if element.EqualSpec(value) {
			return true
		}
	}

	return false
}

// Fetches list of Import models registered in file argument, which are used by Vars field.
func (m *VarGroup) FetchImports(file *File) []*Import {
	result := []*Import{}

	for _, field := range m.Vars {
		result = append(result, field.FetchImports(file)...)
	}

	return uniqImports(result)
}

// Renames import aliases, which are used by Vars field.
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
