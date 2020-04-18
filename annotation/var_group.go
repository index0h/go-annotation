package annotation

// VarGroup represents list of var declaration.
type VarGroup struct {
	Comment     string
	Annotations []interface{}
	Vars        []*Var
}
