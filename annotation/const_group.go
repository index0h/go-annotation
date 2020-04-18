package annotation

// ConstGroup represents list of const declaration.
type ConstGroup struct {
	Comment     string
	Annotations []interface{}
	Consts      []*Const
}
