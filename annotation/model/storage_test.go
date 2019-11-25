package model

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewStorage(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	annotationParser := NewJSONAnnotationParser()
	sourceParser := NewGoSourceParser(annotationParser)

	expected := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
		Namespaces:       []*Namespace{},
	}

	actual := NewStorage()

	ctrl.AssertEqual(expected, actual)
}

func TestStorage_Validate(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Storage{
		AnnotationParser: NewAnnotationParserMock(ctrl),
		SourceParser:     NewSourceParserMock(ctrl),
		Namespaces: []*Namespace{
			{
				Name: "namespace1/alias",
				Path: "/namespace1/path",
			},
			{
				Name: "namespace2/alias",
				Path: "/namespace2/path",
			},
		},
	}

	modelValue.Validate()
}

func TestStorage_Validate_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Storage{
		AnnotationParser: NewAnnotationParserMock(ctrl),
		SourceParser:     NewSourceParserMock(ctrl),
	}

	modelValue.Validate()
}

func TestStorage_Validate_WithNilNamespace(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Storage{
		Namespaces: []*Namespace{
			nil,
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Namespaces[0]' must be not nil"))
}

func TestStorage_Validate_WithInvalidNamespace(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Storage{
		AnnotationParser: NewAnnotationParserMock(ctrl),
		SourceParser:     NewSourceParserMock(ctrl),
		Namespaces: []*Namespace{
			{},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestStorage_Validate_WithDuplicateNamespaceName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Storage{
		AnnotationParser: NewAnnotationParserMock(ctrl),
		SourceParser:     NewSourceParserMock(ctrl),
		Namespaces: []*Namespace{
			{
				Name: "namespace/alias",
				Path: "/namespace1/path",
			},
			{
				Name: "namespace/alias",
				Path: "/namespace2/path",
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Storage has duplicate namespace 'Name': 'namespace/alias'"))
}

func TestStorage_Validate_WithDuplicateNamespacePath(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Storage{
		AnnotationParser: NewAnnotationParserMock(ctrl),
		SourceParser:     NewSourceParserMock(ctrl),
		Namespaces: []*Namespace{
			{
				Name: "namespace1/alias",
				Path: "/namespace/path",
			},
			{
				Name: "namespace2/alias",
				Path: "/namespace/path",
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Storage has duplicate namespace 'Path': '/namespace/path'"))
}

func TestStorage_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Storage{
		AnnotationParser: NewAnnotationParserMock(ctrl),
		SourceParser:     NewSourceParserMock(ctrl),
		Namespaces: []*Namespace{
			{
				Name:      "namespace1/alias",
				Path:      "/namespace1/path",
				IsIgnored: false,
				Files: []*File{
					{
						Name:        "file1Name",
						PackageName: "filePackageName",
						Content:     "package filePackage",
						Comment:     "file1Comment",
						Annotations: []interface{}{
							&TestAnnotation{
								Name: "file1Annotation",
							},
						},
					},
				},
			},
			{
				Name:      "namespace2/alias",
				Path:      "/namespace2/path",
				IsIgnored: true,
			},
		},
	}

	actual := modelValue.Clone()

	ctrl.AssertEqual(
		modelValue,
		actual,
		unit.IgnoreUnexportedOption{Value: AnnotationParserMock{}},
		unit.IgnoreUnexportedOption{Value: SourceParserMock{}},
	)
	ctrl.AssertNotSame(modelValue, actual)
	ctrl.AssertNotSame(modelValue.Namespaces[0], actual.(*Storage).Namespaces[0])
	ctrl.AssertNotSame(modelValue.Namespaces[0].Files[0], actual.(*Storage).Namespaces[0].Files[0])
	ctrl.AssertNotSame(
		modelValue.Namespaces[0].Files[0].Annotations[0],
		actual.(*Storage).Namespaces[0].Files[0].Annotations[0],
	)
	ctrl.AssertNotSame(modelValue.Namespaces[1], actual.(*Storage).Namespaces[1])
}

func TestStorage_FindNamespaceByName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "namespace/packageName"
	expected := &Namespace{
		Name: name,
		Path: "/namespace/path",
	}

	modelValue := &Storage{
		AnnotationParser: NewAnnotationParserMock(ctrl),
		SourceParser:     NewSourceParserMock(ctrl),
		Namespaces: []*Namespace{
			expected,
		},
	}

	actual := modelValue.FindNamespaceByName(name)

	ctrl.AssertSame(expected, actual)
}

func TestStorage_FindNamespaceByName_WithNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "notFound"

	modelValue := &Storage{
		AnnotationParser: NewAnnotationParserMock(ctrl),
		SourceParser:     NewSourceParserMock(ctrl),
		Namespaces: []*Namespace{
			{
				Name: "namespace/packageName",
				Path: "/namespace/path",
			},
		},
	}

	actual := modelValue.FindNamespaceByName(name)

	ctrl.AssertNil(actual)
}

func TestStorage_FindNamespaceByName_WithEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := ""

	modelValue := &Storage{
		AnnotationParser: NewAnnotationParserMock(ctrl),
		SourceParser:     NewSourceParserMock(ctrl),
		Namespaces: []*Namespace{
			{
				Name: "namespace/packageName",
				Path: "/namespace/path",
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.FindNamespaceByName, name).
		ExpectPanic(NewErrorMessageConstraint("Variable 'name' must be not empty"))
}

func TestStorage_ScanRecursive(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	namespace1 := "namespace1"
	namespace2 := "namespace1/ignored"
	namespace3 := "namespace1/ignored/namespace3"
	namespace4 := "namespace1/namespace4"

	file11 := &File{
		Name:        "11.go",
		PackageName: "namespace1",
		Content:     "// 1_1.go\npackage namespace1",
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
		CreateFile(filepath.Join(namespace2, file21.Name), 0666, file11.Content).
		CreateFile(filepath.Join(namespace2, file22.Name), 0666, file12.Content).
		CreateDir(namespace3, 0777).
		CreateFile(filepath.Join(namespace3, file31.Name), 0666, file31.Content).
		CreateFile(filepath.Join(namespace3, file32.Name), 0666, file31.Content).
		CreateDir(namespace4, 0777).
		CreateFile(filepath.Join(namespace4, file41.Name), 0666, file41.Content).
		CreateFile(filepath.Join(namespace4, file42.Name), 0666, file42.Content)

	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	modelValue := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
	}

	modelExpected := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
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
			},
			{
				Name:      namespace3,
				Path:      filepath.Join(fs.RootPath(), namespace3),
				IsIgnored: true,
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
		Parse(file41.Name, file41.Content).
		Return(file41)

	modelValue.ScanRecursive("", fs.RootPath(), "ignored")

	ctrl.AssertEqual(
		modelExpected,
		modelValue,
		unit.IgnoreUnexportedOption{Value: AnnotationParserMock{}},
		unit.IgnoreUnexportedOption{Value: SourceParserMock{}},
	)
	ctrl.AssertSame(file11, modelValue.Namespaces[0].Files[0])
	ctrl.AssertSame(file41, modelValue.Namespaces[3].Files[0])
}

func TestStorage_ScanRecursive_WithoutFiles(t *testing.T) {
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

	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	modelValue := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
	}

	modelExpected := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
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
			},
			{
				Name:      namespace3,
				Path:      filepath.Join(fs.RootPath(), namespace3),
				IsIgnored: true,
			},
			{
				Name:  namespace4,
				Path:  filepath.Join(fs.RootPath(), namespace4),
				Files: []*File{},
			},
		},
	}

	modelValue.ScanRecursive("", fs.RootPath(), "ignored")

	ctrl.AssertEqual(
		modelExpected,
		modelValue,
		unit.IgnoreUnexportedOption{Value: AnnotationParserMock{}},
		unit.IgnoreUnexportedOption{Value: SourceParserMock{}},
	)
}

func TestStorage_ScanRecursive_WithRootNamespace(t *testing.T) {
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
		Content:     "// 1_1.go\npackage namespace1",
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
		CreateFile(filepath.Join(namespace2, file21.Name), 0666, file11.Content).
		CreateFile(filepath.Join(namespace2, file22.Name), 0666, file12.Content).
		CreateDir(namespace3, 0777).
		CreateFile(filepath.Join(namespace3, file31.Name), 0666, file31.Content).
		CreateFile(filepath.Join(namespace3, file32.Name), 0666, file31.Content).
		CreateDir(namespace4, 0777).
		CreateFile(filepath.Join(namespace4, file41.Name), 0666, file41.Content).
		CreateFile(filepath.Join(namespace4, file42.Name), 0666, file42.Content)

	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	modelValue := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
	}

	modelExpected := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
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
			},
			{
				Name:      rootNamespace + "/" + namespace3,
				Path:      filepath.Join(fs.RootPath(), namespace3),
				IsIgnored: true,
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
		Parse(file41.Name, file41.Content).
		Return(file41)

	modelValue.ScanRecursive(rootNamespace, fs.RootPath(), "ignored")

	ctrl.AssertEqual(
		modelExpected,
		modelValue,
		unit.IgnoreUnexportedOption{Value: AnnotationParserMock{}},
		unit.IgnoreUnexportedOption{Value: SourceParserMock{}},
	)
	ctrl.AssertSame(file11, modelValue.Namespaces[1].Files[0])
	ctrl.AssertSame(file41, modelValue.Namespaces[4].Files[0])
}

func TestStorage_ScanRecursive_WithNotReadableFolder(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl).
		CreateDir("folder", 0222)

	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	modelValue := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
	}

	ctrl.Subtest("").
		Call(modelValue.ScanRecursive, fs.RootPath(), "").
		ExpectPanic(ctrl.Type(&os.PathError{}))
}

func TestStorage_ScanFiles(t *testing.T) {
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

	modelValue := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
	}

	expected := []*File{
		file1,
	}

	sourceParser.
		EXPECT().
		Parse(file1.Name, file1.Content).
		Return(file1)

	actual := modelValue.ScanFiles(fs.RootPath())

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(file1, actual[0])
}

