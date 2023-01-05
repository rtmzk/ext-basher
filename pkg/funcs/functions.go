package funcs

type Func struct {
	Name      string
	Type      int32
	Cmd       func(args []string)
	Variables string
}

func (f *Func) builder(name string, t int32) *FuncBuilder {
	return newFuncBuilder(name, t)
}
