package annotation

import (
	"go/parser"
	"go/token"
	"path/filepath"
	"regexp"

	"github.com/pkg/errors"
)

var identRegexp = regexp.MustCompile(`^[\p{L}_][\p{L}\d_]*$`)

type EntityValidator struct {
}

func NewEntityValidator() *EntityValidator {
	return &EntityValidator{}
}

func (v *EntityValidator) Validate(entity interface{}) error {
	switch entity := entity.(type) {
	case *SimpleSpec:
		return v.validateSimpleSpec(entity)
	case *ArraySpec:
		return v.validateArraySpec(entity)
	case *MapSpec:
		return v.validateMapSpec(entity)
	case *Field:
		return v.validateField(entity)
	case *FuncSpec:
		return v.validateFuncSpec(entity)
	case *InterfaceSpec:
		return v.validateInterfaceSpec(entity)
	case *StructSpec:
		return v.validateStructSpec(entity)
	case *Import:
		return v.validateImport(entity)
	case *ImportGroup:
		return v.validateImportGroup(entity)
	case *Const:
		return v.validateConst(entity)
	case *ConstGroup:
		return v.validateConstGroup(entity)
	case *Var:
		return v.validateVar(entity)
	case *VarGroup:
		return v.validateVarGroup(entity)
	case *Type:
		return v.validateType(entity)
	case *TypeGroup:
		return v.validateTypeGroup(entity)
	case *Func:
		return v.validateFunc(entity)
	case *File:
		return v.validateFile(entity)
	case *Namespace:
		return v.validateNamespace(entity)
	case *Storage:
		return v.validateStorage(entity)
	default:
		panic(errors.Errorf("Can't validate entity with type: '%T'", entity))
	}
}

func (v *EntityValidator) validateSimpleSpec(entity *SimpleSpec) error {
	if entity.TypeName == "" {
		return errors.New("Variable 'TypeName' must be not empty")
	}

	if !identRegexp.MatchString(entity.TypeName) {
		return errors.Errorf("Variable 'TypeName' must be valid identifier, actual value: '%s'", entity.TypeName)
	}

	if entity.PackageName != "" && !identRegexp.MatchString(entity.PackageName) {
		return errors.Errorf("Variable 'PackageName' must be valid identifier, actual value: '%s'", entity.PackageName)
	}

	return nil
}

func (v *EntityValidator) validateArraySpec(entity *ArraySpec) error {
	if entity.Value == nil {
		return errors.New("Variable 'Value' must be not nil")
	}

	switch entity.Value.(type) {
	case *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec:
		if err := v.Validate(entity.Value); err != nil {
			return err
		}
	default:
		return errors.Errorf("Variable 'Value' has invalid type: '%T'", entity.Value)
	}

	if entity.Length != "" && entity.Length != "..." {
		if _, err := parser.ParseExpr(entity.Length); err != nil {
			return err
		}
	}

	return nil
}

func (v *EntityValidator) validateMapSpec(entity *MapSpec) error {
	if entity.Key == nil {
		return errors.New("Variable 'Key' must be not nil")
	}

	if entity.Value == nil {
		return errors.New("Variable 'Value' must be not nil")
	}

	switch entity.Key.(type) {
	case *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec:
		if err := v.Validate(entity.Key); err != nil {
			return err
		}
	default:
		return errors.Errorf("Variable 'Key' has invalid type: '%T'", entity.Key)
	}

	switch entity.Value.(type) {
	case *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec:
		if err := v.Validate(entity.Value); err != nil {
			return err
		}
	default:
		return errors.Errorf("Variable 'Value' has invalid type: '%T'", entity.Value)
	}

	return nil
}

func (v *EntityValidator) validateField(entity *Field) error {
	if entity.Name != "" && !identRegexp.MatchString(entity.Name) {
		return errors.Errorf("Variable 'Name' must be valid identifier, actual value: '%s'", entity.Name)
	}

	if entity.Spec == nil {
		return errors.New("Variable 'Spec' must be not nil")
	}

	switch entity.Spec.(type) {
	case *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec:
		if err := v.Validate(entity.Spec); err != nil {
			return err
		}
	default:
		return errors.Errorf("Variable 'Spec' has invalid type: '%T'", entity.Spec)
	}

	return nil
}

