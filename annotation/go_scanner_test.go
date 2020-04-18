package annotation

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewGoScanner(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	sourceParser := NewSourceParserMock(ctrl)
	annotationParser := NewAnnotationParserMock(ctrl)

	actual := NewGoScanner(sourceParser, annotationParser)

	ctrl.AssertNotNil(actual)
	ctrl.AssertSame(actual.sourceParser, sourceParser)
	ctrl.AssertSame(actual.annotationParser, annotationParser)
}

func TestNewGoScanner_WithNilSourceParse(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	annotationParser := NewAnnotationParserMock(ctrl)

	ctrl.Subtest("").
		Call(NewGoScanner, nil, annotationParser).
		ExpectPanic(NewErrorMessageConstraint("Variable 'sourceParser' must be not nil"))
}

func TestNewGoScanner_WithNilAnnotationParse(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	sourceParser := NewSourceParserMock(ctrl)

	ctrl.Subtest("").
		Call(NewGoScanner, sourceParser, nil).
		ExpectPanic(NewErrorMessageConstraint("Variable 'annotationParser' must be not nil"))
}

func TestGoScanner_Scan(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	namespace1 := "namespace1"
	namespace2 := "namespace1/ignored"
	namespace3 := "namespace1/ignored/namespace3"
	namespace4 := "namespace1/namespace4"

	file11 := &File{
		Name:        "11.go",
		PackageName: "namespace1",
		Content:     "// 11.go\npackage namespace1",
	}

	file12 := &File{
		Name:    "12.txt",
		Content: "// 12.txt ignored txt file",
	}

	file21 := &File{
		Name:        "21.go",
		PackageName: "namespace2",
		Content:     "// 21.go\npackage namespace2",
	}

	file22 := &File{
		Name:    "22.txt",
		Content: "// 22.txt ignored txt file",
	}

	file31 := &File{
		Name:        "31.go",
		PackageName: "namespace3",
		Content:     "// 31.go\npackage namespace3",
	}

	file32 := &File{
		Name:    "32.txt",
		Content: "// 32.txt ignored txt file",
	}

	file41 := &File{
		Name:        "41.go",
		PackageName: "namespace4",
		Content:     "// 41.go\npackage namespace4",
	}

	file42 := &File{
		Name:    "42.txt",
		Content: "// 42.txt ignored txt file",
	}

	fs := NewTmpFS(ctrl).
		CreateDir(namespace1, 0777).
		CreateFile(filepath.Join(namespace1, file11.Name), 0666, file11.Content).
		CreateFile(filepath.Join(namespace1, file12.Name), 0666, file12.Content).
		CreateDir(namespace2, 0777).
		CreateFile(filepath.Join(namespace2, file21.Name), 0666, file21.Content).
		CreateFile(filepath.Join(namespace2, file22.Name), 0666, file22.Content).
		CreateDir(namespace3, 0777).
		CreateFile(filepath.Join(namespace3, file31.Name), 0666, file31.Content).
		CreateFile(filepath.Join(namespace3, file32.Name), 0666, file31.Content).
		CreateDir(namespace4, 0777).
		CreateFile(filepath.Join(namespace4, file41.Name), 0666, file41.Content).
		CreateFile(filepath.Join(namespace4, file42.Name), 0666, file42.Content)

	storage := &Storage{}
	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	scanner := &GoScanner{
		sourceParser:     sourceParser,
		annotationParser: annotationParser,
	}

	expected := &Storage{
		Namespaces: []*Namespace{
			{
				Name: namespace1,
				Path: filepath.Join(fs.RootPath(), namespace1),
				Files: []*File{
					file11,
				},
			},
			{
				Name:      namespace2,
				Path:      filepath.Join(fs.RootPath(), namespace2),
				IsIgnored: true,
				Files: []*File{
					file21,
				},
			},
			{
				Name:      namespace3,
				Path:      filepath.Join(fs.RootPath(), namespace3),
				IsIgnored: true,
				Files: []*File{
					file31,
				},
			},
			{
				Name: namespace4,
				Path: filepath.Join(fs.RootPath(), namespace4),
				Files: []*File{
					file41,
				},
			},
		},
	}

	sourceParser.
		EXPECT().
		Parse(file11.Name, file11.Content).
		Return(file11)

	sourceParser.
		EXPECT().
		Parse(file21.Name, file21.Content).
		Return(file21)

	sourceParser.
		EXPECT().
		Parse(file31.Name, file31.Content).
		Return(file31)

	sourceParser.
		EXPECT().
		Parse(file41.Name, file41.Content).
		Return(file41)

	scanner.Scan(storage, "", fs.RootPath(), "ignored")

	ctrl.AssertEqual(expected, storage)
	ctrl.AssertSame(file11, storage.Namespaces[0].Files[0])
	ctrl.AssertSame(file21, storage.Namespaces[1].Files[0])
	ctrl.AssertSame(file31, storage.Namespaces[2].Files[0])
	ctrl.AssertSame(file41, storage.Namespaces[3].Files[0])
}

