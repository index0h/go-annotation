package model

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestSimpleSpec_Validate_WithPackageNameAndTypeNameAndPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
		IsPointer:   true,
	}

	modelValue.Validate()
}

func TestSimpleSpec_Validate_WithPackageNameAndTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	modelValue.Validate()
}

func TestSimpleSpec_Validate_WithTypeNameAndPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &SimpleSpec{
		TypeName:  "typeName",
		IsPointer: true,
	}

	modelValue.Validate()
}

func TestSimpleSpec_Validate_WithTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &SimpleSpec{
		TypeName: "typeName",
	}

	modelValue.Validate()
}

func TestSimpleSpec_Validate_WithEmptyTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &SimpleSpec{
		TypeName: "",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'TypeName' must be not empty"))
}

func TestSimpleSpec_Validate_WithInvalidTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &SimpleSpec{
		TypeName: "+invalid",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'TypeName' must be valid identifier, actual value: '+invalid'"))
}

func TestSimpleSpec_Validate_WithInvalidPackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &SimpleSpec{
		PackageName: "+invalid",
		TypeName:    "typeName",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'PackageName' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestSimpleSpec_String_WithPackageNameAndTypeNameAndPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "*packageName.typeName"

	modelValue := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
		IsPointer:   true,
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestSimpleSpec_String_WithPackageNameAndTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "packageName.typeName"

	modelValue := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestSimpleSpec_String_WithTypeNameAndPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "*typeName"

	modelValue := &SimpleSpec{
		TypeName:  "typeName",
		IsPointer: true,
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestSimpleSpec_String_WithTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "typeName"

	modelValue := &SimpleSpec{
		TypeName: "typeName",
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestSimpleSpec_Clone_WithPackageNameAndTypeNameAndPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
		IsPointer:   true,
	}

	actual := modelValue.Clone()

	ctrl.AssertEqual(modelValue, actual)
	ctrl.AssertNotSame(modelValue, actual)
}

func TestSimpleSpec_Clone_WithPackageNameAndTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	actual := modelValue.Clone()

	ctrl.AssertEqual(modelValue, actual)
	ctrl.AssertNotSame(modelValue, actual)
}

func TestSimpleSpec_Clone_WithTypeNameAndPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &SimpleSpec{
		TypeName:  "typeName",
		IsPointer: true,
	}

	actual := modelValue.Clone()

	ctrl.AssertEqual(modelValue, actual)
	ctrl.AssertNotSame(modelValue, actual)
}

func TestSimpleSpec_Clone_WithTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &SimpleSpec{
		TypeName: "typeName",
	}

	actual := modelValue.Clone()

	ctrl.AssertEqual(modelValue, actual)
	ctrl.AssertNotSame(modelValue, actual)
}

func TestSimpleSpec_FetchImports_FoundImportByAlias(t *testing.T) {
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

	modelValue := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	actual := modelValue.FetchImports(file)

	ctrl.AssertEqual(expected, actual)
}

func TestSimpleSpec_FetchImports_FoundImportByRealAlias(t *testing.T) {
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

	modelValue := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	actual := modelValue.FetchImports(file)

	ctrl.AssertEqual(expected, actual)
}

func TestSimpleSpec_FetchImports_WithoutPackageNameAndNotFound(t *testing.T) {
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

	modelValue := &SimpleSpec{
		TypeName: "typeName",
	}

	actual := modelValue.FetchImports(file)

	ctrl.AssertEmpty(actual)
}

func TestSimpleSpec_FetchImports_NotFoundByAlias(t *testing.T) {
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

	modelValue := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	actual := modelValue.FetchImports(file)

	ctrl.AssertEmpty(actual)
}

func TestSimpleSpec_RenameImports_WithRenamePackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	modelValue := &SimpleSpec{
		PackageName: oldAlias,
		TypeName:    "typeName",
	}

	modelExpected := &SimpleSpec{
		PackageName: newAlias,
		TypeName:    "typeName",
	}

	modelValue.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(modelExpected, modelValue)
}

func TestSimpleSpec_RenameImports_WithNotRenamePackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	modelValue := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	modelExpected := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	modelValue.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(modelExpected, modelValue)
}

func TestSimpleSpec_RenameImports_WithInvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	modelValue := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestSimpleSpec_RenameImports_WithInvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "+invalid"

	modelValue := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
