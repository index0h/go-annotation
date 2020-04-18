package annotation

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewImportUniquer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := &EntityImportUniquer{}

	actual := NewEntityImportUniquer()

	ctrl.AssertEqual(expected, actual)
}

func TestImportUniquer_Unique_WithNilImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	actual := (&EntityImportUniquer{}).Unique(nil)

	ctrl.AssertNil(actual)
}

func TestImportUniquer_Unique_WithEmptyImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	actual := (&EntityImportUniquer{}).Unique([]*Import{})

	ctrl.AssertNotNil(actual)
	ctrl.AssertEmpty(actual)
}

func TestImportUniquer_Unique_WithSamePointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	import1 := &Import{
		Alias:     "import1Alias",
		Namespace: "import1Namespace",
		Comment:   "import1Comment",
	}

	import2 := &Import{
		Alias:     "import2Alias",
		Namespace: "import2Namespace",
		Comment:   "import2Comment",
	}

	imports := []*Import{
		import1,
		import1,
		import2,
		import2,
	}

	importExpected := []*Import{
		import1,
		import2,
	}

	actual := (&EntityImportUniquer{}).Unique(imports)

	ctrl.AssertEqual(importExpected, actual)
	ctrl.AssertSame(import1, actual[0])
	ctrl.AssertSame(import2, actual[1])
}

func TestImportUniquer_Unique_WithSameValues(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	imports := []*Import{
		{
			Alias:     "import1Alias",
			Namespace: "import1Namespace",
			Comment:   "import1Comment",
		},
		{
			Alias:     "import1Alias",
			Namespace: "import1Namespace",
			Comment:   "import1Comment",
		},
		{
			Alias:     "import2Alias",
			Namespace: "import2Namespace",
			Comment:   "import2Comment",
		},
		{
			Alias:     "import2Alias",
			Namespace: "import2Namespace",
			Comment:   "import2Comment",
		},
	}

	importExpected := []*Import{
		imports[0],
		imports[2],
	}

	actual := (&EntityImportUniquer{}).Unique(imports)

	ctrl.AssertEqual(importExpected, actual)
	ctrl.AssertSame(importExpected[0], actual[0])
	ctrl.AssertSame(importExpected[1], actual[1])
}
