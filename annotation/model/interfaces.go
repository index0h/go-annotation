package model

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

type AnnotationParser interface {
	Parse(string) []interface{}
}