package model

import (
	"go/parser"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

type Const struct {
	Name        string
	Value       string
	Comment     string
	Annotations []interface{}
	Spec        *SimpleSpec
}

func (m *Const) Validate() {
	if m.Name == "" {
		panic(errors.New("Variable 'Name' must be not empty"))
	}

	if !identRegexp.MatchString(m.Name) {
		panic(errors.Errorf("Variable 'Name' must be valid identifier, actual value: '%s'", m.Name))
	}

	if m.Spec != nil {
		if m.Spec.IsPointer {
			panic(errors.Errorf("Variable 'Spec.(%T).IsPointer' must be 'false' for %T", m.Spec, m))
		}

		m.Spec.Validate()
	}

	if m.Value != "" {
		if _, err := parser.ParseExpr(m.Value); err != nil {
			panic(err)
		}
	}
}

func (m *Const) String() string {
	if m.Value == "" {
		panic(errors.New("Variable 'Value' must be not empty"))
	}

	result := ""

	if m.Comment != "" {
		result += "// " + strings.Join(strings.Split(strings.TrimSpace(m.Comment), "\n"), "\n// ") + "\n"
	}

	result += "const " + m.Name

	if m.Spec != nil {
		result += " " + m.Spec.String()
	}

	return result + " = " + m.Value + "\n"
}

func (m *Const) Clone() interface{} {
	result := &Const{
		Name:        m.Name,
		Value:       m.Value,
		Comment:     m.Comment,
		Annotations: cloneAnnotations(m.Annotations),
	}

	if m.Spec != nil {
		result.Spec = m.Spec.Clone().(*SimpleSpec)
	}

	return result
}

func (m *Const) FetchImports(file *File) []*Import {
	return m.Spec.FetchImports(file)
}

func (m *Const) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	m.Spec.RenameImports(oldAlias, newAlias)

	if m.Value != "" {
		m.Value = regexp.
			MustCompile("([ \\t\\n&;,!~^=+\\-*/()\\[\\]{}])"+oldAlias+"([ \\t]*\\.)").
			ReplaceAllString(m.Value, "${1}"+newAlias+"${2}")
	}
}
