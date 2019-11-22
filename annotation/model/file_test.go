package model

import (
	"go/scanner"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestFile_Validate(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Content:     "package filePackage",
		Comment:     "fileComment",
		Annotations: []interface{}{
			&SimpleSpec{
				TypeName: "fileAnnotation",
			},
		},
		ImportGroups: []*ImportGroup{
			{
				Comment: "importGroupComment",
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "importGroupAnnotation",
					},
				},
				Imports: []*Import{
					{
						Alias:     "importAlias",
						Namespace: "importNamespace",
						Comment:   "importComment",
						Annotations: []interface{}{
							&SimpleSpec{
								TypeName: "importAnnotation",
							},
						},
					},
				},
			},
		},
		ConstGroups: []*ConstGroup{
			{
				Comment: "constGroupComment",
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "constGroupAnnotation",
					},
				},
				Consts: []*Const{
					{
						Name:    "constName",
						Value:   "constValue",
						Comment: "constComment",
						Annotations: []interface{}{
							&SimpleSpec{
								TypeName: "constAnnotation",
							},
						},
						Spec: &SimpleSpec{
							PackageName: "constSpecPackageName",
							TypeName:    "constSpecTypeName",
						},
					},
				},
			},
		},
		VarGroups: []*VarGroup{
			{
				Comment: "varGroupComment",
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "varGroupAnnotation",
					},
				},
				Vars: []*Var{
					{
						Name:    "varName",
						Value:   "varValue",
						Comment: "varComment",
						Annotations: []interface{}{
							&SimpleSpec{
								TypeName: "varAnnotation",
							},
						},
						Spec: &SimpleSpec{
							PackageName: "varSpecPackageName",
							TypeName:    "varSpecTypeName",
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{
			{
				Comment: "typeGroupComment",
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "typeGroupAnnotation",
					},
				},
				Types: []*Type{
					{
						Name:    "typeName",
						Comment: "typeComment",
						Annotations: []interface{}{
							&SimpleSpec{
								TypeName: "typeAnnotation",
							},
						},
						Spec: &SimpleSpec{
							PackageName: "typeSpecPackageName",
							TypeName:    "typeSpecTypeName",
						},
					},
				},
			},
		},
		Funcs: []*Func{
			{
				Name:    "funcName",
				Content: "funcContent",
				Comment: "funcComment",
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "funcAnnotation",
					},
				},
				Spec: &FuncSpec{},
				Related: &Field{
					Name: "funcRelatedName",
					Spec: &SimpleSpec{
						TypeName: "relatedTypeName",
					},
				},
			},
		},
	}

	modelValue.Validate()
}

func TestFile_Validate_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
	}

	modelValue.Validate()
}

func TestFile_Validate_WithEmptyName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		PackageName: "filePackageName",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestFile_Validate_WithEmptyPackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		Name: "fileName",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'PackageName' must be not empty"))
}

func TestFile_Validate_WithInvalidPackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		Name:        "fileName",
		PackageName: "+invalid",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'PackageName' must be valid identifier, actual value: '+invalid'"))
}

func TestFile_Validate_WithNilImportGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		ImportGroups: []*ImportGroup{
			nil,
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'ImportGroups[0]' must be not nil"))
}

func TestFile_Validate_WithInvalidImportGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Namespace' must be not empty"))
}

func TestFile_Validate_WithNilConstGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		ConstGroups: []*ConstGroup{
			nil,
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'ConstGroups[0]' must be not nil"))
}

func TestFile_Validate_WithInvalidConstGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestFile_Validate_WithNilVarGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		VarGroups: []*VarGroup{
			nil,
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'VarGroups[0]' must be not nil"))
}

func TestFile_Validate_WithInvalidVarGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestFile_Validate_WithNilTypeGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		TypeGroups: []*TypeGroup{
			nil,
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'TypeGroups[0]' must be not nil"))
}

