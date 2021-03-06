package annotation

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type ellipsis struct {
	ArraySpec
}

// Parses golang sources and build File models.
type GoSourceParser struct {
	annotationParser AnnotationParser
}

// Creates new instance of GoSourceParser.
func NewGoSourceParser(annotationParser AnnotationParser) *GoSourceParser {
	if annotationParser == nil {
		panic(errors.New("Variable 'annotationParser' must be not nil"))
	}

	return &GoSourceParser{
		annotationParser: annotationParser,
	}
}

// Generates File model by golang source.
func (p *GoSourceParser) Parse(fileName string, content string) *File {
	fileSet := token.NewFileSet()
	astFile, err := parser.ParseFile(fileSet, fileName, content, parser.ParseComments)

	if err != nil {
		panic(err)
	}

	result := &File{
		Name:         fileName,
		Content:      content,
		PackageName:  astFile.Name.Name,
		Comment:      strings.TrimSpace(astFile.Doc.Text()),
		ImportGroups: []*ImportGroup{},
		ConstGroups:  []*ConstGroup{},
		VarGroups:    []*VarGroup{},
		TypeGroups:   []*TypeGroup{},
		Funcs:        []*Func{},
	}

	if result.Comment != "" {
		result.Annotations = p.annotationParser.Parse(result.Comment)
	}

	for _, node := range astFile.Decls {
		if decl, ok := node.(*ast.GenDecl); ok {
			switch decl.Tok {
			case token.IMPORT:
				result.ImportGroups = append(result.ImportGroups, p.parseImportGroup(decl))
			case token.CONST:
				result.ConstGroups = append(result.ConstGroups, p.parseConstGroup(decl, astFile, fileSet))
			case token.VAR:
				result.VarGroups = append(result.VarGroups, p.parseVarGroup(decl, astFile, fileSet))
			case token.TYPE:
				result.TypeGroups = append(result.TypeGroups, p.parseTypeGroup(decl, astFile, fileSet))
			}
		}

		if decl, ok := node.(*ast.FuncDecl); ok {
			result.Funcs = append(result.Funcs, p.parseFunc(decl, astFile, fileSet))
		}
	}

	return result
}

func (p *GoSourceParser) parseImportGroup(decl *ast.GenDecl) *ImportGroup {
	result := &ImportGroup{
		Comment: strings.TrimSpace(decl.Doc.Text()),
		Imports: []*Import{},
	}

	if result.Comment != "" {
		result.Annotations = p.annotationParser.Parse(result.Comment)
	}

	for _, spec := range decl.Specs {
		// This row is expected to be called with import spec
		importSpec, _ := spec.(*ast.ImportSpec)
		element := &Import{
			Namespace: strings.Trim(importSpec.Path.Value, "\""),
			Comment:   strings.TrimSpace(importSpec.Doc.Text()),
		}

		if element.Comment != "" {
			element.Annotations = p.annotationParser.Parse(element.Comment)
		}

		if importSpec.Name != nil {
			element.Alias = importSpec.Name.Name
		}

		result.Imports = append(result.Imports, element)
	}

	return result
}

func (p *GoSourceParser) parseConstGroup(decl *ast.GenDecl, astFile *ast.File, fileSet *token.FileSet) *ConstGroup {
	result := &ConstGroup{
		Comment: strings.TrimSpace(decl.Doc.Text()),
		Consts:  []*Const{},
	}

	if result.Comment != "" {
		result.Annotations = p.annotationParser.Parse(result.Comment)
	}

	var previousSpec *SimpleSpec

	for _, spec := range decl.Specs {
		// This row is expected to be called with const spec
		constSpec, _ := spec.(*ast.ValueSpec)
		comment := strings.TrimSpace(constSpec.Doc.Text())

		for i, name := range constSpec.Names {
			element := &Const{
				Name:    name.Name,
				Comment: comment,
			}

			if element.Comment != "" {
				element.Annotations = p.annotationParser.Parse(element.Comment)
			}

			if constSpec.Type != nil {
				element.Spec = p.parseSpec(constSpec.Type, astFile, fileSet).(*SimpleSpec)
			}

			var value ast.Expr

			if len(constSpec.Values) > 0 && i <= len(constSpec.Values) {
				value = constSpec.Values[i]

				buffer := bytes.Buffer{}
				// FileSet is not changed after parse
				_ = printer.Fprint(&buffer, fileSet, value)

				element.Value = buffer.String()
			}

			if element.Spec == nil {
				if value != nil {
					if basicLit, ok := value.(*ast.BasicLit); ok {
						switch basicLit.Kind {
						case token.STRING:
							element.Spec = &SimpleSpec{
								TypeName: "string",
							}
						case token.INT:
							element.Spec = &SimpleSpec{
								TypeName: "int",
							}
						case token.FLOAT:
							element.Spec = &SimpleSpec{
								TypeName: "float64",
							}
						}
					}
				} else {
					element.Spec = previousSpec
				}
			}

			result.Consts = append(result.Consts, element)

			previousSpec = element.Spec
		}
	}

	return result
}

