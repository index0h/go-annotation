package annotation

import (
	"encoding/json"
	"reflect"

	"github.com/pkg/errors"
)

type EntityCloner struct {
}

func NewEntityCloner() *EntityCloner {
	return &EntityCloner{}
}

// Creates deep copy of entity
func (c *EntityCloner) Clone(entity interface{}) interface{} {
	switch entity := entity.(type) {
	case *SimpleSpec:
		return c.cloneSimpleSpec(entity)
	case *ArraySpec:
		return c.cloneArraySpec(entity)
	case *MapSpec:
		return c.cloneMapSpec(entity)
	case *Field:
		return c.cloneField(entity)
	case *FuncSpec:
		return c.cloneFuncSpec(entity)
	case *InterfaceSpec:
		return c.cloneInterfaceSpec(entity)
	case *StructSpec:
		return c.cloneStructSpec(entity)
	case *Import:
		return c.cloneImport(entity)
	case *ImportGroup:
		return c.cloneImportGroup(entity)
	case *Const:
		return c.cloneConst(entity)
	case *ConstGroup:
		return c.cloneConstGroup(entity)
	case *Var:
		return c.cloneVar(entity)
	case *VarGroup:
		return c.cloneVarGroup(entity)
	case *Type:
		return c.cloneType(entity)
	case *TypeGroup:
		return c.cloneTypeGroup(entity)
	case *Func:
		return c.cloneFunc(entity)
	case *File:
		return c.cloneFile(entity)
	case *Namespace:
		return c.cloneNamespace(entity)
	case *Storage:
		return c.cloneStorage(entity)
	default:
		panic(errors.Errorf("Can't clone entity with type: '%T'", entity))
	}
}

func (c *EntityCloner) cloneSimpleSpec(entity *SimpleSpec) interface{} {
	return &SimpleSpec{
		PackageName: entity.PackageName,
		TypeName:    entity.TypeName,
		IsPointer:   entity.IsPointer,
	}
}

func (c *EntityCloner) cloneArraySpec(entity *ArraySpec) interface{} {
	return &ArraySpec{
		Value:  c.Clone(entity.Value),
		Length: entity.Length,
	}
}

func (c *EntityCloner) cloneMapSpec(entity *MapSpec) interface{} {
	return &MapSpec{
		Value: c.Clone(entity.Value),
		Key:   c.Clone(entity.Key),
	}
}

func (c *EntityCloner) cloneField(entity *Field) interface{} {
	return &Field{
		Name:        entity.Name,
		Tag:         entity.Tag,
		Comment:     entity.Comment,
		Annotations: c.cloneAnnotations(entity.Annotations),
		Spec:        c.Clone(entity.Spec),
	}
}

func (c *EntityCloner) cloneFuncSpec(entity *FuncSpec) interface{} {
	if entity.Params == nil && entity.Results == nil {
		return &FuncSpec{}
	}

	result := &FuncSpec{
		IsVariadic: entity.IsVariadic,
	}

	if entity.Params != nil {
		result.Params = make([]*Field, len(entity.Params))
	}

	if entity.Results != nil {
		result.Results = make([]*Field, len(entity.Results))
	}

	for i, field := range entity.Params {
		result.Params[i] = c.Clone(field).(*Field)
	}

	for i, field := range entity.Results {
		result.Results[i] = c.Clone(field).(*Field)
	}

	return result
}

func (c *EntityCloner) cloneInterfaceSpec(entity *InterfaceSpec) interface{} {
	if entity.Fields == nil {
		return &InterfaceSpec{}
	}

	result := &InterfaceSpec{}

	if entity.Fields != nil {
		result.Fields = make([]*Field, len(entity.Fields))
	}

	for i, method := range entity.Fields {
		result.Fields[i] = c.Clone(method).(*Field)
	}

	return result
}

func (c *EntityCloner) cloneStructSpec(entity *StructSpec) interface{} {
	if entity.Fields == nil {
		return &StructSpec{}
	}

	result := &StructSpec{}

	if entity.Fields != nil {
		result.Fields = make([]*Field, len(entity.Fields))
	}

	for i, field := range entity.Fields {
		result.Fields[i] = c.Clone(field).(*Field)
	}

	return result
}

func (c *EntityCloner) cloneImport(entity *Import) interface{} {
	return &Import{
		Alias:       entity.Alias,
		Namespace:   entity.Namespace,
		Comment:     entity.Comment,
		Annotations: c.cloneAnnotations(entity.Annotations),
	}
}

func (c *EntityCloner) cloneImportGroup(entity *ImportGroup) interface{} {
	result := &ImportGroup{
		Comment:     entity.Comment,
		Annotations: c.cloneAnnotations(entity.Annotations),
	}

	if entity.Imports != nil {
		result.Imports = make([]*Import, len(entity.Imports))
	}

	for i, element := range entity.Imports {
		result.Imports[i] = c.Clone(element).(*Import)
	}

	return result
}

func (c *EntityCloner) cloneConst(entity *Const) interface{} {
	result := &Const{
		Name:        entity.Name,
		Value:       entity.Value,
		Comment:     entity.Comment,
		Annotations: c.cloneAnnotations(entity.Annotations),
	}

	if entity.Spec != nil {
		result.Spec = c.Clone(entity.Spec).(*SimpleSpec)
	}

	return result
}

