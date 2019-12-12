package annotation

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestVarGroup_Validate_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &VarGroup{}

	model.Validate()
}

func TestVarGroup_Validate_WithSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &VarGroup{
		Comment: "varGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varGroupAnnotation",
			},
		},
		Vars: []*Var{
			{
				Name: "name",
				Spec: &SimpleSpec{
					TypeName: "varName",
				},
			},
		},
	}

	model.Validate()
}

func TestVarGroup_Validate_WithArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &VarGroup{
		Comment: "varGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varGroupAnnotation",
			},
		},
		Vars: []*Var{
			{
				Name: "name",
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "varName",
					},
				},
			},
		},
	}

	model.Validate()
}

func TestVarGroup_Validate_WithMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &VarGroup{
		Comment: "varGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varGroupAnnotation",
			},
		},
		Vars: []*Var{
			{
				Name: "name",
				Spec: &MapSpec{
					Key: &SimpleSpec{
						TypeName: "varName",
					},
					Value: &SimpleSpec{
						TypeName: "varName",
					},
				},
			},
		},
	}

	model.Validate()
}

func TestVarGroup_Validate_WithStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &VarGroup{
		Comment: "varGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varGroupAnnotation",
			},
		},
		Vars: []*Var{
			{
				Name: "name",
				Spec: &StructSpec{},
			},
		},
	}

	model.Validate()
}

func TestVarGroup_Validate_WithInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &VarGroup{
		Comment: "varGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varGroupAnnotation",
			},
		},
		Vars: []*Var{
			{
				Name: "name",
				Spec: &InterfaceSpec{},
			},
		},
	}

	model.Validate()
}

func TestVarGroup_Validate_WithFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &VarGroup{
		Comment: "varGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varGroupAnnotation",
			},
		},
		Vars: []*Var{
			{
				Name: "name",
				Spec: &FuncSpec{},
			},
		},
	}

	model.Validate()
}

