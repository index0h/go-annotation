package model

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewStorage(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := &Storage{
		Namespaces: []*Namespace{},
	}

	actual := NewStorage()

	ctrl.AssertEqual(expected, actual)
}

func TestStorage_Validate(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Storage{
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

	modelValue := &Storage{}

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
							&SimpleSpec{
								TypeName: "file1Annotation",
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

	ctrl.AssertEqual(modelValue, actual)
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
