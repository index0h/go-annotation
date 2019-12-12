package annotation

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestStructSpec_Validate_WithSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	model.Validate()
}

func TestStructSpec_Validate_WithArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
			{
				Name: "name",
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	model.Validate()
}

func TestStructSpec_Validate_WithMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
			{
				Name: "name",
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

	model.Validate()
}

func TestStructSpec_Validate_WithStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
			{
				Name: "name",
				Spec: &StructSpec{},
			},
		},
	}

	model.Validate()
}

func TestStructSpec_Validate_WithInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
			{
				Name: "name",
				Spec: &InterfaceSpec{},
			},
		},
	}

	model.Validate()
}

func TestStructSpec_Validate_WithFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
			{
				Name: "name",
				Spec: &FuncSpec{},
			},
		},
	}

	model.Validate()
}

func TestStructSpec_Validate_WithInvalidName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
			{
				Name: "+invalid",
				Spec: NewSpecMock(ctrl),
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be valid identifier, actual value: '+invalid'"))
}

func TestStructSpec_Validate_WithoutFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{}

	model.Validate()
}

func TestStructSpec_Validate_WithNilField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
			nil,
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Fields[0]' must be not nil"))
}

func TestStructSpec_Validate_WithInvalidSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
			{
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

func TestStructSpec_Validate_WithInvalidArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
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

func TestStructSpec_Validate_WithInvalidMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
			{
				Name: "name",
				Spec: &MapSpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Key' must be not nil"))
}

func TestStructSpec_Validate_WithInvalidStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
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

func TestStructSpec_Validate_WithInvalidInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
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

func TestStructSpec_Validate_WithInvalidFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
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

func TestStructSpec_Validate_WithInvalidTypeValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
			{
				Name: "name",
				Spec: NewSpecMock(ctrl),
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Spec' has invalid type: %T", model.Fields[0].Spec))
}

func TestStructSpec_Validate_WithoutNameAndWithInvalidArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
			{
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "typeName",
					},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(
			NewErrorMessageConstraint(
				"Variable 'Fields[0]' with empty 'Name' has invalid type: %T",
				model.Fields[0].Spec,
			),
		)
}

func TestStructSpec_Validate_WithoutNameAndWithInvalidMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
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

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(
			NewErrorMessageConstraint(
				"Variable 'Fields[0]' with empty 'Name' has invalid type: %T",
				model.Fields[0].Spec,
			),
		)
}

func TestStructSpec_Validate_WithoutNameAndWithInvalidStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
			{
				Spec: &StructSpec{},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(
			NewErrorMessageConstraint(
				"Variable 'Fields[0]' with empty 'Name' has invalid type: %T",
				model.Fields[0].Spec,
			),
		)
}

func TestStructSpec_Validate_WithoutNameAndWithInvalidInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
			{
				Spec: &InterfaceSpec{},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(
			NewErrorMessageConstraint(
				"Variable 'Fields[0]' with empty 'Name' has invalid type: %T",
				model.Fields[0].Spec,
			),
		)
}

func TestStructSpec_Validate_WithoutNameAndWithInvalidFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
			{
				Spec: &FuncSpec{},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(
			NewErrorMessageConstraint(
				"Variable 'Fields[0]' with empty 'Name' has invalid type: %T",
				model.Fields[0].Spec,
			),
		)
}