func TestVarGroup_Validate_WithInvalidVar(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &VarGroup{
		Comment: "varGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varGroupAnnotation",
			},
		},
		Vars: []*Var{
			{
				Spec: &FuncSpec{},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestVarGroup_Validate_WithNilVar(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &VarGroup{
		Comment: "varGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varGroupAnnotation",
			},
		},
		Vars: []*Var{
			nil,
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Vars[0]' must be not nil"))
}

func TestVarGroup_Validate_WithInvalidSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &VarGroup{
		Vars: []*Var{
			{
				Name: "name",
				Spec: &SimpleSpec{
					TypeName: "+invalid",
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'TypeName' must be valid identifier, actual value: '+invalid'"))
}

func TestVarGroup_Validate_WithInvalidArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &VarGroup{
		Vars: []*Var{
			{
				Name: "name",
				Spec: &ArraySpec{},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Value' must be not nil"))
}

func TestVarGroup_Validate_WithInvalidMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &VarGroup{
		Vars: []*Var{
			{
				Name: "name",
				Spec: &MapSpec{
					Value: &SimpleSpec{
						TypeName: "varName",
					},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Key' must be not nil"))
}

func TestVarGroup_Validate_WithInvalidStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &VarGroup{
		Vars: []*Var{
			{
				Name: "name",
				Spec: &StructSpec{
					Fields: []*Field{nil},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestVarGroup_Validate_WithInvalidInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &VarGroup{
		Vars: []*Var{
			{
				Name: "name",
				Spec: &InterfaceSpec{
					Fields: []*Field{nil},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestVarGroup_Validate_WithInvalidFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &VarGroup{
		Vars: []*Var{
			{
				Name: "name",
				Spec: &FuncSpec{
					Params: []*Field{nil},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Params[0]' must be not nil"))
}

func TestVarGroup_Validate_WithInvalidVarValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &VarGroup{
		Vars: []*Var{
			{
				Name: "name",
				Spec: NewSpecMock(ctrl),
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Spec' has invalid type: %T", model.Vars[0].Spec))
}

func TestVarGroup_String(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `var (
)
`

	model := &VarGroup{}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestVarGroup_String_WithVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	varGroupComment := "varGroup\ncomment"

	expected := `// varGroup
// comment
var (
)
`

	model := &VarGroup{
		Comment: varGroupComment,
	}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestVarGroup_String_WithOneVarAndVarSpecAndVarComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	varComment := "var\ncomment"
	varName := "varName"
	varSpecValue := "varSpecValue"

	expected := `var (
// var
// comment
varName varSpecValue
)
`

	varSpec := NewSpecMock(ctrl)

	model := &VarGroup{
		Vars: []*Var{
			{
				Comment: varComment,
				Name:    varName,
				Spec:    varSpec,
			},
		},
	}

	varSpec.
		EXPECT().
		String().
		Return(varSpecValue)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestVarGroup_String_WithOneVarAndVarSpecAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	varGroupComment := "varGroup\ncomment"
	varName := "varName"
	varSpecValue := "varSpecValue"

	expected := `// varGroup
// comment
var varName varSpecValue
`

	varSpec := NewSpecMock(ctrl)

	model := &VarGroup{
		Comment: varGroupComment,
		Vars: []*Var{
			{
				Name: varName,
				Spec: varSpec,
			},
		},
	}

	varSpec.
		EXPECT().
		String().
		Return(varSpecValue)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestVarGroup_String_WithOneVarAndVarSpecAndVarCommentAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	varGroupComment := "varGroup\ncomment"
	varComment := "var\ncomment"
	varName := "varName"
	varSpecValue := "varSpecValue"

	expected := `// varGroup
// comment
var (
// var
// comment
varName varSpecValue
)
`

	varSpec := NewSpecMock(ctrl)

	model := &VarGroup{
		Comment: varGroupComment,
		Vars: []*Var{
			{
				Comment: varComment,
				Name:    varName,
				Spec:    varSpec,
			},
		},
	}

	varSpec.
		EXPECT().
		String().
		Return(varSpecValue)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestVarGroup_String_WithOneVarAndVarSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	varName := "varName"
	varSpecValue := "varSpecValue"

	expected := `var varName varSpecValue
`

	varSpec := NewSpecMock(ctrl)

	model := &VarGroup{
		Vars: []*Var{
			{
				Name: varName,
				Spec: varSpec,
			},
		},
	}

	varSpec.
		EXPECT().
		String().
		Return(varSpecValue)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestVarGroup_String_WithOneVarAndVarValueAndVarComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	varComment := "var\ncomment"
	varName := "varName"
	varValue := "varValue"

	expected := `var (
// var
// comment
varName = varValue
)
`

	model := &VarGroup{
		Vars: []*Var{
			{
				Comment: varComment,
				Name:    varName,
				Value:   varValue,
			},
		},
	}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestVarGroup_String_WithOneVarAndVarValueAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	varGroupComment := "varGroup\ncomment"
	varName := "varName"
	varValue := "varValue"

	expected := `// varGroup
// comment
var varName = varValue
`

	model := &VarGroup{
		Comment: varGroupComment,
		Vars: []*Var{
			{
				Name:  varName,
				Value: varValue,
			},
		},
	}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestVarGroup_String_WithOneVarAndVarValueAndVarCommentAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	varGroupComment := "varGroup\ncomment"
	varComment := "var\ncomment"
	varName := "varName"
	varValue := "varValue"

	expected := `// varGroup
// comment
var (
// var
// comment
varName = varValue
)
`

	model := &VarGroup{
		Comment: varGroupComment,
		Vars: []*Var{
			{
				Comment: varComment,
				Name:    varName,
				Value:   varValue,
			},
		},
	}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestVarGroup_String_WithOneVarAndVarValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	varName := "varName"
	varValue := "varValue"

	expected := `var varName = varValue
`

	model := &VarGroup{
		Vars: []*Var{
			{
				Name:  varName,
				Value: varValue,
			},
		},
	}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestVarGroup_String_WithOneVarAndVarSpecAndVarValueAndVarComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	varComment := "var\ncomment"
	varName := "varName"
	varValue := "varValue"
	varSpecValue := "varSpecValue"

	expected := `var (
// var
// comment
varName varSpecValue = varValue
)
`

	varSpec := NewSpecMock(ctrl)

	model := &VarGroup{
		Vars: []*Var{
			{
				Comment: varComment,
				Name:    varName,
				Value:   varValue,
				Spec:    varSpec,
			},
		},
	}

	varSpec.
		EXPECT().
		String().
		Return(varSpecValue)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestVarGroup_String_WithOneVarAndVarSpecAndVarValueAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	varGroupComment := "varGroup\ncomment"
	varName := "varName"
	varValue := "varValue"
	varSpecValue := "varSpecValue"

	expected := `// varGroup
// comment
var varName varSpecValue = varValue
`

	varSpec := NewSpecMock(ctrl)

	model := &VarGroup{
		Comment: varGroupComment,
		Vars: []*Var{
			{
				Name:  varName,
				Value: varValue,
				Spec:  varSpec,
			},
		},
	}

	varSpec.
		EXPECT().
		String().
		Return(varSpecValue)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestVarGroup_String_WithOneVarAndVarSpecAndVarValueAndVarCommentAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	varGroupComment := "varGroup\ncomment"
	varComment := "var\ncomment"
	varName := "varName"
	varValue := "varValue"
	varSpecValue := "varSpecValue"

	expected := `// varGroup
// comment
var (
// var
// comment
varName varSpecValue = varValue
)
`

	varSpec := NewSpecMock(ctrl)

	model := &VarGroup{
		Comment: varGroupComment,
		Vars: []*Var{
			{
				Comment: varComment,
				Name:    varName,
				Value:   varValue,
				Spec:    varSpec,
			},
		},
	}

	varSpec.
		EXPECT().
		String().
		Return(varSpecValue)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestVarGroup_String_WithOneVarAndVarSpecAndVarValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	varName := "varName"
	varValue := "varValue"
	varSpecValue := "varSpecValue"

	expected := `var varName varSpecValue = varValue
`

	varSpec := NewSpecMock(ctrl)

	model := &VarGroup{
		Vars: []*Var{
			{
				Name:  varName,
				Value: varValue,
				Spec:  varSpec,
			},
		},
	}

	varSpec.
		EXPECT().
		String().
		Return(varSpecValue)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestVarGroup_String_WithMultipleVars(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	var1Name := "var1Name"
	var2Name := "var2Name"
	varSpec1Value := "varSpec1Value"
	varSpec2Value := "varSpec2Value"

	expected := `var (
var1Name varSpec1Value
var2Name varSpec2Value
)
`

	var1Spec := NewSpecMock(ctrl)
	var2Spec := NewSpecMock(ctrl)

	model := &VarGroup{
		Vars: []*Var{
			{
				Name: var1Name,
				Spec: var1Spec,
			},
			{
				Name: var2Name,
				Spec: var2Spec,
			},
		},
	}

	var1Spec.
		EXPECT().
		String().
		Return(varSpec1Value)

	var2Spec.
		EXPECT().
		String().
		Return(varSpec2Value)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestVarGroup_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	var1Spec := NewSpecMock(ctrl)
	var2Spec := NewSpecMock(ctrl)

	clonedVar1Spec := NewSpecMock(ctrl)
	clonedVar2Spec := NewSpecMock(ctrl)

	model := &VarGroup{
		Comment: "varGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "varGroupAnnotation",
			},
		},
		Vars: []*Var{
			{
				Name:    "var1Name",
				Comment: "var1\ncomment",
				Annotations: []interface{}{
					&TestAnnotation{
						Name: "var1Annotation",
					},
				},
				Spec: var1Spec,
			},
			{
				Name:    "var2Name",
				Comment: "var2\ncomment",
				Annotations: []interface{}{
					&TestAnnotation{
						Name: "var2Annotation",
					},
				},
				Spec: var2Spec,
			},
		},
	}

	var1Spec.
		EXPECT().
		Clone().
		Return(clonedVar1Spec)

	var2Spec.
		EXPECT().
		Clone().
		Return(clonedVar2Spec)

	actual := model.Clone()

	ctrl.AssertEqual(model, actual, unit.IgnoreUnexportedOption{Value: SpecMock{}})
	ctrl.AssertNotSame(model, actual)
	ctrl.AssertNotSame(model.Annotations[0], actual.(*VarGroup).Annotations[0])
	ctrl.AssertNotSame(model.Vars[0], actual.(*VarGroup).Vars[0])
	ctrl.AssertNotSame(model.Vars[0].Annotations[0], actual.(*VarGroup).Vars[0].Annotations[0])
	ctrl.AssertSame(clonedVar1Spec, actual.(*VarGroup).Vars[0].Spec)
	ctrl.AssertNotSame(model.Vars[1], actual.(*VarGroup).Vars[1])
	ctrl.AssertNotSame(model.Vars[1].Annotations[0], actual.(*VarGroup).Vars[1].Annotations[0])
	ctrl.AssertSame(clonedVar2Spec, actual.(*VarGroup).Vars[1].Spec)
}

func TestVarGroup_Clone_WithoutFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &VarGroup{}

	actual := model.Clone()

	ctrl.AssertEqual(model, actual)
	ctrl.AssertNotSame(model, actual)
}

func TestVarGroup_EqualSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Spec1 := NewSpecMock(ctrl)
	model1Spec2 := NewSpecMock(ctrl)
	model2Spec1 := NewSpecMock(ctrl)
	model2Spec2 := NewSpecMock(ctrl)

	model1 := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name1",
				Spec:  model1Spec1,
				Value: "value1",
			},
			{
				Name:  "name2",
				Spec:  model1Spec2,
				Value: "value2",
			},
		},
	}

	model2 := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name1",
				Spec:  model2Spec1,
				Value: "value1",
			},
			{
				Name:  "name2",
				Spec:  model2Spec2,
				Value: "value2",
			},
		},
	}

	model1Spec1.
		EXPECT().
		EqualSpec(ctrl.Same(model2Spec1)).
		Return(true)

	model1Spec2.
		EXPECT().
		EqualSpec(ctrl.Same(model2Spec2)).
		Return(true)

	actual := model1.EqualSpec(model2)

	ctrl.AssertTrue(actual)
}

func TestVarGroup_EqualSpec_WithoutValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Spec := NewSpecMock(ctrl)
	model2Spec := NewSpecMock(ctrl)

	model1 := &VarGroup{
		Vars: []*Var{
			{
				Name: "name",
				Spec: model1Spec,
			},
		},
	}

	model2 := &VarGroup{
		Vars: []*Var{
			{
				Name: "name",
				Spec: model2Spec,
			},
		},
	}

	model1Spec.
		EXPECT().
		EqualSpec(ctrl.Same(model2Spec)).
		Return(true)

	actual := model1.EqualSpec(model2)

	ctrl.AssertTrue(actual)
}

func TestVarGroup_EqualSpec_WithoutSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1 := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name",
				Value: "value",
			},
		},
	}

	model2 := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name",
				Value: "value",
			},
		},
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertTrue(actual)
}

func TestVarGroup_EqualSpec_WithEmptyVarGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1 := &VarGroup{}
	model2 := &VarGroup{}

	actual := model1.EqualSpec(model2)

	ctrl.AssertTrue(actual)
}

func TestVarGroup_EqualSpec_WithOrder(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Spec1 := NewSpecMock(ctrl)
	model1Spec2 := NewSpecMock(ctrl)
	model2Spec1 := NewSpecMock(ctrl)
	model2Spec2 := NewSpecMock(ctrl)

	model1 := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name1",
				Spec:  model1Spec1,
				Value: "value1",
			},
			{
				Name:  "name2",
				Spec:  model1Spec2,
				Value: "value2",
			},
		},
	}

	model2 := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name2",
				Spec:  model2Spec2,
				Value: "value2",
			},
			{
				Name:  "name1",
				Spec:  model2Spec1,
				Value: "value1",
			},
		},
	}

	model1Spec1.
		EXPECT().
		EqualSpec(ctrl.Same(model2Spec1)).
		Return(true)

	model1Spec2.
		EXPECT().
		EqualSpec(ctrl.Same(model2Spec2)).
		Return(true)

	actual := model1.EqualSpec(model2)

	ctrl.AssertTrue(actual)
}

