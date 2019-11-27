package annotation

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

// Import represents import declaration.
type Import struct {
	Alias       string
	Namespace   string
	Comment     string
	Annotations []interface{}
}

// Returns Alias field if it's not empty, otherwise base path of Namespace field.
func (m *Import) RealAlias() string {
	if m.Alias != "" {
		return m.Alias
	}

	return filepath.Base(m.Namespace)
}

// Validates Import model fields.
func (m *Import) Validate() {
	if m.Alias != "" && !identRegexp.MatchString(m.Alias) {
		panic(errors.Errorf("Variable 'Alias' must be valid identifier, actual value: '%s'", m.Alias))
	}

	if m.Namespace == "" {
		panic(errors.New("Variable 'Namespace' must be not empty"))
	}
}

// Renders Import model to string.
func (m *Import) String() string {
	result := ""

	if m.Comment != "" {
		result += "// " + strings.Join(strings.Split(strings.TrimSpace(m.Comment), "\n"), "\n// ") + "\n"
	}

	result += "import "

	if m.Alias != "" {
		result += m.Alias + " "
	}

	return result + strconv.Quote(m.Namespace) + "\n"
}

// Creates deep copy of Import model.
func (m *Import) Clone() interface{} {
	return &Import{
		Alias:       m.Alias,
		Namespace:   m.Namespace,
		Comment:     m.Comment,
		Annotations: cloneAnnotations(m.Annotations),
	}
}

// Renames import Alias field if it is equal to oldAlias argument.
func (m *Import) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	if m.RealAlias() == oldAlias {
		m.Alias = newAlias
	}
}
