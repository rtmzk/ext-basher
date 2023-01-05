package funcs

import (
	"fmt"
	"strings"
)

var upperFunction = new(Func).
	builder("upper", ExportFunction).
	setCmd(oneArgumentOrPipe(upper)).
	build()

var lowerFunction = new(Func).
	builder("lower", ExportFunction).
	setCmd(oneArgumentOrPipe(lower)).
	build()

func upper(s string) {
	fmt.Println(strings.ToUpper(s))
}

func lower(s string) {
	fmt.Println(strings.ToLower(s))
}
