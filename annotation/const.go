package annotation

// Const represents const declaration.
type Const struct {
	Name string
	// Expression with, which could be calculated in compilation time, or empty.
	Value       string
	Comment     string
	Annotations []interface{}
	Spec        *SimpleSpec
}