func TestGoScanner_Scan_WithoutFiles(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	namespace1 := "namespace1"
	namespace2 := "namespace1/ignored"
	namespace3 := "namespace1/ignored/namespace3"
	namespace4 := "namespace1/namespace4"

	fs := NewTmpFS(ctrl).
		CreateDir(namespace1, 0777).
		CreateDir(namespace2, 0777).
		CreateDir(namespace3, 0777).
		CreateDir(namespace4, 0777)

	storage := &Storage{}
	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	scanner := &GoScanner{
		sourceParser:     sourceParser,
		annotationParser: annotationParser,
	}

	expected := &Storage{
		Namespaces: []*Namespace{
			{
				Name:  namespace1,
				Path:  filepath.Join(fs.RootPath(), namespace1),
				Files: []*File{},
			},
			{
				Name:      namespace2,
				Path:      filepath.Join(fs.RootPath(), namespace2),
				IsIgnored: true,
				Files:     []*File{},
			},
			{
				Name:      namespace3,
				Path:      filepath.Join(fs.RootPath(), namespace3),
				IsIgnored: true,
				Files:     []*File{},
			},
			{
				Name:  namespace4,
				Path:  filepath.Join(fs.RootPath(), namespace4),
				Files: []*File{},
			},
		},
	}

	scanner.Scan(storage, "", fs.RootPath(), "ignored")

	ctrl.AssertEqual(expected, storage)
}

func TestGoScanner_Scan_WithRootNamespace(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	rootNamespace := "root"
	namespace1 := "namespace1"
	namespace2 := "namespace1/ignored"
	namespace3 := "namespace1/ignored/namespace3"
	namespace4 := "namespace1/namespace4"

	file11 := &File{
		Name:        "11.go",
		PackageName: "namespace1",
		Content:     "// 11.go\npackage namespace1",
	}

	file12 := &File{
		Name:    "12.txt",
		Content: "// 12.txt ignored txt file",
	}

	file21 := &File{
		Name:        "21.go",
		PackageName: "namespace2",
		Content:     "// 21.go\npackage namespace2",
	}

	file22 := &File{
		Name:    "22.txt",
		Content: "// 22.txt ignored txt file",
	}

	file31 := &File{
		Name:        "31.go",
		PackageName: "namespace3",
		Content:     "// 31.go\npackage namespace3",
	}

	file32 := &File{
		Name:    "32.txt",
		Content: "// 32.txt ignored txt file",
	}

	file41 := &File{
		Name:        "41.go",
		PackageName: "namespace4",
		Content:     "// 41.go\npackage namespace4",
	}

	file42 := &File{
		Name:    "42.txt",
		Content: "// 42.txt ignored txt file",
	}

	fs := NewTmpFS(ctrl).
		CreateDir(namespace1, 0777).
		CreateFile(filepath.Join(namespace1, file11.Name), 0666, file11.Content).
		CreateFile(filepath.Join(namespace1, file12.Name), 0666, file12.Content).
		CreateDir(namespace2, 0777).
		CreateFile(filepath.Join(namespace2, file21.Name), 0666, file21.Content).
		CreateFile(filepath.Join(namespace2, file22.Name), 0666, file22.Content).
		CreateDir(namespace3, 0777).
		CreateFile(filepath.Join(namespace3, file31.Name), 0666, file31.Content).
		CreateFile(filepath.Join(namespace3, file32.Name), 0666, file31.Content).
		CreateDir(namespace4, 0777).
		CreateFile(filepath.Join(namespace4, file41.Name), 0666, file41.Content).
		CreateFile(filepath.Join(namespace4, file42.Name), 0666, file42.Content)

	storage := &Storage{}
	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	scanner := &GoScanner{
		sourceParser:     sourceParser,
		annotationParser: annotationParser,
	}

	expected := &Storage{
		Namespaces: []*Namespace{
			{
				Name:  rootNamespace,
				Path:  fs.RootPath(),
				Files: []*File{},
			},
			{
				Name: rootNamespace + "/" + namespace1,
				Path: filepath.Join(fs.RootPath(), namespace1),
				Files: []*File{
					file11,
				},
			},
			{
				Name:      rootNamespace + "/" + namespace2,
				Path:      filepath.Join(fs.RootPath(), namespace2),
				IsIgnored: true,
				Files: []*File{
					file21,
				},
			},
			{
				Name:      rootNamespace + "/" + namespace3,
				Path:      filepath.Join(fs.RootPath(), namespace3),
				IsIgnored: true,
				Files: []*File{
					file31,
				},
			},
			{
				Name: rootNamespace + "/" + namespace4,
				Path: filepath.Join(fs.RootPath(), namespace4),
				Files: []*File{
					file41,
				},
			},
		},
	}

	sourceParser.
		EXPECT().
		Parse(file11.Name, file11.Content).
		Return(file11)

	sourceParser.
		EXPECT().
		Parse(file21.Name, file21.Content).
		Return(file21)

	sourceParser.
		EXPECT().
		Parse(file31.Name, file31.Content).
		Return(file31)

	sourceParser.
		EXPECT().
		Parse(file41.Name, file41.Content).
		Return(file41)

	scanner.Scan(storage, rootNamespace, fs.RootPath(), "ignored")

	ctrl.AssertEqual(expected, storage)
	ctrl.AssertSame(file11, storage.Namespaces[1].Files[0])
	ctrl.AssertSame(file41, storage.Namespaces[4].Files[0])
}

