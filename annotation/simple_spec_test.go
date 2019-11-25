package annotation

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestSimpleSpec_Validate_WithPackageNameAndTypeNameAndPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
		IsPointer:   true,
	}

	model.Validate()
}

func TestSimpleSpec_Validate_WithPackageNameAndTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	model.Validate()
}

func TestSimpleSpec_Validate_WithTypeNameAndPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &SimpleSpec{
		TypeName:  "typeName",
		IsPointer: true,
	}

	model.Validate()
}

func TestSimpleSpec_Validate_WithTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &SimpleSpec{
		TypeName: "typeName",
	}

	model.Validate()
}

func TestSimpleSpec_Validate_WithEmptyTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &SimpleSpec{
		TypeName: "",
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'TypeName' must be not empty"))
}

func TestSimpleSpec_Validate_WithInvalidTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &SimpleSpec{
		TypeName: "+invalid",
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'TypeName' must be valid identifier, actual value: '+invalid'"))
}

func TestSimpleSpec_Validate_WithInvalidPackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &SimpleSpec{
		PackageName: "+invalid",
		TypeName:    "typeName",
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'PackageName' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestSimpleSpec_String_WithPackageNameAndTypeNameAndPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "*packageName.typeName"

	model := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
		IsPointer:   true,
	}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestSimpleSpec_String_WithPackageNameAndTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "packageName.typeName"

	model := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestSimpleSpec_String_WithTypeNameAndPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "*typeName"

	model := &SimpleSpec{
		TypeName:  "typeName",
		IsPointer: true,
	}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestSimpleSpec_String_WithTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "typeName"

	model := &SimpleSpec{
		TypeName: "typeName",
	}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestSimpleSpec_Clone_WithPackageNameAndTypeNameAndPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
		IsPointer:   true,
	}

	actual := model.Clone()

	ctrl.AssertEqual(model, actual)
	ctrl.AssertNotSame(model, actual)
}

func TestSimpleSpec_Clone_WithPackageNameAndTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	actual := model.Clone()

	ctrl.AssertEqual(model, actual)
	ctrl.AssertNotSame(model, actual)
}

func TestSimpleSpec_Clone_WithTypeNameAndPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &SimpleSpec{
		TypeName:  "typeName",
		IsPointer: true,
	}

	actual := model.Clone()

	ctrl.AssertEqual(model, actual)
	ctrl.AssertNotSame(model, actual)
}

func TestSimpleSpec_Clone_WithTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &SimpleSpec{
		TypeName: "typeName",
	}

	actual := model.Clone()

	ctrl.AssertEqual(model, actual)
	ctrl.AssertNotSame(model, actual)
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

	model := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	actual := model.FetchImports(file)

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

	model := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	actual := model.FetchImports(file)

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

	model := &SimpleSpec{
		TypeName: "typeName",
	}

	actual := model.FetchImports(file)

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

	model := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	actual := model.FetchImports(file)

	ctrl.AssertEmpty(actual)
}

func TestSimpleSpec_RenameImports_WithRenamePackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	model := &SimpleSpec{
		PackageName: oldAlias,
		TypeName:    "typeName",
	}

	expected := &SimpleSpec{
		PackageName: newAlias,
		TypeName:    "typeName",
	}

	model.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(expected, model)
}

func TestSimpleSpec_RenameImports_WithNotRenamePackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	model := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	expected := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	model.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(expected, model)
}

func TestSimpleSpec_RenameImports_WithInvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	model := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestSimpleSpec_RenameImports_WithInvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "+invalid"

	model := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