func TestFile_Validate_WithInvalidTypeGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		TypeGroups: []*TypeGroup{
			{
				Types: []*Type{
					{},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestFile_Validate_WithNilFunc(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Funcs: []*Func{
			nil,
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Funcs[0]' must be not nil"))
}

func TestFile_Validate_WithInvalidFunc(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Funcs: []*Func{
			{},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(NewErrorMessageConstraint("Variable 'Name' must be not empty"))
}

func TestFile_Validate_WithInvalidContent(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Content:     "+invalid",
	}

	ctrl.Subtest("").
		Call(modelValue.Validate).
		ExpectPanic(ctrl.Type(scanner.ErrorList{}))
}

func TestFile_String(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	varSpecString := "varSpecString"
	typeSpecString := "typeSpecString"
	funcParamSpecString := "funcParamSpecString"
	funcResultSpecString := "funcResultSpecString"
	funcRelatedSpecString := "funcRelatedSpecString"

	expected := `// file
// Comment
package filePackageName

// importGroup
// Comment
import (
	// import
	// Comment
	importAlias "importNamespace"
)

// constGroup
// Comment
const (
	// const
	// Comment
	constName constPackageName.constTypeName = constValue
)

// varGroup
// Comment
var (
	// var
	// Comment
	varName varSpecString = varValue
)

// typeGroup
// Comment
type (
	// type
	// Comment
	typeName typeSpecString
)

// func
// Comment
func (
	// funcRelated
	// Comment
	funcRelatedName funcRelatedSpecString) funcName(
	// funcParam
	// Comment
	funcParamName func(funcParamSpecString)) (
	// funcResult
	// Comment
	funcResultName func(funcResultSpecString)) {
	funcContent
}
`

	varSpec := NewSpecMock(ctrl)
	typeSpec := NewSpecMock(ctrl)
	funcParamSpec := NewSpecMock(ctrl)
	funcResultSpec := NewSpecMock(ctrl)
	funcRelatedSpec := NewSpecMock(ctrl)

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Comment:     "file\nComment",
		Annotations: []interface{}{
			&SimpleSpec{
				TypeName: "fileAnnotation",
			},
		},
		ImportGroups: []*ImportGroup{
			{
				Comment: "importGroup\nComment",
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "importGroupAnnotation",
					},
				},
				Imports: []*Import{
					{
						Alias:     "importAlias",
						Namespace: "importNamespace",
						Comment:   "import\nComment",
						Annotations: []interface{}{
							&SimpleSpec{
								TypeName: "importAnnotation",
							},
						},
					},
				},
			},
		},
		ConstGroups: []*ConstGroup{
			{
				Comment: "constGroup\nComment",
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "constGroupAnnotation",
					},
				},
				Consts: []*Const{
					{
						Name:    "constName",
						Value:   "constValue",
						Comment: "const\nComment",
						Annotations: []interface{}{
							&SimpleSpec{
								TypeName: "constAnnotation",
							},
						},
						Spec: &SimpleSpec{
							PackageName: "constPackageName",
							TypeName:    "constTypeName",
						},
					},
				},
			},
		},
		VarGroups: []*VarGroup{
			{
				Comment: "varGroup\nComment",
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "varGroupAnnotation",
					},
				},
				Vars: []*Var{
					{
						Name:    "varName",
						Value:   "varValue",
						Comment: "var\nComment",
						Annotations: []interface{}{
							&SimpleSpec{
								TypeName: "varAnnotation",
							},
						},
						Spec: varSpec,
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{
			{
				Comment: "typeGroup\nComment",
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "typeGroupAnnotation",
					},
				},
				Types: []*Type{
					{
						Name:    "typeName",
						Comment: "type\nComment",
						Annotations: []interface{}{
							&SimpleSpec{
								TypeName: "typeAnnotation",
							},
						},
						Spec: typeSpec,
					},
				},
			},
		},
		Funcs: []*Func{
			{
				Name:    "funcName",
				Content: "funcContent",
				Comment: "func\nComment",
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "funcAnnotation",
					},
				},
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Name:    "funcParamName",
							Comment: "funcParam\nComment",
							Annotations: []interface{}{
								&SimpleSpec{
									TypeName: "funcParamAnnotation",
								},
							},
							Spec: &FuncSpec{
								Params: []*Field{
									{
										Spec: funcParamSpec,
									},
								},
							},
						},
					},
					Results: []*Field{
						{
							Name:    "funcResultName",
							Comment: "funcResult\nComment",
							Annotations: []interface{}{
								&SimpleSpec{
									TypeName: "funcResultAnnotation",
								},
							},
							Spec: &FuncSpec{
								Params: []*Field{
									{
										Spec: funcResultSpec,
									},
								},
							},
						},
					},
				},
				Related: &Field{
					Name:    "funcRelatedName",
					Comment: "funcRelated\nComment",
					Annotations: []interface{}{
						&SimpleSpec{
							TypeName: "funcRelatedAnnotation",
						},
					},
					Spec: funcRelatedSpec,
				},
			},
		},
	}

	varSpec.
		EXPECT().
		String().
		Return(varSpecString)

	typeSpec.
		EXPECT().
		String().
		Return(typeSpecString)

	funcParamSpec.
		EXPECT().
		String().
		Return(funcParamSpecString)

	funcResultSpec.
		EXPECT().
		String().
		Return(funcResultSpecString)

	funcRelatedSpec.
		EXPECT().
		String().
		Return(funcRelatedSpecString)

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFile_String_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := `package filePackageName
`

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFile_String_WithExistedContent(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	expected := "fileContent"

	modelValue := &File{
		Name:        "fileName",
		Content:     expected,
		PackageName: "filePackageName",
	}

	actual := modelValue.String()

	ctrl.AssertSame(expected, actual)
}

