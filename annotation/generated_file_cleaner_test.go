package annotation

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewGeneratedFileCleaner(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	actual := NewGeneratedFileCleaner()

	ctrl.AssertNotNil(actual)
}

func TestGeneratedFileCleaner_Clean(t *testing.T) {
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

	storage := &Storage{
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

	expected := &Storage{
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

	(&GeneratedFileCleaner{}).Clean(storage)

	ctrl.AssertEqual(expected, storage)

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

func TestGeneratedFileCleaner_Clean_WithFileNotExists(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl)

	storage := &Storage{
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
		Call((&GeneratedFileCleaner{}).Clean, storage).
		ExpectPanic(ctrl.Type(&os.PathError{}))
}