func TestVarGroup_EqualSpec_WithLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Spec := NewSpecMock(ctrl)

	model1 := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name1",
				Spec:  model1Spec,
				Value: "value1",
			},
		},
	}

	model2 := &VarGroup{}

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestVarGroup_EqualSpec_WithAnotherType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Spec := NewSpecMock(ctrl)

	model1 := &VarGroup{
		Vars: []*Var{
			{
				Name: "name",
				Spec: model1Spec,
			},
		},
	}

	model2 := "model2"

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestVarGroup_EqualSpec_WithName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Spec := NewSpecMock(ctrl)
	model2Spec := NewSpecMock(ctrl)

	model1 := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name1",
				Spec:  model1Spec,
				Value: "value",
			},
		},
	}

	model2 := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name2",
				Spec:  model2Spec,
				Value: "value",
			},
		},
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestVarGroup_EqualSpec_WithSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Spec := NewSpecMock(ctrl)
	model2Spec := NewSpecMock(ctrl)

	model1 := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name",
				Spec:  model1Spec,
				Value: "value",
			},
		},
	}

	model2 := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name",
				Spec:  model2Spec,
				Value: "value",
			},
		},
	}

	model1Spec.
		EXPECT().
		EqualSpec(ctrl.Same(model2Spec)).
		Return(false)

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestVarGroup_EqualSpec_WithValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Spec := NewSpecMock(ctrl)
	model2Spec := NewSpecMock(ctrl)

	model1 := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name",
				Spec:  model1Spec,
				Value: "value1",
			},
		},
	}

	model2 := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name",
				Spec:  model2Spec,
				Value: "value2",
			},
		},
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestVarGroup_ContainsSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	elementSpec := NewSpecMock(ctrl)
	modelSpec1 := NewSpecMock(ctrl)
	modelSpec2 := NewSpecMock(ctrl)

	element := &Var{
		Name:  "name2",
		Spec:  elementSpec,
		Value: "value2",
	}

	model := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name1",
				Spec:  modelSpec1,
				Value: "value1",
			},
			{
				Name:  "name2",
				Spec:  modelSpec2,
				Value: "value2",
			},
		},
	}

	modelSpec2.
		EXPECT().
		EqualSpec(ctrl.Same(elementSpec)).
		Return(true)

	actual := model.ContainsSpec(element)

	ctrl.AssertTrue(actual)
}