func TestStructSpec_Validate_WithoutNameAndWithInvalidTypeValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{
		Fields: []*Field{
			{
				Spec: NewSpecMock(ctrl),
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Spec' has invalid type: %T", model.Fields[0].Spec))
}

func TestStructSpec_String_WithCommentAndNameAndTag(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	comment := "comment\nhere"
	name := "name"
	fieldSpecString := "fieldSpecString"
	tag := "tag"
	expected := `struct{
// comment
// here
name fieldSpecString "tag"
}`

	fieldSpec := NewSpecMock(ctrl)

	model := &StructSpec{
		Fields: []*Field{
			{
				Name:    name,
				Tag:     tag,
				Comment: comment,
				Spec:    fieldSpec,
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

func TestStructSpec_String_WithCommentAndName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	comment := "comment\nhere"
	name := "name"
	fieldSpecString := "fieldSpecString"
	expected := `struct{
// comment
// here
name fieldSpecString
}`

	fieldSpec := NewSpecMock(ctrl)

	model := &StructSpec{
		Fields: []*Field{
			{
				Name:    name,
				Comment: comment,
				Spec:    fieldSpec,
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

func TestStructSpec_String_WithComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	comment := "comment\nhere"
	fieldSpecString := "fieldSpecString"
	expected := `struct{
// comment
// here
 fieldSpecString
}`

	fieldSpec := NewSpecMock(ctrl)

	model := &StructSpec{
		Fields: []*Field{
			{
				Comment: comment,
				Spec:    fieldSpec,
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

func TestStructSpec_String_WithCommentAndTag(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	comment := "comment\nhere"
	fieldSpecString := "fieldSpecString"
	tag := "tag"
	expected := `struct{
// comment
// here
 fieldSpecString "tag"
}`

	fieldSpec := NewSpecMock(ctrl)

	model := &StructSpec{
		Fields: []*Field{
			{
				Tag:     tag,
				Comment: comment,
				Spec:    fieldSpec,
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

func TestStructSpec_String_WithNameAndTag(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "name"
	fieldSpecString := "fieldSpecString"
	tag := "tag"
	expected := `struct{
name fieldSpecString "tag"
}`

	fieldSpec := NewSpecMock(ctrl)

	model := &StructSpec{
		Fields: []*Field{
			{
				Name: name,
				Tag:  tag,
				Spec: fieldSpec,
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

func TestStructSpec_String_WithName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "name"
	fieldSpecString := "fieldSpecString"
	expected := `struct{
name fieldSpecString
}`

	fieldSpec := NewSpecMock(ctrl)

	model := &StructSpec{
		Fields: []*Field{
			{
				Name: name,
				Spec: fieldSpec,
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

func TestStructSpec_String_WithTag(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fieldSpecString := "fieldSpecString"
	tag := "tag"
	expected := `struct{
 fieldSpecString "tag"
}`

	fieldSpec := NewSpecMock(ctrl)

	model := &StructSpec{
		Fields: []*Field{
			{
				Tag:  tag,
				Spec: fieldSpec,
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

func TestStructSpec_String_WithFunSpecFieldAndCommentAndNameAndTag(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	comment := "comment\nhere"
	name := "name"
	fieldSpecString := "fieldSpecString"
	tag := "tag"
	expected := `struct{
// comment
// here
name func (fieldSpecString) "tag"
}`

	fieldSpec := NewSpecMock(ctrl)

	model := &StructSpec{
		Fields: []*Field{
			{
				Name:    name,
				Tag:     tag,
				Comment: comment,
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: fieldSpec,
						},
					},
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

func TestStructSpec_String_WithFunSpecFieldCommentAndName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	comment := "comment\nhere"
	name := "name"
	fieldSpecString := "fieldSpecString"
	expected := `struct{
// comment
// here
name func (fieldSpecString)
}`

	fieldSpec := NewSpecMock(ctrl)

	model := &StructSpec{
		Fields: []*Field{
			{
				Name:    name,
				Comment: comment,
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: fieldSpec,
						},
					},
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

func TestStructSpec_String_WithFunSpecFieldNameAndTag(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "name"
	fieldSpecString := "fieldSpecString"
	tag := "tag"
	expected := `struct{
name func (fieldSpecString) "tag"
}`

	fieldSpec := NewSpecMock(ctrl)

	model := &StructSpec{
		Fields: []*Field{
			{
				Name: name,
				Tag:  tag,
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: fieldSpec,
						},
					},
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

func TestStructSpec_String_WithFunSpecFieldAndName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	name := "name"
	fieldSpecString := "fieldSpecString"
	expected := `struct{
name func (fieldSpecString)
}`

	fieldSpec := NewSpecMock(ctrl)

	model := &StructSpec{
		Fields: []*Field{
			{
				Name: name,
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Spec: fieldSpec,
						},
					},
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

func TestStructSpec_String_WithoutFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "struct{}"

	model := &StructSpec{}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestStructSpec_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fieldSpec := NewSpecMock(ctrl)
	clonedFieldSpec := NewSpecMock(ctrl)

	model := &StructSpec{
		Fields: []*Field{
			{
				Annotations: []interface{}{
					&TestAnnotation{
						Name: "fieldAnnotation",
					},
				},
				Spec: fieldSpec,
			},
		},
	}

	fieldSpec.
		EXPECT().
		Clone().
		Return(clonedFieldSpec)

	actual := model.Clone()

	ctrl.AssertEqual(model, actual, unit.IgnoreUnexportedOption{Value: SpecMock{}})
	ctrl.AssertNotSame(model, actual)
	ctrl.AssertSame(clonedFieldSpec, actual.(*StructSpec).Fields[0].Spec)
	ctrl.AssertNotSame(model.Fields[0].Annotations[0], actual.(*StructSpec).Fields[0].Annotations[0])
}

func TestStructSpec_Clone_WithoutFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &StructSpec{}

	actual := model.Clone()

	ctrl.AssertEqual(model, actual)
}

func TestStructSpec_EqualSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Field1Spec := NewSpecMock(ctrl)
	model1Field2Spec := NewSpecMock(ctrl)
	model1Field3Spec := NewSpecMock(ctrl)
	model2Field1Spec := NewSpecMock(ctrl)
	model2Field2Spec := NewSpecMock(ctrl)
	model2Field3Spec := NewSpecMock(ctrl)

	model1 := &StructSpec{
		Fields: []*Field{
			{
				Spec: model1Field1Spec,
			},
			{
				Spec: model1Field2Spec,
			},
			{
				Spec: model1Field3Spec,
			},
		},
	}

	model2 := &StructSpec{
		Fields: []*Field{
			{
				Spec: model2Field1Spec,
			},
			{
				Spec: model2Field2Spec,
			},
			{
				Spec: model2Field3Spec,
			},
		},
	}

	model1Field1Spec.
		EXPECT().
		EqualSpec(ctrl.Same(model2Field1Spec)).
		Return(true)

	model1Field2Spec.
		EXPECT().
		EqualSpec(ctrl.Same(model2Field2Spec)).
		Return(true)

	model1Field3Spec.
		EXPECT().
		EqualSpec(ctrl.Same(model2Field3Spec)).
		Return(true)

	actual := model1.EqualSpec(model2)

	ctrl.AssertTrue(actual)
}

func TestStructSpec_EqualSpec_WithoutOrder(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Field1Spec := NewSpecMock(ctrl)
	model1Field2Spec := NewSpecMock(ctrl)
	model1Field3Spec := NewSpecMock(ctrl)
	model2Field1Spec := NewSpecMock(ctrl)
	model2Field2Spec := NewSpecMock(ctrl)
	model2Field3Spec := NewSpecMock(ctrl)

	model1 := &StructSpec{
		Fields: []*Field{
			{
				Spec: model1Field1Spec,
			},
			{
				Spec: model1Field2Spec,
			},
			{
				Spec: model1Field3Spec,
			},
		},
	}

	model2 := &StructSpec{
		Fields: []*Field{
			{
				Spec: model2Field3Spec,
			},
			{
				Spec: model2Field2Spec,
			},
			{
				Spec: model2Field1Spec,
			},
		},
	}

	model1Field1Spec.
		EXPECT().
		EqualSpec(ctrl.Same(model2Field3Spec)).
		Return(false)

	model1Field1Spec.
		EXPECT().
		EqualSpec(ctrl.Same(model2Field2Spec)).
		Return(false)

	model1Field1Spec.
		EXPECT().
		EqualSpec(ctrl.Same(model2Field1Spec)).
		Return(true)

	model1Field2Spec.
		EXPECT().
		EqualSpec(ctrl.Same(model2Field3Spec)).
		Return(false)

	model1Field2Spec.
		EXPECT().
		EqualSpec(ctrl.Same(model2Field2Spec)).
		Return(true)

	model1Field3Spec.
		EXPECT().
		EqualSpec(ctrl.Same(model2Field3Spec)).
		Return(true)

	actual := model1.EqualSpec(model2)

	ctrl.AssertTrue(actual)
}

func TestStructSpec_EqualSpec_WithFieldsLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Field1Spec := NewSpecMock(ctrl)
	model1Field2Spec := NewSpecMock(ctrl)
	model1Field3Spec := NewSpecMock(ctrl)
	model2Field2Spec := NewSpecMock(ctrl)
	model2Field3Spec := NewSpecMock(ctrl)

	model1 := &StructSpec{
		Fields: []*Field{
			{
				Spec: model1Field1Spec,
			},
			{
				Spec: model1Field2Spec,
			},
			{
				Spec: model1Field3Spec,
			},
		},
	}

	model2 := &StructSpec{
		Fields: []*Field{
			{
				Spec: model2Field3Spec,
			},
			{
				Spec: model2Field2Spec,
			},
		},
	}

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestStructSpec_EqualSpec_WithSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model1Field1Spec := NewSpecMock(ctrl)
	model1Field2Spec := NewSpecMock(ctrl)
	model1Field3Spec := NewSpecMock(ctrl)
	model2Field1Spec := NewSpecMock(ctrl)
	model2Field2Spec := NewSpecMock(ctrl)
	model2Field3Spec := NewSpecMock(ctrl)

	model1 := &StructSpec{
		Fields: []*Field{
			{
				Spec: model1Field1Spec,
			},
			{
				Spec: model1Field2Spec,
			},
			{
				Spec: model1Field3Spec,
			},
		},
	}

	model2 := &StructSpec{
		Fields: []*Field{
			{
				Spec: model2Field1Spec,
			},
			{
				Spec: model2Field2Spec,
			},
			{
				Spec: model2Field3Spec,
			},
		},
	}

	model1Field1Spec.
		EXPECT().
		EqualSpec(ctrl.Same(model2Field1Spec)).
		Return(false)

	model1Field1Spec.
		EXPECT().
		EqualSpec(ctrl.Same(model2Field2Spec)).
		Return(false)

	model1Field1Spec.
		EXPECT().
		EqualSpec(ctrl.Same(model2Field3Spec)).
		Return(false)

	actual := model1.EqualSpec(model2)

	ctrl.AssertFalse(actual)
}

func TestStructSpec_FetchImports(t *testing.T) {
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

	model := &StructSpec{
		Fields: []*Field{
			{
				Spec: valueSpec,
			},
		},
	}

	valueSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(expected)

	actual := model.FetchImports(file)

	ctrl.AssertSame(expected, actual)
}

func TestStructSpec_RenameImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	valueSpec := NewSpecMock(ctrl)

	model := &StructSpec{
		Fields: []*Field{
			{
				Spec: valueSpec,
			},
		},
	}

	valueSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	model.RenameImports(oldAlias, newAlias)
}

func TestStructSpec_RenameImports_InvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	model := &StructSpec{}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestStructSpec_RenameImports_InvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "packageName"
	newAlias := "+invalid"

	model := &StructSpec{}

	ctrl.Subtest("").
		Call(model.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
