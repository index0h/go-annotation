package annotation

import (
	"strconv"
	"strings"
)

type SpecRenderer struct {
}

func (r *SpecRenderer) RenderSpec(element interface{}) string {
	if element == nil {
		return ""
	}

	switch element.(type) {
	case *ArraySpec:
		return r.renderArraySpec(element.(*ArraySpec))
	case *FuncSpec:
		return r.renderFuncSpec(element.(*FuncSpec))
	case *InterfaceSpec:
		return r.renderInterfaceSpec(element.(*InterfaceSpec))
	case *MapSpec:
		return r.renderMapSpec(element.(*MapSpec))
	case *SimpleSpec:
		return r.renderSimpleSpec(element.(*SimpleSpec))
	case *StructSpec:
		return r.renderStructSpec(element.(*StructSpec))
	case *Func:
		return "func " + r.renderFuncSpec(element.(*Func).Spec)
	default:
		panic(NewErrorf("Unknown spec %T", element))
	}
}

func (r *SpecRenderer) renderComment(comment string) string {
	if comment == "" {
		return ""
	}

	return "// " + strings.Join(strings.Split(strings.TrimSpace(comment), "\n"), "\n// ") + "\n"
}

func (r *SpecRenderer) renderArraySpec(element *ArraySpec) string {
	result := ""

	if element.IsEllipsis {
		if element.IsFixedLength {
			result = "[...]"
		} else {
			result = "..."
		}
	} else if element.IsFixedLength {
		result = "[" + strconv.Itoa(element.Length) + "]"
	} else {
		result = "[]"
	}

	return result + r.RenderSpec(element.Value)
}

func (r *SpecRenderer) renderFuncSpec(element *FuncSpec) string {
	result := "("

	for i, field := range element.Params {
		if i > 0 {
			result += ", "
		}

		result += r.renderComment(field.Comment)

		if field.Name != "" {
			result += field.Name + " "
		}

		result += r.RenderSpec(field.Spec)
	}

	result += ")"

	if len(element.Results) > 0 {
		result += " ("

		for i, field := range element.Results {
			if i > 0 {
				result += ", "
			}

			if field.Comment != "" {
				result += "\n" + r.renderComment(field.Comment)
			}

			if field.Name != "" {
				result += field.Name + " "
			}

			result += r.RenderSpec(field.Spec)
		}

		result += ")"
	}

	return result
}

func (r *SpecRenderer) renderInterfaceSpec(element *InterfaceSpec) string {
	result := "interface{"

	if len(element.Methods) > 0 {
		result += "\n"
	}

	for _, method := range element.Methods {
		if method.Comment != "" {
			result += "\n" + r.renderComment(method.Comment)
		}

		result += method.Name + " " + r.RenderSpec(method.Spec) + "\n"
	}

	return result + "}"
}

func (r *SpecRenderer) renderMapSpec(element *MapSpec) string {
	return "map[" + r.RenderSpec(element.Key) + "]" + r.RenderSpec(element.Value)
}

func (r *SpecRenderer) renderSimpleSpec(element *SimpleSpec) string {
	result := ""

	if element.IsPointer {
		result += "*"
	}

	if element.PackageName != "" {
		result += element.PackageName + "."
	}

	return result + element.TypeName
}

func (r *SpecRenderer) renderStructSpec(element *StructSpec) string {
	result := "struct{\n"

	for _, field := range element.Fields {
		if field.Comment != "" {
			result = "\n" + r.renderComment(field.Comment)
		}

		result += field.Name + " " + r.RenderSpec(field.Spec) + " "

		if field.Tag != "" {
			result += field.Tag + " "
		}

		result += "\n"
	}

	return result + "}"
}