func TestFile_String_WithInvalidContent(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:  "varName",
						Value: "[invalid",
					},
				},
			},
		},
	}

	ctrl.Subtest("").
		Call(modelValue.String).
		ExpectPanic(ctrl.Type(scanner.ErrorList{}))
}

func TestFile_Clone(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	varSpec := NewSpecMock(ctrl)
	typeSpec := NewSpecMock(ctrl)
	funcParamSpec := NewSpecMock(ctrl)
	funcResultSpec := NewSpecMock(ctrl)
	funcRelatedSpec := NewSpecMock(ctrl)

	clonedVarSpec := NewSpecMock(ctrl)
	clonedTypeSpec := NewSpecMock(ctrl)
	clonedFuncParamSpec := NewSpecMock(ctrl)
	clonedFuncResultSpec := NewSpecMock(ctrl)
	clonedFuncRelatedSpec := NewSpecMock(ctrl)

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Comment:     "file\nComment",
		Annotations: []interface{}{
			&SimpleSpec{
				TypeName: "fileAnnotation",
			},
		},
		ImportGroups: []*ImportGroup{
			{
				Comment: "importGroup\nComment",
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "importGroupAnnotation",
					},
				},
				Imports: []*Import{
					{
						Alias:     "importAlias",
						Namespace: "importNamespace",
						Comment:   "import\nComment",
						Annotations: []interface{}{
							&SimpleSpec{
								TypeName: "importAnnotation",
							},
						},
					},
				},
			},
		},
		ConstGroups: []*ConstGroup{
			{
				Comment: "constGroup\nComment",
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "constGroupAnnotation",
					},
				},
				Consts: []*Const{
					{
						Name:    "constName",
						Value:   "constValue",
						Comment: "const\nComment",
						Annotations: []interface{}{
							&SimpleSpec{
								TypeName: "constAnnotation",
							},
						},
						Spec: &SimpleSpec{
							PackageName: "constPackageName",
							TypeName:    "constTypeName",
						},
					},
				},
			},
		},
		VarGroups: []*VarGroup{
			{
				Comment: "varGroup\nComment",
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "varGroupAnnotation",
					},
				},
				Vars: []*Var{
					{
						Name:    "varName",
						Value:   "varValue",
						Comment: "var\nComment",
						Annotations: []interface{}{
							&SimpleSpec{
								TypeName: "varAnnotation",
							},
						},
						Spec: varSpec,
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{
			{
				Comment: "typeGroup\nComment",
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "typeGroupAnnotation",
					},
				},
				Types: []*Type{
					{
						Name:    "typeName",
						Comment: "type\nComment",
						Annotations: []interface{}{
							&SimpleSpec{
								TypeName: "typeAnnotation",
							},
						},
						Spec: typeSpec,
					},
				},
			},
		},
		Funcs: []*Func{
			{
				Name:    "funcName",
				Content: "funcContent",
				Comment: "func\nComment",
				Annotations: []interface{}{
					&SimpleSpec{
						TypeName: "funcAnnotation",
					},
				},
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Name:    "funcParamName",
							Comment: "funcParam\nComment",
							Annotations: []interface{}{
								&SimpleSpec{
									TypeName: "funcParamAnnotation",
								},
							},
							Spec: funcParamSpec,
						},
					},
					Results: []*Field{
						{
							Name:    "funcResultName",
							Comment: "funcResult\nComment",
							Annotations: []interface{}{
								&SimpleSpec{
									TypeName: "funcResultAnnotation",
								},
							},
							Spec: funcResultSpec,
						},
					},
				},
				Related: &Field{
					Name:    "funcRelatedName",
					Comment: "funcRelated\nComment",
					Annotations: []interface{}{
						&SimpleSpec{
							TypeName: "funcRelatedAnnotation",
						},
					},
					Spec: funcRelatedSpec,
				},
			},
		},
	}

	varSpec.
		EXPECT().
		Clone().
		Return(clonedVarSpec)

	typeSpec.
		EXPECT().
		Clone().
		Return(clonedTypeSpec)

	funcParamSpec.
		EXPECT().
		Clone().
		Return(clonedFuncParamSpec)

	funcResultSpec.
		EXPECT().
		Clone().
		Return(clonedFuncResultSpec)

	funcRelatedSpec.
		EXPECT().
		Clone().
		Return(clonedFuncRelatedSpec)

	actual := modelValue.Clone()

	ctrl.AssertEqual(
		modelValue,
		actual,
		unit.IgnoreUnexportedOption{Value: *ctrl},
		unit.IgnoreUnexportedOption{Value: MockCallManager{}},
	)
	ctrl.AssertNotSame(modelValue, actual)
	ctrl.AssertNotSame(modelValue.Annotations[0], actual.(*File).Annotations[0])

	ctrl.AssertNotSame(modelValue.ImportGroups[0], actual.(*File).ImportGroups[0])
	ctrl.AssertNotSame(modelValue.ImportGroups[0].Annotations[0], actual.(*File).ImportGroups[0].Annotations[0])
	ctrl.AssertNotSame(modelValue.ImportGroups[0].Imports[0], actual.(*File).ImportGroups[0].Imports[0])
	ctrl.AssertNotSame(
		modelValue.ImportGroups[0].Imports[0].Annotations[0],
		actual.(*File).ImportGroups[0].Imports[0].Annotations[0],
	)

	ctrl.AssertNotSame(modelValue.ConstGroups[0], actual.(*File).ConstGroups[0])
	ctrl.AssertNotSame(modelValue.ConstGroups[0].Annotations[0], actual.(*File).ConstGroups[0].Annotations[0])
	ctrl.AssertNotSame(modelValue.ConstGroups[0].Consts[0], actual.(*File).ConstGroups[0].Consts[0])
	ctrl.AssertNotSame(modelValue.ConstGroups[0].Consts[0].Spec, actual.(*File).ConstGroups[0].Consts[0].Spec)
	ctrl.AssertNotSame(
		modelValue.ConstGroups[0].Consts[0].Annotations[0],
		actual.(*File).ConstGroups[0].Consts[0].Annotations[0],
	)

	ctrl.AssertNotSame(modelValue.VarGroups[0], actual.(*File).VarGroups[0])
	ctrl.AssertNotSame(modelValue.VarGroups[0].Annotations[0], actual.(*File).VarGroups[0].Annotations[0])
	ctrl.AssertNotSame(modelValue.VarGroups[0].Vars[0], actual.(*File).VarGroups[0].Vars[0])
	ctrl.AssertSame(clonedVarSpec, actual.(*File).VarGroups[0].Vars[0].Spec)
	ctrl.AssertNotSame(
		modelValue.VarGroups[0].Vars[0].Annotations[0],
		actual.(*File).VarGroups[0].Vars[0].Annotations[0],
	)

	ctrl.AssertNotSame(modelValue.TypeGroups[0], actual.(*File).TypeGroups[0])
	ctrl.AssertNotSame(modelValue.TypeGroups[0].Annotations[0], actual.(*File).TypeGroups[0].Annotations[0])
	ctrl.AssertNotSame(modelValue.TypeGroups[0].Types[0], actual.(*File).TypeGroups[0].Types[0])
	ctrl.AssertSame(clonedTypeSpec, actual.(*File).TypeGroups[0].Types[0].Spec)
	ctrl.AssertNotSame(
		modelValue.TypeGroups[0].Types[0].Annotations[0],
		actual.(*File).TypeGroups[0].Types[0].Annotations[0],
	)

	ctrl.AssertNotSame(modelValue.Funcs[0], actual.(*File).Funcs[0])
	ctrl.AssertNotSame(modelValue.Funcs[0].Annotations[0], actual.(*File).Funcs[0].Annotations[0])
	ctrl.AssertNotSame(modelValue.Funcs[0].Spec, actual.(*File).Funcs[0].Spec)
	ctrl.AssertNotSame(
		modelValue.Funcs[0].Spec.Params[0].Annotations[0],
		actual.(*File).Funcs[0].Spec.Params[0].Annotations[0],
	)
	ctrl.AssertSame(clonedFuncParamSpec, actual.(*File).Funcs[0].Spec.Params[0].Spec)
	ctrl.AssertNotSame(
		modelValue.Funcs[0].Spec.Results[0].Annotations[0],
		actual.(*File).Funcs[0].Spec.Results[0].Annotations[0],
	)
	ctrl.AssertSame(clonedFuncResultSpec, actual.(*File).Funcs[0].Spec.Results[0].Spec)
	ctrl.AssertNotSame(modelValue.Funcs[0].Related.Annotations[0], actual.(*File).Funcs[0].Related.Annotations[0])
	ctrl.AssertSame(clonedFuncRelatedSpec, actual.(*File).Funcs[0].Related.Spec)
}