func TestStorage_ScanFiles_WithoutFiles(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl)

	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	modelValue := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
	}

	expected := []*File{}

	actual := modelValue.ScanFiles(fs.RootPath())

	ctrl.AssertEqual(expected, actual)
}

func TestStorage_ScanFiles_WithNotExistsFolder(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	modelValue := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
	}

	ctrl.Subtest("").
		Call(modelValue.ScanFiles, "/NotExistedPathHere").
		ExpectPanic(ctrl.Type(&os.PathError{}))
}

func TestStorage_ScanFiles_WithReadFileError(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl).
		CreateFile("file.go", 0222, "")

	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	modelValue := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
	}

	ctrl.Subtest("").
		Call(modelValue.ScanFiles, fs.RootPath()).
		ExpectPanic(ctrl.Type(&os.PathError{}))
}

func TestStorage_RemoveOldGeneratedFiles(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl).
		CreateDir("namespace1", 0777).
		CreateFile("namespace1/1.go", 0666, "package namespace1").
		CreateFile("namespace1/2.go", 0666, "package namespace1").
		CreateFile("namespace1/3.txt", 0666, "3.txt").
		CreateDir("namespace1/ignored", 0777).
		CreateFile("namespace1/ignored/1.go", 0666, "package ignored").
		CreateFile("namespace1/ignored/2.go", 0666, "package ignored").
		CreateFile("namespace1/ignored/3.txt", 0666, "3.txt").
		CreateDir("namespace1/namespace3", 0777).
		CreateFile("namespace1/namespace3/1.go", 0666, "package namespace3").
		CreateFile("namespace1/namespace3/2.go", 0666, "package namespace3").
		CreateFile("namespace1/namespace3/3.txt", 0666, "3.txt")

	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	modelValue := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
		Namespaces: []*Namespace{
			{
				Name: "namespace1",
				Path: filepath.Join(fs.RootPath(), "namespace1"),
				Files: []*File{
					{
						Name:        "1.go",
						Content:     "package namespace1",
						PackageName: "namespace1",
						Annotations: []interface{}{
							FileIsGeneratedAnnotation(true),
						},
					},
					{
						Name:        "2.go",
						Content:     "package namespace1",
						PackageName: "namespace1",
						Annotations: []interface{}{
							FileIsGeneratedAnnotation(false),
						},
					},
				},
			},
			{
				Name:      "namespace1/ignored",
				Path:      filepath.Join(fs.RootPath(), "namespace1", "ignored"),
				IsIgnored: true,
			},
			{
				Name: "namespace1/namespace3",
				Path: filepath.Join(fs.RootPath(), "namespace1", "namespace3"),
				Files: []*File{
					{
						Name:        "1.go",
						Content:     "package namespace3",
						PackageName: "namespace3",
						Annotations: []interface{}{
							FileIsGeneratedAnnotation(true),
						},
					},
					{
						Name:        "2.go",
						Content:     "package namespace3",
						PackageName: "namespace3",
					},
				},
			},
		},
	}

	modelExpected := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
		Namespaces: []*Namespace{
			{
				Name: "namespace1",
				Path: filepath.Join(fs.RootPath(), "namespace1"),
				Files: []*File{
					{
						Name:        "2.go",
						Content:     "package namespace1",
						PackageName: "namespace1",
						Annotations: []interface{}{
							FileIsGeneratedAnnotation(false),
						},
					},
				},
			},
			{
				Name:      "namespace1/ignored",
				Path:      filepath.Join(fs.RootPath(), "namespace1", "ignored"),
				IsIgnored: true,
			},
			{
				Name: "namespace1/namespace3",
				Path: filepath.Join(fs.RootPath(), "namespace1", "namespace3"),
				Files: []*File{
					{
						Name:        "2.go",
						Content:     "package namespace3",
						PackageName: "namespace3",
					},
				},
			},
		},
	}

	modelValue.RemoveOldGeneratedFiles()

	ctrl.AssertEqual(
		modelExpected,
		modelValue,
		unit.IgnoreUnexportedOption{Value: AnnotationParserMock{}},
		unit.IgnoreUnexportedOption{Value: SourceParserMock{}},
	)
	fs.AssertNotFileExists("namespace1/1.go")
	fs.AssertFileExists("namespace1/2.go")
	fs.AssertFileContent("namespace1/2.go", "package namespace1")
	fs.AssertFileExists("namespace1/3.txt")
	fs.AssertFileContent("namespace1/3.txt", "3.txt")

	fs.AssertFileExists("namespace1/ignored/1.go")
	fs.AssertFileContent("namespace1/ignored/2.go", "package ignored")
	fs.AssertFileExists("namespace1/ignored/2.go")
	fs.AssertFileContent("namespace1/ignored/2.go", "package ignored")
	fs.AssertFileExists("namespace1/ignored/3.txt")
	fs.AssertFileContent("namespace1/ignored/3.txt", "3.txt")

	fs.AssertNotFileExists("namespace1/namespace3/1.go")
	fs.AssertFileExists("namespace1/namespace3/2.go")
	fs.AssertFileContent("namespace1/namespace3/2.go", "package namespace3")
	fs.AssertFileExists("namespace1/namespace3/3.txt")
	fs.AssertFileContent("namespace1/namespace3/3.txt", "3.txt")
}

