package funcs

import (
	"github.com/rtmzk/ext-basher/pkg/utils"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

func oneArgumentOrPipe(handler func(string)) func(args []string) {
	return func(args []string) {
		if len(args) == 1 {
			handler(args[0])
		} else {
			if utils.EnsureStdinReader(os.Stdin) {
				b, _ := io.ReadAll(os.Stdin)
				plainText := strings.Trim(string(b), "\n")
				handler(plainText)
			}
		}
	}
}

// in: -> byte
func unitFormat(in uint64) string {
	var out = ""
	var suffix = ""

	if in > TB {
		suffix = "T"
		out = strconv.FormatUint(in/TB, 10)
	} else if in > GB {
		suffix = "G"
		out = strconv.FormatUint(in/GB, 10)
	} else if in > MB {
		suffix = "M"
		out = strconv.FormatUint(in/MB, 10)
	} else {
		suffix = "K"
		out = strconv.FormatUint(in/KB, 10)
	}

	return out + suffix
}
