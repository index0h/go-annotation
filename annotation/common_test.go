package annotation

import (
	"encoding/json"
	"github.com/pkg/errors"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func Test_cloneAnnotations(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	annotations := []interface{}{
		&TestAnnotation{
			Name: "typeName1",
		},
		&TestAnnotation{
			Name: "typeName2",
		},
	}

	actual := cloneAnnotations(annotations)

	ctrl.AssertEqual(annotations, actual)
	ctrl.AssertNotSame(annotations[0], actual[0])
	ctrl.AssertNotSame(annotations[1], actual[1])
}

func Test_cloneAnnotations_WithNilAnnotations(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	actual := cloneAnnotations(nil)

	ctrl.AssertNil(actual)
}

func Test_cloneAnnotations_WithEmptyAnnotations(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	annotations := []interface{}{}

	actual := cloneAnnotations(annotations)

	ctrl.AssertEqual(annotations, actual)
}

func Test_cloneAnnotations_WithMarshalJSONPanic(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	err := errors.New("TestError")

	annotation := NewMarshalerJSONMock(ctrl)

	annotations := []interface{}{
		annotation,
	}

	annotation.
		EXPECT().
		MarshalJSON().
		Return(nil, err)

	ctrl.Subtest("").
		Call(cloneAnnotations, annotations).
		ExpectPanic(ctrl.Type(&json.MarshalerError{}))
}

func Test_cloneAnnotations_WithUnmarshalJSONPanic(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	marshaled := []byte(`"data"`)

	annotation := NewMarshalerJSONMock(ctrl)

	annotations := []interface{}{
		annotation,
	}

	annotation.
		EXPECT().
		MarshalJSON().
		Return(marshaled, nil)

	ctrl.Subtest("").
		Call(cloneAnnotations, annotations).
		ExpectPanic(ctrl.Type(&json.UnmarshalTypeError{}))
}

func Test_uniqImports_WithNilImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	actual := uniqImports(nil)

	ctrl.AssertNil(actual)
}

func Test_uniqImports_WithEmptyImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	actual := uniqImports([]*Import{})

	ctrl.AssertNotNil(actual)
	ctrl.AssertEmpty(actual)
}

func Test_uniqImports_WithSamePointer(t *testing.T) {
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

	actual := uniqImports(imports)

	ctrl.AssertEqual(importExpected, actual)
	ctrl.AssertSame(import1, actual[0])
	ctrl.AssertSame(import2, actual[1])
}

func Test_uniqImports_WithSameValues(t *testing.T) {
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

	actual := uniqImports(imports)

	ctrl.AssertEqual(importExpected, actual)
	ctrl.AssertSame(importExpected[0], actual[0])
	ctrl.AssertSame(importExpected[1], actual[1])
}
