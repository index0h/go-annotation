package annotation

import (
	"go/parser"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

type Var struct {
	Name        string
	Value       string
	Comment     string
	Annotations []interface{}
	Spec        Spec
}

func (m *Var) Validate() {
	if m.Name == "" {
		panic(errors.New("Variable 'Name' must be not empty"))
	}

	if !identRegexp.MatchString(m.Name) {
		panic(errors.Errorf("Variable 'Name' must be valid identifier, actual value: '%s'", m.Name))
	}

	if m.Spec == nil && m.Value == "" {
		panic(errors.Errorf("%T must have not nil 'Spec' or not empty 'Value'", m))
	}

	if m.Spec != nil {
		switch m.Spec.(type) {
		case *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec:
			m.Spec.(Validator).Validate()
		default:
			panic(errors.Errorf("Variable 'Spec' has invalid type: %T", m.Spec))
		}
	}

	if m.Value != "" {
		if _, err := parser.ParseExpr(m.Value); err != nil {
			panic(err)
		}
	}
}

func (m *Var) String() string {
	result := ""

	if m.Comment != "" {
		result += "// " + strings.Join(strings.Split(strings.TrimSpace(m.Comment), "\n"), "\n// ") + "\n"
	}

	result += "var " + m.Name

	if m.Spec != nil {
		result += " " + m.Spec.String()
	}

	if m.Value != "" {
		result += " = " + m.Value
	}

	return result + "\n"
}

func (m *Var) Clone() interface{} {
	return &Var{
		Name:        m.Name,
		Value:       m.Value,
		Comment:     m.Comment,
		Annotations: cloneAnnotations(m.Annotations),
		Spec:        m.Spec.Clone().(Spec),
	}
}

func (m *Var) FetchImports(file *File) []*Import {
	return m.Spec.FetchImports(file)
}

func (m *Var) RenameImports(oldAlias string, newAlias string) {
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
