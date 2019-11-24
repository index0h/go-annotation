package model

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestMapSpec_Validate_WithSimpleSpecKeyAndSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithSimpleSpecKeyAndArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithSimpleSpecKeyAndMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
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

func TestMapSpec_Validate_WithSimpleSpecKeyAndStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &StructSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithSimpleSpecKeyAndInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &InterfaceSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithSimpleSpecKeyAndFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &FuncSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithArraySpecKeyAndSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithArraySpecKeyAndArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithArraySpecKeyAndMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
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

func TestMapSpec_Validate_WithArraySpecKeyAndStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Value: &StructSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithArraySpecKeyAndInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Value: &InterfaceSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithArraySpecKeyAndFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Value: &FuncSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithMapSpecKeyAndSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithMapSpecKeyAndArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithMapSpecKeyAndMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
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

func TestMapSpec_Validate_WithMapSpecKeyAndStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Value: &StructSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithMapSpecKeyAndInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Value: &InterfaceSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithMapSpecKeyAndFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Value: &FuncSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithStructSpecKeyAndSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &StructSpec{},
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithStructSpecKeyAndArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &StructSpec{},
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithStructSpecKeyAndMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &StructSpec{},
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

func TestMapSpec_Validate_WithStructSpecKeyAndStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key:   &StructSpec{},
		Value: &StructSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithStructSpecKeyAndInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &MapSpec{
			Key: &StructSpec{},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Value: &InterfaceSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithStructSpecKeyAndFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &MapSpec{
			Key: &StructSpec{},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Value: &FuncSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithInterfaceSpecKeyAndSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &InterfaceSpec{},
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithInterfaceSpecKeyAndArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &InterfaceSpec{},
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithInterfaceSpecKeyAndMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &InterfaceSpec{},
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

func TestMapSpec_Validate_WithInterfaceSpecKeyAndStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key:   &InterfaceSpec{},
		Value: &StructSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithInterfaceSpecKeyAndInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key:   &InterfaceSpec{},
		Value: &InterfaceSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithInterfaceSpecKeyAndFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key:   &InterfaceSpec{},
		Value: &FuncSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithFuncSpecKeyAndSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &FuncSpec{},
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithFuncSpecKeyAndArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &FuncSpec{},
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithFuncSpecKeyAndMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &FuncSpec{},
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

func TestMapSpec_Validate_WithFuncSpecKeyAndStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key:   &FuncSpec{},
		Value: &StructSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithFuncSpecKeyAndInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key:   &FuncSpec{},
		Value: &InterfaceSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithFuncSpecKeyAndFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key:   &FuncSpec{},
		Value: &FuncSpec{},
	}

	modelValue.Validate()
}

func TestMapSpec_Validate_WithNilKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: nil,
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Key' must be not nil"))
}

func TestMapSpec_Validate_WithNilValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: nil,
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Value' must be not nil"))
}

func TestMapSpec_Validate_WithInvalidSimpleSpecKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "+invalid",
		},
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'TypeName' must be valid identifier, actual value: '+invalid'"))
}

func TestMapSpec_Validate_WithInvalidArraySpecKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
			Length: -100,
		},
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'Length' must be greater than or equal to 0, actual value: -100"),
		)
}

func TestMapSpec_Validate_WithInvalidMapSpecKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &MapSpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Key' must be not nil"))
}

func TestMapSpec_Validate_WithInvalidStructSpecKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &StructSpec{
			Fields: []*Field{nil},
		},
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestMapSpec_Validate_WithInvalidInterfaceSpecKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &InterfaceSpec{
			Fields: []*Field{nil},
		},
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestMapSpec_Validate_WithInvalidFuncSpecKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &FuncSpec{
			Params: []*Field{nil},
		},
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Params[0]' must be not nil"))
}

func TestMapSpec_Validate_WithInvalidTypeKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: NewSpecMock(ctrl),
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Key' has invalid type: *model.SpecMock"))
}

func TestMapSpec_Validate_WithInvalidSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &SimpleSpec{
			TypeName: "+invalid",
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'TypeName' must be valid identifier, actual value: '+invalid'"))
}

func TestMapSpec_Validate_WithInvalidArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
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

func TestMapSpec_Validate_WithInvalidMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
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

func TestMapSpec_Validate_WithInvalidStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &StructSpec{
			Fields: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestMapSpec_Validate_WithInvalidInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &InterfaceSpec{
			Fields: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestMapSpec_Validate_WithInvalidFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: &FuncSpec{
			Params: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Params[0]' must be not nil"))
}

func TestMapSpec_Validate_WithInvalidTypeValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "typeName",
		},
		Value: NewSpecMock(ctrl),
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Value' has invalid type: *model.SpecMock"))
}

