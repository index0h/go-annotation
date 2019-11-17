package annotation

type Storage struct {
	Namespaces []*Namespace
}

type Namespace struct {
	Name  string
	Path  string
	Files []*File
}

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

type ImportGroup struct {
	Comment     string
	Annotations []interface{}
	Imports     []*Import
}

type Import struct {
	Alias       string
	Namespace   string
	Comment     string
	Annotations []interface{}
}

type ConstGroup struct {
	Comment     string
	Annotations []interface{}
	Consts      []*Const
}

type Const struct {
	Name        string
	Value       string
	Comment     string
	Annotations []interface{}
	Spec        interface{}
}

type VarGroup struct {
	Comment     string
	Annotations []interface{}
	Vars        []*Var
}

type Var struct {
	Name        string
	Value       string
	Comment     string
	Annotations []interface{}
	Spec        interface{}
}

type TypeGroup struct {
	Comment     string
	Annotations []interface{}
	Types       []*Type
}

type Type struct {
	Name        string
	Comment     string
	Annotations []interface{}
	Spec        interface{}
}

type Func struct {
	Name        string
	Content     string
	Comment     string
	Annotations []interface{}
	Spec        *FuncSpec
	Related     *Field
}

type Field struct {
	Name    string
	Tag     string
	Comment string

	Annotations []interface{}
	Spec        interface{}
}

type SimpleSpec struct {
	PackageName string
	TypeName    string
	IsPointer   bool
}

type ArraySpec struct {
	Value         interface{}
	Length        int
	IsFixedLength bool
	IsEllipsis    bool
}

type MapSpec struct {
	Key   interface{}
	Value interface{}
}

type StructSpec struct {
	Fields []*Field
}

type InterfaceSpec struct {
	Methods []*Field
}

type FuncSpec struct {
	Params  []*Field
	Results []*Field
}
