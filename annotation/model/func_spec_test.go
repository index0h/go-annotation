package model

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestFuncSpec_Validate(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{}

	modelValue.Validate()
}

func TestFuncSpec_Validate_WithSimpleSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
		Results: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	modelValue.Validate()
}

func TestFuncSpec_Validate_WithArraySpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
		Results: []*Field{
			{
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	modelValue.Validate()
}

func TestFuncSpec_Validate_WithArraySpecAndVariadic(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		IsVariadic: true,
		Params: []*Field{
			{
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
		Results: []*Field{
			{
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	modelValue.Validate()
}

func TestFuncSpec_Validate_WithMapSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: &MapSpec{
					Key: &SimpleSpec{
						TypeName: "typeName",
					},
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
		Results: []*Field{
			{
				Spec: &MapSpec{
					Key: &SimpleSpec{
						TypeName: "typeName",
					},
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	modelValue.Validate()
}

func TestFuncSpec_Validate_WithStructSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: &StructSpec{},
			},
		},
		Results: []*Field{
			{
				Spec: &StructSpec{},
			},
		},
	}

	modelValue.Validate()
}

func TestFuncSpec_Validate_WithInterfaceSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: &InterfaceSpec{},
			},
		},
		Results: []*Field{
			{
				Spec: &InterfaceSpec{},
			},
		},
	}

	modelValue.Validate()
}

func TestFuncSpec_Validate_WithFuncSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: &FuncSpec{},
			},
		},
		Results: []*Field{
			{
				Spec: &FuncSpec{},
			},
		},
	}

	modelValue.Validate()
}

func TestFuncSpec_Validate_WithInvalidNameParam(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Name: "+invalid",
				Spec: NewSpecMock(ctrl),
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be valid identifier, actual value: '+invalid'"))
}