func (p *GoSourceParser) parseVarGroup(decl *ast.GenDecl, astFile *ast.File, fileSet *token.FileSet) *VarGroup {
	result := &VarGroup{
		Comment: strings.TrimSpace(decl.Doc.Text()),
		Vars:    []*Var{},
	}

	if result.Comment != "" {
		result.Annotations = p.annotationParser.Parse(result.Comment)
	}

	for _, spec := range decl.Specs {
		// This row is expected to be called with value spec
		varSpec, _ := spec.(*ast.ValueSpec)
		comment := strings.TrimSpace(varSpec.Doc.Text())

		for i, name := range varSpec.Names {
			element := &Var{
				Name:    name.Name,
				Comment: comment,
			}

			if element.Comment != "" {
				element.Annotations = p.annotationParser.Parse(element.Comment)
			}

			if varSpec.Type != nil {
				element.Spec = p.parseSpec(varSpec.Type, astFile, fileSet)
			}

			var value ast.Expr

			if len(varSpec.Values) > 0 && i <= len(varSpec.Values) {
				value = varSpec.Values[i]

				buffer := bytes.Buffer{}
				// FileSet is not changed after parse
				_ = printer.Fprint(&buffer, fileSet, value)

				element.Value = buffer.String()
			}

			if element.Spec == nil {
				if value != nil {
					if basicLit, ok := value.(*ast.BasicLit); ok {
						switch basicLit.Kind {
						case token.STRING:
							element.Spec = &SimpleSpec{
								TypeName: "string",
							}
						case token.INT:
							element.Spec = &SimpleSpec{
								TypeName: "int",
							}
						case token.FLOAT:
							element.Spec = &SimpleSpec{
								TypeName: "float64",
							}
						}
					}
				}
			}

			result.Vars = append(result.Vars, element)
		}
	}

	return result
}

func (p *GoSourceParser) parseTypeGroup(decl *ast.GenDecl, astFile *ast.File, fileSet *token.FileSet) *TypeGroup {
	result := &TypeGroup{
		Comment: strings.TrimSpace(decl.Doc.Text()),
		Types:   []*Type{},
	}

	if result.Comment != "" {
		result.Annotations = p.annotationParser.Parse(result.Comment)
	}

	for _, spec := range decl.Specs {
		// This row is expected to be called with type spec
		typeSpec, _ := spec.(*ast.TypeSpec)
		name := ""

		if typeSpec.Name != nil {
			name = typeSpec.Name.Name
		}

		element := &Type{
			Name:    name,
			Comment: strings.TrimSpace(typeSpec.Doc.Text()),
			Spec:    p.parseSpec(typeSpec.Type, astFile, fileSet),
		}

		if element.Comment != "" {
			element.Annotations = p.annotationParser.Parse(element.Comment)
		}

		result.Types = append(result.Types, element)
	}

	return result
}

func (p *GoSourceParser) parseFunc(decl *ast.FuncDecl, astFile *ast.File, fileSet *token.FileSet) *Func {
	result := &Func{
		Comment: strings.TrimSpace(decl.Doc.Text()),
		Spec:    p.parseFuncSpec(decl.Type, astFile, fileSet),
	}

	buffer := bytes.Buffer{}
	// FileSet is not changed after parse
	_ = printer.Fprint(&buffer, fileSet, decl.Body.List)

	result.Content = buffer.String()

	if result.Comment != "" {
		result.Annotations = p.annotationParser.Parse(result.Comment)
	}

	if decl.Name != nil {
		result.Name = decl.Name.Name
	}

	related := p.parseFieldsList(decl.Recv, astFile, fileSet)

	if len(related) == 1 {
		result.Related = related[0]
	}

	return result
}

func (p *GoSourceParser) parseSpec(expression ast.Expr, astFile *ast.File, fileSet *token.FileSet) interface{} {
	if expression == nil {
		return nil
	}

	switch expression := expression.(type) {
	case *ast.Ident:
		return p.parseIdentSpec(expression)
	case *ast.SelectorExpr:
		return p.parseSelectorExprSpec(expression)
	case *ast.StarExpr:
		return p.parseStarExprSpec(expression, astFile, fileSet)
	case *ast.ArrayType:
		return p.parseArraySpec(expression, astFile, fileSet)
	case *ast.Ellipsis:
		return p.parseEllipsisSpec(expression, astFile, fileSet)
	case *ast.MapType:
		return p.parseMapSpec(expression, astFile, fileSet)
	case *ast.FuncType:
		return p.parseFuncSpec(expression, astFile, fileSet)
	case *ast.StructType:
		return p.parseStructSpec(expression, astFile, fileSet)
	case *ast.InterfaceType:
		return p.parseInterfaceSpec(expression, astFile, fileSet)
	default:
		panic(errors.Errorf("Variable 'expression' has not allowed type: %T", expression))
	}
}

