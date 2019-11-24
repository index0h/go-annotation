package model

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestArraySpec_Validate_WithSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	modelValue.Validate()
}

func TestArraySpec_Validate_WithSimpleSpecValueAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
		Length: 10,
	}

	modelValue.Validate()
}

func TestArraySpec_Validate_WithSimpleSpecValueAndIsEllipsis(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
		IsEllipsis: true,
	}

	modelValue.Validate()
}

func TestArraySpec_Validate_WithArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	modelValue.Validate()
}

func TestArraySpec_Validate_WithArraySpecValueAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Length: 10,
	}

	modelValue.Validate()
}

func TestArraySpec_Validate_WithArraySpecValueAndIsEllipsis(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		IsEllipsis: true,
	}

	modelValue.Validate()
}

func TestArraySpec_Validate_WithMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &MapSpec{
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

func TestArraySpec_Validate_WithMapSpecValueAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Length: 10,
	}

	modelValue.Validate()
}

func TestArraySpec_Validate_WithMapSpecValueAndIsEllipsis(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		IsEllipsis: true,
	}

	modelValue.Validate()
}

func TestArraySpec_Validate_WithStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &StructSpec{},
	}

	modelValue.Validate()
}

func TestArraySpec_Validate_WithStructSpecValueAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value:  &StructSpec{},
		Length: 10,
	}

	modelValue.Validate()
}

func TestArraySpec_Validate_WithStructSpecValueAndIsEllipsis(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value:      &StructSpec{},
		IsEllipsis: true,
	}

	modelValue.Validate()
}

func TestArraySpec_Validate_WithInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &InterfaceSpec{},
	}

	modelValue.Validate()
}

func TestArraySpec_Validate_WithInterfaceSpecValueAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value:  &InterfaceSpec{},
		Length: 10,
	}

	modelValue.Validate()
}

func TestArraySpec_Validate_WithInterfaceSpecValueAndIsEllipsis(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value:      &InterfaceSpec{},
		IsEllipsis: true,
	}

	modelValue.Validate()
}

func TestArraySpec_Validate_WithFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &FuncSpec{},
	}

	modelValue.Validate()
}

func TestArraySpec_Validate_WithFuncSpecValueAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value:  &FuncSpec{},
		Length: 10,
	}

	modelValue.Validate()
}

func TestArraySpec_Validate_WithFuncSpecValueAndIsEllipsis(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value:      &FuncSpec{},
		IsEllipsis: true,
	}

	modelValue.Validate()
}

func TestArraySpec_Validate_WithNilValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Value' must be not nil"))
}

func TestArraySpec_Validate_WithNegativeLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "+invalid",
		},
		Length: -100,
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'Length' must be greater than or equal to 0, actual value: -100"),
		)
}

func TestArraySpec_Validate_WithPositiveLengthAndIsEllipsis(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "+invalid",
		},
		Length:     10,
		IsEllipsis: true,
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Array must have only length or only ellipsis"))
}

func TestArraySpec_Validate_WithInvalidSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "+invalid",
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'TypeName' must be valid identifier, actual value: '+invalid'"))
}

func TestArraySpec_Validate_WithInvalidArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &ArraySpec{
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

func TestArraySpec_Validate_WithInvalidMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &MapSpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Key' must be not nil"))
}

