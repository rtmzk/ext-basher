package basher

import (
	"github.com/progrium/go-basher"
	"github.com/rtmzk/ext-basher/pkg/funcs"
	"log"
	"os"
)

func NewBashContext() *basher.Context {
	bash, err := basher.NewContext("/bin/bash", true)
	if err != nil {
		log.Fatalf("error create basher context: %s", err.Error())
	}

	for _, f := range funcs.Functions.Functions {
		if f.Type == funcs.ExportFunction {
			bash.ExportFunc(f.Name, f.Cmd)
		}
		if f.Type == funcs.ExportVariables {
			bash.Export(f.Name, f.Variables)
		}
	}

	if bash.HandleFuncs(os.Args) {
		os.Exit(0)
	}

	return bash
}
