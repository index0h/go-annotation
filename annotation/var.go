package annotation

// Var represents var declaration.
type Var struct {
	Name        string
	Value       string
	Comment     string
	Annotations []interface{}
	// Allowed types: *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec.
	Spec interface{}
}
