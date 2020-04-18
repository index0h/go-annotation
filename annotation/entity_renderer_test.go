package annotation

import (
	"go/scanner"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewEntityRenderer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := &EntityRenderer{}

	actual := NewEntityRenderer()

	ctrl.AssertEqual(expected, actual)
}

func TestEntityRenderer_Render_WithSimpleSpecAndPackageNameAndTypeNameAndIsPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "*packageName.typeName"

	entity := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
		IsPointer:   true,
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithSimpleSpecAndPackageNameAndTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "packageName.typeName"

	entity := &SimpleSpec{
		PackageName: "packageName",
		TypeName:    "typeName",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithSimpleSpecAndTypeNameAndIsPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "*typeName"

	entity := &SimpleSpec{
		TypeName:  "typeName",
		IsPointer: true,
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithSimpleSpecAndTypeName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "typeName"

	entity := &SimpleSpec{
		TypeName: "typeName",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithArraySpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "[]valueTypeName"

	entity := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "valueTypeName",
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithArraySpecAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "[10]valueTypeName"

	entity := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "valueTypeName",
		},
		Length: "10",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithArraySpecAndFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "[]func ()"

	entity := &ArraySpec{
		Value: &FuncSpec{},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithMapSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "map[keyTypeName]valueTypeName"

	entity := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "keyTypeName",
		},
		Value: &SimpleSpec{
			TypeName: "valueTypeName",
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithMapSpecAndFuncSpecKey(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "map[func (keyParamTypeName)]valueTypeName"

	entity := &MapSpec{
		Key: &FuncSpec{
			Params: []*Field{
				{
					Spec: &SimpleSpec{
						TypeName: "keyParamTypeName",
					},
				},
			},
		},
		Value: &SimpleSpec{
			TypeName: "valueTypeName",
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithMapSpecAndFuncSpecValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "map[keyTypeName]func (valueParamTypeName)"

	entity := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "keyTypeName",
		},
		Value: &FuncSpec{
			Params: []*Field{
				{
					Spec: &SimpleSpec{
						TypeName: "valueParamTypeName",
					},
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `()`

	entity := &FuncSpec{}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncSpecAndParamSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `(paramTypeName)`

	entity := &FuncSpec{
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "paramTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncSpecAndParamName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `(paramName paramTypeName)`

	entity := &FuncSpec{
		Params: []*Field{
			{
				Name: "paramName",
				Spec: &SimpleSpec{
					TypeName: "paramTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncSpecAndParamComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `(
// param
// comment
paramTypeName)`

	entity := &FuncSpec{
		Params: []*Field{
			{
				Comment: "param\ncomment",
				Spec: &SimpleSpec{
					TypeName: "paramTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncSpecAndParamNameAndParamComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `(
// param
// comment
paramName paramTypeName)`

	entity := &FuncSpec{
		Params: []*Field{
			{
				Name:    "paramName",
				Comment: "param\ncomment",
				Spec: &SimpleSpec{
					TypeName: "paramTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncSpecAndVariadic(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `(...paramTypeName)`

	entity := &FuncSpec{
		IsVariadic: true,
		Params: []*Field{
			{
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "paramTypeName",
					},
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncSpecAndVariadicAndParamName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `(paramName ...paramTypeName)`

	entity := &FuncSpec{
		IsVariadic: true,
		Params: []*Field{
			{
				Name: "paramName",
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "paramTypeName",
					},
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncSpecAndParamSpecFuncSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `(func ())`

	entity := &FuncSpec{
		Params: []*Field{
			{
				Spec: &FuncSpec{},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncSpecAndParamSpecFuncSpecAndParamName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `(paramName func ())`

	entity := &FuncSpec{
		Params: []*Field{
			{
				Name: "paramName",
				Spec: &FuncSpec{},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncSpecAndMultipleParams(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `(param1Name param1TypeName, param2Name param2TypeName)`

	entity := &FuncSpec{
		Params: []*Field{
			{
				Name: "param1Name",
				Spec: &SimpleSpec{
					TypeName: "param1TypeName",
				},
			},
			{
				Name: "param2Name",
				Spec: &SimpleSpec{
					TypeName: "param2TypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncSpecAndResultSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `() (resultTypeName)`

	entity := &FuncSpec{
		Results: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "resultTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncSpecAndResultName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `() (resultName resultTypeName)`

	entity := &FuncSpec{
		Results: []*Field{
			{
				Name: "resultName",
				Spec: &SimpleSpec{
					TypeName: "resultTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncSpecAndResultComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `() (
// result
// comment
resultTypeName)`

	entity := &FuncSpec{
		Results: []*Field{
			{
				Comment: "result\ncomment",
				Spec: &SimpleSpec{
					TypeName: "resultTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncSpecAndResultNameAndResultComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `() (
// result
// comment
resultName resultTypeName)`

	entity := &FuncSpec{
		Results: []*Field{
			{
				Name:    "resultName",
				Comment: "result\ncomment",
				Spec: &SimpleSpec{
					TypeName: "resultTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncSpecAndResultSpecFuncSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `() (func ())`

	entity := &FuncSpec{
		Results: []*Field{
			{
				Spec: &FuncSpec{},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncSpecAndResultSpecFuncSpecAndResultName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `() (resultName func ())`

	entity := &FuncSpec{
		Results: []*Field{
			{
				Name: "resultName",
				Spec: &FuncSpec{},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncSpecAndMultipleResults(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `() (result1Name result1TypeName, result2Name result2TypeName)`

	entity := &FuncSpec{
		Results: []*Field{
			{
				Name: "result1Name",
				Spec: &SimpleSpec{
					TypeName: "result1TypeName",
				},
			},
			{
				Name: "result2Name",
				Spec: &SimpleSpec{
					TypeName: "result2TypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithInterfaceSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "interface{}"

	entity := &InterfaceSpec{}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithInterfaceSpecAndSimpleSpecField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `interface{
fieldTypeName
}`

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "fieldTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithInterfaceSpecAndSimpleSpecFieldAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `interface{
// field
// comment
fieldTypeName
}`

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Comment: "field\ncomment",
				Spec: &SimpleSpec{
					TypeName: "fieldTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithInterfaceSpecAndFuncSpecField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `interface{
fieldName()
}`

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Name: "fieldName",
				Spec: &FuncSpec{},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithInterfaceSpecAndFuncSpecFieldAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `interface{
// field
// comment
fieldName()
}`

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Name:    "fieldName",
				Comment: "field\ncomment",
				Spec:    &FuncSpec{},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithInterfaceSpecAndMultipleFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `interface{
field1TypeName
field2TypeName
}`

	entity := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "field1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field2TypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithStructSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "struct{}"

	entity := &StructSpec{}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithStructSpecAndSimpleSpecField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `struct{
fieldTypeName
}`

	entity := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "fieldTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithStructSpecAndSimpleSpecFieldAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `struct{
// field
// comment
fieldTypeName
}`

	entity := &StructSpec{
		Fields: []*Field{
			{
				Comment: "field\ncomment",
				Spec: &SimpleSpec{
					TypeName: "fieldTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithStructSpecAndSimpleSpecFieldAndNameField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `struct{
fieldName fieldTypeName
}`

	entity := &StructSpec{
		Fields: []*Field{
			{
				Name: "fieldName",
				Spec: &SimpleSpec{
					TypeName: "fieldTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithStructSpecAndSimpleSpecFieldAndNameFieldAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `struct{
// field
// comment
fieldName fieldTypeName
}`

	entity := &StructSpec{
		Fields: []*Field{
			{
				Name:    "fieldName",
				Comment: "field\ncomment",
				Spec: &SimpleSpec{
					TypeName: "fieldTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithStructSpecAndSimpleSpecFieldAndNameFieldAndTag(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `struct{
fieldName fieldTypeName "fieldTag"
}`

	entity := &StructSpec{
		Fields: []*Field{
			{
				Name: "fieldName",
				Tag:  "fieldTag",
				Spec: &SimpleSpec{
					TypeName: "fieldTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithStructSpecAndFuncSpecField(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `struct{
fieldName func ()
}`

	entity := &StructSpec{
		Fields: []*Field{
			{
				Name: "fieldName",
				Spec: &FuncSpec{},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithStructSpecAndFuncSpecFieldAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `struct{
// field
// comment
fieldName func ()
}`

	entity := &StructSpec{
		Fields: []*Field{
			{
				Name:    "fieldName",
				Comment: "field\ncomment",
				Spec:    &FuncSpec{},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithStructSpecAndMultipleFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `struct{
field1TypeName
field2TypeName
}`

	entity := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "field1TypeName",
				},
			},
			{
				Spec: &SimpleSpec{
					TypeName: "field2TypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `import "namespace"
`

	entity := &Import{
		Namespace: "namespace",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithImportAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `// import
// comment
import "namespace"
`

	entity := &Import{
		Namespace: "namespace",
		Comment:   "import\ncomment",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithImportAndAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `import alias "namespace"
`

	entity := &Import{
		Alias:     "alias",
		Namespace: "namespace",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithImportGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `import ()
`

	entity := &ImportGroup{}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithImportGroupAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `// importGroup
// comment
import ()
`

	entity := &ImportGroup{
		Comment: "importGroup\ncomment",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithImportGroupAndImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `import "namespace"
`

	entity := &ImportGroup{
		Imports: []*Import{
			{
				Namespace: "namespace",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithImportGroupAndImportWithComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `import (
// import
// comment
"namespace"
)
`

	entity := &ImportGroup{
		Imports: []*Import{
			{
				Namespace: "namespace",
				Comment:   "import\ncomment",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithImportGroupAndWithAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `import alias "namespace"
`

	entity := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     "alias",
				Namespace: "namespace",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithImportGroupAndMultipleImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `import (
"namespace1"
"namespace2"
)
`

	entity := &ImportGroup{
		Imports: []*Import{
			{
				Namespace: "namespace1",
			},
			{
				Namespace: "namespace2",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithImportGroupAndMultipleImportsAndAliases(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `import (
alias1 "namespace1"
alias2 "namespace2"
)
`

	entity := &ImportGroup{
		Imports: []*Import{
			{
				Alias:     "alias1",
				Namespace: "namespace1",
			},
			{
				Alias:     "alias2",
				Namespace: "namespace2",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithConst(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `const constName = constValue
`

	entity := &Const{
		Name:  "constName",
		Value: "constValue",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithConstAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `const constName constTypeName = constValue
`

	entity := &Const{
		Name:  "constName",
		Value: "constValue",
		Spec: &SimpleSpec{
			TypeName: "constTypeName",
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithConstAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `// const
// comment
const constName = constValue
`

	entity := &Const{
		Comment: "const\ncomment",
		Name:    "constName",
		Value:   "constValue",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithConstAndWithoutValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &Const{
		Name: "constName",
	}

	ctrl.Subtest("").
		Call((&EntityRenderer{}).Render, entity).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'Value' must be not empty"),
		)
}

func TestEntityRenderer_Render_WithConstGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `const ()
`

	entity := &ConstGroup{}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithConstGroupAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `// constGroup
// comment
const ()
`

	entity := &ConstGroup{
		Comment: "constGroup\ncomment",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithConstGroupAndConst(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `const constName = constValue
`

	entity := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "constName",
				Value: "constValue",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithConstGroupAndConstAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `const constName constTypeName = constValue
`

	entity := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "constName",
				Value: "constValue",
				Spec: &SimpleSpec{
					TypeName: "constTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithConstGroupAndConstAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `const (
// const
// comment
constName = constValue
)
`

	entity := &ConstGroup{
		Consts: []*Const{
			{
				Comment: "const\ncomment",
				Name:    "constName",
				Value:   "constValue",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithConstGroupAndConstAndWithoutValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &ConstGroup{
		Consts: []*Const{
			{
				Name: "constName",
			},
		},
	}

	ctrl.Subtest("").
		Call((&EntityRenderer{}).Render, entity).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'Value' must be not empty"),
		)
}

func TestEntityRenderer_Render_WithConstGroupAndMultipleConst(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `const (
const1Name = const1Value
const2Name
)
`

	entity := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "const1Name",
				Value: "const1Value",
			},
			{
				Name: "const2Name",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithConstGroupAndMultipleConstAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `const (
const1Name const1TypeName = const1Value
const2Name
)
`

	entity := &ConstGroup{
		Consts: []*Const{
			{
				Name:  "const1Name",
				Value: "const1Value",
				Spec: &SimpleSpec{
					TypeName: "const1TypeName",
				},
			},
			{
				Name: "const2Name",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithVar(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `var varName = varValue
`

	entity := &Var{
		Name:  "varName",
		Value: "varValue",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithVarAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `var varName varTypeName
`

	entity := &Var{
		Name: "varName",
		Spec: &SimpleSpec{
			TypeName: "varTypeName",
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithVarAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `// var
// comment
var varName = varValue
`

	entity := &Var{
		Comment: "var\ncomment",
		Name:    "varName",
		Value:   "varValue",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithVarGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `var ()
`

	entity := &VarGroup{}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithVarGroupAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `// varGroup
// comment
var ()
`

	entity := &VarGroup{
		Comment: "varGroup\ncomment",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithVarGroupAndVar(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `var varName = varValue
`

	entity := &VarGroup{
		Vars: []*Var{
			{
				Name:  "varName",
				Value: "varValue",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithVarGroupAndVarAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `var varName varTypeName
`

	entity := &VarGroup{
		Vars: []*Var{
			{
				Name: "varName",
				Spec: &SimpleSpec{
					TypeName: "varTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithVarGroupAndVarAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `var (
// var
// comment
varName = varValue
)
`

	entity := &VarGroup{
		Vars: []*Var{
			{
				Comment: "var\ncomment",
				Name:    "varName",
				Value:   "varValue",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithVarGroupAndMultipleVar(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `var (
var1Name = var1Value
var2Name = var2Value
)
`

	entity := &VarGroup{
		Vars: []*Var{
			{
				Name:  "var1Name",
				Value: "var1Value",
			},
			{
				Name:  "var2Name",
				Value: "var2Value",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithVarGroupAndMultipleVarAndSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `var (
var1Name var1TypeName
var2Name var1TypeName
)
`

	entity := &VarGroup{
		Vars: []*Var{
			{
				Name: "var1Name",
				Spec: &SimpleSpec{
					TypeName: "var1TypeName",
				},
			},
			{
				Name: "var2Name",
				Spec: &SimpleSpec{
					TypeName: "var1TypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `type typeName typeTypeName
`

	entity := &Type{
		Name: "typeName",
		Spec: &SimpleSpec{
			TypeName: "typeTypeName",
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithTypeAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `// type
// comment
type typeName typeTypeName
`

	entity := &Type{
		Comment: "type\ncomment",
		Name:    "typeName",
		Spec: &SimpleSpec{
			TypeName: "typeTypeName",
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithTypeGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `type ()
`

	entity := &TypeGroup{}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithTypeGroupAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `// typeGroup
// comment
type ()
`

	entity := &TypeGroup{
		Comment: "typeGroup\ncomment",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithTypeGroupAndType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `type typeName typeTypeName
`

	entity := &TypeGroup{
		Types: []*Type{
			{
				Name: "typeName",
				Spec: &SimpleSpec{
					TypeName: "typeTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithTypeGroupAndTypeAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `type (
// type
// comment
typeName typeTypeName
)
`

	entity := &TypeGroup{
		Types: []*Type{
			{
				Comment: "type\ncomment",
				Name:    "typeName",
				Spec: &SimpleSpec{
					TypeName: "typeTypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithTypeGroupAndMultipleType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `type (
type1Name type1TypeName
type2Name type2TypeName
)
`

	entity := &TypeGroup{
		Types: []*Type{
			{
				Name: "type1Name",
				Spec: &SimpleSpec{
					TypeName: "type1TypeName",
				},
			},
			{
				Name: "type2Name",
				Spec: &SimpleSpec{
					TypeName: "type2TypeName",
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFunc(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `func funcName() {

}
`

	entity := &Func{
		Name: "funcName",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncAndContent(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `func funcName() {
funcContent()
}
`

	entity := &Func{
		Name:    "funcName",
		Content: "funcContent()",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncAndComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `// func
// comment
func funcName() {

}
`

	entity := &Func{
		Comment: "func\ncomment",
		Name:    "funcName",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncWithRelatedType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `func (funcRelatedType) funcName() {

}
`

	entity := &Func{
		Name: "funcName",
		Related: &Field{
			Spec: &SimpleSpec{
				TypeName: "funcRelatedType",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncWithRelatedTypeAndRelatedName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `func (funcRelatedName funcRelatedType) funcName() {

}
`

	entity := &Func{
		Name: "funcName",
		Related: &Field{
			Name: "funcRelatedName",
			Spec: &SimpleSpec{
				TypeName: "funcRelatedType",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncWithRelatedTypeAndRelatedNameAndRelatedComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `func (
// funcRelated
// comment
funcRelatedName funcRelatedType) funcName() {

}
`

	entity := &Func{
		Name: "funcName",
		Related: &Field{
			Comment: "funcRelated\ncomment",
			Name:    "funcRelatedName",
			Spec: &SimpleSpec{
				TypeName: "funcRelatedType",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFuncWithSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `func funcName(paramTypeName) {

}
`

	entity := &Func{
		Name: "funcName",
		Spec: &FuncSpec{
			Params: []*Field{
				{
					Spec: &SimpleSpec{
						TypeName: "paramTypeName",
					},
				},
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFile(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `package filePackage
`

	entity := &File{
		Name:        "fileName.go",
		PackageName: "filePackage",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFileAndContent(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `// file
// comment
package filePackage
`

	entity := &File{
		Name:        "fileName.go",
		PackageName: "filePackage",
		Content:     expected,
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFileWithComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `// file
// comment
package filePackage
`

	entity := &File{
		Name:        "fileName.go",
		Comment:     "file\ncomment",
		PackageName: "filePackage",
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFileWithImportGroups(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `package filePackage

// importGroup1
// comment
import ()

// importGroup2
// comment
import ()
`

	entity := &File{
		Name:        "fileName.go",
		PackageName: "filePackage",
		ImportGroups: []*ImportGroup{
			{
				Comment: "importGroup1\ncomment",
			},
			{
				Comment: "importGroup2\ncomment",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFileWithConstGroups(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `package filePackage

// constGroup1
// comment
const ()

// constGroup2
// comment
const ()
`

	entity := &File{
		Name:        "fileName.go",
		PackageName: "filePackage",
		ConstGroups: []*ConstGroup{
			{
				Comment: "constGroup1\ncomment",
			},
			{
				Comment: "constGroup2\ncomment",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFileWithVarGroups(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `package filePackage

// varGroup1
// comment
var ()

// varGroup2
// comment
var ()
`

	entity := &File{
		Name:        "fileName.go",
		PackageName: "filePackage",
		VarGroups: []*VarGroup{
			{
				Comment: "varGroup1\ncomment",
			},
			{
				Comment: "varGroup2\ncomment",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFileWithTypeGroups(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `package filePackage

// typeGroup1
// comment
type ()

// typeGroup2
// comment
type ()
`

	entity := &File{
		Name:        "fileName.go",
		PackageName: "filePackage",
		TypeGroups: []*TypeGroup{
			{
				Comment: "typeGroup1\ncomment",
			},
			{
				Comment: "typeGroup2\ncomment",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFileWithFuncs(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `package filePackage

func func1Name() {

}

func func2Name() {

}
`

	entity := &File{
		Name:        "fileName.go",
		PackageName: "filePackage",
		Funcs: []*Func{
			{
				Name: "func1Name",
			},
			{
				Name: "func2Name",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithFileWithAllDeclarations(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `package filePackage

// importGroup
// comment
import ()

// constGroup
// comment
const ()

// varGroup
// comment
var ()

// typeGroup
// comment
type ()

func func1Name() {

}
`

	entity := &File{
		Name:        "fileName.go",
		PackageName: "filePackage",
		ImportGroups: []*ImportGroup{
			{
				Comment: "importGroup\ncomment",
			},
		},
		ConstGroups: []*ConstGroup{
			{
				Comment: "constGroup\ncomment",
			},
		},
		VarGroups: []*VarGroup{
			{
				Comment: "varGroup\ncomment",
			},
		},
		TypeGroups: []*TypeGroup{
			{
				Comment: "typeGroup\ncomment",
			},
		},
		Funcs: []*Func{
			{
				Name: "func1Name",
			},
		},
	}

	actual := (&EntityRenderer{}).Render(entity)

	ctrl.AssertSame(expected, actual)
}

func TestEntityRenderer_Render_WithUnexpectedEntity(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := "data"

	ctrl.Subtest("").
		Call((&EntityRenderer{}).Render, entity).
		ExpectPanic(NewErrorMessageConstraint("Can't render entity with type: '%T'", entity))
}

func TestEntityRenderer_Render_WithInvalidFormattedResult(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	entity := &File{
		Name:        "file.go",
		PackageName: "+invalid",
	}

	ctrl.Subtest("").
		Call((&EntityRenderer{}).Render, entity).
		ExpectPanic(ctrl.Type(scanner.ErrorList{}))
}
