package annotation

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var annotationRegexp = regexp.MustCompile(`(?m)^([a-zA-Z][a-zA-Z0-9_]+)\(((.|\n)*)\)$`)

type Parser struct {
	annotations map[string]interface{}
}

func NewParser() *Parser {
	return &Parser{
		annotations: map[string]interface{}{
			"FileIsGenerated": FileIsGeneratedAnnotation(false),
		},
	}
}

func (p *Parser) AddAnnotation(name string, annotationType interface{}) *Parser {
	if name == "" {
		panic(NewNotEmptyError("name"))
	}

	if _, ok := p.annotations[name]; ok {
		panic(NewErrorf("Annotation '%s' already registered", name))
	}

	p.annotations[name] = annotationType

	return p
}

func (p *Parser) Process(storage *Storage) {
	for _, namespace := range storage.Namespaces {
		for _, file := range namespace.Files {
			fileSet := token.NewFileSet()
			astFile, err := parser.ParseFile(fileSet, file.Name, file.Content, parser.ParseComments)

			if err != nil {
				panic(err)
			}

			file.PackageName = astFile.Name.Name
			file.Comment = strings.TrimSpace(astFile.Doc.Text())
			file.Annotations = p.parseComment(file.Comment)

			for _, node := range astFile.Decls {
				if decl, ok := node.(*ast.GenDecl); ok {
					switch decl.Tok {
					case token.IMPORT:
						if file.ImportGroups == nil {
							file.ImportGroups = []*ImportGroup{}
						}

						p.parseImport(storage, file, decl)
					case token.CONST:
						if file.ConstGroups == nil {
							file.ConstGroups = []*ConstGroup{}
						}

						p.parseConst(storage, file, decl, astFile)
					case token.VAR:
						if file.VarGroups == nil {
							file.VarGroups = []*VarGroup{}
						}

						p.parseVar(storage, file, decl, astFile)
					case token.TYPE:
						if file.TypeGroups == nil {
							file.TypeGroups = []*TypeGroup{}
						}

						p.parseType(storage, file, decl, astFile)
					}
				}

				if decl, ok := node.(*ast.FuncDecl); ok {
					if file.Funcs == nil {
						file.Funcs = []*Func{}
					}

					p.parseFunc(storage, file, decl, astFile)
				}
			}
		}
	}
}

func (p *Parser) parseImport(storage *Storage, file *File, decl *ast.GenDecl) {
	groupComment := strings.TrimSpace(decl.Doc.Text())
	group := &ImportGroup{
		Comment:     strings.TrimSpace(decl.Doc.Text()),
		Annotations: p.parseComment(groupComment),
		Imports:     []*Import{},
	}

	file.ImportGroups = append(file.ImportGroups, group)

	for _, spec := range decl.Specs {
		// This method is expected to be called with import spec
		importSpec, _ := spec.(*ast.ImportSpec)
		entityComment := strings.TrimSpace(importSpec.Doc.Text())
		entity := &Import{
			Namespace:   strings.Trim(importSpec.Path.Value, "\""),
			Comment:     entityComment,
			Annotations: p.parseComment(entityComment),
		}

		if importSpec.Name == nil {
			_, entity.Alias = filepath.Split(entity.Namespace)
		} else {
			entity.Alias = importSpec.Name.Name
		}

		group.Imports = append(group.Imports, entity)
	}
}

func (p *Parser) parseConst(storage *Storage, file *File, decl *ast.GenDecl, astFile *ast.File) {
	groupComment := strings.TrimSpace(decl.Doc.Text())
	group := &ConstGroup{
		Comment:     strings.TrimSpace(decl.Doc.Text()),
		Annotations: p.parseComment(groupComment),
		Consts:      []*Const{},
	}

	file.ConstGroups = append(file.ConstGroups, group)

	var previous interface{}

	for _, spec := range decl.Specs {
		// This method is expected to be called with const spec
		constSpec, _ := spec.(*ast.ValueSpec)
		comment := strings.TrimSpace(constSpec.Doc.Text())
		spec := p.parseSpec(constSpec.Type, astFile)

		if spec == nil {
			if len(constSpec.Values) > 0 {
				basicLit, ok := constSpec.Values[0].(*ast.BasicLit)

				if ok {
					typeName := ""

					switch basicLit.Kind {
					case token.STRING:
						typeName = "string"
					case token.INT:
						typeName = "int"
					case token.FLOAT:
						typeName = "float64"
					}

					spec = &SimpleSpec{
						TypeName: typeName,
					}
				}
			} else if previous != nil {
				spec = previous
			}
		}

		for _, name := range constSpec.Names {
			entity := &Const{
				Name:        name.Name,
				Comment:     comment,
				Annotations: p.parseComment(comment),
				Spec:        spec,
			}

			group.Consts = append(group.Consts, entity)
		}

		previous = spec
	}
}

