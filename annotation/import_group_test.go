package annotation

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestImportGroup_Validate(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ImportGroup{
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

	model.Validate()
}

func TestImportGroup_Validate_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ImportGroup{}

	model.Validate()
}

func TestImportGroup_Validate_WithNilImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ImportGroup{
		Imports: []*Import{
			nil,
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Imports[0]' must be not nil"))
}

func TestImportGroup_Validate_WithInvalidImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ImportGroup{
		Imports: []*Import{
			{},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Namespace' must be not empty"))
}

func TestImportGroup_String_WithOneImportAndAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	alias := "alias"
	namespace := "namespace"
	expected := `import alias "namespace"
`

	model := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     alias,
				Namespace: namespace,
			},
		},
	}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestImportGroup_String_WithOneImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	namespace := "namespace"
	expected := `import  "namespace"
`

	model := &ImportGroup{
		Imports: []*Import{
			{
				Namespace: namespace,
			},
		},
	}

	actual := model.String()

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

	model := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     alias,
				Namespace: namespace,
				Comment:   comment,
			},
		},
	}

	actual := model.String()

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

	model := &ImportGroup{
		Imports: []*Import{
			{
				Namespace: namespace,
				Comment:   comment,
			},
		},
	}

	actual := model.String()

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

	model := &ImportGroup{
		Comment: importGroupComment,
		Imports: []*Import{
			{
				Alias:     alias,
				Namespace: namespace,
			},
		},
	}

	actual := model.String()

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

	model := &ImportGroup{
		Comment: importGroupComment,
		Imports: []*Import{
			{
				Namespace: namespace,
			},
		},
	}

	actual := model.String()

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

	model := &ImportGroup{
		Comment: importGroupComment,
		Imports: []*Import{
			{
				Alias:     alias,
				Namespace: namespace,
				Comment:   comment,
			},
		},
	}

	actual := model.String()

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

	model := &ImportGroup{
		Comment: importGroupComment,
		Imports: []*Import{
			{
				Namespace: namespace,
				Comment:   comment,
			},
		},
	}

	actual := model.String()

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

	model := &ImportGroup{
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

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestImportGroup_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ImportGroup{
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

	actual := model.Clone()

	ctrl.AssertEqual(model, actual)
	ctrl.AssertNotSame(model, actual)
	ctrl.AssertNotSame(model.Annotations[0], actual.(*ImportGroup).Annotations[0])
	ctrl.AssertNotSame(model.Imports[0], actual.(*ImportGroup).Imports[0])
	ctrl.AssertNotSame(model.Imports[0].Annotations[0], actual.(*ImportGroup).Imports[0].Annotations[0])
}

func TestImportGroup_Clone_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ImportGroup{}

	actual := model.Clone()

	ctrl.AssertEqual(model, actual)
	ctrl.AssertNotSame(model, actual)
}

func TestImportGroup_RenameImports_WithRenameAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	model := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     oldAlias,
				Namespace: "namespace",
			},
		},
	}

	expected := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     newAlias,
				Namespace: "namespace",
			},
		},
	}

	model.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(expected, model)
}

func TestImportGroup_RenameImports_WithNotRenameAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	model := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     "alias",
				Namespace: "namespace",
			},
		},
	}

	expected := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     "alias",
				Namespace: "namespace",
			},
		},
	}

	model.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(expected, model)
}

func TestImportGroup_RenameImports_WithInvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	model := &ImportGroup{}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestImportGroup_RenameImports_WithInvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "+invalid"

	model := &ImportGroup{}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