func TestFile_Clone_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
	}

	actual := modelValue.Clone()

	ctrl.AssertEqual(
		modelValue,
		actual,
		unit.IgnoreUnexportedOption{Value: *ctrl},
		unit.IgnoreUnexportedOption{Value: MockCallManager{}},
	)
	ctrl.AssertNotSame(modelValue, actual)
}

func TestFile_RenameImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	varSpec := NewSpecMock(ctrl)
	typeSpec := NewSpecMock(ctrl)
	funcParamSpec := NewSpecMock(ctrl)
	funcResultSpec := NewSpecMock(ctrl)
	funcRelatedSpec := NewSpecMock(ctrl)

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Content:     "package packageName\nconst Const = (oldPackageName.File + 1) * 5",
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "oldPackageName",
						Namespace: "importNamespace",
					},
				},
			},
		},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:  "constName",
						Value: "(oldPackageName.Const + 1) * 5",
						Spec: &SimpleSpec{
							PackageName: "oldPackageName",
							TypeName:    "constTypeName",
						},
					},
				},
			},
		},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:  "varName",
						Value: "(oldPackageName.Var + 1) * 5",
						Spec:  varSpec,
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{
			{
				Types: []*Type{
					{
						Name: "typeName",
						Spec: typeSpec,
					},
				},
			},
		},
		Funcs: []*Func{
			{
				Name:    "funcName",
				Content: "(oldPackageName.Func + 1) * 5",
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Name: "funcParamName",
							Spec: funcParamSpec,
						},
					},
					Results: []*Field{
						{
							Name: "funcResultName",
							Spec: funcResultSpec,
						},
					},
				},
				Related: &Field{
					Name: "funcRelatedName",
					Spec: funcRelatedSpec,
				},
			},
		},
	}

	modelExpected := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Content:     "package packageName\nconst Const = (newPackageName.File + 1) * 5",
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "newPackageName",
						Namespace: "importNamespace",
					},
				},
			},
		},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:  "constName",
						Value: "(newPackageName.Const + 1) * 5",
						Spec: &SimpleSpec{
							PackageName: "newPackageName",
							TypeName:    "constTypeName",
						},
					},
				},
			},
		},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:  "varName",
						Value: "(newPackageName.Var + 1) * 5",
						Spec:  varSpec,
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{
			{
				Types: []*Type{
					{
						Name: "typeName",
						Spec: typeSpec,
					},
				},
			},
		},
		Funcs: []*Func{
			{
				Name:    "funcName",
				Content: "(newPackageName.Func + 1) * 5",
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Name: "funcParamName",
							Spec: funcParamSpec,
						},
					},
					Results: []*Field{
						{
							Name: "funcResultName",
							Spec: funcResultSpec,
						},
					},
				},
				Related: &Field{
					Name: "funcRelatedName",
					Spec: funcRelatedSpec,
				},
			},
		},
	}

	varSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	typeSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	funcParamSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	funcResultSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	funcRelatedSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	modelValue.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(
		modelValue,
		modelExpected,
		unit.IgnoreUnexportedOption{Value: *ctrl},
		unit.IgnoreUnexportedOption{Value: MockCallManager{}},
	)
}

