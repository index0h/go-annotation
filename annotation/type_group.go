package annotation

// TypeGroup represents list of type declaration.
type TypeGroup struct {
	Comment     string
	Annotations []interface{}
	Types       []*Type
}
