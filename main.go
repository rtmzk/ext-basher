package main

import (
	"github.com/rtmzk/ext-basher/pkg/basher"
	"github.com/rtmzk/ext-basher/pkg/utils"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 0 {
		var bctx = basher.NewBashContext()
		err := bctx.Source(os.Args[1], utils.Loader)
		if err != nil {
			log.Fatal(err)
		}
		status, err := bctx.Run("", os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}

		os.Exit(status)
	}
}