func (p *Parser) parseVar(storage *Storage, file *File, decl *ast.GenDecl, astFile *ast.File) {
	groupComment := strings.TrimSpace(decl.Doc.Text())
	group := &VarGroup{
		Comment:     strings.TrimSpace(decl.Doc.Text()),
		Annotations: p.parseComment(groupComment),
		Vars:        []*Var{},
	}

	file.VarGroups = append(file.VarGroups, group)

	for _, spec := range decl.Specs {
		// This method is expected to be called with value spec
		varSpec, _ := spec.(*ast.ValueSpec)
		comment := strings.TrimSpace(varSpec.Doc.Text())
		spec := p.parseSpec(varSpec.Type, astFile)

		if spec == nil {
			if len(varSpec.Values) > 0 {
				basicLit, ok := varSpec.Values[0].(*ast.BasicLit)

				if ok {
					typeName := ""

					switch basicLit.Kind {
					case token.STRING:
						typeName = "string"
					case token.INT:
						typeName = "int"
					case token.FLOAT:
						typeName = "float64"
					}

					spec = &SimpleSpec{
						TypeName: typeName,
					}
				}
			}
		}

		for _, name := range varSpec.Names {
			entity := &Var{
				Name:        name.Name,
				Comment:     comment,
				Annotations: p.parseComment(comment),
				Spec:        spec,
			}

			group.Vars = append(group.Vars, entity)
		}
	}
}

func (p *Parser) parseType(storage *Storage, file *File, decl *ast.GenDecl, astFile *ast.File) {
	groupComment := strings.TrimSpace(decl.Doc.Text())
	group := &TypeGroup{
		Comment:     strings.TrimSpace(decl.Doc.Text()),
		Annotations: p.parseComment(groupComment),
		Types:       []*Type{},
	}

	file.TypeGroups = append(file.TypeGroups, group)

	for _, spec := range decl.Specs {
		// This method is expected to be called with type spec
		typeSpec, _ := spec.(*ast.TypeSpec)
		name := ""

		if typeSpec.Name != nil {
			name = typeSpec.Name.Name
		}

		comment := strings.TrimSpace(typeSpec.Doc.Text())

		entity := &Type{
			Name:        name,
			Comment:     comment,
			Annotations: p.parseComment(comment),
			Spec:        p.parseSpec(typeSpec.Type, astFile),
		}

		group.Types = append(group.Types, entity)
	}
}

func (p *Parser) parseFunc(storage *Storage, file *File, decl *ast.FuncDecl, astFile *ast.File) {
	p.parseFuncSpec(decl.Type, astFile)

	comment := strings.TrimSpace(decl.Doc.Text())
	entity := &Func{
		Comment:     comment,
		Annotations: p.parseComment(comment),
		Spec:        p.parseFuncSpec(decl.Type, astFile),
	}

	file.Funcs = append(file.Funcs, entity)

	if decl.Name != nil {
		entity.Name = decl.Name.Name
	}

	related := p.parseFieldsList(decl.Recv, astFile)

	if len(related) == 1 {
		entity.Related = related[0]
	}
}

func (p *Parser) parseSpec(expression ast.Expr, astFile *ast.File) interface{} {
	if expression == nil {
		return nil
	}

	switch expression.(type) {
	case *ast.Ident:
		return p.parseIdentSpec(expression.(*ast.Ident))
	case *ast.SelectorExpr:
		return p.parseSelectorExprSpec(expression.(*ast.SelectorExpr))
	case *ast.StarExpr:
		return p.parseStarExprSpec(expression.(*ast.StarExpr), astFile)
	case *ast.ArrayType:
		return p.parseArraySpec(expression.(*ast.ArrayType), astFile)
	case *ast.Ellipsis:
		return p.parseEllipsisSpec(expression.(*ast.Ellipsis), astFile)
	case *ast.MapType:
		return p.parseMapSpec(expression.(*ast.MapType), astFile)
	case *ast.FuncType:
		return p.parseFuncSpec(expression.(*ast.FuncType), astFile)
	case *ast.StructType:
		return p.parseStructSpec(expression.(*ast.StructType), astFile)
	case *ast.InterfaceType:
		return p.parseInterfaceSpec(expression.(*ast.InterfaceType), astFile)
	default:
		panic(NewErrorf("Unknown type spec expression %+v", expression))
	}
}