func TestGoScanner_Scan_WithNotReadableFolder(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl).
		CreateDir("folder", 0222)

	storage := &Storage{}
	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	scanner := &GoScanner{
		sourceParser:     sourceParser,
		annotationParser: annotationParser,
	}

	ctrl.Subtest("").
		Call(scanner.Scan, storage, fs.RootPath(), "").
		ExpectPanic(ctrl.Type(&os.PathError{}))
}

func TestStorage_scanFiles(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file1 := &File{
		Name:    "1.go",
		Content: "// 1.go\npackage namespace1",
	}

	file2 := &File{
		Name:    "2.txt",
		Content: "// 2.txt ignored txt file",
	}

	fs := NewTmpFS(ctrl).
		CreateFile(file1.Name, 0666, file1.Content).
		CreateFile(file2.Name, 0666, file2.Content)

	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	scanner := &GoScanner{
		sourceParser:     sourceParser,
		annotationParser: annotationParser,
	}

	expected := []*File{
		file1,
	}

	sourceParser.
		EXPECT().
		Parse(file1.Name, file1.Content).
		Return(file1)

	actual := scanner.scanFiles(fs.RootPath())

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(file1, actual[0])
}

func TestStorage_scanFiles_WithoutFiles(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl)

	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	scanner := &GoScanner{
		sourceParser:     sourceParser,
		annotationParser: annotationParser,
	}

	expected := []*File{}

	actual := scanner.scanFiles(fs.RootPath())

	ctrl.AssertEqual(expected, actual)
}

func TestStorage_scanFiles_WithNotExistsFolder(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	scanner := &GoScanner{
		sourceParser:     sourceParser,
		annotationParser: annotationParser,
	}

	ctrl.Subtest("").
		Call(scanner.scanFiles, "/NotExistedPathHere").
		ExpectPanic(ctrl.Type(&os.PathError{}))
}

func TestStorage_scanFiles_WithReadFileError(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl).
		CreateFile("file.go", 0222, "")

	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	scanner := &GoScanner{
		sourceParser:     sourceParser,
		annotationParser: annotationParser,
	}

	ctrl.Subtest("").
		Call(scanner.scanFiles, fs.RootPath()).
		ExpectPanic(ctrl.Type(&os.PathError{}))
}
