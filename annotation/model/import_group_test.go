package model

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestImportGroup_Validate(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ImportGroup{
		Comment: "importGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "importGroupAnnotation",
			},
		},
		Imports: []*Import{
			{

				Alias:     "alias",
				Namespace: "namespace",
				Comment:   "importComment",
				Annotations: []interface{}{
					&TestAnnotation{
						Name: "importAnnotation",
					},
				},
			},
		},
	}

	modelValue.Validate()
}

func TestImportGroup_Validate_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ImportGroup{}

	modelValue.Validate()
}

func TestImportGroup_Validate_WithNilImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ImportGroup{
		Imports: []*Import{
			nil,
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Imports[0]' must be not nil"))
}

func TestImportGroup_Validate_WithInvalidImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ImportGroup{
		Imports: []*Import{
			{},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Namespace' must be not empty"))
}

func TestImportGroup_String_WithOneImportAndAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	alias := "alias"
	namespace := "namespace"
	expected := `import alias "namespace"
`

	modelValue := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     alias,
				Namespace: namespace,
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestImportGroup_String_WithOneImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	namespace := "namespace"
	expected := `import  "namespace"
`

	modelValue := &ImportGroup{
		Imports: []*Import{
			{
				Namespace: namespace,
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestImportGroup_String_WithOneImportAndAliasAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	alias := "alias"
	namespace := "namespace"
	comment := "import\ncomment"
	expected := `import (
// import
// comment
alias "namespace"
)
`

	modelValue := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     alias,
				Namespace: namespace,
				Comment:   comment,
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestImportGroup_String_WithOneImportAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	namespace := "namespace"
	comment := "import\ncomment"
	expected := `import (
// import
// comment
 "namespace"
)
`

	modelValue := &ImportGroup{
		Imports: []*Import{
			{
				Namespace: namespace,
				Comment:   comment,
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestImportGroup_String_WithImportGroupCommentAndOneImportAndAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	importGroupComment := "importGroup\ncomment"
	alias := "alias"
	namespace := "namespace"
	expected := `// importGroup
// comment
import alias "namespace"
`

	modelValue := &ImportGroup{
		Comment: importGroupComment,
		Imports: []*Import{
			{
				Alias:     alias,
				Namespace: namespace,
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestImportGroup_String_WithImportGroupCommentAndOneImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	importGroupComment := "importGroup\ncomment"
	namespace := "namespace"
	expected := `// importGroup
// comment
import  "namespace"
`

	modelValue := &ImportGroup{
		Comment: importGroupComment,
		Imports: []*Import{
			{
				Namespace: namespace,
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestImportGroup_String_WithImportGroupCommentAndOneImportAndAliasAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	importGroupComment := "importGroup\ncomment"
	alias := "alias"
	namespace := "namespace"
	comment := "import\ncomment"
	expected := `// importGroup
// comment
import (
// import
// comment
alias "namespace"
)
`

	modelValue := &ImportGroup{
		Comment: importGroupComment,
		Imports: []*Import{
			{
				Alias:     alias,
				Namespace: namespace,
				Comment:   comment,
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestImportGroup_String_WithImportGroupCommentAndOneImportAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	importGroupComment := "importGroup\ncomment"
	namespace := "namespace"
	comment := "import\ncomment"
	expected := `// importGroup
// comment
import (
// import
// comment
 "namespace"
)
`

	modelValue := &ImportGroup{
		Comment: importGroupComment,
		Imports: []*Import{
			{
				Namespace: namespace,
				Comment:   comment,
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestImportGroup_String_WithMultipleImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	importGroupComment := "importGroup\ncomment1"
	namespace1 := "namespace1"
	namespace2 := "namespace2"
	comment1 := "import\ncomment1"
	expected := `// importGroup
// comment1
import (
// import
// comment1
 "namespace1"
 "namespace2"
)
`

	modelValue := &ImportGroup{
		Comment: importGroupComment,
		Imports: []*Import{
			{
				Namespace: namespace1,
				Comment:   comment1,
			},
			{
				Namespace: namespace2,
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestImportGroup_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ImportGroup{
		Comment: "importGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "importGroupAnnotation",
			},
		},
		Imports: []*Import{
			{
				Alias:     "alias",
				Namespace: "namespace",
				Comment:   "importComment",
				Annotations: []interface{}{
					&TestAnnotation{
						Name: "importAnnotation",
					},
				},
			},
		},
	}

	actual := modelValue.Clone()

	ctrl.AssertEqual(modelValue, actual)
	ctrl.AssertNotSame(modelValue, actual)
	ctrl.AssertNotSame(modelValue.Annotations[0], actual.(*ImportGroup).Annotations[0])
	ctrl.AssertNotSame(modelValue.Imports[0], actual.(*ImportGroup).Imports[0])
	ctrl.AssertNotSame(modelValue.Imports[0].Annotations[0], actual.(*ImportGroup).Imports[0].Annotations[0])
}

func TestImportGroup_Clone_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ImportGroup{}

	actual := modelValue.Clone()

	ctrl.AssertEqual(modelValue, actual)
	ctrl.AssertNotSame(modelValue, actual)
}

func TestImportGroup_RenameImports_WithRenameAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	modelValue := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     oldAlias,
				Namespace: "namespace",
			},
		},
	}

	modelExpected := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     newAlias,
				Namespace: "namespace",
			},
		},
	}

	modelValue.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(modelExpected, modelValue)
}

func TestImportGroup_RenameImports_WithNotRenameAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	modelValue := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     "alias",
				Namespace: "namespace",
			},
		},
	}

	modelExpected := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     "alias",
				Namespace: "namespace",
			},
		},
	}

	modelValue.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(modelExpected, modelValue)
}

func TestImportGroup_RenameImports_WithInvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	modelValue := &ImportGroup{}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestImportGroup_RenameImports_WithInvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "+invalid"

	modelValue := &ImportGroup{}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
