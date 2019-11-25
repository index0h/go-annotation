package annotation

import (
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestTypeGroup_Validate_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &TypeGroup{}

	model.Validate()
}

func TestTypeGroup_Validate_WithSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &TypeGroup{
		Comment: "typeGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "typeGroupAnnotation",
			},
		},
		Types: []*Type{
			{
				Name: "name",
				Spec: &SimpleSpec{
					TypeName: "typeName",
				},
			},
		},
	}

	model.Validate()
}

func TestTypeGroup_Validate_WithArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &TypeGroup{
		Comment: "typeGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "typeGroupAnnotation",
			},
		},
		Types: []*Type{
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

func TestTypeGroup_Validate_WithMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &TypeGroup{
		Comment: "typeGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "typeGroupAnnotation",
			},
		},
		Types: []*Type{
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

func TestTypeGroup_Validate_WithStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &TypeGroup{
		Comment: "typeGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "typeGroupAnnotation",
			},
		},
		Types: []*Type{
			{
				Name: "name",
				Spec: &StructSpec{},
			},
		},
	}

	model.Validate()
}

func TestTypeGroup_Validate_WithInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &TypeGroup{
		Comment: "typeGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "typeGroupAnnotation",
			},
		},
		Types: []*Type{
			{
				Name: "name",
				Spec: &InterfaceSpec{},
			},
		},
	}

	model.Validate()
}

func TestTypeGroup_Validate_WithFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &TypeGroup{
		Comment: "typeGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "typeGroupAnnotation",
			},
		},
		Types: []*Type{
			{
				Name: "name",
				Spec: &FuncSpec{},
			},
		},
	}

	model.Validate()
}

