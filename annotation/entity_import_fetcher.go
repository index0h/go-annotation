package annotation

import (
	"regexp"

	"github.com/pkg/errors"
)

type EntityImportFetcher struct {
	importUniquer ImportUniquer
}

func NewEntityImportFetcher(importUniquer ImportUniquer) *EntityImportFetcher {
	if importUniquer == nil {
		panic(errors.New("Variable 'importUniquer' must be not nil"))
	}

	return &EntityImportFetcher{
		importUniquer: importUniquer,
	}
}

func (f *EntityImportFetcher) Fetch(file *File, entity interface{}) []*Import {
	switch entity := entity.(type) {
	case *SimpleSpec:
		return f.fetchSimpleSpec(file, entity)
	case *ArraySpec:
		return f.fetchArraySpec(file, entity)
	case *MapSpec:
		return f.fetchMapSpec(file, entity)
	case *Field:
		return f.fetchField(file, entity)
	case *FuncSpec:
		return f.fetchFuncSpec(file, entity)
	case *InterfaceSpec:
		return f.fetchInterfaceSpec(file, entity)
	case *StructSpec:
		return f.fetchStructSpec(file, entity)
	case *Const:
		return f.fetchConst(file, entity)
	case *ConstGroup:
		return f.fetchConstGroup(file, entity)
	case *Var:
		return f.fetchVar(file, entity)
	case *VarGroup:
		return f.fetchVarGroup(file, entity)
	case *Type:
		return f.fetchType(file, entity)
	case *TypeGroup:
		return f.fetchTypeGroup(file, entity)
	case *Func:
		return f.fetchFunc(file, entity)
	default:
		panic(errors.Errorf("Can't fetch entity with type: '%T'", entity))
	}
}

func (f *EntityImportFetcher) fetchSimpleSpec(file *File, entity *SimpleSpec) []*Import {
	if entity.PackageName == "" {
		return []*Import{}
	}

	for _, group := range file.ImportGroups {
		for _, element := range group.Imports {
			if element.RealAlias() == entity.PackageName {
				return []*Import{element}
			}
		}
	}

	return []*Import{}
}

func (f *EntityImportFetcher) fetchArraySpec(file *File, entity *ArraySpec) []*Import {
	result := f.Fetch(file, entity.Value)
	result = append(result, f.fetchFromContent(file, entity.Length)...)

	return f.importUniquer.Unique(result)
}

func (f *EntityImportFetcher) fetchMapSpec(file *File, entity *MapSpec) []*Import {
	result := []*Import{}
	result = append(result, f.Fetch(file, entity.Key)...)
	result = append(result, f.Fetch(file, entity.Value)...)

	return f.importUniquer.Unique(result)
}

func (f *EntityImportFetcher) fetchField(file *File, entity *Field) []*Import {
	return f.Fetch(file, entity.Spec)
}

func (f *EntityImportFetcher) fetchFuncSpec(file *File, entity *FuncSpec) []*Import {
	result := []*Import{}

	for _, field := range entity.Params {
		result = append(result, f.Fetch(file, field)...)
	}

	for _, field := range entity.Results {
		result = append(result, f.Fetch(file, field)...)
	}

	return f.importUniquer.Unique(result)
}

func (f *EntityImportFetcher) fetchInterfaceSpec(file *File, entity *InterfaceSpec) []*Import {
	result := []*Import{}

	for _, field := range entity.Fields {
		result = append(result, f.Fetch(file, field)...)
	}

	return f.importUniquer.Unique(result)
}

func (f *EntityImportFetcher) fetchStructSpec(file *File, entity *StructSpec) []*Import {
	result := []*Import{}

	for _, field := range entity.Fields {
		result = append(result, f.Fetch(file, field)...)
	}

	return f.importUniquer.Unique(result)
}

func (f *EntityImportFetcher) fetchConst(file *File, entity *Const) []*Import {
	result := []*Import{}

	if entity.Spec != nil {
		result = append(result, f.Fetch(file, entity.Spec)...)
	}

	result = append(result, f.fetchFromContent(file, entity.Value)...)

	return f.importUniquer.Unique(result)
}

func (f *EntityImportFetcher) fetchConstGroup(file *File, entity *ConstGroup) []*Import {
	result := []*Import{}

	for _, field := range entity.Consts {
		result = append(result, f.Fetch(file, field)...)
	}

	return f.importUniquer.Unique(result)
}

func (f *EntityImportFetcher) fetchVar(file *File, entity *Var) []*Import {
	result := []*Import{}

	if entity.Spec != nil {
		result = append(result, f.Fetch(file, entity.Spec)...)
	}

	result = append(result, f.fetchFromContent(file, entity.Value)...)

	return f.importUniquer.Unique(result)
}

func (f *EntityImportFetcher) fetchVarGroup(file *File, m *VarGroup) []*Import {
	result := []*Import{}

	for _, field := range m.Vars {
		result = append(result, f.Fetch(file, field)...)
	}

	return f.importUniquer.Unique(result)
}

func (f *EntityImportFetcher) fetchType(file *File, entity *Type) []*Import {
	return f.Fetch(file, entity.Spec)
}

func (f *EntityImportFetcher) fetchTypeGroup(file *File, entity *TypeGroup) []*Import {
	result := []*Import{}

	for _, field := range entity.Types {
		result = append(result, f.Fetch(file, field)...)
	}

	return f.importUniquer.Unique(result)
}

func (f *EntityImportFetcher) fetchFunc(file *File, entity *Func) []*Import {
	result := []*Import{}

	if entity.Spec != nil {
		result = append(result, f.Fetch(file, entity.Spec)...)
	}

	if entity.Related != nil {
		result = append(result, f.Fetch(file, entity.Related)...)
	}

	result = append(result, f.fetchFromContent(file, entity.Content)...)

	return f.importUniquer.Unique(result)
}

func (f *EntityImportFetcher) fetchFromContent(file *File, content string) []*Import {
	if content == "" {
		return nil
	}

	aliases := ""

	for _, importGroup := range file.ImportGroups {
		for _, element := range importGroup.Imports {
			if aliases != "" {
				aliases += "|"
			}

			aliases += element.RealAlias()
		}
	}

	if aliases == "" {
		return nil
	}

	result := []*Import{}

	matches := regexp.
		MustCompile("([ \\t\\n&;,!~^=+\\-*/()\\[\\]{}]|^)("+aliases+")([ \\t]*\\.)").
		FindAllStringSubmatch(content, -1)

	for _, match := range matches {
		alias := match[2]

		for _, importGroup := range file.ImportGroups {
			for _, element := range importGroup.Imports {
				if alias == element.RealAlias() {
					result = append(result, element)
				}
			}
		}
	}

	return result
}
