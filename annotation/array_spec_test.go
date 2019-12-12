package annotation

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestArraySpec_Validate_WithSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
	}

	model.Validate()
}

func TestArraySpec_Validate_WithSimpleSpecValueAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "typeName",
		},
		Length: "10",
	}

	model.Validate()
}

func TestArraySpec_Validate_WithArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	model.Validate()
}

func TestArraySpec_Validate_WithArraySpecValueAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value: &ArraySpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Length: "10",
	}

	model.Validate()
}

func TestArraySpec_Validate_WithMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	model.Validate()
}

func TestArraySpec_Validate_WithMapSpecValueAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value: &MapSpec{
			Key: &SimpleSpec{
				TypeName: "typeName",
			},
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
		Length: "10",
	}

	model.Validate()
}

func TestArraySpec_Validate_WithStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value: &StructSpec{},
	}

	model.Validate()
}

func TestArraySpec_Validate_WithStructSpecValueAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value:  &StructSpec{},
		Length: "10",
	}

	model.Validate()
}

func TestArraySpec_Validate_WithInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value: &InterfaceSpec{},
	}

	model.Validate()
}

func TestArraySpec_Validate_WithInterfaceSpecValueAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value:  &InterfaceSpec{},
		Length: "10",
	}

	model.Validate()
}

func TestArraySpec_Validate_WithFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value: &FuncSpec{},
	}

	model.Validate()
}

func TestArraySpec_Validate_WithFuncSpecValueAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value:  &FuncSpec{},
		Length: "10",
	}

	model.Validate()
}

func TestArraySpec_Validate_WithNilValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Value' must be not nil"))
}

func TestArraySpec_Validate_WithInvalidSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "+invalid",
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'TypeName' must be valid identifier, actual value: '+invalid'"))
}

func TestArraySpec_Validate_WithInvalidArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value: &ArraySpec{
			Value: nil,
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Value' must be not nil"))
}

func TestArraySpec_Validate_WithInvalidMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value: &MapSpec{
			Value: &SimpleSpec{
				TypeName: "typeName",
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Key' must be not nil"))
}

func TestArraySpec_Validate_WithInvalidStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value: &StructSpec{
			Fields: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestArraySpec_Validate_WithInvalidInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value: &InterfaceSpec{
			Fields: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestArraySpec_Validate_WithInvalidFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value: &FuncSpec{
			Params: []*Field{nil},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Params[0]' must be not nil"))
}

func TestArraySpec_Validate_WithInvalidTypeValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &ArraySpec{
		Value: NewSpecMock(ctrl),
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Value' has invalid type: %T", model.Value))
}

func TestArraySpec_String(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	valueSpecString := "valueSpecString"
	expected := "[]valueSpecString"

	valueSpec := NewSpecMock(ctrl)

	model := &ArraySpec{
		Value: valueSpec,
	}

	valueSpec.
		EXPECT().
		String().
		Return(valueSpecString)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestArraySpec_String_WithLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	valueSpecString := "valueSpecString"
	expected := "[10]valueSpecString"

	valueSpec := NewSpecMock(ctrl)

	model := &ArraySpec{
		Value:  valueSpec,
		Length: "10",
	}

	valueSpec.
		EXPECT().
		String().
		Return(valueSpecString)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestArraySpec_String_WithFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fieldSpecString := "fieldSpecString"
	expected := "[]func (fieldSpecString)"

	fieldSpec := NewSpecMock(ctrl)

	model := &ArraySpec{
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

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestArraySpec_String_WithLengthAndFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fieldSpecString := "fieldSpecString"
	expected := "[10]func (fieldSpecString)"

	fieldSpec := NewSpecMock(ctrl)

	model := &ArraySpec{
		Value: &FuncSpec{
			Params: []*Field{
				{
					Spec: fieldSpec,
				},
			},
		},
		Length: "10",
	}

	fieldSpec.
		EXPECT().
		String().
		Return(fieldSpecString)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestArraySpec_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	valueSpec := NewSpecMock(ctrl)
	clonedValueSpec := NewSpecMock(ctrl)

	model := &ArraySpec{
		Value: valueSpec,
	}

	valueSpec.
		EXPECT().
		Clone().
		Return(clonedValueSpec)

	actual := model.Clone()

	ctrl.AssertEqual(model, actual, unit.IgnoreUnexportedOption{Value: SpecMock{}})
	ctrl.AssertNotSame(model, actual)
	ctrl.AssertSame(clonedValueSpec, actual.(*ArraySpec).Value)
}

func TestArraySpec_EqualSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1ValueSpec := NewSpecMock(ctrl)
	model2ValueSpec := NewSpecMock(ctrl)

	model1 := &ArraySpec{
		Value:  model1ValueSpec,
		Length: "length",
	}

	model2 := &ArraySpec{
		Value:  model2ValueSpec,
		Length: "length",
	}

	model1ValueSpec.
		EXPECT().
		EqualSpec(ctrl.Same(model2ValueSpec)).
		Return(true)

	actual := model1.EqualSpec(model2)

	ctrl.AssertTrue(actual)
}

func TestArraySpec_EqualSpec_WithValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1ValueSpec := NewSpecMock(ctrl)
	model2ValueSpec := NewSpecMock(ctrl)

	model1 := &ArraySpec{
		Value:  model1ValueSpec,
		Length: "length",
	}

	model2 := &ArraySpec{
		Value:  model2ValueSpec,
		Length: "length",
	}

	model1ValueSpec.
		EXPECT().
		EqualSpec(ctrl.Same(model2ValueSpec)).
		Return(false)

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestArraySpec_EqualSpec_WithAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1ValueSpec := NewSpecMock(ctrl)

	model1 := &ArraySpec{
		Value:  model1ValueSpec,
		Length: "length",
	}

	model2 := "model2"

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestArraySpec_EqualSpec_WithLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1ValueSpec := NewSpecMock(ctrl)
	model2ValueSpec := NewSpecMock(ctrl)

	model1 := &ArraySpec{
		Value:  model1ValueSpec,
		Length: "length1",
	}

	model2 := &ArraySpec{
		Value:  model2ValueSpec,
		Length: "length2",
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
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

	model := &ArraySpec{
		Value: valueSpec,
	}

	valueSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(expected)

	actual := model.FetchImports(file)

	ctrl.AssertSame(expected, actual)
}

func TestArraySpec_FetchImports_WithLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := []*Import{
		{
			Alias:     "packageName",
			Namespace: "namespace",
		},
	}

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: expected,
			},
		},
	}

	valueSpec := NewSpecMock(ctrl)

	model := &ArraySpec{
		Value:  valueSpec,
		Length: "packageName.MyConst + 1",
	}

	valueSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(nil)

	actual := model.FetchImports(file)

	ctrl.AssertSame(expected, actual)
}

func TestArraySpec_RenameImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	valueSpec := NewSpecMock(ctrl)

	model := &ArraySpec{
		Value: valueSpec,
	}

	valueSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	model.RenameImports(oldAlias, newAlias)
}

func TestArraySpec_RenameImports_InvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	valueSpec := NewSpecMock(ctrl)

	model := &ArraySpec{
		Value: valueSpec,
	}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
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

	model := &ArraySpec{
		Value: valueSpec,
	}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
