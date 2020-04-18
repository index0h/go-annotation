package annotation

type Application struct {
	storage *Storage

	annotationParser AnnotationParser
	cloner           Cloner
	storageCleaner   StorageCleaner
	storageWriter    StorageWriter
	importFetcher    ImportFetcher
	importRenamer    ImportRenamer
	importUniquer    ImportUniquer
	equaler          Equaler
	containsChecker  ContainsChecker
	renderer         Renderer
	scanner          Scanner
	sourceParser     SourceParser
	validator        Validator

	generators []Generator
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Storage() *Storage {
	if a.storage == nil {
		a.storage = &Storage{}
	}

	return a.storage
}

func (a *Application) AnnotationParser() AnnotationParser {
	if a.annotationParser == nil {
		a.annotationParser = NewJSONAnnotationParser()
	}

	return a.annotationParser
}

func (a *Application) Cloner() Cloner {
	if a.cloner == nil {
		a.cloner = NewEntityCloner()
	}

	return a.cloner
}

func (a *Application) StorageCleaner() StorageCleaner {
	if a.storageCleaner == nil {
		a.storageCleaner = NewGeneratedFileCleaner()
	}

	return a.storageCleaner
}

func (a *Application) StorageWriter() StorageWriter {
	if a.storageWriter == nil {
		a.storageWriter = NewGeneratedFileWriter(
			a.Validator(),
			a.Renderer(),
		)
	}

	return a.storageWriter
}

func (a *Application) ImportFetcher() ImportFetcher {
	if a.importFetcher == nil {
		a.importFetcher = NewEntityImportFetcher(a.ImportUniquer())
	}

	return a.importFetcher
}

func (a *Application) ImportRenamer() ImportRenamer {
	if a.importRenamer == nil {
		a.importRenamer = NewEntityImportRenamer()
	}

	return a.importRenamer
}

func (a *Application) ImportUniquer() ImportUniquer {
	if a.importUniquer == nil {
		a.importUniquer = NewEntityImportUniquer()
	}

	return a.importUniquer
}

func (a *Application) Equaler() Equaler {
	if a.equaler == nil {
		a.equaler = NewEntityEqualer()
	}

	return a.equaler
}

func (a *Application) ContainsChecker() ContainsChecker {
	if a.containsChecker == nil {
		a.containsChecker = NewEntityContainsChecker(a.Equaler())
	}

	return a.containsChecker
}

func (a *Application) Scanner() Scanner {
	if a.scanner == nil {
		a.scanner = NewGoScanner(a.SourceParser(), a.AnnotationParser())
	}

	return a.scanner
}

func (a *Application) Renderer() Renderer {
	if a.renderer == nil {
		a.renderer = NewEntityRenderer()
	}

	return a.renderer
}

func (a *Application) SourceParser() SourceParser {
	if a.sourceParser == nil {
		a.sourceParser = NewGoSourceParser(a.AnnotationParser())
	}

	return a.sourceParser
}

func (a *Application) Validator() Validator {
	if a.validator == nil {
		a.validator = NewEntityValidator()
	}

	return a.validator
}

func (a *Application) Scan(rootNamespace string, rootPath string, ignores ...string) {
	a.Scanner().Scan(a.storage, rootNamespace, rootPath, ignores...)
}

func (a *Application) RegisterGenerator(generator Generator) {
	annotationParser := a.AnnotationParser()

	for name, annotation := range generator.Annotations() {
		annotationParser.SetAnnotation(name, annotation)
	}

	if a.generators == nil {
		a.generators = []Generator{}
	}

	a.generators = append(a.generators, generator)
}

func (a *Application) Generate() {
	a.StorageCleaner().Clean(a.storage)

	for _, generator := range a.generators {
		generator.Generate(a)
	}

	a.StorageWriter().Write(a.storage)
}
