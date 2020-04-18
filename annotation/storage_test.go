package annotation

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestStorage_FindNamespaceByName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "namespace/packageName"
	expected := &Namespace{
		Name: name,
		Path: "/namespace/path",
	}

	model := &Storage{
		Namespaces: []*Namespace{
			expected,
		},
	}

	actual := model.FindNamespaceByName(name)

	ctrl.AssertSame(expected, actual)
}

func TestStorage_FindNamespaceByName_WithNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "notFound"

	model := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace/packageName",
				Path: "/namespace/path",
			},
		},
	}

	actual := model.FindNamespaceByName(name)

	ctrl.AssertNil(actual)
}

func TestStorage_FindNamespaceByName_WithEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := ""

	model := &Storage{
		Namespaces: []*Namespace{
			{
				Name: "namespace/packageName",
				Path: "/namespace/path",
			},
		},
	}

	ctrl.Subtest("").
		Call(model.FindNamespaceByName, name).
		ExpectPanic(NewErrorMessageConstraint("Variable 'name' must be not empty"))
}