func (v *EntityValidator) validateFuncSpec(entity *FuncSpec) error {
	for i, param := range entity.Params {
		if param == nil {
			return errors.Errorf("Variable 'Params[%d]' must be not nil", i)
		}

		if err := v.Validate(param); err != nil {
			return err
		}
	}

	if entity.IsVariadic {
		if len(entity.Params) == 0 {
			return errors.Errorf("Variable 'Params' must be not empty for variadic %T", entity)
		}

		if _, ok := entity.Params[len(entity.Params)-1].Spec.(*ArraySpec); !ok {
			return errors.Errorf(
				"Variable 'Params[%d].Spec' has invalid type for variadic '%T'",
				len(entity.Params)-1,
				entity,
			)
		}
	}

	hasName := false

	for i, result := range entity.Results {
		if result == nil {
			return errors.Errorf("Variable 'Results[%d]' must be not nil", i)
		}

		if i == 0 {
			hasName = result.Name != ""
		} else if hasName != (result.Name != "") {
			return errors.New("Variable 'Results' must have all fields with names or all without names")
		}

		if err := v.Validate(result); err != nil {
			return err
		}
	}

	return nil
}

func (v *EntityValidator) validateInterfaceSpec(entity *InterfaceSpec) error {
	for i, field := range entity.Fields {
		if field == nil {
			return errors.Errorf("Variable 'Fields[%d]' must be not nil", i)
		}

		if err := v.Validate(field); err != nil {
			return err
		}

		switch field.Spec.(type) {
		case *SimpleSpec:
			if field.Name != "" {
				return errors.Errorf(
					"Variable 'Fields[%d].Name' must be empty for 'Fields[%d].Spec' type *SimpleSpec",
					i,
					i,
				)
			}

			if field.Spec.(*SimpleSpec).IsPointer {
				return errors.Errorf("Variable 'Fields[%d].Spec.(%T).IsPointer' must be 'false'", i, field.Spec)
			}
		case *FuncSpec:
			if field.Name == "" {
				return errors.Errorf(
					"Variable 'Fields[%d].Name' must be not empty for 'Fields[%d].Spec' type *FuncSpec",
					i,
					i,
				)
			}
		default:
			return errors.Errorf("Variable 'Fields[%d]' has invalid type '%T'", i, field.Spec)
		}
	}

	return nil
}

func (v *EntityValidator) validateStructSpec(entity *StructSpec) error {
	for i, field := range entity.Fields {
		if field == nil {
			return errors.Errorf("Variable 'Fields[%d]' must be not nil", i)
		}

		if err := v.Validate(field); err != nil {
			return err
		}

		if _, ok := field.Spec.(*SimpleSpec); field.Name == "" && !ok {
			return errors.Errorf("Variable 'Fields[%d]' with empty 'Name' has invalid type: '%T'", i, field.Spec)
		}
	}

	return nil
}

func (v *EntityValidator) validateImport(entity *Import) error {
	if entity.Alias != "" && !identRegexp.MatchString(entity.Alias) {
		return errors.Errorf("Variable 'Alias' must be valid identifier, actual value: '%s'", entity.Alias)
	}

	if entity.Namespace == "" {
		return errors.New("Variable 'Namespace' must be not empty")
	}

	return nil
}

func (v *EntityValidator) validateImportGroup(entity *ImportGroup) error {
	for i, element := range entity.Imports {
		if element == nil {
			return errors.Errorf("Variable 'Imports[%d]' must be not nil", i)
		}

		if err := v.Validate(element); err != nil {
			return err
		}
	}

	return nil
}

