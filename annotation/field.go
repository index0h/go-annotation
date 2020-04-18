package annotation

// Field represents field of struct, interface, func param, func result or func related field.
type Field struct {
	Name string
	// Used for render *StructSpec
	Tag         string
	Comment     string
	Annotations []interface{}
	// Allowed types: *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec.
	Spec interface{}
}
