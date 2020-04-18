package annotation

type EntityEqualer struct {
}

func NewEntityEqualer() *EntityEqualer {
	return &EntityEqualer{}
}

func (c *EntityEqualer) Equal(x interface{}, y interface{}) bool {
	switch x := x.(type) {
	case *SimpleSpec:
		if yValue, ok := y.(*SimpleSpec); ok {
			return c.equalSimpleSpec(x, yValue)
		}
	case *ArraySpec:
		if yValue, ok := y.(*ArraySpec); ok {
			return c.equalArraySpec(x, yValue)
		}
	case *MapSpec:
		if yValue, ok := y.(*MapSpec); ok {
			return c.equalMapSpec(x, yValue)
		}
	case *Field:
		if yValue, ok := y.(*Field); ok {
			return c.equalField(x, yValue)
		}
	case *FuncSpec:
		if yValue, ok := y.(*FuncSpec); ok {
			return c.equalFuncSpec(x, yValue)
		}
	case *InterfaceSpec:
		if yValue, ok := y.(*InterfaceSpec); ok {
			return c.equalInterfaceSpec(x, yValue)
		}
	case *StructSpec:
		if yValue, ok := y.(*StructSpec); ok {
			return c.equalStructSpec(x, yValue)
		}
	case *Import:
		if yValue, ok := y.(*Import); ok {
			return c.equalImport(x, yValue)
		}
	case *ImportGroup:
		if yValue, ok := y.(*ImportGroup); ok {
			return c.equalImportGroup(x, yValue)
		}
	case *Const:
		if yValue, ok := y.(*Const); ok {
			return c.equalConst(x, yValue)
		}
	case *ConstGroup:
		if yValue, ok := y.(*ConstGroup); ok {
			return c.equalConstGroup(x, yValue)
		}
	case *Var:
		if yValue, ok := y.(*Var); ok {
			return c.equalVar(x, yValue)
		}
	case *VarGroup:
		if yValue, ok := y.(*VarGroup); ok {
			return c.equalVarGroup(x, yValue)
		}
	case *Type:
		if yValue, ok := y.(*Type); ok {
			return c.equalType(x, yValue)
		}
	case *TypeGroup:
		if yValue, ok := y.(*TypeGroup); ok {
			return c.equalTypeGroup(x, yValue)
		}
	case *Func:
		if yValue, ok := y.(*Func); ok {
			return c.equalFunc(x, yValue)
		}
	}

	return false
}

func (c *EntityEqualer) equalSimpleSpec(x *SimpleSpec, y *SimpleSpec) bool {
	return y.PackageName == x.PackageName && y.TypeName == x.TypeName && y.IsPointer == x.IsPointer
}

func (c *EntityEqualer) equalArraySpec(x *ArraySpec, y *ArraySpec) bool {
	return y.Length == x.Length && c.Equal(x.Value, y.Value)
}

func (c *EntityEqualer) equalMapSpec(x *MapSpec, y *MapSpec) bool {
	return c.Equal(x.Key, y.Key) && c.Equal(x.Value, y.Value)
}

func (c *EntityEqualer) equalField(x *Field, y *Field) bool {
	return x.Name == y.Name && x.Tag == y.Tag && c.Equal(x.Spec, y.Spec)
}

func (c *EntityEqualer) equalFuncSpec(x *FuncSpec, y *FuncSpec) bool {
	if x.IsVariadic != y.IsVariadic || len(x.Params) != len(y.Params) || len(x.Results) != len(y.Results) {
		return false
	}

	for i, field := range x.Params {
		if !c.Equal(field, y.Params[i]) {
			return false
		}
	}

	for i, field := range x.Results {
		if !c.Equal(field, y.Results[i]) {
			return false
		}
	}

	return true
}

func (c *EntityEqualer) equalInterfaceSpec(x *InterfaceSpec, y *InterfaceSpec) bool {
	if len(x.Fields) != len(y.Fields) {
		return false
	}

	checkedYFields := make([]bool, len(x.Fields))

	for _, field := range x.Fields {
		fieldEqual := false

		for j, yField := range y.Fields {
			if checkedYFields[j] {
				continue
			}

			if c.Equal(field, yField) {
				fieldEqual = true
				checkedYFields[j] = true

				break
			}
		}

		if !fieldEqual {
			return false
		}
	}

	return true
}