func TestFile_RenameImports_WithNotRenamedImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	varSpec := NewSpecMock(ctrl)
	typeSpec := NewSpecMock(ctrl)
	funcParamSpec := NewSpecMock(ctrl)
	funcResultSpec := NewSpecMock(ctrl)
	funcRelatedSpec := NewSpecMock(ctrl)

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Content:     "package packageName\nconst Const = (alias.File + 1) * 5",
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "alias",
						Namespace: "importNamespace",
					},
				},
			},
		},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:  "constName",
						Value: "(alias.Const + 1) * 5",
						Spec: &SimpleSpec{
							PackageName: "alias",
							TypeName:    "constTypeName",
						},
					},
				},
			},
		},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:  "varName",
						Value: "(alias.Var + 1) * 5",
						Spec:  varSpec,
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{
			{
				Types: []*Type{
					{
						Name: "typeName",
						Spec: typeSpec,
					},
				},
			},
		},
		Funcs: []*Func{
			{
				Name:    "funcName",
				Content: "(alias.Func + 1) * 5",
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Name: "funcParamName",
							Spec: funcParamSpec,
						},
					},
					Results: []*Field{
						{
							Name: "funcResultName",
							Spec: funcResultSpec,
						},
					},
				},
				Related: &Field{
					Name: "funcRelatedName",
					Spec: funcRelatedSpec,
				},
			},
		},
	}

	modelExpected := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
		Content:     "package packageName\nconst Const = (alias.File + 1) * 5",
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     "alias",
						Namespace: "importNamespace",
					},
				},
			},
		},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:  "constName",
						Value: "(alias.Const + 1) * 5",
						Spec: &SimpleSpec{
							PackageName: "alias",
							TypeName:    "constTypeName",
						},
					},
				},
			},
		},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:  "varName",
						Value: "(alias.Var + 1) * 5",
						Spec:  varSpec,
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{
			{
				Types: []*Type{
					{
						Name: "typeName",
						Spec: typeSpec,
					},
				},
			},
		},
		Funcs: []*Func{
			{
				Name:    "funcName",
				Content: "(alias.Func + 1) * 5",
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Name: "funcParamName",
							Spec: funcParamSpec,
						},
					},
					Results: []*Field{
						{
							Name: "funcResultName",
							Spec: funcResultSpec,
						},
					},
				},
				Related: &Field{
					Name: "funcRelatedName",
					Spec: funcRelatedSpec,
				},
			},
		},
	}

	varSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	typeSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	funcParamSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	funcResultSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	funcRelatedSpec.
		EXPECT().
		RenameImports(oldAlias, newAlias).
		Return()

	modelValue.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(
		modelValue,
		modelExpected,
		unit.IgnoreUnexportedOption{Value: *ctrl},
		unit.IgnoreUnexportedOption{Value: MockCallManager{}},
	)
}

func TestFile_RenameImports_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "newPackageName"

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
	}

	modelExpected := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
	}

	modelValue.RenameImports(oldAlias, newAlias)

	ctrl.AssertEqual(
		modelValue,
		modelExpected,
		unit.IgnoreUnexportedOption{Value: *ctrl},
		unit.IgnoreUnexportedOption{Value: MockCallManager{}},
	)
}

func TestFile_RenameImports_WithInvalidOldAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "+invalid"
	newAlias := "newPackageName"

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'oldAlias' must be valid identifier, actual value: '+invalid'"),
		)
}

func TestFile_RenameImports_WithInvalidNewAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	oldAlias := "oldPackageName"
	newAlias := "+invalid"

	modelValue := &File{
		Name:        "fileName",
		PackageName: "filePackageName",
	}

	ctrl.Subtest("").
		Call(modelValue.RenameImports, oldAlias, newAlias).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'newAlias' must be valid identifier, actual value: '+invalid'"),
		)
}
