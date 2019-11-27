package annotation

import (
	"go/ast"
	"go/scanner"
	"testing"

	"github.com/index0h/go-unit/unit"
)

func TestNewGoSourceParser(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	annotationParser := NewAnnotationParserMock(ctrl)

	expected := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := NewGoSourceParser(annotationParser)

	ctrl.AssertEqual(expected, actual, unit.IgnoreUnexportedOption{Value: AnnotationParserMock{}})
	ctrl.AssertSame(annotationParser, actual.annotationParser)
}

func TestNewSourceParser_WithNilAnnotationParser(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	ctrl.Subtest("").
		Call(NewGoSourceParser, nil).
		ExpectPanic(NewErrorMessageConstraint("Variable 'annotationParser' must be not nil"))
}

func TestSourceParser_Parse_WithFileAnnotations(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	fileComment := "file\ncomment"
	fileAnnotations := []interface{}{
		&TestAnnotation{
			Name: "fileAnnotation",
		},
	}
	fileContent := `//         file
// comment       
package filePackageName

`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Comment:      fileComment,
		Content:      fileContent,
		Annotations:  fileAnnotations,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups:    []*VarGroup{},
		TypeGroups:   []*TypeGroup{},
		Funcs:        []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(fileComment).
		Return(fileAnnotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(fileAnnotations, actual.Annotations)
}

func TestSourceParser_Parse_WithEmptyFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	fileComment := ""
	fileContent := `package filePackageName`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Comment:      fileComment,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups:    []*VarGroup{},
		TypeGroups:   []*TypeGroup{},
		Funcs:        []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithEmptyImportGroupAndImportGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	importGroupComment := "importGroup\ncomment"
	importGroupAnnotations := []interface{}{
		&TestAnnotation{
			Name: "importGroupAnnotation",
		},
	}
	fileContent := `package filePackageName
// importGroup
// comment
import (
)`
	expected := &File{
		Name:        fileName,
		PackageName: filePackageName,
		Content:     fileContent,
		ImportGroups: []*ImportGroup{
			{
				Comment:     importGroupComment,
				Annotations: importGroupAnnotations,
				Imports:     []*Import{},
			},
		},
		ConstGroups: []*ConstGroup{},
		VarGroups:   []*VarGroup{},
		TypeGroups:  []*TypeGroup{},
		Funcs:       []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(importGroupComment).
		Return(importGroupAnnotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(importGroupAnnotations, actual.ImportGroups[0].Annotations)
}

func TestSourceParser_Parse_WithEmptyImportGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	fileContent := `package filePackageName

import (
)`
	expected := &File{
		Name:        fileName,
		PackageName: filePackageName,
		Content:     fileContent,
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{},
			},
		},
		ConstGroups: []*ConstGroup{},
		VarGroups:   []*VarGroup{},
		TypeGroups:  []*TypeGroup{},
		Funcs:       []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithOneImportAndImportAliasAndImportGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	importGroupComment := "importGroup\ncomment"
	importGroupAnnotations := []interface{}{
		&TestAnnotation{
			Name: "importGroupAnnotation",
		},
	}
	importAlias := "importAlias"
	importNamespace := "importNamespace/path"
	fileContent := `package filePackageName
// importGroup
// comment
import importAlias "importNamespace/path"`
	expected := &File{
		Name:        fileName,
		PackageName: filePackageName,
		Content:     fileContent,
		ImportGroups: []*ImportGroup{
			{
				Comment:     importGroupComment,
				Annotations: importGroupAnnotations,
				Imports: []*Import{
					{
						Alias:     importAlias,
						Namespace: importNamespace,
					},
				},
			},
		},
		ConstGroups: []*ConstGroup{},
		VarGroups:   []*VarGroup{},
		TypeGroups:  []*TypeGroup{},
		Funcs:       []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(importGroupComment).
		Return(importGroupAnnotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(importGroupAnnotations, actual.ImportGroups[0].Annotations)
}

func TestSourceParser_Parse_WithOneImportAndImportAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	importAlias := "importAlias"
	importNamespace := "importNamespace/path"
	fileContent := `package filePackageName

import importAlias "importNamespace/path"`
	expected := &File{
		Name:        fileName,
		PackageName: filePackageName,
		Content:     fileContent,
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     importAlias,
						Namespace: importNamespace,
					},
				},
			},
		},
		ConstGroups: []*ConstGroup{},
		VarGroups:   []*VarGroup{},
		TypeGroups:  []*TypeGroup{},
		Funcs:       []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithOneImportAndImportGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	importGroupComment := "importGroup\ncomment"
	importGroupAnnotations := []interface{}{
		&TestAnnotation{
			Name: "importGroupAnnotation",
		},
	}
	importNamespace := "importNamespace/path"
	fileContent := `package filePackageName
// importGroup
// comment
import "importNamespace/path"`
	expected := &File{
		Name:        fileName,
		PackageName: filePackageName,
		Content:     fileContent,
		ImportGroups: []*ImportGroup{
			{
				Comment:     importGroupComment,
				Annotations: importGroupAnnotations,
				Imports: []*Import{
					{
						Namespace: importNamespace,
					},
				},
			},
		},
		ConstGroups: []*ConstGroup{},
		VarGroups:   []*VarGroup{},
		TypeGroups:  []*TypeGroup{},
		Funcs:       []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(importGroupComment).
		Return(importGroupAnnotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(importGroupAnnotations, actual.ImportGroups[0].Annotations)
}

func TestSourceParser_Parse_WithOneImport(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	importNamespace := "importNamespace/path"
	fileContent := `package filePackageName

import "importNamespace/path"`
	expected := &File{
		Name:        fileName,
		PackageName: filePackageName,
		Content:     fileContent,
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: importNamespace,
					},
				},
			},
		},
		ConstGroups: []*ConstGroup{},
		VarGroups:   []*VarGroup{},
		TypeGroups:  []*TypeGroup{},
		Funcs:       []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultipleImportAndImportAliasAndImportGroupCommentAndImportComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	importGroupComment := "importGroup\ncomment"
	importGroupAnnotations := []interface{}{
		&TestAnnotation{
			Name: "importGroupAnnotation",
		},
	}
	import1Alias := "import1Alias"
	import1Namespace := "import1Namespace/path"
	import1Comment := "import1\ncomment"
	import1Annotations := []interface{}{
		&TestAnnotation{
			Name: "import1Annotation",
		},
	}
	import2Alias := "import2Alias"
	import2Namespace := "import2Namespace/path"
	import2Comment := "import2\ncomment"
	import2Annotations := []interface{}{
		&TestAnnotation{
			Name: "import2Annotation",
		},
	}
	content := `package filePackageName
// importGroup
// comment
import (
	// import1
	// comment
	import1Alias "import1Namespace/path"
	// import2
	// comment
	import2Alias "import2Namespace/path"
)
`
	expected := &File{
		Name:        fileName,
		PackageName: filePackageName,
		Content:     content,
		ImportGroups: []*ImportGroup{
			{
				Comment:     importGroupComment,
				Annotations: importGroupAnnotations,
				Imports: []*Import{
					{
						Alias:       import1Alias,
						Namespace:   import1Namespace,
						Comment:     import1Comment,
						Annotations: import1Annotations,
					},
					{
						Alias:       import2Alias,
						Namespace:   import2Namespace,
						Comment:     import2Comment,
						Annotations: import2Annotations,
					},
				},
			},
		},
		ConstGroups: []*ConstGroup{},
		VarGroups:   []*VarGroup{},
		TypeGroups:  []*TypeGroup{},
		Funcs:       []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(importGroupComment).
		Return(importGroupAnnotations)

	annotationParser.
		EXPECT().
		Parse(import1Comment).
		Return(import1Annotations)

	annotationParser.
		EXPECT().
		Parse(import2Comment).
		Return(import2Annotations)

	actual := parser.Parse(fileName, content)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(importGroupAnnotations, actual.ImportGroups[0].Annotations)
	ctrl.AssertSame(import1Annotations, actual.ImportGroups[0].Imports[0].Annotations)
	ctrl.AssertSame(import2Annotations, actual.ImportGroups[0].Imports[1].Annotations)
}