func TestStorage_RemoveOldGeneratedFiles_WithFileNotExists(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl)

	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	modelValue := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
		Namespaces: []*Namespace{
			{
				Name: "namespace1",
				Path: fs.RootPath(),
				Files: []*File{
					{
						Name:        "not_existed.go",
						Content:     "package namespace1",
						PackageName: "namespace1",
						Annotations: []interface{}{
							FileIsGeneratedAnnotation(true),
						},
					},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.RemoveOldGeneratedFiles).
		ExpectPanic(ctrl.Type(&os.PathError{}))
}

func TestStorage_WriteGeneratedFiles(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl).
		CreateDir("root", 0777).
		CreateFile("root/do_not_override.go", 0666, "// do not override\npackage namespace")

	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	modelValue := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: filepath.Join(fs.RootPath(), "root"),
				Files: []*File{
					{
						Name:        "file.go",
						PackageName: "namespace",
					},
					{
						Name:        "do_not_override.go",
						PackageName: "namespace",
						Content:     "// do not override\npackage namespace",
					},
				},
			},
			{
				Name: "namespace/folder1/folder2/folder3",
				Path: filepath.Join(fs.RootPath(), "root", "folder1", "folder2", "folder3"),
				Files: []*File{
					{
						Name:        "second_file.go",
						PackageName: "folder3",
					},
				},
			},
			{
				Name:      "namespace/ignored",
				Path:      filepath.Join(fs.RootPath(), "ignored"),
				IsIgnored: true,
			},
		},
	}

	modelValue.WriteGeneratedFiles()

	fs.AssertFileExists("root/file.go")
	fs.AssertFileContent("root/file.go", modelValue.Namespaces[0].Files[0].String())
	fs.AssertFileExists("root/do_not_override.go")
	fs.AssertFileContent("root/do_not_override.go", modelValue.Namespaces[0].Files[1].Content)
	fs.AssertFileExists("root/folder1/folder2/folder3/second_file.go")
	fs.AssertFileContent("root/folder1/folder2/folder3/second_file.go", modelValue.Namespaces[1].Files[0].String())
}

