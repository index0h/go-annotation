package annotation

// FuncSpec represents specification of func.
type FuncSpec struct {
	Params     []*Field
	Results    []*Field
	IsVariadic bool
}