func TestSourceParser_Parse_WithMultipleImportsAndImportAlias(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	import1Alias := "import1Alias"
	import1Namespace := "import1Namespace/path"
	import2Alias := "import2Alias"
	import2Namespace := "import2Namespace/path"
	fileContent := `package filePackageName

import (
	import1Alias "import1Namespace/path"
	import2Alias "import2Namespace/path"
)`
	expected := &File{
		Name:        fileName,
		PackageName: filePackageName,
		Content:     fileContent,
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Alias:     import1Alias,
						Namespace: import1Namespace,
					},
					{
						Alias:     import2Alias,
						Namespace: import2Namespace,
					},
				},
			},
		},
		ConstGroups: []*ConstGroup{},
		VarGroups:   []*VarGroup{},
		TypeGroups:  []*TypeGroup{},
		Funcs:       []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultipleImportsAndImportGroupCommentAndImportComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	importGroupComment := "importGroup\ncomment"
	importGroupAnnotations := []interface{}{
		&TestAnnotation{
			Name: "importGroupAnnotation",
		},
	}
	import1Namespace := "import1Namespace/path1"
	import1Comment := "import1\ncomment"
	import1Annotations := []interface{}{
		&TestAnnotation{
			Name: "import1Annotation",
		},
	}
	import2Namespace := "import2Namespace/path2"
	import2Comment := "import2\ncomment"
	import2Annotations := []interface{}{
		&TestAnnotation{
			Name: "import2Annotation",
		},
	}
	content := `package filePackageName
// importGroup
// comment
import (
	// import1
	// comment
	"import1Namespace/path1"
	// import2
	// comment
	"import2Namespace/path2"
)
`
	expected := &File{
		Name:        fileName,
		PackageName: filePackageName,
		Content:     content,
		ImportGroups: []*ImportGroup{
			{
				Comment:     importGroupComment,
				Annotations: importGroupAnnotations,
				Imports: []*Import{
					{
						Namespace:   import1Namespace,
						Comment:     import1Comment,
						Annotations: import1Annotations,
					},
					{
						Namespace:   import2Namespace,
						Comment:     import2Comment,
						Annotations: import2Annotations,
					},
				},
			},
		},
		ConstGroups: []*ConstGroup{},
		VarGroups:   []*VarGroup{},
		TypeGroups:  []*TypeGroup{},
		Funcs:       []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(importGroupComment).
		Return(importGroupAnnotations)

	annotationParser.
		EXPECT().
		Parse(import1Comment).
		Return(import1Annotations)

	annotationParser.
		EXPECT().
		Parse(import2Comment).
		Return(import2Annotations)

	actual := parser.Parse(fileName, content)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(importGroupAnnotations, actual.ImportGroups[0].Annotations)
	ctrl.AssertSame(import1Annotations, actual.ImportGroups[0].Imports[0].Annotations)
	ctrl.AssertSame(import2Annotations, actual.ImportGroups[0].Imports[1].Annotations)
}

