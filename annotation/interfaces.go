package annotation

type Stringer interface {
	String() string
}

type Validator interface {
	Validate()
}

type Cloner interface {
	Clone() interface{}
}

type ImportsFetcher interface {
	FetchImports(*File) []*Import
}

type ImportsRenamer interface {
	RenameImports(oldAlias string, newAlias string)
}

type Spec interface {
	Validator
	Stringer
	Cloner
	ImportsFetcher
	ImportsRenamer
}

//noinspection GoNameStartsWithPackageName
type AnnotationParser interface {
	Parse(string) []interface{}
}

type SourceParser interface {
	Parse(fileName string, content string) *File
}