func TestFuncSpec_Validate_WithInvalidNameResult(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Results: []*Field{
			{
				Name: "+invalid",
				Spec: NewSpecMock(ctrl),
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be valid identifier, actual value: '+invalid'"))
}

func TestFuncSpec_Validate_WithNilParam(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Params: []*Field{
			nil,
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Params[0]' must be not nil"))
}

func TestFuncSpec_Validate_WithNilResult(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Results: []*Field{
			nil,
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Results[0]' must be not nil"))
}

func TestFuncSpec_Validate_WithInvalidSimpleSpecParamSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "+invalid",
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'TypeName' must be valid identifier, actual value: '+invalid'"))
}

func TestFuncSpec_Validate_WithInvalidSimpleSpecResultSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Results: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "+invalid",
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'TypeName' must be valid identifier, actual value: '+invalid'"))
}

func TestFuncSpec_Validate_WithInvalidArraySpecParamSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: &ArraySpec{},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Value' must be not nil"))
}
func TestFuncSpec_Validate_WithInvalidArraySpecResultSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Results: []*Field{
			{
				Spec: &ArraySpec{},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Value' must be not nil"))
}

func TestFuncSpec_Validate_WithInvalidMapSpecParamSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: &MapSpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Key' must be not nil"))
}

func TestFuncSpec_Validate_WithInvalidMapSpecResultSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Results: []*Field{
			{
				Spec: &MapSpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Key' must be not nil"))
}

func TestFuncSpec_Validate_WithInvalidStructSpecParamSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: &StructSpec{
					Fields: []*Field{nil},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestFuncSpec_Validate_WithInvalidStructSpecResultSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Results: []*Field{
			{
				Spec: &StructSpec{
					Fields: []*Field{nil},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestFuncSpec_Validate_WithInvalidInterfaceSpecParamSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: &InterfaceSpec{
					Fields: []*Field{nil},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestFuncSpec_Validate_WithInvalidInterfaceSpecResultSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Results: []*Field{
			{
				Spec: &InterfaceSpec{
					Fields: []*Field{nil},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestFuncSpec_Validate_WithInvalidFuncSpecParamSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: &FuncSpec{
					Params: []*Field{nil},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Params[0]' must be not nil"))
}

func TestFuncSpec_Validate_WithInvalidFuncSpecResultSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Results: []*Field{
			{
				Spec: &FuncSpec{
					Params: []*Field{nil},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Params[0]' must be not nil"))
}

func TestFuncSpec_Validate_WithInvalidTypeParamSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: NewSpecMock(ctrl),
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Spec' has invalid type: *model.SpecMock"))
}

func TestFuncSpec_Validate_WithEmptyParamsAndIsVariadic(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		IsVariadic: true,
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Params' must be not empty for variadic *model.FuncSpec"))
}

func TestFuncSpec_Validate_WithInvalidTypeParamSpecAndIsVariadic(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		IsVariadic: true,
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'Params[0].Spec' has invalid type for variadic *model.FuncSpec"),
		)
}

func TestFuncSpec_Validate_WithInvalidTypeResultSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Results: []*Field{
			{
				Spec: NewSpecMock(ctrl),
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Spec' has invalid type: *model.SpecMock"))
}

func TestFuncSpec_Validate_WithOneResultWithNameAndSecondResultWithoutName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Results: []*Field{
			{
				Name: "name",
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'Results' must have all fields with names or all without names"),
		)
}

func TestFuncSpec_String_WithParamCommentAndParamNameAndResultCommentAndResultNameAndIsVariadic(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	funcSpecParamComment := "funcSpecParamComment\nhere"
	funcSpecParamName := "funcSpecParamName"
	funcSpecParamSpecString := "funcSpecParamSpecString"
	funcSpecResultComment := "funcSpecResultComment\nhere"
	funcSpecResultName := "funcSpecResultName"
	funcSpecResultString := "funcSpecResultString"
	expected := `(
// funcSpecParamComment
// here
funcSpecParamName ...funcSpecParamSpecString) (
// funcSpecResultComment
// here
funcSpecResultName funcSpecResultString)`

	funcSpecParamSpec := NewSpecMock(ctrl)
	funcSpecResult := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		IsVariadic: true,
		Params: []*Field{
			{
				Name:    funcSpecParamName,
				Comment: funcSpecParamComment,
				Spec: &ArraySpec{
					Value: funcSpecParamSpec,
				},
			},
		},
		Results: []*Field{
			{
				Name:    funcSpecResultName,
				Comment: funcSpecResultComment,
				Spec:    funcSpecResult,
			},
		},
	}

	funcSpecParamSpec.
		EXPECT().
		String().
		Return(funcSpecParamSpecString)

	funcSpecResult.
		EXPECT().
		String().
		Return(funcSpecResultString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_WithParamCommentAndParamNameAndResultCommentAndResultName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	funcSpecParamComment := "funcSpecParamComment\nhere"
	funcSpecParamName := "funcSpecParamName"
	funcSpecParamSpecString := "funcSpecParamSpecString"
	funcSpecResultComment := "funcSpecResultComment\nhere"
	funcSpecResultName := "funcSpecResultName"
	funcSpecResultString := "funcSpecResultString"
	expected := `(
// funcSpecParamComment
// here
funcSpecParamName funcSpecParamSpecString) (
// funcSpecResultComment
// here
funcSpecResultName funcSpecResultString)`

	funcSpecParamSpec := NewSpecMock(ctrl)
	funcSpecResult := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Name:    funcSpecParamName,
				Comment: funcSpecParamComment,
				Spec:    funcSpecParamSpec,
			},
		},
		Results: []*Field{
			{
				Name:    funcSpecResultName,
				Comment: funcSpecResultComment,
				Spec:    funcSpecResult,
			},
		},
	}

	funcSpecParamSpec.
		EXPECT().
		String().
		Return(funcSpecParamSpecString)

	funcSpecResult.
		EXPECT().
		String().
		Return(funcSpecResultString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_WithParamNameAndResultName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	funcSpecParamName := "funcSpecParamName"
	funcSpecParamSpecString := "funcSpecParamSpecString"
	funcSpecResultName := "funcSpecResultName"
	funcSpecResultString := "funcSpecResultString"
	expected := `(funcSpecParamName funcSpecParamSpecString) (funcSpecResultName funcSpecResultString)`

	funcSpecParamSpec := NewSpecMock(ctrl)
	funcSpecResult := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Name: funcSpecParamName,
				Spec: funcSpecParamSpec,
			},
		},
		Results: []*Field{
			{
				Name: funcSpecResultName,
				Spec: funcSpecResult,
			},
		},
	}

	funcSpecParamSpec.
		EXPECT().
		String().
		Return(funcSpecParamSpecString)

	funcSpecResult.
		EXPECT().
		String().
		Return(funcSpecResultString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	funcSpecParamSpecString := "funcSpecParamSpecString"
	funcSpecResultString := "funcSpecResultString"
	expected := `(funcSpecParamSpecString) (funcSpecResultString)`

	funcSpecParamSpec := NewSpecMock(ctrl)
	funcSpecResult := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: funcSpecParamSpec,
			},
		},
		Results: []*Field{
			{
				Spec: funcSpecResult,
			},
		},
	}

	funcSpecParamSpec.
		EXPECT().
		String().
		Return(funcSpecParamSpecString)

	funcSpecResult.
		EXPECT().
		String().
		Return(funcSpecResultString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_WithParamCommentAndParamNameAndResultCommentAndResultNameAndBothFuncSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	funcSpecParamComment := "funcSpecParamComment\nhere"
	funcSpecParamName := "funcSpecParamName"
	funcSpecParamSpecString := "funcSpecParamSpecString"
	funcSpecResultComment := "funcSpecResultComment\nhere"
	funcSpecResultName := "funcSpecResultName"
	funcSpecResultString := "funcSpecResultString"
	expected := `(
// funcSpecParamComment
// here
funcSpecParamName func (funcSpecParamSpecString)) (
// funcSpecResultComment
// here
funcSpecResultName func (funcSpecResultString))`

	funcSpecParamSpec := NewSpecMock(ctrl)
	funcSpecResult := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Name:    funcSpecParamName,
				Comment: funcSpecParamComment,
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: funcSpecParamSpec,
						},
					},
				},
			},
		},
		Results: []*Field{
			{
				Name:    funcSpecResultName,
				Comment: funcSpecResultComment,
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: funcSpecResult,
						},
					},
				},
			},
		},
	}

	funcSpecParamSpec.
		EXPECT().
		String().
		Return(funcSpecParamSpecString)

	funcSpecResult.
		EXPECT().
		String().
		Return(funcSpecResultString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_WithParamNameAndFuncSpecParamSpecAndResultNameAndFuncSpecResultSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	funcSpecParamName := "funcSpecParamName"
	funcSpecParamSpecString := "funcSpecParamSpecString"
	funcSpecResultName := "funcSpecResultName"
	funcSpecResultString := "funcSpecResultString"
	expected := `(funcSpecParamName func (funcSpecParamSpecString)) (funcSpecResultName func (funcSpecResultString))`

	funcSpecParamSpec := NewSpecMock(ctrl)
	funcSpecResult := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Name: funcSpecParamName,
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: funcSpecParamSpec,
						},
					},
				},
			},
		},
		Results: []*Field{
			{
				Name: funcSpecResultName,
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: funcSpecResult,
						},
					},
				},
			},
		},
	}

	funcSpecParamSpec.
		EXPECT().
		String().
		Return(funcSpecParamSpecString)

	funcSpecResult.
		EXPECT().
		String().
		Return(funcSpecResultString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_WithFuncSpecParamSpecAndFuncSpecResultSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	funcSpecParamSpecString := "funcSpecParamSpecString"
	funcSpecResultString := "funcSpecResultString"
	expected := `(func (funcSpecParamSpecString)) (func (funcSpecResultString))`

	funcSpecParamSpec := NewSpecMock(ctrl)
	funcSpecResult := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: funcSpecParamSpec,
						},
					},
				},
			},
		},
		Results: []*Field{
			{
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: funcSpecResult,
						},
					},
				},
			},
		},
	}

	funcSpecParamSpec.
		EXPECT().
		String().
		Return(funcSpecParamSpecString)

	funcSpecResult.
		EXPECT().
		String().
		Return(funcSpecResultString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_WithParamCommentAndParamNameAndResultCommentAndResultNameAndIsVariadicMultiple(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	funcSpecParam1Comment := "funcSpecParam1Comment\nhere"
	funcSpecParam2Comment := "funcSpecParam2Comment\nhere"
	funcSpecParam1Name := "funcSpecParam1Name"
	funcSpecParam2Name := "funcSpecParam2Name"
	funcSpecParam1SpecString := "funcSpecParam1SpecString"
	funcSpecParam2SpecString := "funcSpecParam2SpecString"
	funcSpecResult1Comment := "funcSpecResult1Comment\nhere"
	funcSpecResult2Comment := "funcSpecResult2Comment\nhere"
	funcSpecResult1Name := "funcSpecResult1Name"
	funcSpecResult2Name := "funcSpecResult2Name"
	funcSpecResult1SpecString := "funcSpecResult1SpecString"
	funcSpecResult2SpecString := "funcSpecResult2SpecString"

	expected := `(
// funcSpecParam1Comment
// here
funcSpecParam1Name []funcSpecParam1SpecString, 
// funcSpecParam2Comment
// here
funcSpecParam2Name ...funcSpecParam2SpecString) (
// funcSpecResult1Comment
// here
funcSpecResult1Name funcSpecResult1SpecString, 
// funcSpecResult2Comment
// here
funcSpecResult2Name funcSpecResult2SpecString)`

	funcSpecParam1Spec := NewSpecMock(ctrl)
	funcSpecParam2Spec := NewSpecMock(ctrl)
	funcSpecResult1Spec := NewSpecMock(ctrl)
	funcSpecResult2Spec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		IsVariadic: true,
		Params: []*Field{
			{
				Name:    funcSpecParam1Name,
				Comment: funcSpecParam1Comment,
				Spec: &ArraySpec{
					Value: funcSpecParam1Spec,
				},
			},
			{
				Name:    funcSpecParam2Name,
				Comment: funcSpecParam2Comment,
				Spec: &ArraySpec{
					Value: funcSpecParam2Spec,
				},
			},
		},
		Results: []*Field{
			{
				Name:    funcSpecResult1Name,
				Comment: funcSpecResult1Comment,
				Spec:    funcSpecResult1Spec,
			},
			{
				Name:    funcSpecResult2Name,
				Comment: funcSpecResult2Comment,
				Spec:    funcSpecResult2Spec,
			},
		},
	}

	funcSpecParam1Spec.
		EXPECT().
		String().
		Return(funcSpecParam1SpecString)

	funcSpecParam2Spec.
		EXPECT().
		String().
		Return(funcSpecParam2SpecString)

	funcSpecResult1Spec.
		EXPECT().
		String().
		Return(funcSpecResult1SpecString)

	funcSpecResult2Spec.
		EXPECT().
		String().
		Return(funcSpecResult2SpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_WithParamCommentAndParamNameAndResultCommentAndResultNameAndMultipleValues(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	funcSpecParam1Comment := "funcSpecParam1Comment\nhere"
	funcSpecParam2Comment := "funcSpecParam2Comment\nhere"
	funcSpecParam1Name := "funcSpecParam1Name"
	funcSpecParam2Name := "funcSpecParam2Name"
	funcSpecParam1SpecString := "funcSpecParam1SpecString"
	funcSpecParam2SpecString := "funcSpecParam2SpecString"
	funcSpecResult1Comment := "funcSpecResult1Comment\nhere"
	funcSpecResult2Comment := "funcSpecResult2Comment\nhere"
	funcSpecResult1Name := "funcSpecResult1Name"
	funcSpecResult2Name := "funcSpecResult2Name"
	funcSpecResult1SpecString := "funcSpecResult1SpecString"
	funcSpecResult2SpecString := "funcSpecResult2SpecString"

	expected := `(
// funcSpecParam1Comment
// here
funcSpecParam1Name funcSpecParam1SpecString, 
// funcSpecParam2Comment
// here
funcSpecParam2Name funcSpecParam2SpecString) (
// funcSpecResult1Comment
// here
funcSpecResult1Name funcSpecResult1SpecString, 
// funcSpecResult2Comment
// here
funcSpecResult2Name funcSpecResult2SpecString)`

	funcSpecParam1Spec := NewSpecMock(ctrl)
	funcSpecParam2Spec := NewSpecMock(ctrl)
	funcSpecResult1Spec := NewSpecMock(ctrl)
	funcSpecResult2Spec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Name:    funcSpecParam1Name,
				Comment: funcSpecParam1Comment,
				Spec:    funcSpecParam1Spec,
			},
			{
				Name:    funcSpecParam2Name,
				Comment: funcSpecParam2Comment,
				Spec:    funcSpecParam2Spec,
			},
		},
		Results: []*Field{
			{
				Name:    funcSpecResult1Name,
				Comment: funcSpecResult1Comment,
				Spec:    funcSpecResult1Spec,
			},
			{
				Name:    funcSpecResult2Name,
				Comment: funcSpecResult2Comment,
				Spec:    funcSpecResult2Spec,
			},
		},
	}

	funcSpecParam1Spec.
		EXPECT().
		String().
		Return(funcSpecParam1SpecString)

	funcSpecParam2Spec.
		EXPECT().
		String().
		Return(funcSpecParam2SpecString)

	funcSpecResult1Spec.
		EXPECT().
		String().
		Return(funcSpecResult1SpecString)

	funcSpecResult2Spec.
		EXPECT().
		String().
		Return(funcSpecResult2SpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_WithParamNameAndResultNameAndMultipleValues(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	funcSpecParam1Name := "funcSpecParam1Name"
	funcSpecParam2Name := "funcSpecParam2Name"
	funcSpecParam1SpecString := "funcSpecParam1SpecString"
	funcSpecParam2SpecString := "funcSpecParam2SpecString"
	funcSpecResult1Name := "funcSpecResult1Name"
	funcSpecResult2Name := "funcSpecResult2Name"
	funcSpecResult1SpecString := "funcSpecResult1SpecString"
	funcSpecResult2SpecString := "funcSpecResult2SpecString"

	expected := `(funcSpecParam1Name funcSpecParam1SpecString, funcSpecParam2Name funcSpecParam2SpecString) ` +
		`(funcSpecResult1Name funcSpecResult1SpecString, funcSpecResult2Name funcSpecResult2SpecString)`

	funcSpecParam1Spec := NewSpecMock(ctrl)
	funcSpecParam2Spec := NewSpecMock(ctrl)
	funcSpecResult1Spec := NewSpecMock(ctrl)
	funcSpecResult2Spec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Name: funcSpecParam1Name,
				Spec: funcSpecParam1Spec,
			},
			{
				Name: funcSpecParam2Name,
				Spec: funcSpecParam2Spec,
			},
		},
		Results: []*Field{
			{
				Name: funcSpecResult1Name,
				Spec: funcSpecResult1Spec,
			},
			{
				Name: funcSpecResult2Name,
				Spec: funcSpecResult2Spec,
			},
		},
	}

	funcSpecParam1Spec.
		EXPECT().
		String().
		Return(funcSpecParam1SpecString)

	funcSpecParam2Spec.
		EXPECT().
		String().
		Return(funcSpecParam2SpecString)

	funcSpecResult1Spec.
		EXPECT().
		String().
		Return(funcSpecResult1SpecString)

	funcSpecResult2Spec.
		EXPECT().
		String().
		Return(funcSpecResult2SpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_AndMultipleValues(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	funcSpecParam1SpecString := "funcSpecParam1SpecString"
	funcSpecParam2SpecString := "funcSpecParam2SpecString"
	funcSpecResult1SpecString := "funcSpecResult1SpecString"
	funcSpecResult2SpecString := "funcSpecResult2SpecString"

	expected := `(funcSpecParam1SpecString, funcSpecParam2SpecString) (funcSpecResult1SpecString, funcSpecResult2SpecString)`

	funcSpecParam1Spec := NewSpecMock(ctrl)
	funcSpecParam2Spec := NewSpecMock(ctrl)
	funcSpecResult1Spec := NewSpecMock(ctrl)
	funcSpecResult2Spec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: funcSpecParam1Spec,
			},
			{
				Spec: funcSpecParam2Spec,
			},
		},
		Results: []*Field{
			{
				Spec: funcSpecResult1Spec,
			},
			{
				Spec: funcSpecResult2Spec,
			},
		},
	}

	funcSpecParam1Spec.
		EXPECT().
		String().
		Return(funcSpecParam1SpecString)

	funcSpecParam2Spec.
		EXPECT().
		String().
		Return(funcSpecParam2SpecString)

	funcSpecResult1Spec.
		EXPECT().
		String().
		Return(funcSpecResult1SpecString)

	funcSpecResult2Spec.
		EXPECT().
		String().
		Return(funcSpecResult2SpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	funcSpecParamSpec := NewSpecMock(ctrl)
	funcResultSpec := NewSpecMock(ctrl)

	clonedFuncParamSpec := NewSpecMock(ctrl)
	clonedFuncResultSpec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Annotations: []interface{}{
					&TestAnnotation{
						Name: "funcParamAnnotation",
					},
				},
				Spec: funcSpecParamSpec,
			},
		},
		Results: []*Field{
			{
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "funcSpecResultAnnotation",
					},
				},
				Spec: funcResultSpec,
			},
		},
	}

	funcSpecParamSpec.
		EXPECT().
		Clone().
		Return(clonedFuncParamSpec)

	funcResultSpec.
		EXPECT().
		Clone().
		Return(clonedFuncResultSpec)

	actual := modelValue.Clone()

	ctrl.AssertEqual(modelValue, actual, unit.IgnoreUnexportedOption{Value: SpecMock{}})
	ctrl.AssertNotSame(modelValue, actual)
	ctrl.AssertSame(clonedFuncParamSpec, actual.(*FuncSpec).Params[0].Spec)
	ctrl.AssertNotSame(modelValue.Params[0].Annotations[0], actual.(*FuncSpec).Params[0].Annotations[0])
	ctrl.AssertSame(clonedFuncResultSpec, actual.(*FuncSpec).Results[0].Spec)
	ctrl.AssertNotSame(modelValue.Results[0].Annotations[0], actual.(*FuncSpec).Results[0].Annotations[0])
}