func (v *EntityValidator) validateConst(entity *Const) error {
	if entity.Name == "" {
		return errors.New("Variable 'Name' must be not empty")
	}

	if !identRegexp.MatchString(entity.Name) {
		return errors.Errorf("Variable 'Name' must be valid identifier, actual value: '%s'", entity.Name)
	}

	if entity.Spec != nil {
		if entity.Spec.IsPointer {
			return errors.Errorf("Variable 'Spec.(%T).IsPointer' must be 'false' for %T", entity.Spec, entity)
		}

		if err := v.Validate(entity.Spec); err != nil {
			return err
		}
	}

	if entity.Value != "" {
		if _, err := parser.ParseExpr(entity.Value); err != nil {
			return err
		}
	}

	return nil
}

func (v *EntityValidator) validateConstGroup(entity *ConstGroup) error {
	for i, element := range entity.Consts {
		if element == nil {
			return errors.Errorf("Variable 'Consts[%d]' must be not nil", i)
		}

		if err := v.Validate(element); err != nil {
			return err
		}
	}

	return nil
}

func (v *EntityValidator) validateVar(entity *Var) error {
	if entity.Name == "" {
		return errors.New("Variable 'Name' must be not empty")
	}

	if !identRegexp.MatchString(entity.Name) {
		return errors.Errorf("Variable 'Name' must be valid identifier, actual value: '%s'", entity.Name)
	}

	if entity.Spec == nil && entity.Value == "" {
		return errors.Errorf("%T must have not nil 'Spec' or not empty 'Value'", entity)
	}

	if entity.Spec != nil {
		switch entity.Spec.(type) {
		case *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec:
			if err := v.Validate(entity.Spec); err != nil {
				return err
			}
		default:
			return errors.Errorf("Variable 'Spec' has invalid type: '%T'", entity.Spec)
		}
	}

	if entity.Value != "" {
		if _, err := parser.ParseExpr(entity.Value); err != nil {
			return err
		}
	}

	return nil
}

func (v *EntityValidator) validateVarGroup(entity *VarGroup) error {
	for i, element := range entity.Vars {
		if element == nil {
			return errors.Errorf("Variable 'Vars[%d]' must be not nil", i)
		}

		if err := v.Validate(element); err != nil {
			return err
		}
	}

	return nil
}

func (v *EntityValidator) validateType(entity *Type) error {
	if entity.Name == "" {
		return errors.New("Variable 'Name' must be not empty")
	}

	if !identRegexp.MatchString(entity.Name) {
		return errors.Errorf("Variable 'Name' must be valid identifier, actual value: '%s'", entity.Name)
	}

	if entity.Spec == nil {
		return errors.New("Variable 'Spec' must be not nil")
	}

	switch entity.Spec.(type) {
	case *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec:
		if err := v.Validate(entity.Spec); err != nil {
			return err
		}
	default:
		return errors.Errorf("Variable 'Spec' has invalid type: '%T'", entity.Spec)
	}

	return nil
}

func (v *EntityValidator) validateTypeGroup(entity *TypeGroup) error {
	for i, element := range entity.Types {
		if element == nil {
			return errors.Errorf("Variable 'Types[%d]' must be not nil", i)
		}

		if err := v.Validate(element); err != nil {
			return err
		}
	}

	return nil
}

func (v *EntityValidator) validateFunc(entity *Func) error {
	if entity.Name == "" {
		return errors.New("Variable 'Name' must be not empty")
	}

	if !identRegexp.MatchString(entity.Name) {
		return errors.Errorf("Variable 'Name' must be valid identifier, actual value: '%s'", entity.Name)
	}

	if entity.Spec != nil {
		if err := v.Validate(entity.Spec); err != nil {
			return err
		}
	}

	if entity.Related != nil {
		if err := v.Validate(entity.Related); err != nil {
			return err
		}

		if related, ok := entity.Related.Spec.(*SimpleSpec); !ok {
			return errors.Errorf("Variable 'Related.Spec.(%T)' has invalid type for '%T'", entity.Related.Spec, entity)
		} else if related.PackageName != "" {
			return errors.Errorf("Variable 'Related.Spec.(%T).PackageName' must be empty for '%T'", related, entity)
		}
	}

	if entity.Content != "" {
		content := "func(){\n" + entity.Content + "\n}"

		if _, err := parser.ParseExpr(content); err != nil {
			return err
		}
	}

	return nil
}

