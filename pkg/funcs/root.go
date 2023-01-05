package funcs

type functionSlice struct {
	Functions []Func
}

var Functions = new(functionSlice)

func (fs *functionSlice) addCommand(f Func) *functionSlice {
	fs.Functions = append(fs.Functions, f)
	return fs
}

func init() {
	Functions.
		addCommand(*upperFunction).
		addCommand(*lowerFunction).
		addCommand(*memoryAvailable).
		addCommand(*memoryUsage)
}