func TestFuncSpec_Clone_WithoutParamsAndWithoutResults(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{}

	actual := modelValue.Clone()

	ctrl.AssertEqual(modelValue, actual)
}

func TestFuncSpec_FetchImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{}

	expected := []*Import{
		{
			Alias:     "funcSpecParamPackageName",
			Namespace: "funcSpecParamNamespace",
		},
		{
			Alias:     "funcSpecResultPackageName",
			Namespace: "funcSpecResultNamespace",
		},
	}

	funcSpecParamSpec := NewSpecMock(ctrl)
	funcResultSpec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: funcSpecParamSpec,
			},
		},
		Results: []*Field{
			{
				Spec: funcResultSpec,
			},
		},
	}

	funcSpecParamSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return([]*Import{expected[0]})

	funcResultSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return([]*Import{expected[1]})

	actual := modelValue.FetchImports(file)

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_FetchImports_WithoutImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{}

	funcSpecParamSpec := NewSpecMock(ctrl)
	funcResultSpec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: funcSpecParamSpec,
			},
		},
		Results: []*Field{
			{
				Spec: funcResultSpec,
			},
		},
	}

	funcSpecParamSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(nil)

	funcResultSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(nil)

	actual := modelValue.FetchImports(file)

	ctrl.AssertEmpty(actual)
}

func TestFuncSpec_FetchImports_WithEmptyFuncSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{}

	modelValue := &FuncSpec{}

	actual := modelValue.FetchImports(file)

	ctrl.AssertEmpty(actual)
}

func TestFuncSpec_RenameImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	funcSpecParamSpec := NewSpecMock(ctrl)
	funcSpecResult := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: funcSpecParamSpec,
			},
		},
		Results: []*Field{
			{
				Spec: funcSpecResult,
			},
		},
	}

	funcSpecParamSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	funcSpecResult.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	modelValue.RenameImports(oldAlias, newAlias)
}

func TestFuncSpec_RenameImports_WithEmptyFuncSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	modelValue := &FuncSpec{}

	modelValue.RenameImports(oldAlias, newAlias)
}

func TestFuncSpec_RenameImports_InvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	modelValue := &FuncSpec{}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestFuncSpec_RenameImports_InvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "packageName"
	newAlias := "+invalid"

	modelValue := &FuncSpec{}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
