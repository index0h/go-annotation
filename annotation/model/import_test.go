package model

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestImport_RealAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "alias"

	modelValue := &Import{
		Namespace: "namespace/alias",
	}

	actual := modelValue.RealAlias()

	ctrl.AssertSame(expected, actual)
}

func TestImport_RealAlias_WithAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "alias"

	modelValue := &Import{
		Alias:     expected,
		Namespace: "namespace",
	}

	actual := modelValue.RealAlias()

	ctrl.AssertSame(expected, actual)
}

func TestImport_Validate_WithAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Import{
		Alias:     "alias",
		Namespace: "namespace",
		Comment:   "importComment",
		Annotations: []interface{}{
			&SimpleSpec{
				TypeName: "importAnnotation",
			},
		},
	}

	modelValue.Validate()
}

func TestImport_Validate(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Import{
		Namespace: "namespace",
	}

	modelValue.Validate()
}

func TestImport_Validate_WithInvalidAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Import{
		Alias:     "+invalid",
		Namespace: "namespace",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Alias' must be valid identifier, actual value: '+invalid'"))
}

func TestImport_Validate_WithEmptyNamespace(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Import{}

	ctrl.Subtest("").
		Call(modelValue.Validate).
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

	modelValue := &Import{
		Alias:     alias,
		Namespace: namespace,
		Comment:   comment,
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestImport_String_WithAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	alias := "alias"
	namespace := "namespace"
	expected := `import alias "namespace"
`

	modelValue := &Import{
		Alias:     alias,
		Namespace: namespace,
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestImport_String_WithComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	namespace := "namespace"
	comment := "import\ncomment"
	expected := `// import
// comment
import  "namespace"
`

	modelValue := &Import{
		Namespace: namespace,
		Comment:   comment,
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestImport_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Import{
		Alias:     "alias",
		Namespace: "namespace",
		Comment:   "importComment",
		Annotations: []interface{}{
			&SimpleSpec{
				TypeName: "importAnnotation",
			},
		},
	}

	actual := modelValue.Clone()

	ctrl.AssertEqual(modelValue, actual)
	ctrl.AssertNotSame(modelValue, actual)
	ctrl.AssertNotSame(modelValue.Annotations[0], actual.(*Import).Annotations[0])
}

func TestImport_Clone_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Import{
		Namespace: "namespace",
	}

	actual := modelValue.Clone()

	ctrl.AssertEqual(modelValue, actual)
	ctrl.AssertNotSame(modelValue, actual)
}

func TestImport_RenameImports_WithRenameAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	modelValue := &Import{
		Alias:     oldAlias,
		Namespace: "namespace",
	}

	modelExpected := &Import{
		Alias:     newAlias,
		Namespace: "namespace",
	}

	modelValue.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(modelExpected, modelValue)
}

func TestImport_RenameImports_WithNotRenameAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	modelValue := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	modelExpected := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	modelValue.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(modelExpected, modelValue)
}

func TestImport_RenameImports_WithInvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	modelValue := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestImport_RenameImports_WithInvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "+invalid"

	modelValue := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