func (v *EntityValidator) validateFile(entity *File) error {
	if entity.Name == "" {
		return errors.New("Variable 'Name' must be not empty")
	}

	if entity.PackageName == "" {
		return errors.New("Variable 'PackageName' must be not empty")
	}

	if !identRegexp.MatchString(entity.PackageName) {
		return errors.Errorf("Variable 'PackageName' must be valid identifier, actual value: '%s'", entity.PackageName)
	}

	for i, element := range entity.ImportGroups {
		if element == nil {
			return errors.Errorf("Variable 'ImportGroups[%d]' must be not nil", i)
		}

		if err := v.Validate(element); err != nil {
			return err
		}
	}

	for i, element := range entity.ConstGroups {
		if element == nil {
			return errors.Errorf("Variable 'ConstGroups[%d]' must be not nil", i)
		}

		if err := v.Validate(element); err != nil {
			return err
		}
	}

	for i, element := range entity.VarGroups {
		if element == nil {
			return errors.Errorf("Variable 'VarGroups[%d]' must be not nil", i)
		}

		if err := v.Validate(element); err != nil {
			return err
		}
	}

	for i, element := range entity.TypeGroups {
		if element == nil {
			return errors.Errorf("Variable 'TypeGroups[%d]' must be not nil", i)
		}

		if err := v.Validate(element); err != nil {
			return err
		}
	}

	for i, element := range entity.Funcs {
		if element == nil {
			return errors.Errorf("Variable 'Funcs[%d]' must be not nil", i)
		}

		if err := v.Validate(element); err != nil {
			return err
		}
	}

	if entity.Content != "" {
		if _, err := parser.ParseFile(token.NewFileSet(), "", entity.Content, 0); err != nil {
			return err
		}
	}

	return nil
}

func (v *EntityValidator) validateNamespace(entity *Namespace) error {
	if entity.Name == "" {
		return errors.New("Variable 'Name' must be not empty")
	}

	if entity.Path == "" {
		return errors.New("Variable 'Path' must be not empty")
	}

	if !filepath.IsAbs(entity.Path) {
		return errors.Errorf("Variable 'Path' must be absolute path, actual value: '%s'", entity.Path)
	}

	fileNames := map[string]bool{}
	packageName := ""

	for i, element := range entity.Files {
		if element == nil {
			return errors.Errorf("Variable 'Files[%d]' must be not nil", i)
		}

		if err := v.Validate(element); err != nil {
			return err
		}

		if _, ok := fileNames[element.Name]; ok {
			return errors.Errorf("Namespace has duplicate file name: %s", element.Name)
		} else {
			fileNames[element.Name] = true
		}

		if i == 0 {
			packageName = element.PackageName
		} else if element.PackageName != packageName {
			return errors.New("Namespace has different packages")
		}
	}

	return nil
}

func (v *EntityValidator) validateStorage(entity *Storage) error {
	namespaceNames := map[string]bool{}
	namespacePaths := map[string]bool{}

	for i, element := range entity.Namespaces {
		if element == nil {
			return errors.Errorf("Variable 'Namespaces[%d]' must be not nil", i)
		}

		if err := v.Validate(element); err != nil {
			return err
		}

		if _, ok := namespaceNames[element.Name]; ok {
			return errors.Errorf("Storage has duplicate namespace 'Name': '%s'", element.Name)
		} else {
			namespaceNames[element.Name] = true
		}

		if _, ok := namespacePaths[element.Path]; ok {
			return errors.Errorf("Storage has duplicate namespace 'Path': '%s'", element.Path)
		} else {
			namespacePaths[element.Path] = true
		}
	}

	return nil
}
