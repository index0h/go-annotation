package annotation

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewCleaner(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	ctrl.Subtest("").
		Call(NewCleaner).
		ExpectResult(&Cleaner{})
}

func TestCleaner_Process(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	rootDir := NewTempDir(ctrl, "", "go-annotation-cleaner-fixtures-root")
	rootFile1 := NewFile(ctrl, rootDir, "remove_file.go")
	rootFile2 := NewFile(ctrl, rootDir, "dont_remove_file1.go")
	rootFile3 := NewFile(ctrl, rootDir, "dont_remove_file2.go")

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: rootDir,
				Files: []*File{
					{
						Name: filepath.Base(rootFile1.Name()),
						Annotations: []interface{}{
							FileIsGeneratedAnnotation(true),
						},
					},
					{
						Name: filepath.Base(rootFile2.Name()),
					},
					{
						Name: filepath.Base(rootFile3.Name()),
						Annotations: []interface{}{
							FileIsGeneratedAnnotation(false),
						},
					},
				},
			},
		},
	}

	cleaner := &Cleaner{}

	cleaner.Process(storage)

	if _, err := os.Stat(rootFile1.Name()); !os.IsNotExist(err) {
		t.Errorf("File '%s' must removed", rootFile1.Name())
	}

	if _, err := os.Stat(rootFile2.Name()); err != nil {
		t.Errorf("File '%s' must exist", rootFile2.Name())
	}

	if _, err := os.Stat(rootFile3.Name()); err != nil {
		t.Errorf("File '%s' must exist", rootFile3.Name())
	}
}

func TestCleaner_Process_WithNotExistedFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	rootDir := NewTempDir(ctrl, "", "go-annotation-cleaner-fixtures-root")

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: rootDir,
				Files: []*File{
					{
						Name: "not_existed.go",
						Annotations: []interface{}{
							FileIsGeneratedAnnotation(true),
						},
					},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call((&Cleaner{}).Process, storage).
		ExpectPanic(ctrl.Type(&os.PathError{}))
}
