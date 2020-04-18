package annotation

// ArraySpec represents specification of array or slice type.
type ArraySpec struct {
	// Allowed types: *SimpleSpec, *ArraySpec, *MapSpec, *StructSpec, *InterfaceSpec, *FuncSpec.
	Value interface{}
	// Expression with int result, which could be calculated at compilation time, or "...".
	Length string
}