func (p *Parser) parseIdentSpec(node *ast.Ident) *SimpleSpec {
	return &SimpleSpec{
		TypeName: node.Name,
	}
}

func (p *Parser) parseSelectorExprSpec(node *ast.SelectorExpr) *SimpleSpec {
	return &SimpleSpec{
		TypeName:    node.Sel.Name,
		PackageName: node.X.(*ast.Ident).Name,
	}
}

func (p *Parser) parseStarExprSpec(node *ast.StarExpr, astFile *ast.File) *SimpleSpec {
	result := p.parseSpec(node.X, astFile).(*SimpleSpec)
	result.IsPointer = true

	return result
}

func (p *Parser) parseArraySpec(node *ast.ArrayType, astFile *ast.File) *ArraySpec {
	value := p.parseSpec(node.Elt, astFile)

	result := &ArraySpec{
		Value: value,
	}

	if node.Len != nil {
		result.IsFixedLength = true

		if basicLit, ok := node.Len.(*ast.BasicLit); ok {
			result.Length, _ = strconv.Atoi(basicLit.Value)
		}
	}

	return result
}

func (p *Parser) parseEllipsisSpec(node *ast.Ellipsis, astFile *ast.File) *ArraySpec {
	value := p.parseSpec(node.Elt, astFile)

	return &ArraySpec{
		Value:      value,
		IsEllipsis: true,
	}
}

func (p *Parser) parseMapSpec(node *ast.MapType, astFile *ast.File) *MapSpec {
	key := p.parseSpec(node.Key, astFile)
	value := p.parseSpec(node.Value, astFile)

	return &MapSpec{
		Key:   key,
		Value: value,
	}
}

func (p *Parser) parseFuncSpec(node *ast.FuncType, astFile *ast.File) *FuncSpec {
	params := p.parseFieldsList(node.Params, astFile)
	results := p.parseFieldsList(node.Results, astFile)

	return &FuncSpec{
		Params:  params,
		Results: results,
	}
}

func (p *Parser) parseFieldsList(node *ast.FieldList, astFile *ast.File) []*Field {
	result := []*Field{}

	if node == nil {
		return result
	}

	beforeCommentPosition := node.Pos()
	afterCommentPosition := beforeCommentPosition

	for _, astField := range node.List {
		tag := ""

		if astField.Tag != nil {
			// Invalid tag could not be parsed
			tag, _ = strconv.Unquote(astField.Tag.Value)
		}

		spec := p.parseSpec(astField.Type, astFile)
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
				Spec:        spec,
				Tag:         tag,
				Comment:     comment,
				Annotations: p.parseComment(comment),
			}

			result = append(result, field)
		} else {
			for _, name := range astField.Names {
				entity := &Field{
					Name:        name.Name,
					Spec:        spec,
					Tag:         tag,
					Comment:     comment,
					Annotations: p.parseComment(comment),
				}

				result = append(result, entity)
			}
		}
	}

	return result
}

func (p *Parser) parseStructSpec(node *ast.StructType, astFile *ast.File) *StructSpec {
	fields := p.parseFieldsList(node.Fields, astFile)

	return &StructSpec{
		Fields: fields,
	}
}

func (p *Parser) parseInterfaceSpec(node *ast.InterfaceType, astFile *ast.File) *InterfaceSpec {
	methods := p.parseFieldsList(node.Methods, astFile)

	return &InterfaceSpec{
		Methods: methods,
	}
}

func (p *Parser) parseComment(comment string) []interface{} {
	result := []interface{}{}

	parts := strings.Split("\n"+comment, "\n@")

	for i := 1; i < len(parts); i++ {
		for _, elements := range annotationRegexp.FindAllStringSubmatch(parts[i], -1) {
			name := elements[1]
			content := strings.TrimSpace(elements[2])

			if annotation, ok := p.annotations[name]; !ok {
				panic(NewErrorf("Unknown annotation '%s'", name))
			} else {
				value := reflect.New(reflect.TypeOf(annotation)).Interface()

				if len(content) > 0 {
					if err := json.Unmarshal([]byte(content), &value); err != nil {
						panic(err)
					}
				}

				result = append(result, reflect.ValueOf(value).Elem().Interface())
			}
		}
	}

	return result
}
