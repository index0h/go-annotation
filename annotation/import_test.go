package annotation

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestImport_RealAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "alias"

	model := &Import{
		Namespace: "namespace/alias",
	}

	actual := model.RealAlias()

	ctrl.AssertSame(expected, actual)
}

func TestImport_RealAlias_WithAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "alias"

	model := &Import{
		Alias:     expected,
		Namespace: "namespace",
	}

	actual := model.RealAlias()

	ctrl.AssertSame(expected, actual)
}
