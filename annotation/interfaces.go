package annotation

type Cloner interface {
	Clone(interface{}) interface{}
}

type Scanner interface {
	Scan(storage *Storage, rootNamespace string, rootPath string, ignores ...string)
}

type Renderer interface {
	Render(entity interface{}) string
}

type Validator interface {
	Validate(entity interface{}) error
}

type AnnotationParser interface {
	SetAnnotation(name string, annotationType interface{})
	Parse(content string) (annotations []interface{})
}

type SourceParser interface {
	Parse(fileName string, content string) *File
}

type StorageWriter interface {
	Write(storage *Storage)
}

type StorageCleaner interface {
	Clean(storage *Storage)
}

type Generator interface {
	Annotations() map[string]interface{}
	Generate(application *Application)
}

type ImportFetcher interface {
	Fetch(file *File, entity interface{}) []*Import
}

type ImportRenamer interface {
	Rename(entity interface{}, oldAlias string, newAlias string)
}

type ImportUniquer interface {
	Unique(list []*Import) []*Import
}

type Equaler interface {
	Equal(x interface{}, y interface{}) bool
}

type ContainsChecker interface {
	Contains(collection interface{}, entity interface{}) bool
}
