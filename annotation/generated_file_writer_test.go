package annotation

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewGeneratedFileWriter(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	validator := NewEntityValidator()
	renderer := NewEntityRenderer()

	expected := &GeneratedFileWriter{
		validator: validator,
		renderer:  renderer,
	}

	actual := NewGeneratedFileWriter(validator, renderer)

	ctrl.AssertEqual(expected, actual)
}

func TestNewGeneratedFileWriter_WithNilValidator(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	renderer := NewEntityRenderer()

	ctrl.Subtest("").
		Call(NewGeneratedFileWriter, nil, renderer).
		ExpectPanic(NewErrorMessageConstraint("Variable 'validator' must be not nil"))
}

func TestNewGeneratedFileWriter_WithNilRenderer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	validator := NewEntityValidator()

	ctrl.Subtest("").
		Call(NewGeneratedFileWriter, validator, nil).
		ExpectPanic(NewErrorMessageConstraint("Variable 'renderer' must be not nil"))
}

func TestGeneratedFileWriter_Write(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl).
		CreateDir("root", 0777).
		CreateFile("root/do_not_override.go", 0666, "// do not override\npackage namespace")

	storage := &Storage{
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

	validator := NewEntityValidator()
	renderer := NewEntityRenderer()

	generatedFileWriter := &GeneratedFileWriter{validator: validator, renderer: renderer}

	generatedFileWriter.Write(storage)

	fs.AssertFileExists("root/file.go")
	fs.AssertFileContent("root/file.go", storage.Namespaces[0].Files[0].Content)
	fs.AssertFileExists("root/do_not_override.go")
	fs.AssertFileContent("root/do_not_override.go", storage.Namespaces[0].Files[1].Content)
	fs.AssertFileExists("root/folder1/folder2/folder3/second_file.go")
	fs.AssertFileContent("root/folder1/folder2/folder3/second_file.go", storage.Namespaces[1].Files[0].Content)
}

func TestGeneratedFileWriter_Write_WithInvalidStorage(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl).
		CreateDir("root", 0000)

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "",
				Path: filepath.Join(fs.RootPath(), "root", "folder"),
			},
		},
	}

	validator := NewEntityValidator()
	renderer := NewEntityRenderer()

	generatedFileWriter := &GeneratedFileWriter{validator: validator, renderer: renderer}

	ctrl.Subtest("").
		Call(generatedFileWriter.Write, storage).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestGeneratedFileWriter_Write_WithCreateFolderError(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl).
		CreateDir("root", 0000)

	storage := &Storage{
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

	validator := NewEntityValidator()
	renderer := NewEntityRenderer()

	generatedFileWriter := &GeneratedFileWriter{validator: validator, renderer: renderer}

	ctrl.Subtest("").
		Call(generatedFileWriter.Write, storage).
		ExpectPanic(ctrl.Type(&os.PathError{}))
}

func TestGeneratedFileWriter_Write_WithFileOverrideError(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl).
		CreateDir("root", 0777).
		CreateFile("root/file.go", 0666, "package root")

	storage := &Storage{
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

	validator := NewEntityValidator()
	renderer := NewEntityRenderer()

	generatedFileWriter := &GeneratedFileWriter{validator: validator, renderer: renderer}

	ctrl.Subtest("").
		Call(generatedFileWriter.Write, storage).
		ExpectPanic(
			NewErrorMessageConstraint("File '%s' already exists", filepath.Join(fs.RootPath(), "root", "file.go")),
		)
}

func TestGeneratedFileWriter_Write_WithFileWriteError(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fs := NewTmpFS(ctrl).
		CreateDir("root", 0111)

	storage := &Storage{
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

	validator := NewEntityValidator()
	renderer := NewEntityRenderer()

	generatedFileWriter := &GeneratedFileWriter{validator: validator, renderer: renderer}

	ctrl.Subtest("").
		Call(generatedFileWriter.Write, storage).
		ExpectPanic(ctrl.Type(&os.PathError{}))
}
