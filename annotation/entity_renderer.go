package annotation

import (
	"go/format"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type EntityRenderer struct {
}

func NewEntityRenderer() *EntityRenderer {
	return &EntityRenderer{}
}

func (r *EntityRenderer) Render(entity interface{}) string {
	switch entity := entity.(type) {
	case *SimpleSpec:
		return r.renderSimpleSpec(entity)
	case *ArraySpec:
		return r.renderArraySpec(entity)
	case *MapSpec:
		return r.renderMapSpec(entity)
	case *FuncSpec:
		return r.renderFuncSpec(entity)
	case *InterfaceSpec:
		return r.renderInterfaceSpec(entity)
	case *StructSpec:
		return r.renderStructSpec(entity)
	case *Import:
		return r.renderImport(entity)
	case *ImportGroup:
		return r.renderImportGroup(entity)
	case *Const:
		return r.renderConst(entity)
	case *ConstGroup:
		return r.renderConstGroup(entity)
	case *Var:
		return r.renderVar(entity)
	case *VarGroup:
		return r.renderVarGroup(entity)
	case *Type:
		return r.renderType(entity)
	case *TypeGroup:
		return r.renderTypeGroup(entity)
	case *Func:
		return r.renderFunc(entity)
	case *File:
		return r.renderFile(entity)
	default:
		panic(errors.Errorf("Can't render entity with type: '%T'", entity))
	}
}

func (r *EntityRenderer) renderSimpleSpec(entity *SimpleSpec) string {
	result := ""

	if entity.IsPointer {
		result += "*"
	}

	if entity.PackageName != "" {
		result += entity.PackageName + "."
	}

	return result + entity.TypeName
}

func (r *EntityRenderer) renderArraySpec(entity *ArraySpec) string {
	result := ""

	if entity.Length != "" {
		result = "[" + entity.Length + "]"
	} else {
		result = "[]"
	}

	if _, ok := entity.Value.(*FuncSpec); ok {
		result += "func "
	}

	return result + r.Render(entity.Value)
}

func (r *EntityRenderer) renderMapSpec(entity *MapSpec) string {
	result := "map["

	if _, ok := entity.Key.(*FuncSpec); ok {
		result += "func "
	}

	result += r.Render(entity.Key) + "]"

	if _, ok := entity.Value.(*FuncSpec); ok {
		result += "func "
	}

	return result + r.Render(entity.Value)
}

func (r *EntityRenderer) renderFuncSpec(entity *FuncSpec) string {
	result := "("

	for i, field := range entity.Params {
		if i > 0 {
			result += ", "
		}

		if field.Comment != "" {
			result += "\n" + r.renderComment(field.Comment)
		}

		if field.Name != "" {
			result += field.Name + " "
		}

		if _, ok := field.Spec.(*FuncSpec); ok {
			result += "func "
		}

		if entity.IsVariadic && i == len(entity.Params)-1 {
			result += "..." + r.Render(field.Spec.(*ArraySpec).Value)
		} else {
			result += r.Render(field.Spec)
		}
	}

	result += ")"

	if len(entity.Results) == 0 {
		return result
	}

	result += " "

	if len(entity.Results) > 0 || entity.Results[0].Name != "" {
		result += "("
	}

	for i, field := range entity.Results {
		if i > 0 {
			result += ", "
		}

		if field.Comment != "" {
			result += "\n" + r.renderComment(field.Comment)
		}

		if field.Name != "" {
			result += field.Name + " "
		}

		if _, ok := field.Spec.(*FuncSpec); ok {
			result += "func "
		}

		result += r.Render(field.Spec)
	}

	if len(entity.Results) > 0 || entity.Results[0].Name != "" {
		result += ")"
	}

	return result
}

func (r *EntityRenderer) renderInterfaceSpec(entity *InterfaceSpec) string {
	if len(entity.Fields) == 0 {
		return "interface{}"
	}

	result := "interface{\n"

	for _, field := range entity.Fields {
		result += r.renderComment(field.Comment) +
			field.Name + r.Render(field.Spec) + "\n"
	}

	return result + "}"
}

func (r *EntityRenderer) renderStructSpec(entity *StructSpec) string {
	if len(entity.Fields) == 0 {
		return "struct{}"
	}

	result := "struct{\n"

	for _, field := range entity.Fields {
		result += r.renderComment(field.Comment)

		if field.Name != "" {
			result += field.Name + " "
		}

		if _, ok := field.Spec.(*FuncSpec); ok {
			result += "func "
		}

		result += r.Render(field.Spec)

		if field.Tag != "" {
			result += " " + strconv.Quote(field.Tag)
		}

		result += "\n"
	}

	return result + "}"
}

func (r *EntityRenderer) renderImport(entity *Import) string {
	result := r.renderComment(entity.Comment) +
		"import "

	if entity.Alias != "" {
		result += entity.Alias + " "
	}

	return result + strconv.Quote(entity.Namespace) + "\n"
}

func (r *EntityRenderer) renderImportGroup(entity *ImportGroup) string {
	result := r.renderComment(entity.Comment)

	if len(entity.Imports) == 0 {
		return result + "import ()\n"
	}

	if len(entity.Imports) == 1 && entity.Imports[0].Comment == "" {
		return result + r.Render(entity.Imports[0])
	}

	result += "import (\n"

	for _, element := range entity.Imports {
		result += r.renderComment(element.Comment)

		if element.Alias != "" {
			result += element.Alias + " "
		}

		result += strconv.Quote(element.Namespace) + "\n"
	}

	return result + ")\n"
}

func (r *EntityRenderer) renderConst(entity *Const) string {
	if entity.Value == "" {
		panic(errors.New("Variable 'Value' must be not empty"))
	}

	result := r.renderComment(entity.Comment) +
		"const " + entity.Name

	if entity.Spec != nil {
		result += " " + r.Render(entity.Spec)
	}

	return result + " = " + entity.Value + "\n"
}

func (r *EntityRenderer) renderConstGroup(entity *ConstGroup) string {
	result := r.renderComment(entity.Comment)

	if len(entity.Consts) == 0 {
		return result + "const ()\n"
	}

	if len(entity.Consts) == 1 && entity.Consts[0].Comment == "" {
		return result + r.Render(entity.Consts[0])
	}

	result += "const (\n"

	for _, element := range entity.Consts {
		result += r.renderComment(element.Comment) +
			element.Name

		if element.Spec != nil {
			result += " " + r.Render(element.Spec)
		}

		if element.Value != "" {
			result += " = " + element.Value
		}

		result += "\n"
	}

	return result + ")\n"
}

func (r *EntityRenderer) renderVar(entity *Var) string {
	result := r.renderComment(entity.Comment) +
		"var " + entity.Name

	if entity.Spec != nil {
		result += " " + r.Render(entity.Spec)
	}

	if entity.Value != "" {
		result += " = " + entity.Value
	}

	return result + "\n"
}

func (r *EntityRenderer) renderVarGroup(entity *VarGroup) string {
	result := r.renderComment(entity.Comment)

	if len(entity.Vars) == 0 {
		return result + "var ()\n"
	}

	if len(entity.Vars) == 1 && entity.Vars[0].Comment == "" {
		return result + r.Render(entity.Vars[0])
	}

	result += "var (\n"

	for _, element := range entity.Vars {
		result += r.renderComment(element.Comment) +
			element.Name

		if element.Spec != nil {
			result += " " + r.Render(element.Spec)
		}

		if element.Value != "" {
			result += " = " + element.Value
		}

		result += "\n"
	}

	return result + ")\n"
}

func (r *EntityRenderer) renderType(entity *Type) string {
	return r.renderComment(entity.Comment) +
		"type " + entity.Name + " " + r.Render(entity.Spec) + "\n"
}

func (r *EntityRenderer) renderTypeGroup(entity *TypeGroup) string {
	result := r.renderComment(entity.Comment)

	if len(entity.Types) == 0 {
		return result + "type ()\n"
	}

	if len(entity.Types) == 1 && entity.Types[0].Comment == "" {
		return result + r.Render(entity.Types[0])
	}

	result += "type (\n"

	for _, element := range entity.Types {
		result += r.renderComment(element.Comment)
		result += element.Name + " " + r.Render(element.Spec) + "\n"
	}

	return result + ")\n"
}

func (r *EntityRenderer) renderFunc(entity *Func) string {
	result := r.renderComment(entity.Comment) +
		"func "

	if entity.Related != nil {
		result += "("

		if entity.Related.Comment != "" {
			result += "\n" + r.renderComment(entity.Related.Comment)
		}

		if entity.Related.Name != "" {
			result += entity.Related.Name + " "
		}

		result += r.Render(entity.Related.Spec) + ") "
	}

	result += entity.Name

	if entity.Spec == nil {
		result += "()"
	} else {
		result += r.Render(entity.Spec)
	}

	return result + " {\n" + entity.Content + "\n}\n"
}

func (r *EntityRenderer) renderFile(entity *File) string {
	if entity.Content != "" {
		return entity.Content
	}

	result := r.renderComment(entity.Comment)

	result += "package " + entity.PackageName + "\n\n"

	for _, element := range entity.ImportGroups {
		result += r.Render(element) + "\n"
	}

	for _, element := range entity.ConstGroups {
		result += r.Render(element) + "\n"
	}

	for _, element := range entity.VarGroups {
		result += r.Render(element) + "\n"
	}

	for _, element := range entity.TypeGroups {
		result += r.Render(element) + "\n"
	}

	for _, element := range entity.Funcs {
		result += r.Render(element) + "\n"
	}

	formattedResult, err := format.Source([]byte(result))

	if err != nil {
		panic(err)
	}

	return string(formattedResult)
}

func (r *EntityRenderer) renderComment(comment string) string {
	if comment == "" {
		return ""
	}

	return "// " + strings.Join(strings.Split(strings.TrimSpace(comment), "\n"), "\n// ") + "\n"
}