func (c *EntityCloner) cloneConstGroup(entity *ConstGroup) interface{} {
	result := &ConstGroup{
		Comment:     entity.Comment,
		Annotations: c.cloneAnnotations(entity.Annotations),
	}

	if entity.Consts != nil {
		result.Consts = make([]*Const, len(entity.Consts))
	}

	for i, element := range entity.Consts {
		result.Consts[i] = c.Clone(element).(*Const)
	}

	return result
}

func (c *EntityCloner) cloneVar(entity *Var) interface{} {
	result := &Var{
		Name:        entity.Name,
		Value:       entity.Value,
		Comment:     entity.Comment,
		Annotations: c.cloneAnnotations(entity.Annotations),
	}

	if entity.Spec != nil {
		result.Spec = c.Clone(entity.Spec).(*SimpleSpec)
	}

	return result
}

func (c *EntityCloner) cloneVarGroup(entity *VarGroup) interface{} {
	result := &VarGroup{
		Comment:     entity.Comment,
		Annotations: c.cloneAnnotations(entity.Annotations),
	}

	if entity.Vars != nil {
		result.Vars = make([]*Var, len(entity.Vars))
	}

	for i, element := range entity.Vars {
		result.Vars[i] = c.Clone(element).(*Var)
	}

	return result
}

func (c *EntityCloner) cloneType(entity *Type) interface{} {
	return &Type{
		Name:        entity.Name,
		Comment:     entity.Comment,
		Annotations: c.cloneAnnotations(entity.Annotations),
		Spec:        c.Clone(entity.Spec),
	}
}

func (c *EntityCloner) cloneTypeGroup(entity *TypeGroup) interface{} {
	result := &TypeGroup{
		Comment:     entity.Comment,
		Annotations: c.cloneAnnotations(entity.Annotations),
	}

	if entity.Types != nil {
		result.Types = make([]*Type, len(entity.Types))
	}

	for i, element := range entity.Types {
		result.Types[i] = c.Clone(element).(*Type)
	}

	return result
}

func (c *EntityCloner) cloneFunc(entity *Func) interface{} {
	result := &Func{
		Name:        entity.Name,
		Content:     entity.Content,
		Comment:     entity.Comment,
		Annotations: c.cloneAnnotations(entity.Annotations),
	}

	if entity.Spec != nil {
		result.Spec = c.Clone(entity.Spec).(*FuncSpec)
	}

	if entity.Related != nil {
		result.Related = c.Clone(entity.Related).(*Field)
	}

	return result
}

func (c *EntityCloner) cloneFile(entity *File) interface{} {
	result := &File{
		Name:        entity.Name,
		Content:     entity.Content,
		PackageName: entity.PackageName,
		Comment:     entity.Comment,
		Annotations: c.cloneAnnotations(entity.Annotations),
	}

	if entity.ImportGroups != nil {
		result.ImportGroups = make([]*ImportGroup, len(entity.ImportGroups))
	}

	if entity.ConstGroups != nil {
		result.ConstGroups = make([]*ConstGroup, len(entity.ConstGroups))
	}

	if entity.VarGroups != nil {
		result.VarGroups = make([]*VarGroup, len(entity.VarGroups))
	}

	if entity.TypeGroups != nil {
		result.TypeGroups = make([]*TypeGroup, len(entity.TypeGroups))
	}

	if entity.Funcs != nil {
		result.Funcs = make([]*Func, len(entity.Funcs))
	}

	for i, element := range entity.ImportGroups {
		result.ImportGroups[i] = c.Clone(element).(*ImportGroup)
	}

	for i, element := range entity.ConstGroups {
		result.ConstGroups[i] = c.Clone(element).(*ConstGroup)
	}

	for i, element := range entity.VarGroups {
		result.VarGroups[i] = c.Clone(element).(*VarGroup)
	}

	for i, element := range entity.TypeGroups {
		result.TypeGroups[i] = c.Clone(element).(*TypeGroup)
	}

	for i, element := range entity.Funcs {
		result.Funcs[i] = c.Clone(element).(*Func)
	}

	return result
}

func (c *EntityCloner) cloneNamespace(entity *Namespace) interface{} {
	result := &Namespace{
		Name:      entity.Name,
		Path:      entity.Path,
		IsIgnored: entity.IsIgnored,
	}

	if entity.Files != nil {
		result.Files = make([]*File, len(entity.Files))
	}

	for i, element := range entity.Files {
		result.Files[i] = c.Clone(element).(*File)
	}

	return result
}

func (c *EntityCloner) cloneStorage(entity *Storage) interface{} {
	result := &Storage{}

	if entity.Namespaces != nil {
		result.Namespaces = make([]*Namespace, len(entity.Namespaces))
	}

	for i, element := range entity.Namespaces {
		result.Namespaces[i] = c.Clone(element).(*Namespace)
	}

	return result
}

func (c *EntityCloner) cloneAnnotations(annotations []interface{}) []interface{} {
	if annotations == nil {
		return nil
	}

	result := make([]interface{}, len(annotations))

	for i, annotation := range annotations {
		data, err := json.Marshal(annotation)

		if err != nil {
			panic(err)
		}

		annotationCopy := reflect.New(reflect.TypeOf(annotation)).Interface()

		if len(data) > 0 {
			if err = json.Unmarshal(data, &annotationCopy); err != nil {
				panic(err)
			}
		}

		result[i] = reflect.ValueOf(annotationCopy).Elem().Interface()
	}

	return result
}