func TestMapSpec_String(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	keySpecString := "keySpecString"
	valueSpecString := "valueSpecString"
	expected := "map[keySpecString]valueSpecString"

	keySpec := NewSpecMock(ctrl)
	valueSpec := NewSpecMock(ctrl)

	modelValue := &MapSpec{
		Key:   keySpec,
		Value: valueSpec,
	}

	keySpec.
		EXPECT().
		String().
		Return(keySpecString)

	valueSpec.
		EXPECT().
		String().
		Return(valueSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestMapSpec_String_WithFuncSpecKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fieldSpecString := "fieldSpecString"
	valueSpecString := "valueSpecString"
	expected := "map[func (fieldSpecString)]valueSpecString"

	fieldSpec := NewSpecMock(ctrl)
	valueSpec := NewSpecMock(ctrl)

	modelValue := &MapSpec{
		Key: &FuncSpec{
			Params: []*Field{
				{
					Spec: fieldSpec,
				},
			},
		},
		Value: valueSpec,
	}

	fieldSpec.
		EXPECT().
		String().
		Return(fieldSpecString)

	valueSpec.
		EXPECT().
		String().
		Return(valueSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestMapSpec_String_WithFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	keySpecString := "keySpecString"
	fieldSpecString := "fieldSpecString"
	expected := "map[keySpecString]func (fieldSpecString)"

	keySpec := NewSpecMock(ctrl)
	fieldSpec := NewSpecMock(ctrl)

	modelValue := &MapSpec{
		Key: keySpec,
		Value: &FuncSpec{
			Params: []*Field{
				{
					Spec: fieldSpec,
				},
			},
		},
	}

	keySpec.
		EXPECT().
		String().
		Return(keySpecString)

	fieldSpec.
		EXPECT().
		String().
		Return(fieldSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestMapSpec_String_WithFuncSpecKeyAndFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	keyFieldSpecString := "keyFieldSpecString"
	valueFieldSpecString := "valueFieldSpecString"
	expected := "map[func (keyFieldSpecString)]func (valueFieldSpecString)"

	keyFieldSpec := NewSpecMock(ctrl)
	valueFieldSpec := NewSpecMock(ctrl)

	modelValue := &MapSpec{
		Key: &FuncSpec{
			Params: []*Field{
				{
					Spec: keyFieldSpec,
				},
			},
		},
		Value: &FuncSpec{
			Params: []*Field{
				{
					Spec: valueFieldSpec,
				},
			},
		},
	}

	keyFieldSpec.
		EXPECT().
		String().
		Return(keyFieldSpecString)

	valueFieldSpec.
		EXPECT().
		String().
		Return(valueFieldSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestMapSpec_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	keySpec := NewSpecMock(ctrl)
	valueSpec := NewSpecMock(ctrl)

	clonedKeySpec := NewSpecMock(ctrl)
	clonedValueSpec := NewSpecMock(ctrl)

	modelValue := &MapSpec{
		Key:   keySpec,
		Value: valueSpec,
	}

	keySpec.
		EXPECT().
		Clone().
		Return(clonedKeySpec)

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
	ctrl.AssertSame(clonedKeySpec, actual.(*MapSpec).Key)
	ctrl.AssertSame(clonedValueSpec, actual.(*MapSpec).Value)
}

func TestMapSpec_FetchImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{}

	expected := []*Import{
		{
			Alias:     "keyPackageName",
			Namespace: "keyNamespace",
		},
		{
			Alias:     "valuePackageName",
			Namespace: "valueNamespace",
		},
	}

	keySpec := NewSpecMock(ctrl)
	valueSpec := NewSpecMock(ctrl)

	modelValue := &MapSpec{
		Key:   keySpec,
		Value: valueSpec,
	}

	keySpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return([]*Import{expected[0]})

	valueSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return([]*Import{expected[1]})

	actual := modelValue.FetchImports(file)

	ctrl.AssertEqual(
		expected,
		actual,
		unit.IgnoreUnexportedOption{Value: *ctrl},
		unit.IgnoreUnexportedOption{Value: MockCallManager{}},
	)
}

func TestMapSpec_RenameImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	keySpec := NewSpecMock(ctrl)
	valueSpec := NewSpecMock(ctrl)

	modelValue := &MapSpec{
		Key:   keySpec,
		Value: valueSpec,
	}

	keySpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	valueSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	modelValue.RenameImports(oldAlias, newAlias)
}

func TestMapSpec_RenameImports_InvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	keySpec := NewSpecMock(ctrl)
	valueSpec := NewSpecMock(ctrl)

	modelValue := &MapSpec{
		Key:   keySpec,
		Value: valueSpec,
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestMapSpec_RenameImports_InvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "packageName"
	newAlias := "+invalid"

	keySpec := NewSpecMock(ctrl)
	valueSpec := NewSpecMock(ctrl)

	modelValue := &MapSpec{
		Key:   keySpec,
		Value: valueSpec,
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
