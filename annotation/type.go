package annotation

// Type represents type declaration.
type Type struct {
	Name        string
	Comment     string
	Annotations []interface{}
	// Allowed types: *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec.
	Spec interface{}
}
