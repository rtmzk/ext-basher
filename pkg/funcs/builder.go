package funcs

const (
	ExportFunction = iota + 1
	ExportVariables
)

type FuncBuilder struct {
	Name      string
	Type      int32
	Cmd       func(args []string)
	Variables string
}

func (fb *FuncBuilder) setCmd(cmd func(args []string)) *FuncBuilder {
	fb.Cmd = cmd
	return fb
}

func (fb *FuncBuilder) setVariables(v string) *FuncBuilder {
	fb.Variables = v
	return fb
}

func (fb *FuncBuilder) build() *Func {
	function := &Func{
		Name: fb.Name,
		Type: fb.Type,
	}
	if fb.Type == ExportFunction {
		function.Cmd = fb.Cmd
		function.Variables = ""
	}

	if fb.Type == ExportVariables {
		function.Cmd = nil
		function.Variables = fb.Variables
	}

	return function
}

func newFuncBuilder(name string, t int32) *FuncBuilder {
	return &FuncBuilder{
		Name: name,
		Type: t,
	}
}