func TestStorage_WriteGeneratedFiles_WithCreateFolderError(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl).
		CreateDir("root", 0000)

	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	modelValue := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: filepath.Join(fs.RootPath(), "root", "folder"),
				Files: []*File{
					{
						Name:        "file.go",
						PackageName: "namespace",
					},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.WriteGeneratedFiles).
		ExpectPanic(ctrl.Type(&os.PathError{}))
}

func TestStorage_WriteGeneratedFiles_WithFileOverrideError(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl).
		CreateDir("root", 0777).
		CreateFile("root/file.go", 0666, "package root")

	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	modelValue := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: filepath.Join(fs.RootPath(), "root"),
				Files: []*File{
					{
						Name:        "file.go",
						PackageName: "namespace",
					},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.WriteGeneratedFiles).
		ExpectPanic(
			NewErrorMessageConstraint("File '%s' already exists", filepath.Join(fs.RootPath(), "root", "file.go")),
		)
}

func TestStorage_WriteGeneratedFiles_WithFileWriteError(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl).
		CreateDir("root", 0111)

	annotationParser := NewAnnotationParserMock(ctrl)
	sourceParser := NewSourceParserMock(ctrl)

	modelValue := &Storage{
		AnnotationParser: annotationParser,
		SourceParser:     sourceParser,
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: filepath.Join(fs.RootPath(), "root"),
				Files: []*File{
					{
						Name:        "file.go",
						PackageName: "namespace",
					},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.WriteGeneratedFiles).
		ExpectPanic(ctrl.Type(&os.PathError{}))
}
