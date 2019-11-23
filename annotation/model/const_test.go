package model

import (
	"go/scanner"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestConst_Validate(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Const{
		Name:    "constName",
		Value:   "iota",
		Comment: "const\ncomment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "constAnnotation",
			},
		},
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	modelValue.Validate()
}

func TestConst_Validate_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Const{
		Name: "constName",
	}

	modelValue.Validate()
}

func TestConst_Validate_WithEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Const{
		Name: "",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestConst_Validate_WithInvalidName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Const{
		Name: "+invalid",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be valid identifier, actual value: '+invalid'"))
}

func TestConst_Validate_WithPointerSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			IsPointer: true,
			TypeName:  "typeName",
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'Spec.(*model.SimpleSpec).IsPointer' must be 'false' for *model.Const"),
		)
}

func TestConst_Validate_WithInvalidValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Const{
		Name:  "constName",
		Value: "[invalid",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(ctrl.Type(scanner.ErrorList{}))
}

func TestConst_String_WithCommentAndSpecAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "constName"
	value := "value"
	comment := "const\ncomment"

	expected := `// const
// comment
const constName packageName.typeName = value
`

	modelValue := &Const{
		Name:    name,
		Value:   value,
		Comment: comment,
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestConst_String_WithCommentAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "constName"
	value := "value"
	comment := "const\ncomment"

	expected := `// const
// comment
const constName = value
`

	modelValue := &Const{
		Name:    name,
		Value:   value,
		Comment: comment,
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestConst_String_WithSpecAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "constName"
	value := "value"

	expected := `const constName packageName.typeName = value
`

	modelValue := &Const{
		Name:  name,
		Value: value,
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestConst_String_WithValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "constName"
	value := "value"

	expected := `const constName = value
`

	modelValue := &Const{
		Name:  name,
		Value: value,
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestConst_String_WithEmptyValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "constName"

	modelValue := &Const{
		Name: name,
	}

	ctrl.Subtest("").
		Call(modelValue.String).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'Value' must be not empty"),
		)
}

func TestConst_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Const{
		Name:    "constName",
		Value:   "iota",
		Comment: "const\ncomment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "constAnnotation",
			},
		},
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
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
	ctrl.AssertNotSame(modelValue.Annotations[0], actual.(*Const).Annotations[0])
}

func TestConst_Clone_WithoutFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Const{
		Name: "constName",
	}

	actual := modelValue.Clone()

	ctrl.AssertEqual(modelValue, actual)
	ctrl.AssertNotSame(modelValue, actual)
}

func TestConst_FetchImports_FoundImportByAlias(t *testing.T) {
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

	modelValue := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	actual := modelValue.FetchImports(file)

	ctrl.AssertEqual(expected, actual)
}

func TestConst_FetchImports_FoundImportByRealAlias(t *testing.T) {
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

	modelValue := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	actual := modelValue.FetchImports(file)

	ctrl.AssertEqual(expected, actual)
}

func TestConst_FetchImports_WithoutPackageNameAndNotFound(t *testing.T) {
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

	modelValue := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := modelValue.FetchImports(file)

	ctrl.AssertEmpty(actual)
}

func TestConst_FetchImports_NotFoundByAlias(t *testing.T) {
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

	modelValue := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	actual := modelValue.FetchImports(file)

	ctrl.AssertEmpty(actual)
}

func TestConst_RenameImports_WithRenamePackageNameAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	modelValue := &Const{
		Name:  "constName",
		Value: "(oldPackageName.value + 5) + iota",
		Spec: &SimpleSpec{
			PackageName: oldAlias,
			TypeName:    "typeName",
		},
	}

	modelExpected := &Const{
		Name:  "constName",
		Value: "(newPackageName.value + 5) + iota",
		Spec: &SimpleSpec{
			PackageName: newAlias,
			TypeName:    "typeName",
		},
	}

	modelValue.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(modelExpected, modelValue)
}

func TestConst_RenameImports_WithValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	modelValue := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	modelExpected := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	modelValue.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(modelExpected, modelValue)
}

func TestConst_RenameImports_WithNotRenamePackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	modelValue := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	modelExpected := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	modelValue.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(modelExpected, modelValue)
}

func TestConst_RenameImports_WithInvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	modelValue := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestConst_RenameImports_WithInvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "+invalid"

	modelValue := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
