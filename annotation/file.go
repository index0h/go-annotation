package annotation

type File struct {
	Name         string
	Content      string
	PackageName  string
	Comment      string
	Annotations  []interface{}
	ImportGroups []*ImportGroup
	ConstGroups  []*ConstGroup
	VarGroups    []*VarGroup
	TypeGroups   []*TypeGroup
	Funcs        []*Func
}