func TestSourceParser_Parse_WithMultipleImports(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	import1Namespace := "import1Namespace/path1"
	import2Namespace := "import2Namespace/path2"
	content := `package filePackageName

import (
	"import1Namespace/path1"
	"import2Namespace/path2"
)
`
	expected := &File{
		Name:        fileName,
		PackageName: filePackageName,
		Content:     content,
		ImportGroups: []*ImportGroup{
			{
				Imports: []*Import{
					{
						Namespace: import1Namespace,
					},
					{
						Namespace: import2Namespace,
					},
				},
			},
		},
		ConstGroups: []*ConstGroup{},
		VarGroups:   []*VarGroup{},
		TypeGroups:  []*TypeGroup{},
		Funcs:       []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, content)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithEmptyConstGroupAndConstGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	constGroupComment := "constGroup\ncomment"
	constGroupAnnotations := []interface{}{
		&TestAnnotation{
			Name: "constGroupAnnotation",
		},
	}
	fileContent := `package filePackageName
// constGroup
// comment
const (
)`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Comment:     constGroupComment,
				Annotations: constGroupAnnotations,
				Consts:      []*Const{},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(constGroupComment).
		Return(constGroupAnnotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(constGroupAnnotations, actual.ConstGroups[0].Annotations)
}

func TestSourceParser_Parse_WithEmptyConstGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	fileContent := `package filePackageName

const (
)`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithOneConstAndConstSpecAndConstValueAndConstCommentAndConstGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	constGroupComment := "constGroup\ncomment"
	constGroupAnnotations := []interface{}{
		&TestAnnotation{
			Name: "constGroupAnnotation",
		},
	}
	constComment := "const\ncomment"
	constAnnotations := []interface{}{
		&TestAnnotation{
			Name: "constAnnotation",
		},
	}
	constName := "constName"
	constPackageName := "constPackageName"
	constTypeName := "constTypeName"
	constValue := "constValue"
	fileContent := `package filePackageName
// constGroup
// comment
const (
	// const
	// comment
	constName constPackageName.constTypeName = constValue
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Comment:     constGroupComment,
				Annotations: constGroupAnnotations,
				Consts: []*Const{
					{
						Name:        constName,
						Value:       constValue,
						Comment:     constComment,
						Annotations: constAnnotations,
						Spec: &SimpleSpec{
							PackageName: constPackageName,
							TypeName:    constTypeName,
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(constGroupComment).
		Return(constGroupAnnotations)

	annotationParser.
		EXPECT().
		Parse(constComment).
		Return(constAnnotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(constGroupAnnotations, actual.ConstGroups[0].Annotations)
	ctrl.AssertSame(constAnnotations, actual.ConstGroups[0].Consts[0].Annotations)
}

func TestSourceParser_Parse_WithOneConstAndConstSpecAndConstValueAndConstComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	constComment := "const\ncomment"
	constAnnotations := []interface{}{
		&TestAnnotation{
			Name: "constAnnotation",
		},
	}
	constName := "constName"
	constPackageName := "constPackageName"
	constTypeName := "constTypeName"
	constValue := "constValue"
	fileContent := `package filePackageName
const (
	// const
	// comment
	constName constPackageName.constTypeName = constValue
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:        constName,
						Value:       constValue,
						Comment:     constComment,
						Annotations: constAnnotations,
						Spec: &SimpleSpec{
							PackageName: constPackageName,
							TypeName:    constTypeName,
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(constComment).
		Return(constAnnotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(constAnnotations, actual.ConstGroups[0].Consts[0].Annotations)
}

func TestSourceParser_Parse_WithOneConstAndConstSpecAndConstValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	constName := "constName"
	constPackageName := "constPackageName"
	constTypeName := "constTypeName"
	constValue := "constValue"
	fileContent := `package filePackageName

const constName constPackageName.constTypeName = constValue
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:  constName,
						Value: constValue,
						Spec: &SimpleSpec{
							PackageName: constPackageName,
							TypeName:    constTypeName,
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithOneConstAndConstSpecByIntValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	constName := "constName"
	constValue := "10"
	fileContent := `package filePackageName

const constName = 10
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:  constName,
						Value: constValue,
						Spec: &SimpleSpec{
							TypeName: "int",
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithOneConstAndConstSpecByFloat64Value(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	constName := "constName"
	constValue := "10.20"
	fileContent := `package filePackageName

const constName = 10.20
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:  constName,
						Value: constValue,
						Spec: &SimpleSpec{
							TypeName: "float64",
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithOneConstAndConstSpecByStringValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	constName := "constName"
	constValue := `"data"`
	fileContent := `package filePackageName

const constName = "data"
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:  constName,
						Value: constValue,
						Spec: &SimpleSpec{
							TypeName: "string",
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultiStmtAndConstSpecAndConstValueAndConstCommentAndConstGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	constGroupComment := "constGroup\ncomment"
	constGroupAnnotations := []interface{}{
		&TestAnnotation{
			Name: "constGroupAnnotation",
		},
	}
	constComment := "const\ncomment"
	const1Annotations := []interface{}{
		&TestAnnotation{
			Name: "const1Annotation",
		},
	}
	const2Annotations := []interface{}{
		&TestAnnotation{
			Name: "const1Annotation",
		},
	}
	constPackageName := "constPackageName"
	constTypeName := "constTypeName"
	const1Name := "const1Name"
	const1Value := "const1Value"
	const2Name := "const2Name"
	const2Value := "const2Value"
	fileContent := `package filePackageName
// constGroup
// comment
const (
	// const
	// comment
	const1Name, const2Name constPackageName.constTypeName = const1Value, const2Value
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Comment:     constGroupComment,
				Annotations: constGroupAnnotations,
				Consts: []*Const{
					{
						Name:        const1Name,
						Value:       const1Value,
						Comment:     constComment,
						Annotations: const1Annotations,
						Spec: &SimpleSpec{
							PackageName: constPackageName,
							TypeName:    constTypeName,
						},
					},
					{
						Name:        const2Name,
						Value:       const2Value,
						Comment:     constComment,
						Annotations: const2Annotations,
						Spec: &SimpleSpec{
							PackageName: constPackageName,
							TypeName:    constTypeName,
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(constGroupComment).
		Return(constGroupAnnotations)

	annotationParser.
		EXPECT().
		Parse(constComment).
		Return(const1Annotations)

	annotationParser.
		EXPECT().
		Parse(constComment).
		Return(const2Annotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(constGroupAnnotations, actual.ConstGroups[0].Annotations)
	ctrl.AssertSame(const1Annotations, actual.ConstGroups[0].Consts[0].Annotations)
	ctrl.AssertSame(const2Annotations, actual.ConstGroups[0].Consts[1].Annotations)
}

func TestSourceParser_Parse_WithMultiStmtAndConstSpecAndConstValueAndConstComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	constComment := "const\ncomment"
	const1Annotations := []interface{}{
		&TestAnnotation{
			Name: "const1Annotation",
		},
	}
	const2Annotations := []interface{}{
		&TestAnnotation{
			Name: "const1Annotation",
		},
	}
	constPackageName := "constPackageName"
	constTypeName := "constTypeName"
	const1Name := "const1Name"
	const1Value := "const1Value"
	const2Name := "const2Name"
	const2Value := "const2Value"
	fileContent := `package filePackageName
const (
	// const
	// comment
	const1Name, const2Name constPackageName.constTypeName = const1Value, const2Value
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:        const1Name,
						Value:       const1Value,
						Comment:     constComment,
						Annotations: const1Annotations,
						Spec: &SimpleSpec{
							PackageName: constPackageName,
							TypeName:    constTypeName,
						},
					},
					{
						Name:        const2Name,
						Value:       const2Value,
						Comment:     constComment,
						Annotations: const2Annotations,
						Spec: &SimpleSpec{
							PackageName: constPackageName,
							TypeName:    constTypeName,
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(constComment).
		Return(const1Annotations)

	annotationParser.
		EXPECT().
		Parse(constComment).
		Return(const2Annotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(const1Annotations, actual.ConstGroups[0].Consts[0].Annotations)
	ctrl.AssertSame(const2Annotations, actual.ConstGroups[0].Consts[1].Annotations)
}

func TestSourceParser_Parse_WithMultiStmtAndConstSpecAndConstValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	constPackageName := "constPackageName"
	constTypeName := "constTypeName"
	const1Name := "const1Name"
	const1Value := "const1Value"
	const2Name := "const2Name"
	const2Value := "const2Value"
	fileContent := `package filePackageName

const const1Name, const2Name constPackageName.constTypeName = const1Value, const2Value
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:  const1Name,
						Value: const1Value,
						Spec: &SimpleSpec{
							PackageName: constPackageName,
							TypeName:    constTypeName,
						},
					},
					{
						Name:  const2Name,
						Value: const2Value,
						Spec: &SimpleSpec{
							PackageName: constPackageName,
							TypeName:    constTypeName,
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultiStmtAndConstSpecByIntValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	const1Name := "const1Name"
	const1Value := "10"
	const2Name := "const2Name"
	const2Value := "20"
	fileContent := `package filePackageName

const const1Name, const2Name = 10, 20
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:  const1Name,
						Value: const1Value,
						Spec: &SimpleSpec{
							TypeName: "int",
						},
					},
					{
						Name:  const2Name,
						Value: const2Value,
						Spec: &SimpleSpec{
							TypeName: "int",
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultiStmtAndConstSpecByFloat64Value(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	const1Name := "const1Name"
	const1Value := "10.20"
	const2Name := "const2Name"
	const2Value := "20.40"
	fileContent := `package filePackageName

const const1Name, const2Name = 10.20, 20.40
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:  const1Name,
						Value: const1Value,
						Spec: &SimpleSpec{
							TypeName: "float64",
						},
					},
					{
						Name:  const2Name,
						Value: const2Value,
						Spec: &SimpleSpec{
							TypeName: "float64",
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultiStmtAndConstSpecByStringValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	const1Name := "const1Name"
	const1Value := `"data1"`
	const2Name := "const2Name"
	const2Value := `"data2"`
	fileContent := `package filePackageName

const const1Name, const2Name = "data1", "data2"
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:  const1Name,
						Value: const1Value,
						Spec: &SimpleSpec{
							TypeName: "string",
						},
					},
					{
						Name:  const2Name,
						Value: const2Value,
						Spec: &SimpleSpec{
							TypeName: "string",
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultiConstsAndConstSpecAndConstValueAndConstCommentAndConstGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	constGroupComment := "constGroup\ncomment"
	constGroupAnnotations := []interface{}{
		&TestAnnotation{
			Name: "constGroupAnnotation",
		},
	}
	const1Comment := "const\ncomment"
	const1Annotations := []interface{}{
		&TestAnnotation{
			Name: "const1Annotation",
		},
	}
	constPackageName := "constPackageName"
	constTypeName := "constTypeName"
	const1Name := "const1Name"
	const1Value := "const1Value"
	const2Name := "const2Name"
	fileContent := `package filePackageName
// constGroup
// comment
const (
	// const
	// comment
	const1Name constPackageName.constTypeName = const1Value
	const2Name
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Comment:     constGroupComment,
				Annotations: constGroupAnnotations,
				Consts: []*Const{
					{
						Name:        const1Name,
						Value:       const1Value,
						Comment:     const1Comment,
						Annotations: const1Annotations,
						Spec: &SimpleSpec{
							PackageName: constPackageName,
							TypeName:    constTypeName,
						},
					},
					{
						Name: const2Name,
						Spec: &SimpleSpec{
							PackageName: constPackageName,
							TypeName:    constTypeName,
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(constGroupComment).
		Return(constGroupAnnotations)

	annotationParser.
		EXPECT().
		Parse(const1Comment).
		Return(const1Annotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(constGroupAnnotations, actual.ConstGroups[0].Annotations)
	ctrl.AssertSame(const1Annotations, actual.ConstGroups[0].Consts[0].Annotations)
}

func TestSourceParser_Parse_WithMultiConstsAndConstSpecAndConstValueAndConstComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	const1Comment := "const1\ncomment"
	const1Annotations := []interface{}{
		&TestAnnotation{
			Name: "const1Annotation",
		},
	}
	constPackageName := "constPackageName"
	constTypeName := "constTypeName"
	const1Name := "const1Name"
	const1Value := "const1Value"
	const2Name := "const2Name"
	fileContent := `package filePackageName
const (
	// const1
	// comment
	const1Name constPackageName.constTypeName = const1Value
	const2Name
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:        const1Name,
						Value:       const1Value,
						Comment:     const1Comment,
						Annotations: const1Annotations,
						Spec: &SimpleSpec{
							PackageName: constPackageName,
							TypeName:    constTypeName,
						},
					},
					{
						Name: const2Name,
						Spec: &SimpleSpec{
							PackageName: constPackageName,
							TypeName:    constTypeName,
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(const1Comment).
		Return(const1Annotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(const1Annotations, actual.ConstGroups[0].Consts[0].Annotations)
}

func TestSourceParser_Parse_WithMultiConstsAndConstSpecAndConstValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	constPackageName := "constPackageName"
	constTypeName := "constTypeName"
	const1Name := "const1Name"
	const1Value := "const1Value"
	const2Name := "const2Name"
	fileContent := `package filePackageName

const (
	const1Name constPackageName.constTypeName = const1Value
	const2Name
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:  const1Name,
						Value: const1Value,
						Spec: &SimpleSpec{
							PackageName: constPackageName,
							TypeName:    constTypeName,
						},
					},
					{
						Name: const2Name,
						Spec: &SimpleSpec{
							PackageName: constPackageName,
							TypeName:    constTypeName,
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultiConstsAndConstSpecByIntValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	const1Name := "const1Name"
	const1Value := "10"
	const2Name := "const2Name"
	fileContent := `package filePackageName

const (
	const1Name = 10
	const2Name
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:  const1Name,
						Value: const1Value,
						Spec: &SimpleSpec{
							TypeName: "int",
						},
					},
					{
						Name: const2Name,
						Spec: &SimpleSpec{
							TypeName: "int",
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultiConstsAndConstSpecByFloat64Value(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	const1Name := "const1Name"
	const1Value := "10.20"
	const2Name := "const2Name"
	fileContent := `package filePackageName

const (
	const1Name = 10.20
	const2Name
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:  const1Name,
						Value: const1Value,
						Spec: &SimpleSpec{
							TypeName: "float64",
						},
					},
					{
						Name: const2Name,
						Spec: &SimpleSpec{
							TypeName: "float64",
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultiConstsAndConstSpecByStringValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	const1Name := "const1Name"
	const1Value := `"data1"`
	const2Name := "const2Name"
	fileContent := `package filePackageName

const (
	const1Name = "data1"
	const2Name
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups: []*ConstGroup{
			{
				Consts: []*Const{
					{
						Name:  const1Name,
						Value: const1Value,
						Spec: &SimpleSpec{
							TypeName: "string",
						},
					},
					{
						Name: const2Name,
						Spec: &SimpleSpec{
							TypeName: "string",
						},
					},
				},
			},
		},
		VarGroups:  []*VarGroup{},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithEmptyVarGroupAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	varGroupComment := "varGroup\ncomment"
	varGroupAnnotations := []interface{}{
		&TestAnnotation{
			Name: "varGroupAnnotation",
		},
	}
	fileContent := `package filePackageName
// varGroup
// comment
var (
)`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Comment:     varGroupComment,
				Annotations: varGroupAnnotations,
				Vars:        []*Var{},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(varGroupComment).
		Return(varGroupAnnotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(varGroupAnnotations, actual.VarGroups[0].Annotations)
}

func TestSourceParser_Parse_WithEmptyVarGroup(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	fileContent := `package filePackageName

var (
)`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithOneVarAndVatSpecAndVarValueAndVarCommentAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	varGroupComment := "varGroup\ncomment"
	varGroupAnnotations := []interface{}{
		&TestAnnotation{
			Name: "varGroupAnnotation",
		},
	}
	varComment := "var\ncomment"
	varAnnotations := []interface{}{
		&TestAnnotation{
			Name: "varAnnotation",
		},
	}
	varName := "varName"
	varPackageName := "varPackageName"
	varTypeName := "varTypeName"
	varValue := "varValue"
	fileContent := `package filePackageName
// varGroup
// comment
var (
	// var
	// comment
	varName varPackageName.varTypeName = varValue
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Comment:     varGroupComment,
				Annotations: varGroupAnnotations,
				Vars: []*Var{
					{
						Name:        varName,
						Value:       varValue,
						Comment:     varComment,
						Annotations: varAnnotations,
						Spec: &SimpleSpec{
							PackageName: varPackageName,
							TypeName:    varTypeName,
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(varGroupComment).
		Return(varGroupAnnotations)

	annotationParser.
		EXPECT().
		Parse(varComment).
		Return(varAnnotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(varGroupAnnotations, actual.VarGroups[0].Annotations)
	ctrl.AssertSame(varAnnotations, actual.VarGroups[0].Vars[0].Annotations)
}

func TestSourceParser_Parse_WithOneVarAndVarSpecAndVarValueAndVarComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	varComment := "var\ncomment"
	varAnnotations := []interface{}{
		&TestAnnotation{
			Name: "varAnnotation",
		},
	}
	varName := "varName"
	varPackageName := "varPackageName"
	varTypeName := "varTypeName"
	varValue := "varValue"
	fileContent := `package filePackageName
var (
	// var
	// comment
	varName varPackageName.varTypeName = varValue
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:        varName,
						Value:       varValue,
						Comment:     varComment,
						Annotations: varAnnotations,
						Spec: &SimpleSpec{
							PackageName: varPackageName,
							TypeName:    varTypeName,
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(varComment).
		Return(varAnnotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(varAnnotations, actual.VarGroups[0].Vars[0].Annotations)
}

func TestSourceParser_Parse_WithOneVarAndVarSpecTypeAndVarValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	varName := "varName"
	varPackageName := "varPackageName"
	varTypeName := "varTypeName"
	varValue := "varValue"
	fileContent := `package filePackageName

var varName varPackageName.varTypeName = varValue
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:  varName,
						Value: varValue,
						Spec: &SimpleSpec{
							PackageName: varPackageName,
							TypeName:    varTypeName,
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithOneVarAndVarSpecByIntValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	varName := "varName"
	varValue := "10"
	fileContent := `package filePackageName

var varName = 10
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:  varName,
						Value: varValue,
						Spec: &SimpleSpec{
							TypeName: "int",
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithOneVarAndVarSpecByFloat64Value(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	varName := "varName"
	varValue := "10.20"
	fileContent := `package filePackageName

var varName = 10.20
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:  varName,
						Value: varValue,
						Spec: &SimpleSpec{
							TypeName: "float64",
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithOneVarAndVarSpecByStringValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	varName := "varName"
	varValue := `"data"`
	fileContent := `package filePackageName

var varName = "data"
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:  varName,
						Value: varValue,
						Spec: &SimpleSpec{
							TypeName: "string",
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultiStmtAndVarSpecAndVarValueAndVarCommentAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	varGroupComment := "varGroup\ncomment"
	varGroupAnnotations := []interface{}{
		&TestAnnotation{
			Name: "varGroupAnnotation",
		},
	}
	varComment := "var\ncomment"
	var1Annotations := []interface{}{
		&TestAnnotation{
			Name: "var1Annotation",
		},
	}
	var2Annotations := []interface{}{
		&TestAnnotation{
			Name: "var1Annotation",
		},
	}
	varPackageName := "varPackageName"
	varTypeName := "varTypeName"
	var1Name := "var1Name"
	var1Value := "var1Value"
	var2Name := "var2Name"
	var2Value := "var2Value"
	fileContent := `package filePackageName
// varGroup
// comment
var (
	// var
	// comment
	var1Name, var2Name varPackageName.varTypeName = var1Value, var2Value
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Comment:     varGroupComment,
				Annotations: varGroupAnnotations,
				Vars: []*Var{
					{
						Name:        var1Name,
						Value:       var1Value,
						Comment:     varComment,
						Annotations: var1Annotations,
						Spec: &SimpleSpec{
							PackageName: varPackageName,
							TypeName:    varTypeName,
						},
					},
					{
						Name:        var2Name,
						Value:       var2Value,
						Comment:     varComment,
						Annotations: var2Annotations,
						Spec: &SimpleSpec{
							PackageName: varPackageName,
							TypeName:    varTypeName,
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(varGroupComment).
		Return(varGroupAnnotations)

	annotationParser.
		EXPECT().
		Parse(varComment).
		Return(var1Annotations)

	annotationParser.
		EXPECT().
		Parse(varComment).
		Return(var2Annotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(varGroupAnnotations, actual.VarGroups[0].Annotations)
	ctrl.AssertSame(var1Annotations, actual.VarGroups[0].Vars[0].Annotations)
	ctrl.AssertSame(var2Annotations, actual.VarGroups[0].Vars[1].Annotations)
}

func TestSourceParser_Parse_WithMultiStmtAndVarSpecAndVarValueAndVarComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	varComment := "var\ncomment"
	var1Annotations := []interface{}{
		&TestAnnotation{
			Name: "var1Annotation",
		},
	}
	var2Annotations := []interface{}{
		&TestAnnotation{
			Name: "var1Annotation",
		},
	}
	varPackageName := "varPackageName"
	varTypeName := "varTypeName"
	var1Name := "var1Name"
	var1Value := "var1Value"
	var2Name := "var2Name"
	var2Value := "var2Value"
	fileContent := `package filePackageName
var (
	// var
	// comment
	var1Name, var2Name varPackageName.varTypeName = var1Value, var2Value
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:        var1Name,
						Value:       var1Value,
						Comment:     varComment,
						Annotations: var1Annotations,
						Spec: &SimpleSpec{
							PackageName: varPackageName,
							TypeName:    varTypeName,
						},
					},
					{
						Name:        var2Name,
						Value:       var2Value,
						Comment:     varComment,
						Annotations: var2Annotations,
						Spec: &SimpleSpec{
							PackageName: varPackageName,
							TypeName:    varTypeName,
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(varComment).
		Return(var1Annotations)

	annotationParser.
		EXPECT().
		Parse(varComment).
		Return(var2Annotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(var1Annotations, actual.VarGroups[0].Vars[0].Annotations)
	ctrl.AssertSame(var2Annotations, actual.VarGroups[0].Vars[1].Annotations)
}

func TestSourceParser_Parse_WithMultiStmtAndVarSpecAndVarValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	varPackageName := "varPackageName"
	varTypeName := "varTypeName"
	var1Name := "var1Name"
	var1Value := "var1Value"
	var2Name := "var2Name"
	var2Value := "var2Value"
	fileContent := `package filePackageName

var var1Name, var2Name varPackageName.varTypeName = var1Value, var2Value
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:  var1Name,
						Value: var1Value,
						Spec: &SimpleSpec{
							PackageName: varPackageName,
							TypeName:    varTypeName,
						},
					},
					{
						Name:  var2Name,
						Value: var2Value,
						Spec: &SimpleSpec{
							PackageName: varPackageName,
							TypeName:    varTypeName,
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultiStmtAndVarSpecByIntValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	var1Name := "var1Name"
	var1Value := "10"
	var2Name := "var2Name"
	var2Value := "20"
	fileContent := `package filePackageName

var var1Name, var2Name = 10, 20
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:  var1Name,
						Value: var1Value,
						Spec: &SimpleSpec{
							TypeName: "int",
						},
					},
					{
						Name:  var2Name,
						Value: var2Value,
						Spec: &SimpleSpec{
							TypeName: "int",
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultiStmtAndVarSpecByFloat64Value(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	var1Name := "var1Name"
	var1Value := "10.20"
	var2Name := "var2Name"
	var2Value := "20.40"
	fileContent := `package filePackageName

var var1Name, var2Name = 10.20, 20.40
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:  var1Name,
						Value: var1Value,
						Spec: &SimpleSpec{
							TypeName: "float64",
						},
					},
					{
						Name:  var2Name,
						Value: var2Value,
						Spec: &SimpleSpec{
							TypeName: "float64",
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultiStmtAndVarSpecByStringValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	var1Name := "var1Name"
	var1Value := `"data1"`
	var2Name := "var2Name"
	var2Value := `"data2"`
	fileContent := `package filePackageName

var var1Name, var2Name = "data1", "data2"
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:  var1Name,
						Value: var1Value,
						Spec: &SimpleSpec{
							TypeName: "string",
						},
					},
					{
						Name:  var2Name,
						Value: var2Value,
						Spec: &SimpleSpec{
							TypeName: "string",
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultiVarsAndVarSpecAndVarValueAndVarCommentAndVarGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	varGroupComment := "varGroup\ncomment"
	varGroupAnnotations := []interface{}{
		&TestAnnotation{
			Name: "varGroupAnnotation",
		},
	}
	var1Comment := "var\ncomment"
	var1Annotations := []interface{}{
		&TestAnnotation{
			Name: "var1Annotation",
		},
	}
	varPackageName := "varPackageName"
	varTypeName := "varTypeName"
	var1Name := "var1Name"
	var1Value := "var1Value"
	var2Name := "var2Name"
	fileContent := `package filePackageName
// varGroup
// comment
var (
	// var
	// comment
	var1Name varPackageName.varTypeName = var1Value
	var2Name varPackageName.varTypeName
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Comment:     varGroupComment,
				Annotations: varGroupAnnotations,
				Vars: []*Var{
					{
						Name:        var1Name,
						Value:       var1Value,
						Comment:     var1Comment,
						Annotations: var1Annotations,
						Spec: &SimpleSpec{
							PackageName: varPackageName,
							TypeName:    varTypeName,
						},
					},
					{
						Name: var2Name,
						Spec: &SimpleSpec{
							PackageName: varPackageName,
							TypeName:    varTypeName,
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(varGroupComment).
		Return(varGroupAnnotations)

	annotationParser.
		EXPECT().
		Parse(var1Comment).
		Return(var1Annotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(varGroupAnnotations, actual.VarGroups[0].Annotations)
	ctrl.AssertSame(var1Annotations, actual.VarGroups[0].Vars[0].Annotations)
}

func TestSourceParser_Parse_WithMultiVarsAndVarSpecAndVarValueAndVarComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	var1Comment := "var1\ncomment"
	var1Annotations := []interface{}{
		&TestAnnotation{
			Name: "var1Annotation",
		},
	}
	varPackageName := "varPackageName"
	varTypeName := "varTypeName"
	var1Name := "var1Name"
	var1Value := "var1Value"
	var2Name := "var2Name"
	fileContent := `package filePackageName
var (
	// var1
	// comment
	var1Name varPackageName.varTypeName = var1Value
	var2Name varPackageName.varTypeName
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:        var1Name,
						Value:       var1Value,
						Comment:     var1Comment,
						Annotations: var1Annotations,
						Spec: &SimpleSpec{
							PackageName: varPackageName,
							TypeName:    varTypeName,
						},
					},
					{
						Name: var2Name,
						Spec: &SimpleSpec{
							PackageName: varPackageName,
							TypeName:    varTypeName,
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(var1Comment).
		Return(var1Annotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(var1Annotations, actual.VarGroups[0].Vars[0].Annotations)
}

func TestSourceParser_Parse_WithMultiVarsAndVarSpecAndVarValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	varPackageName := "varPackageName"
	varTypeName := "varTypeName"
	var1Name := "var1Name"
	var1Value := "var1Value"
	var2Name := "var2Name"
	fileContent := `package filePackageName

var (
	var1Name varPackageName.varTypeName = var1Value
	var2Name varPackageName.varTypeName
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:  var1Name,
						Value: var1Value,
						Spec: &SimpleSpec{
							PackageName: varPackageName,
							TypeName:    varTypeName,
						},
					},
					{
						Name: var2Name,
						Spec: &SimpleSpec{
							PackageName: varPackageName,
							TypeName:    varTypeName,
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultiVarsAndVarSpecByIntValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	var1Name := "var1Name"
	var1Value := "10"
	var2Name := "var2Name"
	var2Value := "20"
	fileContent := `package filePackageName

var (
	var1Name = 10
	var2Name = 20
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:  var1Name,
						Value: var1Value,
						Spec: &SimpleSpec{
							TypeName: "int",
						},
					},
					{
						Name:  var2Name,
						Value: var2Value,
						Spec: &SimpleSpec{
							TypeName: "int",
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultiVarsAndVarSpecByFloat64Value(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	var1Name := "var1Name"
	var1Value := "10.20"
	var2Name := "var2Name"
	var2Value := "30.40"
	fileContent := `package filePackageName

var (
	var1Name = 10.20
	var2Name = 30.40
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:  var1Name,
						Value: var1Value,
						Spec: &SimpleSpec{
							TypeName: "float64",
						},
					},
					{
						Name:  var2Name,
						Value: var2Value,
						Spec: &SimpleSpec{
							TypeName: "float64",
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultiVarsAndVarSpecByStringValue(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	var1Name := "var1Name"
	var1Value := `"data1"`
	var2Name := "var2Name"
	var2Value := `"data2"`
	fileContent := `package filePackageName

var (
	var1Name = "data1"
	var2Name = "data2"
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups: []*VarGroup{
			{
				Vars: []*Var{
					{
						Name:  var1Name,
						Value: var1Value,
						Spec: &SimpleSpec{
							TypeName: "string",
						},
					},
					{
						Name:  var2Name,
						Value: var2Value,
						Spec: &SimpleSpec{
							TypeName: "string",
						},
					},
				},
			},
		},
		TypeGroups: []*TypeGroup{},
		Funcs:      []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithOneTypeAndTypeCommentAndTypeGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	typeGroupComment := "typeGroup\ncomment"
	typeGroupAnnotations := []interface{}{
		&TestAnnotation{
			Name: "typeGroupAnnotation",
		},
	}
	typeComment := "type\ncomment"
	typeAnnotations := []interface{}{
		&TestAnnotation{
			Name: "typeAnnotation",
		},
	}
	typeName := "typeName"
	typePackageName := "typePackageName"
	typeTypeName := "typeTypeName"
	fileContent := `package filePackageName
// typeGroup
// comment
type (
	// type
	// comment
	typeName typePackageName.typeTypeName
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups:    []*VarGroup{},
		TypeGroups: []*TypeGroup{
			{
				Comment:     typeGroupComment,
				Annotations: typeGroupAnnotations,
				Types: []*Type{
					{
						Name:        typeName,
						Comment:     typeComment,
						Annotations: typeAnnotations,
						Spec: &SimpleSpec{
							PackageName: typePackageName,
							TypeName:    typeTypeName,
						},
					},
				},
			},
		},
		Funcs: []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(typeGroupComment).
		Return(typeGroupAnnotations)

	annotationParser.
		EXPECT().
		Parse(typeComment).
		Return(typeAnnotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(typeGroupAnnotations, actual.TypeGroups[0].Annotations)
	ctrl.AssertSame(typeAnnotations, actual.TypeGroups[0].Types[0].Annotations)
}

func TestSourceParser_Parse_WithOneTypeAndTypeComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	typeComment := "type\ncomment"
	typeAnnotations := []interface{}{
		&TestAnnotation{
			Name: "typeAnnotation",
		},
	}
	typeName := "typeName"
	typePackageName := "typePackageName"
	typeTypeName := "typeTypeName"
	fileContent := `package filePackageName
type (
	// type
	// comment
	typeName typePackageName.typeTypeName
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups:    []*VarGroup{},
		TypeGroups: []*TypeGroup{
			{
				Types: []*Type{
					{
						Name:        typeName,
						Comment:     typeComment,
						Annotations: typeAnnotations,
						Spec: &SimpleSpec{
							PackageName: typePackageName,
							TypeName:    typeTypeName,
						},
					},
				},
			},
		},
		Funcs: []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(typeComment).
		Return(typeAnnotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(typeAnnotations, actual.TypeGroups[0].Types[0].Annotations)
}

func TestSourceParser_Parse_WithOneType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	typeName := "typeName"
	typePackageName := "typePackageName"
	typeTypeName := "typeTypeName"
	fileContent := `package filePackageName

type typeName typePackageName.typeTypeName
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups:    []*VarGroup{},
		TypeGroups: []*TypeGroup{
			{
				Types: []*Type{
					{
						Name: typeName,
						Spec: &SimpleSpec{
							PackageName: typePackageName,
							TypeName:    typeTypeName,
						},
					},
				},
			},
		},
		Funcs: []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithMultiTypesAndTypeCommentAndTypeGroupComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	typeGroupComment := "typeGroup\ncomment"
	typeGroupAnnotations := []interface{}{
		&TestAnnotation{
			Name: "typeGroupAnnotation",
		},
	}
	type1Comment := "type\ncomment"
	type1Annotations := []interface{}{
		&TestAnnotation{
			Name: "type1Annotation",
		},
	}
	typePackageName := "typePackageName"
	typeTypeName := "typeTypeName"
	type1Name := "type1Name"
	type2Name := "type2Name"
	fileContent := `package filePackageName
// typeGroup
// comment
type (
	// type
	// comment
	type1Name typePackageName.typeTypeName
	type2Name typePackageName.typeTypeName
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups:    []*VarGroup{},
		TypeGroups: []*TypeGroup{
			{
				Comment:     typeGroupComment,
				Annotations: typeGroupAnnotations,
				Types: []*Type{
					{
						Name:        type1Name,
						Comment:     type1Comment,
						Annotations: type1Annotations,
						Spec: &SimpleSpec{
							PackageName: typePackageName,
							TypeName:    typeTypeName,
						},
					},
					{
						Name: type2Name,
						Spec: &SimpleSpec{
							PackageName: typePackageName,
							TypeName:    typeTypeName,
						},
					},
				},
			},
		},
		Funcs: []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(typeGroupComment).
		Return(typeGroupAnnotations)

	annotationParser.
		EXPECT().
		Parse(type1Comment).
		Return(type1Annotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(typeGroupAnnotations, actual.TypeGroups[0].Annotations)
	ctrl.AssertSame(type1Annotations, actual.TypeGroups[0].Types[0].Annotations)
}

func TestSourceParser_Parse_WithMultiTypesAndTypeComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	type1Comment := "type1\ncomment"
	type1Annotations := []interface{}{
		&TestAnnotation{
			Name: "type1Annotation",
		},
	}
	typePackageName := "typePackageName"
	typeTypeName := "typeTypeName"
	type1Name := "type1Name"
	type2Name := "type2Name"
	fileContent := `package filePackageName
type (
	// type1
	// comment
	type1Name typePackageName.typeTypeName
	type2Name typePackageName.typeTypeName
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups:    []*VarGroup{},
		TypeGroups: []*TypeGroup{
			{
				Types: []*Type{
					{
						Name:        type1Name,
						Comment:     type1Comment,
						Annotations: type1Annotations,
						Spec: &SimpleSpec{
							PackageName: typePackageName,
							TypeName:    typeTypeName,
						},
					},
					{
						Name: type2Name,
						Spec: &SimpleSpec{
							PackageName: typePackageName,
							TypeName:    typeTypeName,
						},
					},
				},
			},
		},
		Funcs: []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(type1Comment).
		Return(type1Annotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(type1Annotations, actual.TypeGroups[0].Types[0].Annotations)
}

func TestSourceParser_Parse_WithMultiTypes(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	typePackageName := "typePackageName"
	typeTypeName := "typeTypeName"
	type1Name := "type1Name"
	type2Name := "type2Name"
	fileContent := `package filePackageName

type (
	type1Name typePackageName.typeTypeName
	type2Name typePackageName.typeTypeName
)
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups:    []*VarGroup{},
		TypeGroups: []*TypeGroup{
			{
				Types: []*Type{
					{
						Name: type1Name,
						Spec: &SimpleSpec{
							PackageName: typePackageName,
							TypeName:    typeTypeName,
						},
					},
					{
						Name: type2Name,
						Spec: &SimpleSpec{
							PackageName: typePackageName,
							TypeName:    typeTypeName,
						},
					},
				},
			},
		},
		Funcs: []*Func{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithFuncAndFuncRelatedAndFuncCommentAndFuncRelatedComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	funcComment := "func\ncomment"
	funcAnnotations := []interface{}{
		&TestAnnotation{
			Name: "funcAnnotation",
		},
	}
	funcName := "funcName"
	funcContent := "funcContent"
	funcRelatedName := "funcRelatedName"
	funcRelatedType := "funcRelatedType"
	funcRelatedComment := "funcRelated\ncomment"
	funcRelatedAnnotations := []interface{}{
		&TestAnnotation{
			Name: "funcRelatedAnnotation",
		},
	}
	funcArgumentName := "funcArgumentName"
	funcArgumentType := "funcArgumentType"
	fileContent := `package filePackageName
// func
// comment
func (
	// funcRelated
	// comment
	funcRelatedName *funcRelatedType,
) funcName(funcArgumentName funcArgumentType) {
	funcContent
}
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups:    []*VarGroup{},
		TypeGroups:   []*TypeGroup{},
		Funcs: []*Func{
			{
				Name:        funcName,
				Content:     funcContent,
				Comment:     funcComment,
				Annotations: funcAnnotations,
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Name: funcArgumentName,
							Spec: &SimpleSpec{
								TypeName: funcArgumentType,
							},
						},
					},
					Results: []*Field{},
				},
				Related: &Field{
					Name:        funcRelatedName,
					Comment:     funcRelatedComment,
					Annotations: funcRelatedAnnotations,
					Spec: &SimpleSpec{
						TypeName:  funcRelatedType,
						IsPointer: true,
					},
				},
			},
		},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(funcComment).
		Return(funcAnnotations)

	annotationParser.
		EXPECT().
		Parse(funcRelatedComment).
		Return(funcRelatedAnnotations)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
	ctrl.AssertSame(funcAnnotations, actual.Funcs[0].Annotations)
	ctrl.AssertSame(funcRelatedAnnotations, actual.Funcs[0].Related.Annotations)
}

func TestSourceParser_Parse_WithFuncAndFuncRelated(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	funcName := "funcName"
	funcContent := "funcContent"
	funcRelatedName := "funcRelatedName"
	funcRelatedType := "funcRelatedType"
	funcArgumentName := "funcArgumentName"
	funcArgumentType := "funcArgumentType"
	fileContent := `package filePackageName

func (funcRelatedName *funcRelatedType) funcName(funcArgumentName funcArgumentType) {
	funcContent
}
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups:    []*VarGroup{},
		TypeGroups:   []*TypeGroup{},
		Funcs: []*Func{
			{
				Name:    funcName,
				Content: funcContent,
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Name: funcArgumentName,
							Spec: &SimpleSpec{
								TypeName: funcArgumentType,
							},
						},
					},
					Results: []*Field{},
				},
				Related: &Field{
					Name: funcRelatedName,
					Spec: &SimpleSpec{
						TypeName:  funcRelatedType,
						IsPointer: true,
					},
				},
			},
		},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithFunc(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	filePackageName := "filePackageName"
	funcName := "funcName"
	funcContent := "funcContent"
	funcArgumentName := "funcArgumentName"
	funcArgumentType := "funcArgumentType"
	fileContent := `package filePackageName

func funcName(funcArgumentName funcArgumentType) {funcContent}
`
	expected := &File{
		Name:         fileName,
		PackageName:  filePackageName,
		Content:      fileContent,
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups:    []*VarGroup{},
		TypeGroups:   []*TypeGroup{},
		Funcs: []*Func{
			{
				Name:    funcName,
				Content: funcContent,
				Spec: &FuncSpec{
					Params: []*Field{
						{
							Name: funcArgumentName,
							Spec: &SimpleSpec{
								TypeName: funcArgumentType,
							},
						},
					},
					Results: []*Field{},
				},
			},
		},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertEqual(expected, actual)
}

func TestSourceParser_Parse_WithInvalidFileContent(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `[invalid`

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	ctrl.Subtest("").
		Call(parser.Parse, fileName, fileContent).
		ExpectPanic(ctrl.Type(scanner.ErrorList{}))
}

func TestSourceParser_Parse_WithSimpleSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `package filePackageName

type typeName valueSpec
`
	expected := &SimpleSpec{
		TypeName: "valueSpec",
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
}

func TestSourceParser_Parse_WithSimpleSpecAndPackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `package filePackageName

type typeName valueSpecPackageName.valueSpec
`
	expected := &SimpleSpec{
		PackageName: "valueSpecPackageName",
		TypeName:    "valueSpec",
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
}

func TestSourceParser_Parse_WithSimpleSpecAndIsPointer(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `package filePackageName

type typeName *valueSpec
`
	expected := &SimpleSpec{
		TypeName:  "valueSpec",
		IsPointer: true,
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
}

func TestSourceParser_Parse_WithSimpleSpecAndIsPointerAndPackageName(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `package filePackageName

type typeName *valueSpecPackageName.valueSpec
`
	expected := &SimpleSpec{
		PackageName: "valueSpecPackageName",
		TypeName:    "valueSpec",
		IsPointer:   true,
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
}

func TestSourceParser_Parse_WithArraySpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `package filePackageName

type typeName []valueSpec
`
	expected := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "valueSpec",
		},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
}

func TestSourceParser_Parse_WithArraySpecAndLength(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `package filePackageName

type typeName [5]valueSpec
`
	expected := &ArraySpec{
		Value: &SimpleSpec{
			TypeName: "valueSpec",
		},
		Length: "5",
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
}

func TestSourceParser_Parse_WithMapSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `package filePackageName

type typeName map[keySpec]valueSpec
`
	expected := &MapSpec{
		Key: &SimpleSpec{
			TypeName: "keySpec",
		},
		Value: &SimpleSpec{
			TypeName: "valueSpec",
		},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
}

func TestSourceParser_Parse_WithStructSpecAndWithoutFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `package filePackageName

type typeName struct {}
`
	expected := &StructSpec{
		Fields: []*Field{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
}

func TestSourceParser_Parse_WithStructSpecAndNameAndTagAndFieldComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fieldComment := "field\ncomment"
	fieldAnnotation := []interface{}{
		TestAnnotation{
			Name: "fieldAnnotation",
		},
	}
	fileContent := `package filePackageName

type typeName struct {
	// field
	// comment
	fieldName fieldType "fieldTag"
}
`
	expected := &StructSpec{
		Fields: []*Field{
			{
				Name:        "fieldName",
				Tag:         "fieldTag",
				Comment:     "field\ncomment",
				Annotations: fieldAnnotation,
				Spec: &SimpleSpec{
					TypeName: "fieldType",
				},
			},
		},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(fieldComment).
		Return(fieldAnnotation)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertSame(
		expected.Fields[0].Annotations,
		actual.TypeGroups[0].Types[0].Spec.(*StructSpec).Fields[0].Annotations,
	)
}

func TestSourceParser_Parse_WithStructSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `package filePackageName

type typeName struct {
	fieldType
}
`
	expected := &StructSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "fieldType",
				},
			},
		},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertSame(
		expected.Fields[0].Annotations,
		actual.TypeGroups[0].Types[0].Spec.(*StructSpec).Fields[0].Annotations,
	)
}

func TestSourceParser_Parse_WithStructSpecAndMultiStmtAndNameAndTagAndFieldComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fieldComment := "field\ncomment"
	fieldAnnotation := []interface{}{
		TestAnnotation{
			Name: "fieldAnnotation",
		},
	}
	fileContent := `package filePackageName

type typeName struct {
	// field
	// comment
	field1Name, field2Name fieldType "fieldTag"
}
`
	expected := &StructSpec{
		Fields: []*Field{
			{
				Name:        "field1Name",
				Tag:         "fieldTag",
				Comment:     "field\ncomment",
				Annotations: fieldAnnotation,
				Spec: &SimpleSpec{
					TypeName: "fieldType",
				},
			},
			{
				Name:        "field2Name",
				Tag:         "fieldTag",
				Comment:     "field\ncomment",
				Annotations: fieldAnnotation,
				Spec: &SimpleSpec{
					TypeName: "fieldType",
				},
			},
		},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(fieldComment).
		Return(fieldAnnotation)

	annotationParser.
		EXPECT().
		Parse(fieldComment).
		Return(fieldAnnotation)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertSame(
		expected.Fields[0].Annotations,
		actual.TypeGroups[0].Types[0].Spec.(*StructSpec).Fields[0].Annotations,
	)
	ctrl.AssertSame(
		expected.Fields[1].Annotations,
		actual.TypeGroups[0].Types[0].Spec.(*StructSpec).Fields[1].Annotations,
	)
}

func TestSourceParser_Parse_WithStructSpecAndMultiStmt(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `package filePackageName

type typeName struct {
	field1Name, field2Name fieldType
}
`
	expected := &StructSpec{
		Fields: []*Field{
			{
				Name: "field1Name",
				Spec: &SimpleSpec{
					TypeName: "fieldType",
				},
			},
			{
				Name: "field2Name",
				Spec: &SimpleSpec{
					TypeName: "fieldType",
				},
			},
		},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
}

func TestSourceParser_Parse_WithInterfaceSpecAndWithoutFields(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `package filePackageName

type typeName interface {}
`
	expected := &InterfaceSpec{
		Fields: []*Field{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
}

func TestSourceParser_Parse_WithInterfaceSpecAndSimpleSpecFieldTypeAndNameAndFieldComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fieldComment := "field\ncomment"
	fieldAnnotation := []interface{}{
		TestAnnotation{
			Name: "fieldAnnotation",
		},
	}
	fileContent := `package filePackageName

type typeName interface {
	// field
	// comment
	fieldName
}
`
	expected := &InterfaceSpec{
		Fields: []*Field{
			{
				Name:        "",
				Comment:     "field\ncomment",
				Annotations: fieldAnnotation,
				Spec: &SimpleSpec{
					TypeName: "fieldName",
				},
			},
		},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(fieldComment).
		Return(fieldAnnotation)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertSame(
		expected.Fields[0].Annotations,
		actual.TypeGroups[0].Types[0].Spec.(*InterfaceSpec).Fields[0].Annotations,
	)
}

func TestSourceParser_Parse_WithInterfaceSpecAndSimpleSpecFieldType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `package filePackageName

type typeName interface {
	fieldName
}
`
	expected := &InterfaceSpec{
		Fields: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "fieldName",
				},
			},
		},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertSame(
		expected.Fields[0].Annotations,
		actual.TypeGroups[0].Types[0].Spec.(*InterfaceSpec).Fields[0].Annotations,
	)
}

func TestSourceParser_Parse_WithInterfaceSpecAndFuncSpecFieldTypeAndNameAndFieldComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fieldComment := "field\ncomment"
	fieldAnnotation := []interface{}{
		TestAnnotation{
			Name: "fieldAnnotation",
		},
	}
	fileContent := `package filePackageName

type typeName interface {
	// field
	// comment
	fieldName()
}
`
	expected := &InterfaceSpec{
		Fields: []*Field{
			{
				Name:        "fieldName",
				Comment:     "field\ncomment",
				Annotations: fieldAnnotation,
				Spec: &FuncSpec{
					Params:  []*Field{},
					Results: []*Field{},
				},
			},
		},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(fieldComment).
		Return(fieldAnnotation)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertSame(
		expected.Fields[0].Annotations,
		actual.TypeGroups[0].Types[0].Spec.(*InterfaceSpec).Fields[0].Annotations,
	)
}

func TestSourceParser_Parse_WithInterfaceSpecAndFuncSpecFieldType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `package filePackageName

type typeName interface {
	fieldName()
}
`
	expected := &InterfaceSpec{
		Fields: []*Field{
			{
				Name: "fieldName",
				Spec: &FuncSpec{
					Params:  []*Field{},
					Results: []*Field{},
				},
			},
		},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertSame(
		expected.Fields[0].Annotations,
		actual.TypeGroups[0].Types[0].Spec.(*InterfaceSpec).Fields[0].Annotations,
	)
}

func TestSourceParser_Parse_WithFuncSpecAndWithoutParamsAndWithoutResults(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `package filePackageName

type typeName func ()
`
	expected := &FuncSpec{
		Params:  []*Field{},
		Results: []*Field{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
}

func TestSourceParser_Parse_WithFuncSpecAndParamNameAndParamCommentAndIsVariadic(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	paramComment := "param\ncomment"
	paramAnnotation := []interface{}{
		TestAnnotation{
			Name: "paramAnnotation",
		},
	}
	fileContent := `package filePackageName

type typeName func (
	// param
	// comment
	paramName ...paramType,
)
`
	expected := &FuncSpec{
		Params: []*Field{
			{
				Name:        "paramName",
				Comment:     "param\ncomment",
				Annotations: paramAnnotation,
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "paramType",
					},
				},
			},
		},
		Results:    []*Field{},
		IsVariadic: true,
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(paramComment).
		Return(paramAnnotation)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertSame(
		expected.Params[0].Annotations,
		actual.TypeGroups[0].Types[0].Spec.(*FuncSpec).Params[0].Annotations,
	)
}

func TestSourceParser_Parse_WithFuncSpecIsVariadic(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `package filePackageName

type typeName func (...paramType)
`
	expected := &FuncSpec{
		Params: []*Field{
			{
				Spec: &ArraySpec{
					Value: &SimpleSpec{
						TypeName: "paramType",
					},
				},
			},
		},
		Results:    []*Field{},
		IsVariadic: true,
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
}

func TestSourceParser_Parse_WithFuncSpec(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `package filePackageName

type typeName func (paramType)
`
	expected := &FuncSpec{
		Params: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "paramType",
				},
			},
		},
		Results: []*Field{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
}

func TestSourceParser_Parse_WithFuncSpecAndMultiStmtParamAndParamNameAndParamComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	paramComment := "param\ncomment"
	paramAnnotation := []interface{}{
		TestAnnotation{
			Name: "paramAnnotation",
		},
	}
	fileContent := `package filePackageName

type typeName func (
	// param
	// comment
	param1Name, param2Name paramType,
)
`
	expected := &FuncSpec{
		Params: []*Field{
			{
				Name:        "param1Name",
				Comment:     "param\ncomment",
				Annotations: paramAnnotation,
				Spec: &SimpleSpec{
					TypeName: "paramType",
				},
			},
			{
				Name:        "param2Name",
				Comment:     "param\ncomment",
				Annotations: paramAnnotation,
				Spec: &SimpleSpec{
					TypeName: "paramType",
				},
			},
		},
		Results: []*Field{},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(paramComment).
		Return(paramAnnotation)

	annotationParser.
		EXPECT().
		Parse(paramComment).
		Return(paramAnnotation)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertSame(
		expected.Params[0].Annotations,
		actual.TypeGroups[0].Types[0].Spec.(*FuncSpec).Params[0].Annotations,
	)
}

func TestSourceParser_Parse_WithFuncSpecAnResultNameAndResultComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	resultComment := "result\ncomment"
	resultAnnotation := []interface{}{
		TestAnnotation{
			Name: "resultAnnotation",
		},
	}
	fileContent := `package filePackageName

type typeName func () (
	// result
	// comment
	resultName resultType,
)
`
	expected := &FuncSpec{
		Params: []*Field{},
		Results: []*Field{
			{
				Name:        "resultName",
				Comment:     "result\ncomment",
				Annotations: resultAnnotation,
				Spec: &SimpleSpec{
					TypeName: "resultType",
				},
			},
		},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(resultComment).
		Return(resultAnnotation)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertSame(
		expected.Results[0].Annotations,
		actual.TypeGroups[0].Types[0].Spec.(*FuncSpec).Results[0].Annotations,
	)
}

func TestSourceParser_Parse_WithFuncSpecAndResultType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	fileContent := `package filePackageName

type typeName func () resultType
`
	expected := &FuncSpec{
		Params: []*Field{},
		Results: []*Field{
			{
				Spec: &SimpleSpec{
					TypeName: "resultType",
				},
			},
		},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
}

func TestSourceParser_Parse_WithFuncSpecAndMultiStmtResultAndResultNameAndResultComment(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	fileName := "fileName"
	resultComment := "result\ncomment"
	resultAnnotation := []interface{}{
		TestAnnotation{
			Name: "resultAnnotation",
		},
	}
	fileContent := `package filePackageName

type typeName func () (
	// result
	// comment
	result1Name, result2Name resultType,
)
`
	expected := &FuncSpec{
		Params: []*Field{},
		Results: []*Field{
			{
				Name:        "result1Name",
				Comment:     "result\ncomment",
				Annotations: resultAnnotation,
				Spec: &SimpleSpec{
					TypeName: "resultType",
				},
			},
			{
				Name:        "result2Name",
				Comment:     "result\ncomment",
				Annotations: resultAnnotation,
				Spec: &SimpleSpec{
					TypeName: "resultType",
				},
			},
		},
	}

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	annotationParser.
		EXPECT().
		Parse(resultComment).
		Return(resultAnnotation)

	annotationParser.
		EXPECT().
		Parse(resultComment).
		Return(resultAnnotation)

	actual := parser.Parse(fileName, fileContent)

	ctrl.AssertNotNil(actual)
	ctrl.AssertNotEmpty(actual.TypeGroups)
	ctrl.AssertNotNil(actual.TypeGroups[0])
	ctrl.AssertNotEmpty(actual.TypeGroups[0].Types)
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0])
	ctrl.AssertNotNil(actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertEqual(expected, actual.TypeGroups[0].Types[0].Spec)
	ctrl.AssertSame(
		expected.Results[0].Annotations,
		actual.TypeGroups[0].Types[0].Spec.(*FuncSpec).Results[0].Annotations,
	)
}

func TestSourceParser_parseSpec_WithNilExpression(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	actual := parser.parseSpec(nil, nil, nil)

	ctrl.AssertNil(actual)
}

func TestSourceParser_parseSpec_WithUnknownExpressionType(t *testing.T) {
	ctrl := unit.NewController(t)
	defer ctrl.Finish()

	annotationParser := NewAnnotationParserMock(ctrl)

	parser := &GoSourceParser{
		annotationParser: annotationParser,
	}

	ctrl.Subtest("").
		Call(parser.parseSpec, &ast.BadExpr{}, nil, nil).
		ExpectPanic(
			NewErrorMessageConstraint("Variable 'expression' has not allowed type: *ast.BadExpr"),
		)
}