func TestVarGroup_ContainsSpec_WithEmptyVars(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelSpec := NewSpecMock(ctrl)

	element := &Var{
		Name:  "name2",
		Spec:  modelSpec,
		Value: "value2",
	}

	model := &VarGroup{}

	actual := model.ContainsSpec(element)

	ctrl.AssertFalse(actual)
}

func TestVarGroup_ContainsSpec_WithNotContains(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	elementSpec := NewSpecMock(ctrl)
	modelSpec1 := NewSpecMock(ctrl)
	modelSpec2 := NewSpecMock(ctrl)

	element := &Var{
		Name:  "name3",
		Spec:  elementSpec,
		Value: "value3",
	}

	model := &VarGroup{
		Vars: []*Var{
			{
				Name:  "name1",
				Spec:  modelSpec1,
				Value: "value1",
			},
			{
				Name:  "name2",
				Spec:  modelSpec2,
				Value: "value2",
			},
		},
	}

	actual := model.ContainsSpec(element)

	ctrl.AssertFalse(actual)
}

func TestVarGroup_FetchImports_WithFoundImport(t *testing.T) {
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

	varSpec := NewSpecMock(ctrl)

	model := &VarGroup{
		Vars: []*Var{
			{
				Name: "varName",
				Spec: varSpec,
			},
		},
	}

	varSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(expected)

	actual := model.FetchImports(file)

	ctrl.AssertEqual(expected, actual)
}

func TestVarGroup_FetchImports_WithNotFoundImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	file := &File{
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "packageName",
						Namespace: "namespace",
					},
				},
			},
		},
	}

	varSpec := NewSpecMock(ctrl)

	model := &VarGroup{
		Vars: []*Var{
			{
				Name: "varName",
				Spec: varSpec,
			},
		},
	}

	varSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(nil)

	actual := model.FetchImports(file)

	ctrl.AssertEmpty(actual)
}

func TestVarGroup_RenameImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	varSpec := NewSpecMock(ctrl)

	model := &VarGroup{
		Vars: []*Var{
			{
				Name: "varName",
				Spec: varSpec,
			},
		},
	}

	varSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	model.RenameImports(oldAlias, newAlias)
}

func TestVarGroup_RenameImports_WithInvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	model := &VarGroup{
		Vars: []*Var{
			{
				Name: "varName",
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "varName",
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestVarGroup_RenameImports_WithInvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "+invalid"

	model := &VarGroup{
		Vars: []*Var{
			{
				Name: "varName",
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "varName",
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
