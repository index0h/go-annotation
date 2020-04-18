package annotation

// Func represents declaration of function.
type Func struct {
	Name        string
	Content     string
	Comment     string
	Annotations []interface{}
	Spec        *FuncSpec
	Related     *Field
}
