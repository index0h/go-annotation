package annotation

import (
	"go/format"
	"go/parser"
	"go/token"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

type File struct {
	Name         string
	Content      string
	PackageName  string
	Comment      string
	Annotations  []interface{}
	ImportGroups []*ImportGroup
	ConstGroups  []*ConstGroup
	VarGroups    []*VarGroup
	TypeGroups   []*TypeGroup
	Funcs        []*Func
}

func (m *File) Validate() {
	if m.Name == "" {
		panic(errors.New("Variable 'Name' must be not empty"))
	}

	if m.PackageName == "" {
		panic(errors.New("Variable 'PackageName' must be not empty"))
	}

	if !identRegexp.MatchString(m.PackageName) {
		panic(errors.Errorf("Variable 'PackageName' must be valid identifier, actual value: '%s'", m.PackageName))
	}

	for i, element := range m.ImportGroups {
		if element == nil {
			panic(errors.Errorf("Variable 'ImportGroups[%d]' must be not nil", i))
		}

		element.Validate()
	}

	for i, element := range m.ConstGroups {
		if element == nil {
			panic(errors.Errorf("Variable 'ConstGroups[%d]' must be not nil", i))
		}

		element.Validate()
	}

	for i, element := range m.VarGroups {
		if element == nil {
			panic(errors.Errorf("Variable 'VarGroups[%d]' must be not nil", i))
		}

		element.Validate()
	}

	for i, element := range m.TypeGroups {
		if element == nil {
			panic(errors.Errorf("Variable 'TypeGroups[%d]' must be not nil", i))
		}

		element.Validate()
	}

	for i, element := range m.Funcs {
		if element == nil {
			panic(errors.Errorf("Variable 'Funcs[%d]' must be not nil", i))
		}

		element.Validate()
	}

	if m.Content != "" {
		if _, err := parser.ParseFile(token.NewFileSet(), "", m.Content, 0); err != nil {
			panic(err)
		}
	}
}

func (m *File) String() string {
	if m.Content != "" {
		return m.Content
	}

	result := Header

	if m.Comment != "" {
		result += "// " + strings.Join(strings.Split(strings.TrimSpace(m.Comment), "\n"), "\n// ") + "\n"
	}

	result += "package " + m.PackageName + "\n\n"

	for _, element := range m.ImportGroups {
		result += element.String() + "\n"
	}

	for _, element := range m.ConstGroups {
		result += element.String() + "\n"
	}

	for _, element := range m.VarGroups {
		result += element.String() + "\n"
	}

	for _, element := range m.TypeGroups {
		result += element.String() + "\n"
	}

	for _, element := range m.Funcs {
		result += element.String() + "\n"
	}

	formattedResult, err := format.Source([]byte(result))

	if err != nil {
		panic(err)
	}

	return string(formattedResult)
}

func (m *File) Clone() interface{} {
	result := &File{
		Name:        m.Name,
		Content:     m.Content,
		PackageName: m.PackageName,
		Comment:     m.Comment,
		Annotations: cloneAnnotations(m.Annotations),
	}

	if m.ImportGroups != nil {
		result.ImportGroups = make([]*ImportGroup, len(m.ImportGroups))
	}

	if m.ConstGroups != nil {
		result.ConstGroups = make([]*ConstGroup, len(m.ConstGroups))
	}

	if m.VarGroups != nil {
		result.VarGroups = make([]*VarGroup, len(m.VarGroups))
	}

	if m.TypeGroups != nil {
		result.TypeGroups = make([]*TypeGroup, len(m.TypeGroups))
	}

	if m.Funcs != nil {
		result.Funcs = make([]*Func, len(m.Funcs))
	}

	for i, element := range m.ImportGroups {
		result.ImportGroups[i] = element.Clone().(*ImportGroup)
	}

	for i, element := range m.ConstGroups {
		result.ConstGroups[i] = element.Clone().(*ConstGroup)
	}

	for i, element := range m.VarGroups {
		result.VarGroups[i] = element.Clone().(*VarGroup)
	}

	for i, element := range m.TypeGroups {
		result.TypeGroups[i] = element.Clone().(*TypeGroup)
	}

	for i, element := range m.Funcs {
		result.Funcs[i] = element.Clone().(*Func)
	}

	return result
}

func (m *File) RenameImports(oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	for _, element := range m.ImportGroups {
		element.RenameImports(oldAlias, newAlias)
	}

	for _, element := range m.ConstGroups {
		element.RenameImports(oldAlias, newAlias)
	}

	for _, element := range m.VarGroups {
		element.RenameImports(oldAlias, newAlias)
	}

	for _, element := range m.TypeGroups {
		element.RenameImports(oldAlias, newAlias)
	}

	for _, element := range m.Funcs {
		element.RenameImports(oldAlias, newAlias)
	}

	if m.Content != "" {
		m.Content = regexp.
			MustCompile("([ \\t\\n&;,!~^=+\\-*/()\\[\\]{}])"+oldAlias+"([ \\t]*\\.)").
			ReplaceAllString(m.Content, "${1}"+newAlias+"${2}")
	}
}