func (p *GoSourceParser) parseIdentSpec(node *ast.Ident) *SimpleSpec {
	return &SimpleSpec{
		TypeName: node.Name,
	}
}

func (p *GoSourceParser) parseSelectorExprSpec(node *ast.SelectorExpr) *SimpleSpec {
	return &SimpleSpec{
		TypeName:    node.Sel.Name,
		PackageName: node.X.(*ast.Ident).Name,
	}
}

func (p *GoSourceParser) parseStarExprSpec(node *ast.StarExpr, astFile *ast.File, fileSet *token.FileSet) *SimpleSpec {
	result := p.parseSpec(node.X, astFile, fileSet).(*SimpleSpec)
	result.IsPointer = true

	return result
}

func (p *GoSourceParser) parseArraySpec(node *ast.ArrayType, astFile *ast.File, fileSet *token.FileSet) *ArraySpec {
	value := p.parseSpec(node.Elt, astFile, fileSet)

	result := &ArraySpec{
		Value: value,
	}

	if node.Len != nil {
		buffer := bytes.Buffer{}
		// FileSet is not changed after parse
		_ = printer.Fprint(&buffer, fileSet, node.Len)

		result.Length = buffer.String()
	}

	return result
}

func (p *GoSourceParser) parseEllipsisSpec(node *ast.Ellipsis, astFile *ast.File, fileSet *token.FileSet) *ellipsis {
	value := p.parseSpec(node.Elt, astFile, fileSet)

	return &ellipsis{
		ArraySpec{
			Value: value,
		},
	}
}

func (p *GoSourceParser) parseMapSpec(node *ast.MapType, astFile *ast.File, fileSet *token.FileSet) *MapSpec {
	key := p.parseSpec(node.Key, astFile, fileSet)
	value := p.parseSpec(node.Value, astFile, fileSet)

	return &MapSpec{
		Key:   key,
		Value: value,
	}
}

func (p *GoSourceParser) parseFieldsList(node *ast.FieldList, astFile *ast.File, fileSet *token.FileSet) []*Field {
	result := []*Field{}

	if node == nil {
		return result
	}

	beforeCommentPosition := node.Pos()
	afterCommentPosition := node.Pos()

	for _, astField := range node.List {
		tag := ""

		if astField.Tag != nil {
			// Invalid tag could not be parsed
			tag, _ = strconv.Unquote(astField.Tag.Value)
		}

		spec := p.parseSpec(astField.Type, astFile, fileSet)
		comment := strings.TrimSpace(astField.Doc.Text())

		if comment == "" {
			if len(astField.Names) == 0 {
				afterCommentPosition = astField.Type.Pos()
			} else {
				afterCommentPosition = astField.Names[0].Pos()
			}

			for _, commentGroup := range astFile.Comments {
				position := commentGroup.Pos()

				if position >= beforeCommentPosition && position <= afterCommentPosition {
					comment = strings.TrimSpace(commentGroup.Text())
				}
			}

			if len(astField.Names) == 0 {
				beforeCommentPosition = astField.Type.End()
			} else {
				beforeCommentPosition = astField.Names[0].End()
			}
		}

		if len(astField.Names) == 0 {
			field := &Field{
				Spec:    spec,
				Tag:     tag,
				Comment: comment,
			}

			if comment != "" {
				field.Annotations = p.annotationParser.Parse(comment)
			}

			result = append(result, field)
		} else {
			for _, name := range astField.Names {
				field := &Field{
					Name:    name.Name,
					Spec:    spec,
					Tag:     tag,
					Comment: comment,
				}

				if comment != "" {
					field.Annotations = p.annotationParser.Parse(comment)
				}

				result = append(result, field)
			}
		}
	}

	return result
}

func (p *GoSourceParser) parseStructSpec(node *ast.StructType, astFile *ast.File, fileSet *token.FileSet) *StructSpec {
	fields := p.parseFieldsList(node.Fields, astFile, fileSet)

	return &StructSpec{
		Fields: fields,
	}
}

func (p *GoSourceParser) parseInterfaceSpec(
	node *ast.InterfaceType,
	astFile *ast.File,
	fileSet *token.FileSet,
) *InterfaceSpec {
	methods := p.parseFieldsList(node.Methods, astFile, fileSet)

	return &InterfaceSpec{
		Fields: methods,
	}
}

func (p *GoSourceParser) parseFuncSpec(node *ast.FuncType, astFile *ast.File, fileSet *token.FileSet) *FuncSpec {
	params := p.parseFieldsList(node.Params, astFile, fileSet)
	results := p.parseFieldsList(node.Results, astFile, fileSet)
	isVariadic := false

	if len(params) > 0 {
		id := len(params) - 1

		if spec, ok := params[id].Spec.(*ellipsis); ok {
			params[id].Spec = &spec.ArraySpec
			isVariadic = true
		}
	}

	return &FuncSpec{
		Params:     params,
		Results:    results,
		IsVariadic: isVariadic,
	}
}
