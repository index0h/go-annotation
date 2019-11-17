package annotation

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewScanner(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	ctrl.Subtest("").
		Call(NewScanner, "namespace", "path", "path1", "path2").
		ExpectResult(
			&Scanner{
				namespace:    "namespace",
				path:         "path",
				ignoredPaths: map[string]bool{"path1": true, "path2": true},
			},
		)
}

func TestScanner_Process(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	rootDir := NewTempDir(ctrl, "", "go-annotation-scanner-fixtures-root")
	rootFile1 := NewFile(ctrl, rootDir, "1_file.go")
	rootFile2 := NewFile(ctrl, rootDir, "2_file.txt")

	if _, err := rootFile1.WriteString("rootFile1Content"); err != nil {
		t.Error(err)
	}

	if _, err := rootFile2.WriteString("rootFile2Content"); err != nil {
		t.Error(err)
	}

	srcDir := NewDir(ctrl, rootDir, "annotation")
	srcFile1 := NewFile(ctrl, srcDir, "1_file.go")
	srcFile2 := NewFile(ctrl, srcDir, "2_file.go")

	if _, err := srcFile1.WriteString("src1File1Content"); err != nil {
		t.Error(err)
	}

	if _, err := srcFile2.WriteString("src1File2Content"); err != nil {
		t.Error(err)
	}

	_ = NewDir(ctrl, rootDir, "empty")

	withoutGoFilesDir := NewDir(ctrl, rootDir, "withoutGoFiles")
	withoutGoFilesFile1 := NewFile(ctrl, withoutGoFilesDir, "1_file.txt")
	withoutGoFilesFile2 := NewFile(ctrl, withoutGoFilesDir, "2_file.txt")

	if _, err := withoutGoFilesFile1.WriteString("withoutGoFilesFile1Content"); err != nil {
		t.Error(err)
	}

	if _, err := withoutGoFilesFile2.WriteString("withoutGoFilesFile2Content"); err != nil {
		t.Error(err)
	}

	src2Dir := NewDir(ctrl, withoutGoFilesDir, "annotation")
	src2File1 := NewFile(ctrl, src2Dir, "1_file.go")
	src2File2 := NewFile(ctrl, src2Dir, "2_file.go")

	if _, err := src2File1.WriteString("src2File1Content"); err != nil {
		t.Error(err)
	}

	if _, err := src2File2.WriteString("src2File2Content"); err != nil {
		t.Error(err)
	}

	ignoredDir := NewDir(ctrl, withoutGoFilesDir, "ignored")
	ignoredFile1 := NewFile(ctrl, ignoredDir, "1_file.go")
	ignoredFile2 := NewFile(ctrl, ignoredDir, "2_file.go")

	if _, err := ignoredFile1.WriteString("ignored1File2Content"); err != nil {
		t.Error(err)
	}

	if _, err := ignoredFile2.WriteString("ignored2File2Content"); err != nil {
		t.Error(err)
	}

	expected := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "root.namespace",
				Path: rootDir,
				Files: []*File{
					{
						Name:    filepath.Base(rootFile1.Name()),
						Content: "rootFile1Content",
					},
				},
			},
			{
				Name: "root.namespace/" + filepath.Base(srcDir),
				Path: srcDir,
				Files: []*File{
					{
						Name:    filepath.Base(srcFile1.Name()),
						Content: "src1File1Content",
					},
					{
						Name:    filepath.Base(srcFile2.Name()),
						Content: "src1File2Content",
					},
				},
			},
			{
				Name: "root.namespace/" + filepath.Base(withoutGoFilesDir) + "/" + filepath.Base(src2Dir),
				Path: src2Dir,
				Files: []*File{
					{
						Name:    filepath.Base(src2File1.Name()),
						Content: "src2File1Content",
					},
					{
						Name:    filepath.Base(src2File2.Name()),
						Content: "src2File2Content",
					},
				},
			},
		},
	}

	actual := &Storage{}

	(&Scanner{
		namespace:    "root.namespace",
		path:         rootDir,
		ignoredPaths: map[string]bool{ignoredDir: true},
	}).Process(actual)

	ctrl.AssertEqual(expected, actual)
}

func TestScanner_Process_WithNotExistedFolder(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	ctrl.Subtest("").Call(
		(&Scanner{namespace: "root.namespace", path: "/NotExistedPathHere", ignoredPaths: map[string]bool{}}).Process,
		&Storage{},
	).
		ExpectPanic(ctrl.Type(&os.PathError{}))
}

func TestScanner_Process_WithDeclinedFileRead(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	rootDir := NewTempDir(ctrl, "", "go-annotation-scanner-fixtures-root")
	rootFile1 := NewFile(ctrl, rootDir, "file.go")

	if err := rootFile1.Chmod(0222); err != nil {
		t.Error(err)
	}

	ctrl.Subtest("").Call(
		(&Scanner{namespace: "root.namespace", path: rootDir, ignoredPaths: map[string]bool{}}).Process,
		&Storage{},
	).
		ExpectPanic(ctrl.Type(&os.PathError{}))
}
