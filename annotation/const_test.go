package annotation

import (
	"go/scanner"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestConst_Validate(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Const{
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

	model.Validate()
}

func TestConst_Validate_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Const{
		Name: "constName",
	}

	model.Validate()
}

func TestConst_Validate_WithEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Const{
		Name: "",
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestConst_Validate_WithInvalidName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Const{
		Name: "+invalid",
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be valid identifier, actual value: '+invalid'"))
}

func TestConst_Validate_WithPointerSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			IsPointer: true,
			TypeName:  "typeName",
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'Spec.(%T).IsPointer' must be 'false' for %T", &SimpleSpec{}, model),
		)
}

func TestConst_Validate_WithInvalidValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Const{
		Name:  "constName",
		Value: "[invalid",
	}

	ctrl.Subtest("").
		Call(model.Validate).
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

	model := &Const{
		Name:    name,
		Value:   value,
		Comment: comment,
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	actual := model.String()

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

	model := &Const{
		Name:    name,
		Value:   value,
		Comment: comment,
	}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestConst_String_WithSpecAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "constName"
	value := "value"

	expected := `const constName packageName.typeName = value
`

	model := &Const{
		Name:  name,
		Value: value,
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestConst_String_WithValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "constName"
	value := "value"

	expected := `const constName = value
`

	model := &Const{
		Name:  name,
		Value: value,
	}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestConst_String_WithEmptyValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "constName"

	model := &Const{
		Name: name,
	}

	ctrl.Subtest("").
		Call(model.String).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'Value' must be not empty"),
		)
}

func TestConst_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Const{
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

	actual := model.Clone()

	ctrl.AssertEqual(model, actual, unit.IgnoreUnexportedOption{Value: SpecMock{}})
	ctrl.AssertNotSame(model, actual)
	ctrl.AssertNotSame(model.Annotations[0], actual.(*Const).Annotations[0])
}

func TestConst_Clone_WithoutFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &Const{
		Name: "constName",
	}

	actual := model.Clone()

	ctrl.AssertEqual(model, actual)
	ctrl.AssertNotSame(model, actual)
}

func TestConst_EqualSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1 := &Const{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: "value",
	}

	model2 := &Const{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: "value",
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertTrue(actual)
}

func TestConst_EqualSpec_WithoutValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1 := &Const{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	model2 := &Const{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertTrue(actual)
}

func TestConst_EqualSpec_WithoutSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1 := &Const{
		Name:  "name",
		Value: "value",
	}

	model2 := &Const{
		Name:  "name",
		Value: "value",
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertTrue(actual)
}

func TestConst_EqualSpec_WithAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1 := &Const{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	model2 := "model2"

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestConst_EqualSpec_WithName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1 := &Const{
		Name: "name1",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: "value",
	}

	model2 := &Const{
		Name: "name2",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: "value",
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestConst_EqualSpec_WithSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1 := &Const{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName1",
		},
		Value: "value",
	}

	model2 := &Const{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName2",
		},
		Value: "value",
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestConst_EqualSpec_WithValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1 := &Const{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: "value1",
	}

	model2 := &Const{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: "value2",
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
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

	model := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	actual := model.FetchImports(file)

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

	model := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	actual := model.FetchImports(file)

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

	model := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	actual := model.FetchImports(file)

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

	model := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	actual := model.FetchImports(file)

	ctrl.AssertEmpty(actual)
}

func TestConst_RenameImports_WithRenamePackageNameAndValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	model := &Const{
		Name:  "constName",
		Value: "(oldPackageName.value + 5) + iota",
		Spec: &SimpleSpec{
			PackageName: oldAlias,
			TypeName:    "typeName",
		},
	}

	expected := &Const{
		Name:  "constName",
		Value: "(newPackageName.value + 5) + iota",
		Spec: &SimpleSpec{
			PackageName: newAlias,
			TypeName:    "typeName",
		},
	}

	model.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(expected, model)
}

func TestConst_RenameImports_WithValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	model := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	expected := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	model.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(expected, model)
}

func TestConst_RenameImports_WithNotRenamePackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	model := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	expected := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	model.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(expected, model)
}

func TestConst_RenameImports_WithInvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	model := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestConst_RenameImports_WithInvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "+invalid"

	model := &Const{
		Name: "constName",
		Spec: &SimpleSpec{
			PackageName: "packageName",
			TypeName:    "typeName",
		},
	}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
