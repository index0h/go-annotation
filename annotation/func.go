package annotation

import (
	"go/parser"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

type Func struct {
	Name        string
	Content     string
	Comment     string
	Annotations []interface{}
	Spec        *FuncSpec
	Related     *Field
}

func (m *Func) Validate() {
	if m.Name == "" {
		panic(errors.New("Variable 'Name' must be not empty"))
	}

	if !identRegexp.MatchString(m.Name) {
		panic(errors.Errorf("Variable 'Name' must be valid identifier, actual value: '%s'", m.Name))
	}

	if m.Spec == nil {
		panic(errors.New("Variable 'Spec' must be not nil"))
	}

	m.Spec.Validate()

	if m.Related != nil {
		m.Related.Validate()

		if related, ok := m.Related.Spec.(*SimpleSpec); !ok {
			panic(errors.Errorf("Variable 'Related.Spec.(%T)' has invalid type for %T", m.Related.Spec, m))
		} else if related.PackageName != "" {
			panic(errors.Errorf("Variable 'Related.Spec.(%T).PackageName' must be empty for %T", related, m))
		}
	}

	if m.Content != "" {
		content := "func(){\n" + m.Content + "\n}"

		if _, err := parser.ParseExpr(content); err != nil {
			panic(err)
		}
	}
}

func (m *Func) String() string {
	result := ""

	if m.Comment != "" {
		result += "// " + strings.Join(strings.Split(strings.TrimSpace(m.Comment), "\n"), "\n// ") + "\n"
	}

	result += "func "

	if m.Related != nil {
		result += "("

		if m.Related.Comment != "" {
			result += "\n// " + strings.Join(strings.Split(strings.TrimSpace(m.Related.Comment), "\n"), "\n// ") + "\n"
		}

		result += m.Related.Name + " " + m.Related.Spec.String() + ") "
	}

	return result + m.Name + m.Spec.String() + " {\n" + m.Content + "\n}\n"
}

func (m *Func) Clone() interface{} {
	result := &Func{
		Name:        m.Name,
		Content:     m.Content,
		Comment:     m.Comment,
		Annotations: cloneAnnotations(m.Annotations),
		Spec:        m.Spec.Clone().(*FuncSpec),
	}

	if m.Related != nil {
		result.Related = m.Related.Clone().(*Field)
	}

	return result
}

func (m *Func) FetchImports(file *File) []*Import {
	result := []*Import{}
	result = append(result, m.Spec.FetchImports(file)...)

	if m.Related != nil {
		result = append(result, m.Related.FetchImports(file)...)
	}

	return uniqImports(result)
}

func (m *Func) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	m.Spec.RenameImports(oldAlias, newAlias)

	if m.Related != nil {
		m.Related.RenameImports(oldAlias, newAlias)
	}

	if m.Content != "" {
		m.Content = regexp.
			MustCompile("([ \\t\\n&;,!~^=+\\-*/()\\[\\]{}])"+oldAlias+"([ \\t]*\\.)").
			ReplaceAllString(m.Content, "${1}"+newAlias+"${2}")
	}
}
