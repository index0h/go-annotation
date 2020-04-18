package annotation

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewApplication(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	actual := NewApplication()

	ctrl.AssertNotNil(actual)
	ctrl.AssertNil(actual.storage)
	ctrl.AssertNil(actual.annotationParser)
	ctrl.AssertNil(actual.cloner)
	ctrl.AssertNil(actual.storageCleaner)
	ctrl.AssertNil(actual.storageWriter)
	ctrl.AssertNil(actual.importFetcher)
	ctrl.AssertNil(actual.importRenamer)
	ctrl.AssertNil(actual.importUniquer)
	ctrl.AssertNil(actual.equaler)
	ctrl.AssertNil(actual.containsChecker)
	ctrl.AssertNil(actual.renderer)
	ctrl.AssertNil(actual.scanner)
	ctrl.AssertNil(actual.sourceParser)
	ctrl.AssertNil(actual.validator)
	ctrl.AssertNil(actual.generators)
}

func TestApplication_Storage(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	application := &Application{}

	actual := application.Storage()

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(application.storage, actual)
}

func TestApplication_AnnotationParser(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	application := &Application{}

	actual := application.AnnotationParser()

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(application.annotationParser, actual)
}

func TestApplication_Cloner(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	application := &Application{}

	actual := application.Cloner()

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(application.cloner, actual)
}

func TestApplication_StorageCleaner(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	application := &Application{}

	actual := application.StorageCleaner()

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(application.storageCleaner, actual)
}

func TestApplication_StorageWriter(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	application := &Application{}

	actual := application.StorageWriter()

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(application.storageWriter, actual)
	ctrl.AssertSame(application.storageWriter.(*GeneratedFileWriter).validator, actual.(*GeneratedFileWriter).validator)
	ctrl.AssertSame(application.storageWriter.(*GeneratedFileWriter).renderer, actual.(*GeneratedFileWriter).renderer)
}

func TestApplication_ImportFetcher(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	application := &Application{}

	actual := application.ImportFetcher()

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(application.importFetcher, actual)
	ctrl.AssertSame(
		application.importFetcher.(*EntityImportFetcher).importUniquer,
		actual.(*EntityImportFetcher).importUniquer,
	)
}

func TestApplication_ImportRenamer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	application := &Application{}

	actual := application.ImportRenamer()

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(application.importRenamer, actual)
}

func TestApplication_ImportUniquer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	application := &Application{}

	actual := application.ImportUniquer()

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(application.importUniquer, actual)
}

func TestApplication_Equaler(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	application := &Application{}

	actual := application.Equaler()

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(application.equaler, actual)
}

func TestApplication_ContainsChecker(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	application := &Application{}

	actual := application.ContainsChecker()

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(application.containsChecker, actual)
	ctrl.AssertSame(application.equaler, actual.(*EntityContainsChecker).equaler)
}

func TestApplication_Scanner(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	application := &Application{}

	actual := application.Scanner()

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(application.scanner, actual)
	ctrl.AssertSame(application.scanner.(*GoScanner).annotationParser, actual.(*GoScanner).annotationParser)
	ctrl.AssertSame(application.scanner.(*GoScanner).sourceParser, actual.(*GoScanner).sourceParser)
}

func TestApplication_Renderer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	application := &Application{}

	actual := application.Renderer()

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(application.renderer, actual)
}

func TestApplication_SourceParser(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	application := &Application{}

	actual := application.SourceParser()

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(application.sourceParser, actual)
	ctrl.AssertSame(
		application.sourceParser.(*GoSourceParser).annotationParser,
		actual.(*GoSourceParser).annotationParser,
	)
}

func TestApplication_Validator(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	application := &Application{}

	actual := application.Validator()

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(application.validator, actual)
}

func TestApplication_Scan(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	rootNamespace := "rootNamespace"
	rootPath := "rootPath"
	ignores := []string{"ignore1", "ignore2"}

	storage := &Storage{}
	scanner := NewScannerMock(ctrl)

	application := &Application{
		storage: storage,
		scanner: scanner,
	}

	scanner.
		EXPECT().
		Scan(ctrl.Same(storage), rootNamespace, rootPath, ignores[0], ignores[1]).
		Return()

	application.Scan(rootNamespace, rootPath, ignores...)
}

func TestApplication_RegisterGenerator(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	annotation1Name := "Annotation1"
	annotation1Type := TestAnnotation{Name: annotation1Name}
	annotation2Name := "Annotation2"
	annotation2Type := TestAnnotation{Name: annotation2Name}

	annotations := map[string]interface{}{
		annotation1Name: annotation1Type,
		annotation2Name: annotation2Type,
	}

	generator := NewGeneratorMock(ctrl)
	annotationParser := NewAnnotationParserMock(ctrl)

	application := &Application{
		annotationParser: annotationParser,
	}

	generator.
		EXPECT().
		Annotations().
		Return(annotations)

	annotationParser.
		EXPECT().
		SetAnnotation(
			ctrl.Or(ctrl.Same(annotation1Name), ctrl.Same(annotation2Name)),
			ctrl.Or(ctrl.Same(annotation1Type), ctrl.Same(annotation2Type)),
		).
		Return()

	annotationParser.
		EXPECT().
		SetAnnotation(
			ctrl.Or(ctrl.Same(annotation1Name), ctrl.Same(annotation2Name)),
			ctrl.Or(ctrl.Same(annotation1Type), ctrl.Same(annotation2Type)),
		).
		Return()

	application.RegisterGenerator(generator)

	ctrl.AssertLength(1, application.generators)
	ctrl.AssertSame(generator, application.generators[0])
}

func TestApplication_Generate(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	storage := &Storage{}
	storageCleaner := NewStorageCleanerMock(ctrl)
	generator1 := NewGeneratorMock(ctrl)
	generator2 := NewGeneratorMock(ctrl)
	storageWriter := NewStorageWriterMock(ctrl)

	application := &Application{
		storage:        storage,
		storageCleaner: storageCleaner,
		storageWriter:  storageWriter,
		generators: []Generator{
			generator1,
			generator2,
		},
	}

	storageCleaner.
		EXPECT().
		Clean(ctrl.Same(storage)).
		Return()

	generator1.
		EXPECT().
		Generate(ctrl.Same(application)).
		Return()

	generator2.
		EXPECT().
		Generate(ctrl.Same(application)).
		Return()

	storageWriter.
		EXPECT().
		Write(ctrl.Same(storage)).
		Return()

	application.Generate()
}