func TestArraySpec_Validate_WithInvalidStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &StructSpec{
			Fields: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestArraySpec_Validate_WithInvalidInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &InterfaceSpec{
			Fields: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestArraySpec_Validate_WithInvalidFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: &FuncSpec{
			Params: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Params[0]' must be not nil"))
}

func TestArraySpec_Validate_WithInvalidTypeValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &ArraySpec{
		Value: NewSpecMock(ctrl),
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Value' has invalid type: *model.SpecMock"))
}

func TestArraySpec_String(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	valueSpecString := "valueSpecString"
	expected := "[]valueSpecString"

	valueSpec := NewSpecMock(ctrl)

	modelValue := &ArraySpec{
		Value: valueSpec,
	}

	valueSpec.
		EXPECT().
		String().
		Return(valueSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestArraySpec_String_WithLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	valueSpecString := "valueSpecString"
	expected := "[10]valueSpecString"

	valueSpec := NewSpecMock(ctrl)

	modelValue := &ArraySpec{
		Value:  valueSpec,
		Length: 10,
	}

	valueSpec.
		EXPECT().
		String().
		Return(valueSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestArraySpec_String_WithIsEllipsis(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	valueSpecString := "valueSpecString"
	expected := "[...]valueSpecString"

	valueSpec := NewSpecMock(ctrl)

	modelValue := &ArraySpec{
		Value:      valueSpec,
		IsEllipsis: true,
	}

	valueSpec.
		EXPECT().
		String().
		Return(valueSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestArraySpec_String_WithFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fieldSpecString := "fieldSpecString"
	expected := "[]func (fieldSpecString)"

	fieldSpec := NewSpecMock(ctrl)

	modelValue := &ArraySpec{
		Value: &FuncSpec{
			Params: []*Field{
				{
					Spec: fieldSpec,
				},
			},
		},
	}

	fieldSpec.
		EXPECT().
		String().
		Return(fieldSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestArraySpec_String_WithLengthAndFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fieldSpecString := "fieldSpecString"
	expected := "[10]func (fieldSpecString)"

	fieldSpec := NewSpecMock(ctrl)

	modelValue := &ArraySpec{
		Value: &FuncSpec{
			Params: []*Field{
				{
					Spec: fieldSpec,
				},
			},
		},
		Length: 10,
	}

	fieldSpec.
		EXPECT().
		String().
		Return(fieldSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestArraySpec_String_WithIsEllipsisAndFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fieldSpecString := "fieldSpecString"
	expected := "[...]func (fieldSpecString)"

	fieldSpec := NewSpecMock(ctrl)

	modelValue := &ArraySpec{
		Value: &FuncSpec{
			Params: []*Field{
				{
					Spec: fieldSpec,
				},
			},
		},
		IsEllipsis: true,
	}

	fieldSpec.
		EXPECT().
		String().
		Return(fieldSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestArraySpec_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	valueSpec := NewSpecMock(ctrl)
	clonedValueSpec := NewSpecMock(ctrl)

	modelValue := &ArraySpec{
		Value: valueSpec,
	}

	valueSpec.
		EXPECT().
		Clone().
		Return(clonedValueSpec)

	actual := modelValue.Clone()

	ctrl.AssertEqual(
		modelValue,
		actual,
		unit.IgnoreUnexportedOption{Value: *ctrl},
		unit.IgnoreUnexportedOption{Value: MockCallManager{}},
	)
	ctrl.AssertNotSame(modelValue, actual)
	ctrl.AssertSame(clonedValueSpec, actual.(*ArraySpec).Value)
}

func TestArraySpec_FetchImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{}

	expected := []*Import{
		{
			Alias:     "packageName",
			Namespace: "namespace",
		},
	}

	valueSpec := NewSpecMock(ctrl)

	modelValue := &ArraySpec{
		Value: valueSpec,
	}

	valueSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(expected)

	actual := modelValue.FetchImports(file)

	ctrl.AssertSame(expected, actual)
}

func TestArraySpec_RenameImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	valueSpec := NewSpecMock(ctrl)

	modelValue := &ArraySpec{
		Value: valueSpec,
	}

	valueSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	modelValue.RenameImports(oldAlias, newAlias)
}

func TestArraySpec_RenameImports_InvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	valueSpec := NewSpecMock(ctrl)

	modelValue := &ArraySpec{
		Value: valueSpec,
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestArraySpec_RenameImports_InvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "packageName"
	newAlias := "+invalid"

	valueSpec := NewSpecMock(ctrl)

	modelValue := &ArraySpec{
		Value: valueSpec,
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
