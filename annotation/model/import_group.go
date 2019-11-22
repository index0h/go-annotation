package model

import (
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type ImportGroup struct {
	Comment     string
	Annotations []interface{}
	Imports     []*Import
}

func (m *ImportGroup) Validate() {
	for i, element := range m.Imports {
		if element == nil {
			panic(errors.Errorf("Variable 'Imports[%d]' must be not nil", i))
		}

		element.Validate()
	}
}

func (m *ImportGroup) String() string {
	if len(m.Imports) == 1 && m.Imports[0].Comment == "" {
		result := ""

		if m.Comment != "" {
			result += "// " + strings.Join(strings.Split(strings.TrimSpace(m.Comment), "\n"), "\n// ") + "\n"
		}

		return result + m.Imports[0].String()
	}

	result := ""

	if m.Comment != "" {
		result += "// " + strings.Join(strings.Split(strings.TrimSpace(m.Comment), "\n"), "\n// ") + "\n"
	}

	result += "import (\n"

	for _, element := range m.Imports {
		if element.Comment != "" {
			result += "// " + strings.Join(strings.Split(strings.TrimSpace(element.Comment), "\n"), "\n// ") + "\n"
		}

		result += element.Alias + " " + strconv.Quote(element.Namespace) + "\n"
	}

	return result + ")\n"
}

func (m *ImportGroup) Clone() interface{} {
	result := &ImportGroup{
		Comment:     m.Comment,
		Annotations: cloneAnnotations(m.Annotations),
	}

	if m.Imports != nil {
		result.Imports = make([]*Import, len(m.Imports))
	}

	for i, element := range m.Imports {
		result.Imports[i] = element.Clone().(*Import)
	}

	return result
}

func (m *ImportGroup) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	for _, element := range m.Imports {
		element.RenameImports(oldAlias, newAlias)
	}
}
