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

func TestImport_Validate_WithAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Import{
		Alias:     "alias",
		Namespace: "namespace",
		Comment:   "importComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "importAnnotation",
			},
		},
	}

	model.Validate()
}

func TestImport_Validate(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Import{
		Namespace: "namespace",
	}

	model.Validate()
}

func TestImport_Validate_WithInvalidAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Import{
		Alias:     "+invalid",
		Namespace: "namespace",
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Alias' must be valid identifier, actual value: '+invalid'"))
}

func TestImport_Validate_WithEmptyNamespace(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Import{}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Namespace' must be not empty"))
}

func TestImport_String_WithAliasAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	alias := "alias"
	namespace := "namespace"
	comment := "import\ncomment"
	expected := `// import
// comment
import alias "namespace"
`

	model := &Import{
		Alias:     alias,
		Namespace: namespace,
		Comment:   comment,
	}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestImport_String_WithAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	alias := "alias"
	namespace := "namespace"
	expected := `import alias "namespace"
`

	model := &Import{
		Alias:     alias,
		Namespace: namespace,
	}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestImport_String_WithComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	namespace := "namespace"
	comment := "import\ncomment"
	expected := `// import
// comment
import "namespace"
`

	model := &Import{
		Namespace: namespace,
		Comment:   comment,
	}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestImport_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Import{
		Alias:     "alias",
		Namespace: "namespace",
		Comment:   "importComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "importAnnotation",
			},
		},
	}

	actual := model.Clone()

	ctrl.AssertEqual(model, actual)
	ctrl.AssertNotSame(model, actual)
	ctrl.AssertNotSame(model.Annotations[0], actual.(*Import).Annotations[0])
}

func TestImport_Clone_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Import{
		Namespace: "namespace",
	}

	actual := model.Clone()

	ctrl.AssertEqual(model, actual)
	ctrl.AssertNotSame(model, actual)
}

func TestImport_EqualSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1 := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	model2 := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertTrue(actual)
}

func TestImport_EqualSpec_WithRealAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1 := &Import{
		Namespace: "namespace/alias",
	}

	model2 := &Import{
		Alias:     "alias",
		Namespace: "namespace/alias",
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertTrue(actual)
}

func TestImport_EqualSpec_WithAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1 := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	model2 := "model2"

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestImport_EqualSpec_WithAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1 := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	model2 := &Import{
		Alias:     "another",
		Namespace: "namespace",
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestImport_EqualSpec_WithNamespace(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1 := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	model2 := &Import{
		Alias:     "alias",
		Namespace: "another",
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestImport_RenameImports_WithRenameAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	model := &Import{
		Alias:     oldAlias,
		Namespace: "namespace",
	}

	expected := &Import{
		Alias:     newAlias,
		Namespace: "namespace",
	}

	model.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(expected, model)
}

func TestImport_RenameImports_WithNotRenameAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	model := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	expected := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	model.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(expected, model)
}

func TestImport_RenameImports_WithInvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	model := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestImport_RenameImports_WithInvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "+invalid"

	model := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