func TestTypeGroup_Validate_WithInvalidType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &TypeGroup{
		Comment: "typeGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "typeGroupAnnotation",
			},
		},
		Types: []*Type{
			{
				Spec: &FuncSpec{},
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestTypeGroup_Validate_WithNilType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &TypeGroup{
		Comment: "typeGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "typeGroupAnnotation",
			},
		},
		Types: []*Type{
			nil,
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Types[0]' must be not nil"))
}

func TestTypeGroup_Validate_WithInvalidSimpleSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &TypeGroup{
		Types: []*Type{
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

func TestTypeGroup_Validate_WithInvalidArraySpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &TypeGroup{
		Types: []*Type{
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

func TestTypeGroup_Validate_WithInvalidMapSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &TypeGroup{
		Types: []*Type{
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

func TestTypeGroup_Validate_WithInvalidStructSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &TypeGroup{
		Types: []*Type{
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

func TestTypeGroup_Validate_WithInvalidInterfaceSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &TypeGroup{
		Types: []*Type{
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

func TestTypeGroup_Validate_WithInvalidFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &TypeGroup{
		Types: []*Type{
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

func TestTypeGroup_Validate_WithInvalidTypeValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &TypeGroup{
		Types: []*Type{
			{
				Name: "name",
				Spec: NewSpecMock(ctrl),
			},
		},
	}

	ctrl.Subtest("").
		Call(model.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Spec' has invalid type: %T", model.Types[0].Spec))
}

func TestTypeGroup_String(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `type (
)
`

	model := &TypeGroup{}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestTypeGroup_String_WithTypeGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	typeGroupComment := "typeGroup\ncomment"

	expected := `// typeGroup
// comment
type (
)
`

	model := &TypeGroup{
		Comment: typeGroupComment,
	}

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestTypeGroup_String_WithOneTypeAndTypeComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	typeComment := "type\ncomment"
	typeName := "typeName"
	typeSpecValue := "typeSpecValue"

	expected := `type (
// type
// comment
typeName typeSpecValue
)
`

	typeSpec := NewSpecMock(ctrl)

	model := &TypeGroup{
		Types: []*Type{
			{
				Comment: typeComment,
				Name:    typeName,
				Spec:    typeSpec,
			},
		},
	}

	typeSpec.
		EXPECT().
		String().
		Return(typeSpecValue)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestTypeGroup_String_WithOneTypeAndTypeGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	typeGroupComment := "typeGroup\ncomment"
	typeName := "typeName"
	typeSpecValue := "typeSpecValue"

	expected := `// typeGroup
// comment
type typeName typeSpecValue
`

	typeSpec := NewSpecMock(ctrl)

	model := &TypeGroup{
		Comment: typeGroupComment,
		Types: []*Type{
			{
				Name: typeName,
				Spec: typeSpec,
			},
		},
	}

	typeSpec.
		EXPECT().
		String().
		Return(typeSpecValue)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestTypeGroup_String_WithOneTypeAndTypeCommentAndTypeGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	typeGroupComment := "typeGroup\ncomment"
	typeComment := "type\ncomment"
	typeName := "typeName"
	typeSpecValue := "typeSpecValue"

	expected := `// typeGroup
// comment
type (
// type
// comment
typeName typeSpecValue
)
`

	typeSpec := NewSpecMock(ctrl)

	model := &TypeGroup{
		Comment: typeGroupComment,
		Types: []*Type{
			{
				Comment: typeComment,
				Name:    typeName,
				Spec:    typeSpec,
			},
		},
	}

	typeSpec.
		EXPECT().
		String().
		Return(typeSpecValue)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestTypeGroup_String_WithOneType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	typeName := "typeName"
	typeSpecValue := "typeSpecValue"

	expected := `type typeName typeSpecValue
`

	typeSpec := NewSpecMock(ctrl)

	model := &TypeGroup{
		Types: []*Type{
			{
				Name: typeName,
				Spec: typeSpec,
			},
		},
	}

	typeSpec.
		EXPECT().
		String().
		Return(typeSpecValue)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestTypeGroup_String_WithMultipleTypes(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type1Name := "type1Name"
	type2Name := "type2Name"
	typeSpec1Value := "typeSpec1Value"
	typeSpec2Value := "typeSpec2Value"

	expected := `type (
type1Name typeSpec1Value
type2Name typeSpec2Value
)
`

	type1Spec := NewSpecMock(ctrl)
	type2Spec := NewSpecMock(ctrl)

	model := &TypeGroup{
		Types: []*Type{
			{
				Name: type1Name,
				Spec: type1Spec,
			},
			{
				Name: type2Name,
				Spec: type2Spec,
			},
		},
	}

	type1Spec.
		EXPECT().
		String().
		Return(typeSpec1Value)

	type2Spec.
		EXPECT().
		String().
		Return(typeSpec2Value)

	actual := model.String()

	ctrl.AssertSame(expected, actual)
}

func TestTypeGroup_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	type1Spec := NewSpecMock(ctrl)
	type2Spec := NewSpecMock(ctrl)

	clonedType1Spec := NewSpecMock(ctrl)
	clonedType2Spec := NewSpecMock(ctrl)

	model := &TypeGroup{
		Comment: "typeGroupComment",
		Annotations: []interface{}{
			&TestAnnotation{
				Name: "typeGroupAnnotation",
			},
		},
		Types: []*Type{
			{
				Name:    "type1Name",
				Comment: "type1\ncomment",
				Annotations: []interface{}{
					&TestAnnotation{
						Name: "type1Annotation",
					},
				},
				Spec: type1Spec,
			},
			{
				Name:    "type2Name",
				Comment: "type2\ncomment",
				Annotations: []interface{}{
					&TestAnnotation{
						Name: "type2Annotation",
					},
				},
				Spec: type2Spec,
			},
		},
	}

	type1Spec.
		EXPECT().
		Clone().
		Return(clonedType1Spec)

	type2Spec.
		EXPECT().
		Clone().
		Return(clonedType2Spec)

	actual := model.Clone()

	ctrl.AssertEqual(model, actual, unit.IgnoreUnexportedOption{Value: SpecMock{}})
	ctrl.AssertNotSame(model, actual)
	ctrl.AssertNotSame(model.Annotations[0], actual.(*TypeGroup).Annotations[0])
	ctrl.AssertNotSame(model.Types[0], actual.(*TypeGroup).Types[0])
	ctrl.AssertNotSame(model.Types[0].Annotations[0], actual.(*TypeGroup).Types[0].Annotations[0])
	ctrl.AssertSame(clonedType1Spec, actual.(*TypeGroup).Types[0].Spec)
	ctrl.AssertNotSame(model.Types[1], actual.(*TypeGroup).Types[1])
	ctrl.AssertNotSame(model.Types[1].Annotations[0], actual.(*TypeGroup).Types[1].Annotations[0])
	ctrl.AssertSame(clonedType2Spec, actual.(*TypeGroup).Types[1].Spec)
}

func TestTypeGroup_Clone_WithoutFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	model := &TypeGroup{}

	actual := model.Clone()

	ctrl.AssertEqual(model, actual)
	ctrl.AssertNotSame(model, actual)
}

func TestTypeGroup_FetchImports_WithFoundImport(t *testing.T) {
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

	typeSpec := NewSpecMock(ctrl)

	model := &TypeGroup{
		Types: []*Type{
			{
				Name: "typeName",
				Spec: typeSpec,
			},
		},
	}

	typeSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(expected)

	actual := model.FetchImports(file)

	ctrl.AssertEqual(expected, actual)
}

func TestTypeGroup_FetchImports_WithNotFoundImport(t *testing.T) {
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

	typeSpec := NewSpecMock(ctrl)

	model := &TypeGroup{
		Types: []*Type{
			{
				Name: "typeName",
				Spec: typeSpec,
			},
		},
	}

	typeSpec.
		EXPECT().
		FetchImports(ctrl.Same(file)).
		Return(nil)

	actual := model.FetchImports(file)

	ctrl.AssertEmpty(actual)
}

func TestTypeGroup_RenameImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	typeSpec := NewSpecMock(ctrl)

	model := &TypeGroup{
		Types: []*Type{
			{
				Name: "typeName",
				Spec: typeSpec,
			},
		},
	}

	typeSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	model.RenameImports(oldAlias, newAlias)
}

func TestTypeGroup_RenameImports_WithInvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	model := &TypeGroup{
		Types: []*Type{
			{
				Name: "typeName",
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
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

func TestTypeGroup_RenameImports_WithInvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "+invalid"

	model := &TypeGroup{
		Types: []*Type{
			{
				Name: "typeName",
				Spec: &SimpleSpec{
					PackageName: "packageName",
					TypeName:    "typeName",
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
