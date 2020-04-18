package annotation

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNamespace_PackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "packageName"

	model := &Namespace{
		Name: "namespace/packageName",
		Path: "/namespace/path",
	}

	actual := model.PackageName()

	ctrl.AssertSame(expected, actual)
}

func TestNamespace_PackageName_WithFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "filePackageName"

	model := &Namespace{
		Name: "namespace/packageName",
		Path: "/namespace/path",
		Files: []*File{
			{
				Name:        "fileName",
				PackageName: "filePackageName",
			},
		},
	}

	actual := model.PackageName()

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

	model := &Namespace{
		Name: "namespace/packageName",
		Path: "/namespace/path",
		Files: []*File{
			expected,
		},
	}

	actual := model.FindFileByName(name)

	ctrl.AssertSame(expected, actual)
}

func TestNamespace_FindFileByName_WithNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "notFound"

	model := &Namespace{
		Name: "namespace/packageName",
		Path: "/namespace/path",
		Files: []*File{
			{
				Name:        "fileName",
				PackageName: "filePackageName",
			},
		},
	}

	actual := model.FindFileByName(name)

	ctrl.AssertNil(actual)
}

func TestNamespace_FindFileByName_WithEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := ""

	model := &Namespace{
		Name: "namespace/packageName",
		Path: "/namespace/path",
	}

	ctrl.Subtest("").
		Call(model.FindFileByName, name).
		ExpectPanic(NewErrorMessageConstraint("Variable 'name' must be not empty"))
}
