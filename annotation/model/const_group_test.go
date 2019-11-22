package model

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestConstGroup_Validate(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ConstGroup{
		Comment: "comment",
		Annotations: []interface{}{
			&SimpleSpec{
				TypeName: "typeName",
			},
		},
		Consts: []*Const{
			{
				Name: "constName",
			},
		},
	}

	modelValue.Validate()
}

func TestConstGroup_Validate_WithNilConst(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ConstGroup{
		Comment: "comment",
		Annotations: []interface{}{
			&SimpleSpec{
				TypeName: "typeName",
			},
		},
		Consts: []*Const{nil},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Consts[0]' must be not nil"))
}

func TestConstGroup_Validate_WithInvalidConst(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ConstGroup{
		Comment: "comment",
		Annotations: []interface{}{
			&SimpleSpec{
				TypeName: "typeName",
			},
		},
		Consts: []*Const{
			{},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestConstGroup_String_WithOneConstAndSpecAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "constName"
	value := "value"

	expected := `const constName packageName.typeName = value
`

	modelValue := &ConstGroup{
		Consts: []*Const{
			{
				Name:  name,
				Value: value,
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestConstGroup_String_WithOneConstAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "constName"
	value := "value"

	expected := `const constName = value
`

	modelValue := &ConstGroup{
		Consts: []*Const{
			{
				Name:  name,
				Value: value,
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestConstGroup_String_WithOneConstAndEmptyValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "constName"

	modelValue := &ConstGroup{
		Consts: []*Const{
			{
				Name: name,
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.String).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'Value' must be not empty"),
		)
}

func TestConstGroup_String_WithOneConstAndCommentAndSpecAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "constName"
	value := "value"
	comment := "const\ncomment"

	expected := `const (
// const
// comment
constName packageName.typeName = value
)
`

	modelValue := &ConstGroup{
		Consts: []*Const{
			{
				Name:    name,
				Value:   value,
				Comment: comment,
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestConstGroup_String_WithOneConstAndCommentAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "constName"
	value := "value"
	comment := "const\ncomment"

	expected := `const (
// const
// comment
constName = value
)
`

	modelValue := &ConstGroup{
		Consts: []*Const{
			{
				Name:    name,
				Value:   value,
				Comment: comment,
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestConstGroup_String_WithOneConstAndConstGroupCommentAndSpecAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	constGroupComment := "constGroup\ncomment"
	name := "constName"
	value := "value"

	expected := `// constGroup
// comment
const constName packageName.typeName = value
`

	modelValue := &ConstGroup{
		Comment: constGroupComment,
		Consts: []*Const{
			{
				Name:  name,
				Value: value,
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestConstGroup_String_WithOneConstAndConstGroupCommentAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	constGroupComment := "constGroup\ncomment"
	name := "constName"
	value := "value"

	expected := `// constGroup
// comment
const constName = value
`

	modelValue := &ConstGroup{
		Comment: constGroupComment,
		Consts: []*Const{
			{
				Name:  name,
				Value: value,
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestConstGroup_String_WithOneConstAndConstGroupCommentAndCommentAndSpecAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	constGroupComment := "constGroup\ncomment"
	name := "constName"
	value := "value"
	comment := "const\ncomment"

	expected := `// constGroup
// comment
const (
// const
// comment
constName packageName.typeName = value
)
`

	modelValue := &ConstGroup{
		Comment: constGroupComment,
		Consts: []*Const{
			{
				Name:    name,
				Value:   value,
				Comment: comment,
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestConstGroup_String_WithOneConstAndConstGroupCommentAndCommentAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	constGroupComment := "constGroup\ncomment"
	name := "constName"
	value := "value"
	comment := "const\ncomment"

	expected := `// constGroup
// comment
const (
// const
// comment
constName = value
)
`

	modelValue := &ConstGroup{
		Comment: constGroupComment,
		Consts: []*Const{
			{
				Name:    name,
				Value:   value,
				Comment: comment,
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestConstGroup_String_WithMultipleConsts(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	constGroupComment := "constGroup\ncomment"
	name1 := "const1Name"
	value := "value"
	name2 := "const2Name"

	expected := `// constGroup
// comment
const (
const1Name = value
const2Name
)
`

	modelValue := &ConstGroup{
		Comment: constGroupComment,
		Consts: []*Const{
			{
				Name:  name1,
				Value: value,
			},
			{
				Name: name2,
			},
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestConstGroup_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ConstGroup{
		Comment: "constGroupComment",
		Annotations: []interface{}{
			&SimpleSpec{
				TypeName: "constGroupAnnotation",
			},
		},
		Consts: []*Const{
			{
				Name:    "const1Name",
				Value:   "iota",
				Comment: "const1\ncomment",
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "const1Annotation",
					},
				},
				Spec: &SimpleSpec{
					PackageName: "package1Name",
					TypeName:    "type1Name",
				},
			},
			{
				Name:    "const2Name",
				Value:   "iota",
				Comment: "const2\ncomment",
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "const2Annotation",
					},
				},
				Spec: &SimpleSpec{
					PackageName: "package2Name",
					TypeName:    "type2Name",
				},
			},
		},
	}

	actual := modelValue.Clone()

	ctrl.AssertEqual(
		modelValue,
		actual,
		unit.IgnoreUnexportedOption{Value: *ctrl},
		unit.IgnoreUnexportedOption{Value: MockCallManager{}},
	)
	ctrl.AssertNotSame(modelValue, actual)
	ctrl.AssertNotSame(modelValue.Annotations[0], actual.(*ConstGroup).Annotations[0])
	ctrl.AssertNotSame(modelValue.Consts[0], actual.(*ConstGroup).Consts[0])
	ctrl.AssertNotSame(modelValue.Consts[0].Annotations[0], actual.(*ConstGroup).Consts[0].Annotations[0])
	ctrl.AssertNotSame(modelValue.Consts[1], actual.(*ConstGroup).Consts[1])
	ctrl.AssertNotSame(modelValue.Consts[1].Annotations[0], actual.(*ConstGroup).Consts[1].Annotations[0])
}

func TestConstGroup_Clone_WithoutFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ConstGroup{}

	actual := modelValue.Clone()

	ctrl.AssertEqual(modelValue, actual)
	ctrl.AssertNotSame(modelValue, actual)
}

func TestConstGroup_FetchImports_FoundImportByAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expectedImport := &Import{
		Alias:     "packageName",
		Namespace: "namespace",
	}

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{expectedImport},
			},
		},
	}

	expected := []*Import{expectedImport}

	modelValue := &ConstGroup{
		Consts: []*Const{
			{
				Name: "constName",
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
			},
		},
	}

	actual := modelValue.FetchImports(file)

	ctrl.AssertEqual(expected, actual)
}

func TestConstGroup_FetchImports_FoundImportByRealAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expectedImport := &Import{
		Namespace: "namespace/packageName",
	}

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{expectedImport},
			},
		},
	}

	expected := []*Import{expectedImport}

	modelValue := &ConstGroup{
		Consts: []*Const{
			{
				Name: "constName",
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
			},
		},
	}

	actual := modelValue.FetchImports(file)

	ctrl.AssertEqual(expected, actual)
}

func TestConstGroup_FetchImports_WithoutPackageNameAndNotFound(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "alias",
						Namespace: "namespace",
					},
				},
			},
		},
	}

	modelValue := &ConstGroup{
		Consts: []*Const{
			{
				Name: "constName",
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	actual := modelValue.FetchImports(file)

	ctrl.AssertEmpty(actual)
}

func TestConstGroup_FetchImports_NotFoundByAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "alias",
						Namespace: "namespace",
					},
				},
			},
		},
	}

	modelValue := &ConstGroup{
		Consts: []*Const{
			{
				Name: "constName",
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
			},
		},
	}

	actual := modelValue.FetchImports(file)

	ctrl.AssertEmpty(actual)
}

func TestConstGroup_RenameImports_WithRenamePackageNameAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	modelValue := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "constName",
				Value: "(oldPackageName.value + 5) + iota",
				Spec: &SimpleSpec{
					PackageName: oldAlias,
					TypeName:    "typeName",
				},
			},
		},
	}

	modelExpected := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "constName",
				Value: "(newPackageName.value + 5) + iota",
				Spec: &SimpleSpec{
					PackageName: newAlias,
					TypeName:    "typeName",
				},
			},
		},
	}

	modelValue.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(modelExpected, modelValue)
}

func TestConstGroup_RenameImports_WithValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	modelValue := &ConstGroup{
		Consts: []*Const{
			{
				Name: "constName",
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
			},
		},
	}

	modelExpected := &ConstGroup{
		Consts: []*Const{
			{
				Name: "constName",
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
			},
		},
	}

	modelValue.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(modelExpected, modelValue)
}

func TestConstGroup_RenameImports_WithNotRenamePackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	modelValue := &ConstGroup{
		Consts: []*Const{
			{
				Name: "constName",
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
			},
		},
	}

	modelExpected := &ConstGroup{
		Consts: []*Const{
			{
				Name: "constName",
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
			},
		},
	}

	modelValue.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(modelExpected, modelValue)
}

func TestConstGroup_RenameImports_WithInvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	modelValue := &ConstGroup{
		Consts: []*Const{
			{
				Name: "constName",
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestConstGroup_RenameImports_WithInvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "+invalid"

	modelValue := &ConstGroup{
		Consts: []*Const{
			{
				Name: "constName",
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}