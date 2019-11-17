package annotation

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewWriter(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	ctrl.Subtest("").
		Call(NewWriter).
		ExpectResult(&Writer{})
}

func TestWriter_Process(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	rootDir1 := NewTempDir(ctrl, "", "go-annotation-writer-fixtures-root")
	rootDir2 := rootDir1 + "/path1/path2"

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: rootDir1,
				Files: []*File{
					{
						Name:    "first.go",
						Content: "first",
						Annotations: []interface{}{
							FileIsGeneratedAnnotation(true),
						},
					},
					{
						Name:    "second.go",
						Content: "second",
						Annotations: []interface{}{
							FileIsGeneratedAnnotation(false),
						},
					},
					{
						Name:    "third.go",
						Content: "third",
					},
				},
			},
			{
				Name: "namespace/path1/path2",
				Path: rootDir2,
				Files: []*File{
					{
						Name:    "fourth.go",
						Content: "fourth",
						Annotations: []interface{}{
							FileIsGeneratedAnnotation(true),
						},
					},
					{
						Name:    "fifth.go",
						Content: "fifth",
						Annotations: []interface{}{
							FileIsGeneratedAnnotation(true),
						},
					},
					{
						Name:    "sixth.go",
						Content: "sixth",
					},
				},
			},
		},
	}

	writer := &Writer{}

	writer.Process(storage)

	if content, err := ioutil.ReadFile(filepath.Join(rootDir1, "first.go")); err != nil {
		t.Error(err)
	} else {
		ctrl.AssertSame("first", string(content))
	}

	if _, err := os.Stat(filepath.Join(rootDir1, "second.go")); !os.IsNotExist(err) {
		t.Errorf("File '%s' must not exist", filepath.Join(rootDir1, "second.go"))
	}

	if _, err := os.Stat(filepath.Join(rootDir1, "third.go")); !os.IsNotExist(err) {
		t.Errorf("File '%s' must not exist", filepath.Join(rootDir1, "third.go"))
	}

	if content, err := ioutil.ReadFile(filepath.Join(rootDir2, "fourth.go")); err != nil {
		t.Error(err)
	} else {
		ctrl.AssertSame("fourth", string(content))
	}

	if _, err := os.Stat(filepath.Join(rootDir1, "fifth.go")); !os.IsNotExist(err) {
		t.Errorf("File '%s' must not exist", filepath.Join(rootDir1, "fifth.go"))
	}

	if _, err := os.Stat(filepath.Join(rootDir1, "sixth.go")); !os.IsNotExist(err) {
		t.Errorf("File '%s' must not exist", filepath.Join(rootDir1, "sixth.go"))
	}
}

func TestWriter_Process_WithExistFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	rootDir := NewTempDir(ctrl, "", "go-annotation-writer-fixtures-root")
	rootFile := NewFile(ctrl, rootDir, "file.go")

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: rootDir,
				Files: []*File{
					{
						Name:    "file.go",
						Content: "file",
						Annotations: []interface{}{
							FileIsGeneratedAnnotation(true),
						},
					},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call((&Writer{}).Process, storage).
		ExpectPanic(NewErrorf("File '%s' already exists", rootFile.Name()))
}

func TestWriter_Process_WithFolderCreateError(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	rootDir := NewTempDir(ctrl, "", "go-annotation-writer-fixtures-root")

	if err := os.Chmod(rootDir, 0111); err != nil {
		t.Error(err)
	}

	storage := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace",
				Path: rootDir + "/sub_path",
				Files: []*File{
					{
						Name: "file.go",
						Annotations: []interface{}{
							FileIsGeneratedAnnotation(true),
						},
					},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call((&Writer{}).Process, storage).
		ExpectPanic(ctrl.Type(&os.PathError{}))
}
