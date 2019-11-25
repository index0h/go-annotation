package model

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNamespace_Validate(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Namespace{
		Name: "namespace/alias",
		Path: "/namespace/path",
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
			{
				Name:        "file2Name",
				PackageName: "filePackageName",
				Content:     "package filePackage",
				Comment:     "file2Comment",
				Annotations: []interface{}{
					&TestAnnotation{
						Name: "file2Annotation",
					},
				},
			},
		},
	}

	modelValue.Validate()
}

func TestNamespace_Validate_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Namespace{
		Name: "namespace/alias",
		Path: "/namespace/path",
	}

	modelValue.Validate()
}

func TestNamespace_Validate_WithEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Namespace{
		Path: "/namespace/path",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestNamespace_Validate_WithEmptyPath(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Namespace{
		Name: "/namespace/alias",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Path' must be not empty"))
}

func TestNamespace_Validate_WithInvalidPath(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Namespace{
		Name: "namespace/alias",
		Path: "namespace/path",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Path' must be absolute path, actual value: 'namespace/path'"))
}

func TestNamespace_Validate_WithIgnoredAndFiles(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Namespace{
		Name:      "namespace/alias",
		Path:      "/namespace/path",
		IsIgnored: true,
		Files: []*File{
			{
				Name:        "fileName",
				PackageName: "filePackageName",
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Ignored namespace with name: 'namespace/alias' must have no files"))
}

func TestNamespace_Validate_WithNilFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Namespace{
		Name: "namespace/alias",
		Path: "/namespace/path",
		Files: []*File{
			nil,
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Files[0]' must be not nil"))
}

func TestNamespace_Validate_WithDuplicateFileName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Namespace{
		Name: "namespace/alias",
		Path: "/namespace/path",
		Files: []*File{
			{
				Name:        "fileName",
				PackageName: "filePackageName",
			},
			{
				Name:        "fileName",
				PackageName: "filePackageName",
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Namespace has duplicate file name: fileName"))
}

func TestNamespace_Validate_WithDifferentPackageNames(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Namespace{
		Name: "namespace/alias",
		Path: "/namespace/path",
		Files: []*File{
			{
				Name:        "fileName1",
				PackageName: "filePackageName1",
			},
			{
				Name:        "fileName2",
				PackageName: "filePackageName2",
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Namespace has different packages"))
}

func TestNamespace_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Namespace{
		Name:      "namespace/alias",
		Path:      "/namespace/path",
		IsIgnored: true,
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
			{
				Name:        "file2Name",
				PackageName: "filePackageName",
				Content:     "package filePackage",
				Comment:     "file2Comment",
				Annotations: []interface{}{
					&TestAnnotation{
						Name: "file2Annotation",
					},
				},
			},
		},
	}

	actual := modelValue.Clone()

	ctrl.AssertEqual(modelValue, actual)
	ctrl.AssertNotSame(modelValue, actual)
	ctrl.AssertNotSame(modelValue.Files[0], actual.(*Namespace).Files[0])
	ctrl.AssertNotSame(modelValue.Files[1], actual.(*Namespace).Files[1])
	ctrl.AssertNotSame(modelValue.Files[0].Annotations[0], actual.(*Namespace).Files[0].Annotations[0])
	ctrl.AssertNotSame(modelValue.Files[1].Annotations[0], actual.(*Namespace).Files[1].Annotations[0])
}

func TestNamespace_Clone_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Namespace{
		Name: "namespace/alias",
		Path: "/namespace/path",
	}

	actual := modelValue.Clone()

	ctrl.AssertEqual(modelValue, actual)
	ctrl.AssertNotSame(modelValue, actual)
}

func TestNamespace_PackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "packageName"

	modelValue := &Namespace{
		Name: "namespace/packageName",
		Path: "/namespace/path",
	}

	actual := modelValue.PackageName()

	ctrl.AssertSame(expected, actual)
}

func TestNamespace_PackageName_WithFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "filePackageName"

	modelValue := &Namespace{
		Name: "namespace/packageName",
		Path: "/namespace/path",
		Files: []*File{
			{
				Name:        "fileName",
				PackageName: "filePackageName",
			},
		},
	}

	actual := modelValue.PackageName()

	ctrl.AssertSame(expected, actual)
}

func TestNamespace_FindFileByName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "fileName"
	expected := &File{
		Name:        name,
		PackageName: "filePackageName",
	}

	modelValue := &Namespace{
		Name: "namespace/packageName",
		Path: "/namespace/path",
		Files: []*File{
			expected,
		},
	}

	actual := modelValue.FindFileByName(name)

	ctrl.AssertSame(expected, actual)
}

func TestNamespace_FindFileByName_WithNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "notFound"

	modelValue := &Namespace{
		Name: "namespace/packageName",
		Path: "/namespace/path",
		Files: []*File{
			{
				Name:        "fileName",
				PackageName: "filePackageName",
			},
		},
	}

	actual := modelValue.FindFileByName(name)

	ctrl.AssertNil(actual)
}

func TestNamespace_FindFileByName_WithEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := ""

	modelValue := &Namespace{
		Name: "namespace/packageName",
		Path: "/namespace/path",
	}

	ctrl.Subtest("").
		Call(modelValue.FindFileByName, name).
		ExpectPanic(NewErrorMessageConstraint("Variable 'name' must be not empty"))
}
