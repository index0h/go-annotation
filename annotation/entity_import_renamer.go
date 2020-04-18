package annotation

import (
	"regexp"

	"github.com/pkg/errors"
)

type EntityImportRenamer struct {
}

func NewEntityImportRenamer() *EntityImportRenamer {
	return &EntityImportRenamer{}
}

func (r *EntityImportRenamer) Rename(entity interface{}, oldAlias string, newAlias string) {
	if !identRegexp.MatchString(oldAlias) {
		panic(errors.Errorf("Variable 'oldAlias' must be valid identifier, actual value: '%s'", oldAlias))
	}

	if !identRegexp.MatchString(newAlias) {
		panic(errors.Errorf("Variable 'newAlias' must be valid identifier, actual value: '%s'", newAlias))
	}

	switch entity := entity.(type) {
	case *SimpleSpec:
		r.renameInSimpleSpec(entity, oldAlias, newAlias)
	case *ArraySpec:
		r.renameInArraySpec(entity, oldAlias, newAlias)
	case *MapSpec:
		r.renameInMapSpec(entity, oldAlias, newAlias)
	case *Field:
		r.renameInField(entity, oldAlias, newAlias)
	case *FuncSpec:
		r.renameInFuncSpec(entity, oldAlias, newAlias)
	case *InterfaceSpec:
		r.renameInInterfaceSpec(entity, oldAlias, newAlias)
	case *StructSpec:
		r.renameInStructSpec(entity, oldAlias, newAlias)
	case *Import:
		r.renameInImport(entity, oldAlias, newAlias)
	case *ImportGroup:
		r.renameInImportGroup(entity, oldAlias, newAlias)
	case *Const:
		r.renameInConst(entity, oldAlias, newAlias)
	case *ConstGroup:
		r.renameInConstGroup(entity, oldAlias, newAlias)
	case *Var:
		r.renameInVar(entity, oldAlias, newAlias)
	case *VarGroup:
		r.renameInVarGroup(entity, oldAlias, newAlias)
	case *Type:
		r.renameInType(entity, oldAlias, newAlias)
	case *TypeGroup:
		r.renameInTypeGroup(entity, oldAlias, newAlias)
	case *Func:
		r.renameInFunc(entity, oldAlias, newAlias)
	case *File:
		r.renameInFile(entity, oldAlias, newAlias)
	default:
		panic(errors.Errorf("Can't rename entity with type: '%T'", entity))
	}
}
func (r *EntityImportRenamer) renameInSimpleSpec(entity *SimpleSpec, oldAlias string, newAlias string) {
	if entity.PackageName == oldAlias {
		entity.PackageName = newAlias
	}
}

func (r *EntityImportRenamer) renameInArraySpec(entity *ArraySpec, oldAlias string, newAlias string) {
	r.Rename(entity.Value, oldAlias, newAlias)

	entity.Length = r.renameInContent(entity.Length, oldAlias, newAlias)
}

func (r *EntityImportRenamer) renameInMapSpec(entity *MapSpec, oldAlias string, newAlias string) {
	r.Rename(entity.Key, oldAlias, newAlias)
	r.Rename(entity.Value, oldAlias, newAlias)
}

func (r *EntityImportRenamer) renameInField(entity *Field, oldAlias string, newAlias string) {
	r.Rename(entity.Spec, oldAlias, newAlias)
}

func (r *EntityImportRenamer) renameInFuncSpec(entity *FuncSpec, oldAlias string, newAlias string) {
	for _, field := range entity.Params {
		r.renameInField(field, oldAlias, newAlias)
	}

	for _, field := range entity.Results {
		r.renameInField(field, oldAlias, newAlias)
	}
}

func (r *EntityImportRenamer) renameInInterfaceSpec(entity *InterfaceSpec, oldAlias string, newAlias string) {
	for _, field := range entity.Fields {
		r.renameInField(field, oldAlias, newAlias)
	}
}

func (r *EntityImportRenamer) renameInStructSpec(entity *StructSpec, oldAlias string, newAlias string) {
	for _, field := range entity.Fields {
		r.renameInField(field, oldAlias, newAlias)
	}
}

func (r *EntityImportRenamer) renameInImport(entity *Import, oldAlias string, newAlias string) {
	if entity.RealAlias() == oldAlias {
		entity.Alias = newAlias
	}
}

func (r *EntityImportRenamer) renameInImportGroup(entity *ImportGroup, oldAlias string, newAlias string) {
	for _, element := range entity.Imports {
		r.renameInImport(element, oldAlias, newAlias)
	}
}

func (r *EntityImportRenamer) renameInConst(entity *Const, oldAlias string, newAlias string) {
	if entity.Spec != nil {
		r.Rename(entity.Spec, oldAlias, newAlias)
	}

	entity.Value = r.renameInContent(entity.Value, oldAlias, newAlias)
}

func (r *EntityImportRenamer) renameInConstGroup(entity *ConstGroup, oldAlias string, newAlias string) {
	for _, element := range entity.Consts {
		r.renameInConst(element, oldAlias, newAlias)
	}
}

func (r *EntityImportRenamer) renameInVar(entity *Var, oldAlias string, newAlias string) {
	if entity.Spec != nil {
		r.Rename(entity.Spec, oldAlias, newAlias)
	}

	entity.Value = r.renameInContent(entity.Value, oldAlias, newAlias)
}

func (r *EntityImportRenamer) renameInVarGroup(entity *VarGroup, oldAlias string, newAlias string) {
	for _, element := range entity.Vars {
		r.renameInVar(element, oldAlias, newAlias)
	}
}

func (r *EntityImportRenamer) renameInType(entity *Type, oldAlias string, newAlias string) {
	r.Rename(entity.Spec, oldAlias, newAlias)
}

func (r *EntityImportRenamer) renameInTypeGroup(entity *TypeGroup, oldAlias string, newAlias string) {
	for _, element := range entity.Types {
		r.renameInType(element, oldAlias, newAlias)
	}
}

func (r *EntityImportRenamer) renameInFunc(entity *Func, oldAlias string, newAlias string) {
	if entity.Spec != nil {
		r.Rename(entity.Spec, oldAlias, newAlias)
	}

	if entity.Related != nil {
		r.renameInField(entity.Related, oldAlias, newAlias)
	}

	entity.Content = r.renameInContent(entity.Content, oldAlias, newAlias)
}

func (r *EntityImportRenamer) renameInFile(entity *File, oldAlias string, newAlias string) {
	for _, element := range entity.ImportGroups {
		r.renameInImportGroup(element, oldAlias, newAlias)
	}

	for _, element := range entity.ConstGroups {
		r.renameInConstGroup(element, oldAlias, newAlias)
	}

	for _, element := range entity.VarGroups {
		r.renameInVarGroup(element, oldAlias, newAlias)
	}

	for _, element := range entity.TypeGroups {
		r.renameInTypeGroup(element, oldAlias, newAlias)
	}

	for _, element := range entity.Funcs {
		r.renameInFunc(element, oldAlias, newAlias)
	}

	entity.Content = r.renameInContent(entity.Content, oldAlias, newAlias)
}

func (r *EntityImportRenamer) renameInContent(content string, oldAlias string, newAlias string) string {
	if content == "" {
		return ""
	}

	return regexp.
		MustCompile("([ \\t\\n&;,!~^=+\\-*/()\\[\\]{}]|^)"+oldAlias+"([ \\t]*\\.)").
		ReplaceAllString(content, "${1}"+newAlias+"${2}")
}
