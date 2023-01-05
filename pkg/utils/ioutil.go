package utils

import (
	"os"
)

func EnsureStdinReader(reader *os.File) bool {
	state, _ := reader.Stat()
	return (state.Mode() & os.ModeCharDevice) == 0
}
