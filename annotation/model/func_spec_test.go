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
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
					Length: -100,
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'Length' must be greater than or equal to 0, actual value: -100"),
		)
}
func TestFuncSpec_Validate_WithInvalidArraySpecResultSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Results: []*Field{
			{
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
					Length: -100,
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'Length' must be greater than or equal to 0, actual value: -100"),
		)
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
					Methods: []*Field{nil},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Methods[0]' must be not nil"))
}

func TestFuncSpec_Validate_WithInvalidInterfaceSpecResultSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &FuncSpec{
		Results: []*Field{
			{
				Spec: &InterfaceSpec{
					Methods: []*Field{nil},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Methods[0]' must be not nil"))
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

func TestFuncSpec_Validate_WithArraySpecParamSpecWithEllipsisAndIsVariadic(t *testing.T) {
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
					IsEllipsis: true,
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(
			NewErrorMessageConstraint(
				"Variable 'Params[0].Spec.(*model.ArraySpec).IsEllipsis' must be 'false' for variadic *model.FuncSpec",
			),
		)
}

func TestFuncSpec_Validate_WithArraySpecParamSpecWithLengthAndIsVariadic(t *testing.T) {
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
					Length: 10,
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(
			NewErrorMessageConstraint(
				"Variable 'Params[0].Spec.(*model.ArraySpec).Length' must be '0' for variadic *model.FuncSpec",
			),
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

	paramComment := "paramComment\nhere"
	paramName := "paramName"
	paramSpecString := "paramSpecString"
	resultComment := "resultComment\nhere"
	resultName := "resultName"
	resultSpecString := "resultSpecString"
	expected := `(
// paramComment
// here
paramName ...paramSpecString) (
// resultComment
// here
resultName resultSpecString)`

	paramSpec := NewSpecMock(ctrl)
	resultSpec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		IsVariadic: true,
		Params: []*Field{
			{
				Name:    paramName,
				Comment: paramComment,
				Spec: &ArraySpec{
					Value: paramSpec,
				},
			},
		},
		Results: []*Field{
			{
				Name:    resultName,
				Comment: resultComment,
				Spec:    resultSpec,
			},
		},
	}

	paramSpec.
		EXPECT().
		String().
		Return(paramSpecString)

	resultSpec.
		EXPECT().
		String().
		Return(resultSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_WithParamCommentAndParamNameAndResultCommentAndResultName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	paramComment := "paramComment\nhere"
	paramName := "paramName"
	paramSpecString := "paramSpecString"
	resultComment := "resultComment\nhere"
	resultName := "resultName"
	resultSpecString := "resultSpecString"
	expected := `(
// paramComment
// here
paramName paramSpecString) (
// resultComment
// here
resultName resultSpecString)`

	paramSpec := NewSpecMock(ctrl)
	resultSpec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Name:    paramName,
				Comment: paramComment,
				Spec:    paramSpec,
			},
		},
		Results: []*Field{
			{
				Name:    resultName,
				Comment: resultComment,
				Spec:    resultSpec,
			},
		},
	}

	paramSpec.
		EXPECT().
		String().
		Return(paramSpecString)

	resultSpec.
		EXPECT().
		String().
		Return(resultSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_WithParamNameAndResultName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	paramName := "paramName"
	paramSpecString := "paramSpecString"
	resultName := "resultName"
	resultSpecString := "resultSpecString"
	expected := `(paramName paramSpecString) (resultName resultSpecString)`

	paramSpec := NewSpecMock(ctrl)
	resultSpec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Name: paramName,
				Spec: paramSpec,
			},
		},
		Results: []*Field{
			{
				Name: resultName,
				Spec: resultSpec,
			},
		},
	}

	paramSpec.
		EXPECT().
		String().
		Return(paramSpecString)

	resultSpec.
		EXPECT().
		String().
		Return(resultSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	paramSpecString := "paramSpecString"
	resultSpecString := "resultSpecString"
	expected := `(paramSpecString) (resultSpecString)`

	paramSpec := NewSpecMock(ctrl)
	resultSpec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: paramSpec,
			},
		},
		Results: []*Field{
			{
				Spec: resultSpec,
			},
		},
	}

	paramSpec.
		EXPECT().
		String().
		Return(paramSpecString)

	resultSpec.
		EXPECT().
		String().
		Return(resultSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_WithParamCommentAndParamNameAndResultCommentAndResultNameAndBothFuncSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	paramComment := "paramComment\nhere"
	paramName := "paramName"
	paramSpecString := "paramSpecString"
	resultComment := "resultComment\nhere"
	resultName := "resultName"
	resultSpecString := "resultSpecString"
	expected := `(
// paramComment
// here
paramName func (paramSpecString)) (
// resultComment
// here
resultName func (resultSpecString))`

	paramSpec := NewSpecMock(ctrl)
	resultSpec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Name:    paramName,
				Comment: paramComment,
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: paramSpec,
						},
					},
				},
			},
		},
		Results: []*Field{
			{
				Name:    resultName,
				Comment: resultComment,
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: resultSpec,
						},
					},
				},
			},
		},
	}

	paramSpec.
		EXPECT().
		String().
		Return(paramSpecString)

	resultSpec.
		EXPECT().
		String().
		Return(resultSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_WithParamNameAndFuncSpecParamSpecAndResultNameAndFuncSpecResultSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	paramName := "paramName"
	paramSpecString := "paramSpecString"
	resultName := "resultName"
	resultSpecString := "resultSpecString"
	expected := `(paramName func (paramSpecString)) (resultName func (resultSpecString))`

	paramSpec := NewSpecMock(ctrl)
	resultSpec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Name: paramName,
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: paramSpec,
						},
					},
				},
			},
		},
		Results: []*Field{
			{
				Name: resultName,
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: resultSpec,
						},
					},
				},
			},
		},
	}

	paramSpec.
		EXPECT().
		String().
		Return(paramSpecString)

	resultSpec.
		EXPECT().
		String().
		Return(resultSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_WithFuncSpecParamSpecAndFuncSpecResultSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	paramSpecString := "paramSpecString"
	resultSpecString := "resultSpecString"
	expected := `(func (paramSpecString)) (func (resultSpecString))`

	paramSpec := NewSpecMock(ctrl)
	resultSpec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: paramSpec,
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
							Spec: resultSpec,
						},
					},
				},
			},
		},
	}

	paramSpec.
		EXPECT().
		String().
		Return(paramSpecString)

	resultSpec.
		EXPECT().
		String().
		Return(resultSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_WithParamCommentAndParamNameAndResultCommentAndResultNameAndIsVariadicMultiple(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	param1Comment := "param1Comment\nhere"
	param2Comment := "param2Comment\nhere"
	param1Name := "param1Name"
	param2Name := "param2Name"
	param1SpecString := "param1SpecString"
	param2SpecString := "param2SpecString"
	result1Comment := "result1Comment\nhere"
	result2Comment := "result2Comment\nhere"
	result1Name := "result1Name"
	result2Name := "result2Name"
	result1SpecString := "result1SpecString"
	result2SpecString := "result2SpecString"

	expected := `(
// param1Comment
// here
param1Name []param1SpecString, 
// param2Comment
// here
param2Name ...param2SpecString) (
// result1Comment
// here
result1Name result1SpecString, 
// result2Comment
// here
result2Name result2SpecString)`

	param1Spec := NewSpecMock(ctrl)
	param2Spec := NewSpecMock(ctrl)
	result1Spec := NewSpecMock(ctrl)
	result2Spec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		IsVariadic: true,
		Params: []*Field{
			{
				Name:    param1Name,
				Comment: param1Comment,
				Spec: &ArraySpec{
					Value: param1Spec,
				},
			},
			{
				Name:    param2Name,
				Comment: param2Comment,
				Spec: &ArraySpec{
					Value: param2Spec,
				},
			},
		},
		Results: []*Field{
			{
				Name:    result1Name,
				Comment: result1Comment,
				Spec:    result1Spec,
			},
			{
				Name:    result2Name,
				Comment: result2Comment,
				Spec:    result2Spec,
			},
		},
	}

	param1Spec.
		EXPECT().
		String().
		Return(param1SpecString)

	param2Spec.
		EXPECT().
		String().
		Return(param2SpecString)

	result1Spec.
		EXPECT().
		String().
		Return(result1SpecString)

	result2Spec.
		EXPECT().
		String().
		Return(result2SpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_WithParamCommentAndParamNameAndResultCommentAndResultNameAndMultipleValues(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	param1Comment := "param1Comment\nhere"
	param2Comment := "param2Comment\nhere"
	param1Name := "param1Name"
	param2Name := "param2Name"
	param1SpecString := "param1SpecString"
	param2SpecString := "param2SpecString"
	result1Comment := "result1Comment\nhere"
	result2Comment := "result2Comment\nhere"
	result1Name := "result1Name"
	result2Name := "result2Name"
	result1SpecString := "result1SpecString"
	result2SpecString := "result2SpecString"

	expected := `(
// param1Comment
// here
param1Name param1SpecString, 
// param2Comment
// here
param2Name param2SpecString) (
// result1Comment
// here
result1Name result1SpecString, 
// result2Comment
// here
result2Name result2SpecString)`

	param1Spec := NewSpecMock(ctrl)
	param2Spec := NewSpecMock(ctrl)
	result1Spec := NewSpecMock(ctrl)
	result2Spec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Name:    param1Name,
				Comment: param1Comment,
				Spec:    param1Spec,
			},
			{
				Name:    param2Name,
				Comment: param2Comment,
				Spec:    param2Spec,
			},
		},
		Results: []*Field{
			{
				Name:    result1Name,
				Comment: result1Comment,
				Spec:    result1Spec,
			},
			{
				Name:    result2Name,
				Comment: result2Comment,
				Spec:    result2Spec,
			},
		},
	}

	param1Spec.
		EXPECT().
		String().
		Return(param1SpecString)

	param2Spec.
		EXPECT().
		String().
		Return(param2SpecString)

	result1Spec.
		EXPECT().
		String().
		Return(result1SpecString)

	result2Spec.
		EXPECT().
		String().
		Return(result2SpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_WithParamNameAndResultNameAndMultipleValues(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	param1Name := "param1Name"
	param2Name := "param2Name"
	param1SpecString := "param1SpecString"
	param2SpecString := "param2SpecString"
	result1Name := "result1Name"
	result2Name := "result2Name"
	result1SpecString := "result1SpecString"
	result2SpecString := "result2SpecString"

	expected := `(param1Name param1SpecString, param2Name param2SpecString) ` +
		`(result1Name result1SpecString, result2Name result2SpecString)`

	param1Spec := NewSpecMock(ctrl)
	param2Spec := NewSpecMock(ctrl)
	result1Spec := NewSpecMock(ctrl)
	result2Spec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Name: param1Name,
				Spec: param1Spec,
			},
			{
				Name: param2Name,
				Spec: param2Spec,
			},
		},
		Results: []*Field{
			{
				Name: result1Name,
				Spec: result1Spec,
			},
			{
				Name: result2Name,
				Spec: result2Spec,
			},
		},
	}

	param1Spec.
		EXPECT().
		String().
		Return(param1SpecString)

	param2Spec.
		EXPECT().
		String().
		Return(param2SpecString)

	result1Spec.
		EXPECT().
		String().
		Return(result1SpecString)

	result2Spec.
		EXPECT().
		String().
		Return(result2SpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_String_AndMultipleValues(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	param1SpecString := "param1SpecString"
	param2SpecString := "param2SpecString"
	result1SpecString := "result1SpecString"
	result2SpecString := "result2SpecString"

	expected := `(param1SpecString, param2SpecString) (result1SpecString, result2SpecString)`

	param1Spec := NewSpecMock(ctrl)
	param2Spec := NewSpecMock(ctrl)
	result1Spec := NewSpecMock(ctrl)
	result2Spec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: param1Spec,
			},
			{
				Spec: param2Spec,
			},
		},
		Results: []*Field{
			{
				Spec: result1Spec,
			},
			{
				Spec: result2Spec,
			},
		},
	}

	param1Spec.
		EXPECT().
		String().
		Return(param1SpecString)

	param2Spec.
		EXPECT().
		String().
		Return(param2SpecString)

	result1Spec.
		EXPECT().
		String().
		Return(result1SpecString)

	result2Spec.
		EXPECT().
		String().
		Return(result2SpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFuncSpec_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	paramSpec := NewSpecMock(ctrl)
	resultSpec := NewSpecMock(ctrl)

	clonedParamSpec := NewSpecMock(ctrl)
	clonedResultSpec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "paramAnnotation",
					},
				},
				Spec: paramSpec,
			},
		},
		Results: []*Field{
			{
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "resultAnnotation",
					},
				},
				Spec: resultSpec,
			},
		},
	}

	paramSpec.
		EXPECT().
		Clone().
		Return(clonedParamSpec)

	resultSpec.
		EXPECT().
		Clone().
		Return(clonedResultSpec)

	actual := modelValue.Clone()

	ctrl.AssertEqual(
		modelValue,
		actual,
		unit.IgnoreUnexportedOption{Value: *ctrl},
		unit.IgnoreUnexportedOption{Value: MockCallManager{}},
	)
	ctrl.AssertNotSame(modelValue, actual)
	ctrl.AssertSame(clonedParamSpec, actual.(*FuncSpec).Params[0].Spec)
	ctrl.AssertNotSame(modelValue.Params[0].Annotations[0], actual.(*FuncSpec).Params[0].Annotations[0])
	ctrl.AssertSame(clonedResultSpec, actual.(*FuncSpec).Results[0].Spec)
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
			Alias:     "paramPackageName",
			Namespace: "paramNamespace",
		},
		{
			Alias:     "resultPackageName",
			Namespace: "resultNamespace",
		},
	}

	paramSpec := NewSpecMock(ctrl)
	resultSpec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: paramSpec,
			},
		},
		Results: []*Field{
			{
				Spec: resultSpec,
			},
		},
	}

	paramSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return([]*Import{expected[0]})

	resultSpec.
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

	paramSpec := NewSpecMock(ctrl)
	resultSpec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: paramSpec,
			},
		},
		Results: []*Field{
			{
				Spec: resultSpec,
			},
		},
	}

	paramSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(nil)

	resultSpec.
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

	paramSpec := NewSpecMock(ctrl)
	resultSpec := NewSpecMock(ctrl)

	modelValue := &FuncSpec{
		Params: []*Field{
			{
				Spec: paramSpec,
			},
		},
		Results: []*Field{
			{
				Spec: resultSpec,
			},
		},
	}

	paramSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	resultSpec.
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
