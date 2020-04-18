package annotation

import "github.com/pkg/errors"

type EntityContainsChecker struct {
	equaler Equaler
}

func NewEntityContainsChecker(equaler Equaler) *EntityContainsChecker {
	if equaler == nil {
		panic(errors.New("Variable 'equaler' must be not nil"))
	}

	return &EntityContainsChecker{
		equaler: equaler,
	}
}

func (c *EntityContainsChecker) Contains(collection interface{}, entity interface{}) bool {
	switch collection := collection.(type) {
	case *ImportGroup:
		if entityValue, ok := entity.(*Import); ok {
			return c.containsImport(collection, entityValue)
		}
	case *ConstGroup:
		if entityValue, ok := entity.(*Const); ok {
			return c.containsConst(collection, entityValue)
		}
	case *VarGroup:
		if entityValue, ok := entity.(*Var); ok {
			return c.containsVar(collection, entityValue)
		}
	case *TypeGroup:
		if entityValue, ok := entity.(*Type); ok {
			return c.containsType(collection, entityValue)
		}
	case *File:
		return c.containsFile(collection, entity)
	}

	return false
}

func (c *EntityContainsChecker) containsImport(collection *ImportGroup, entity *Import) bool {
	for _, element := range collection.Imports {
		if c.equaler.Equal(element, entity) {
			return true
		}
	}

	return false
}

func (c *EntityContainsChecker) containsConst(collection *ConstGroup, entity *Const) bool {
	for _, element := range collection.Consts {
		if c.equaler.Equal(element, entity) {
			return true
		}
	}

	return false
}

func (c *EntityContainsChecker) containsVar(collection *VarGroup, entity *Var) bool {
	for _, element := range collection.Vars {
		if c.equaler.Equal(element, entity) {
			return true
		}
	}

	return false
}

func (c *EntityContainsChecker) containsType(collection *TypeGroup, entity *Type) bool {
	for _, element := range collection.Types {
		if c.equaler.Equal(element, entity) {
			return true
		}
	}

	return false
}

func (c *EntityContainsChecker) containsFile(file *File, entity interface{}) bool {
	switch entity := entity.(type) {
	case *Import:
		for _, group := range file.ImportGroups {
			if c.containsImport(group, entity) {
				return true
			}
		}
	case *Const:
		for _, group := range file.ConstGroups {
			if c.containsConst(group, entity) {
				return true
			}
		}
	case *Var:
		for _, group := range file.VarGroups {
			if c.containsVar(group, entity) {
				return true
			}
		}
	case *Type:
		for _, group := range file.TypeGroups {
			if c.containsType(group, entity) {
				return true
			}
		}
	case *Func:
		for _, element := range file.Funcs {
			if c.equaler.Equal(element, entity) {
				return true
			}
		}
	}

	return false
}
