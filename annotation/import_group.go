package annotation

// ImportGroup represents list of import declaration.
type ImportGroup struct {
	Comment     string
	Annotations []interface{}
	Imports     []*Import
}