func (c *EntityEqualer) equalStructSpec(x *StructSpec, y *StructSpec) bool {
	if len(x.Fields) != len(y.Fields) {
		return false
	}

	checkedYFields := make([]bool, len(x.Fields))

	for _, field := range x.Fields {
		fieldEqual := false

		for j, yField := range y.Fields {
			if checkedYFields[j] {
				continue
			}

			if c.Equal(field, yField) {
				fieldEqual = true
				checkedYFields[j] = true

				break
			}
		}

		if !fieldEqual {
			return false
		}
	}

	return true
}

func (c *EntityEqualer) equalImport(x *Import, y *Import) bool {
	return y.RealAlias() == x.RealAlias() && y.Namespace == x.Namespace
}

func (c *EntityEqualer) equalImportGroup(x *ImportGroup, y *ImportGroup) bool {
	if len(x.Imports) != len(y.Imports) {
		return false
	}

	checkedYElements := make([]bool, len(x.Imports))

	for _, element := range x.Imports {
		elementEqual := false

		for j, yElement := range y.Imports {
			if checkedYElements[j] {
				continue
			}

			if c.Equal(element, yElement) {
				elementEqual = true
				checkedYElements[j] = true

				break
			}
		}

		if !elementEqual {
			return false
		}
	}

	return true
}

func (c *EntityEqualer) equalConst(x *Const, y *Const) bool {
	if y.Name != x.Name || y.Value != x.Value || ((x.Spec == nil) != (y.Spec == nil)) {
		return false
	}

	if x.Spec != nil && !c.Equal(x.Spec, y.Spec) {
		return false
	}

	return true
}

func (c *EntityEqualer) equalConstGroup(x *ConstGroup, y *ConstGroup) bool {
	if len(x.Consts) != len(y.Consts) {
		return false
	}

	for i, element := range x.Consts {
		if !c.Equal(element, y.Consts[i]) {
			return false
		}
	}

	return true
}

func (c *EntityEqualer) equalVar(x *Var, y *Var) bool {
	if y.Name != x.Name || y.Value != x.Value || ((x.Spec == nil) != (y.Spec == nil)) {
		return false
	}

	if x.Spec != nil && !c.Equal(x.Spec, y.Spec) {
		return false
	}

	return true
}

func (c *EntityEqualer) equalVarGroup(x *VarGroup, y *VarGroup) bool {
	if len(x.Vars) != len(y.Vars) {
		return false
	}

	checkedYElements := make([]bool, len(x.Vars))

	for _, element := range x.Vars {
		elementEqual := false

		for j, yElement := range y.Vars {
			if checkedYElements[j] {
				continue
			}

			if c.Equal(element, yElement) {
				elementEqual = true
				checkedYElements[j] = true

				break
			}
		}

		if !elementEqual {
			return false
		}
	}

	return true
}

func (c *EntityEqualer) equalType(x *Type, y *Type) bool {
	if y.Name != x.Name {
		return false
	}

	return c.Equal(x.Spec, y.Spec)
}

func (c *EntityEqualer) equalTypeGroup(x *TypeGroup, y *TypeGroup) bool {
	if len(x.Types) != len(y.Types) {
		return false
	}

	checkedYElements := make([]bool, len(x.Types))

	for _, element := range x.Types {
		elementEqual := false

		for j, yElement := range y.Types {
			if checkedYElements[j] {
				continue
			}

			if c.Equal(element, yElement) {
				elementEqual = true
				checkedYElements[j] = true

				break
			}
		}

		if !elementEqual {
			return false
		}
	}

	return true
}

func (c *EntityEqualer) equalFunc(x *Func, y *Func) bool {
	if y.Name != x.Name ||
		y.Content != x.Content ||
		((x.Spec == nil) != (y.Spec == nil)) ||
		((x.Related == nil) != (y.Related == nil)) {
		return false
	}

	if x.Spec != nil && !c.Equal(x.Spec, y.Spec) {
		return false
	}

	if x.Related != nil && !c.Equal(x.Related, y.Related) {
		return false
	}

	return true
}
