package model

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestType_Validate_WithSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Type{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	modelValue.Validate()
}

func TestType_Validate_WithArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Type{
		Name: "name",
		Spec: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	modelValue.Validate()
}

func TestType_Validate_WithMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Type{
		Name: "name",
		Spec: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	modelValue.Validate()
}

func TestType_Validate_WithStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Type{
		Name: "name",
		Spec: &StructSpec{},
	}

	modelValue.Validate()
}

func TestType_Validate_WithInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Type{
		Name: "name",
		Spec: &InterfaceSpec{},
	}

	modelValue.Validate()
}

func TestType_Validate_WithFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Type{
		Name: "name",
		Spec: &FuncSpec{},
	}

	modelValue.Validate()
}

func TestType_Validate_WithEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Type{
		Spec: &FuncSpec{},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestType_Validate_WithInvalidName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Type{
		Name: "+invalid",
		Spec: &FuncSpec{},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be valid identifier, actual value: '+invalid'"))
}

func TestType_Validate_WithNilSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Type{
		Name: "name",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Spec' must be not nil"))
}

func TestType_Validate_WithInvalidSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Type{
		Name: "name",
		Spec: &SimpleSpec{
			TypeName: "+invalid",
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'TypeName' must be valid identifier, actual value: '+invalid'"))
}

func TestType_Validate_WithInvalidArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Type{
		Name: "name",
		Spec: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
			Length: -100,
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'Length' must be greater than or equal to 0, actual value: -100"),
		)
}

func TestType_Validate_WithInvalidMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Type{
		Name: "name",
		Spec: &MapSpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Key' must be not nil"))
}

func TestType_Validate_WithInvalidStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Type{
		Name: "name",
		Spec: &StructSpec{
			Fields: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestType_Validate_WithInvalidInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Type{
		Name: "name",
		Spec: &InterfaceSpec{
			Methods: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Methods[0]' must be not nil"))
}

func TestType_Validate_WithInvalidFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Type{
		Name: "name",
		Spec: &FuncSpec{
			Params: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Params[0]' must be not nil"))
}

func TestType_Validate_WithInvalidTypeValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &Type{
		Name: "name",
		Spec: NewSpecMock(ctrl),
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Spec' has invalid type: *model.SpecMock"))
}

func TestType_String_WithComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "name"
	comment := "type\ncomment"
	specSpecString := "specSpecString"
	expected := `// type
// comment
type name specSpecString
`

	specSpec := NewSpecMock(ctrl)

	modelValue := &Type{
		Name:    name,
		Comment: comment,
		Annotations: []interface{}{
			&SimpleSpec{
				TypeName: "typeAnnotation",
			},
		},
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		String().
		Return(specSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestType_String(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "name"
	specSpecString := "specSpecString"
	expected := `type name specSpecString
`

	specSpec := NewSpecMock(ctrl)

	modelValue := &Type{
		Name: name,
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		String().
		Return(specSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestType_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	specSpec := NewSpecMock(ctrl)
	clonedSpecSpec := NewSpecMock(ctrl)

	modelValue := &Type{
		Name:    "name",
		Comment: "comment",
		Annotations: []interface{}{
			&SimpleSpec{
				TypeName: "typeAnnotation",
			},
		},
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		Clone().
		Return(clonedSpecSpec)

	actual := modelValue.Clone()

	ctrl.AssertEqual(
		modelValue,
		actual,
		unit.IgnoreUnexportedOption{Value: *ctrl},
		unit.IgnoreUnexportedOption{Value: MockCallManager{}},
	)
	ctrl.AssertNotSame(modelValue, actual)
	ctrl.AssertSame(clonedSpecSpec, actual.(*Type).Spec)
	ctrl.AssertNotSame(modelValue.Annotations[0], actual.(*Type).Annotations[0])
}

func TestType_Clone_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	specSpec := NewSpecMock(ctrl)
	clonedSpecSpec := NewSpecMock(ctrl)

	modelValue := &Type{
		Name: "name",
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		Clone().
		Return(clonedSpecSpec)

	actual := modelValue.Clone()

	ctrl.AssertEqual(
		modelValue,
		actual,
		unit.IgnoreUnexportedOption{Value: *ctrl},
		unit.IgnoreUnexportedOption{Value: MockCallManager{}},
	)
	ctrl.AssertNotSame(modelValue, actual)
	ctrl.AssertSame(clonedSpecSpec, actual.(*Type).Spec)
}

func TestType_FetchImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{}

	expected := []*Import{
		{
			Alias:     "packageName",
			Namespace: "namespace",
		},
	}

	specSpec := NewSpecMock(ctrl)

	modelValue := &Type{
		Name: "name",
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(expected)

	actual := modelValue.FetchImports(file)

	ctrl.AssertSame(expected, actual)
}

func TestType_RenameImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	specSpec := NewSpecMock(ctrl)

	modelValue := &Type{
		Name: "name",
		Spec: specSpec,
	}

	specSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	modelValue.RenameImports(oldAlias, newAlias)
}

func TestType_RenameImports_InvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	specSpec := NewSpecMock(ctrl)

	modelValue := &Type{
		Name: "name",
		Spec: specSpec,
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestType_RenameImports_InvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "packageName"
	newAlias := "+invalid"

	specSpec := NewSpecMock(ctrl)

	modelValue := &Type{
		Name: "name",
		Spec: specSpec,
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
