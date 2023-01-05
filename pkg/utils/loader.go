package utils

import (
	"os"
	"regexp"
)

func Loader(f string) ([]byte, error) {
	b, err := os.ReadFile(f)

	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`#!.*\n`)
	b = re.ReplaceAll(b, []byte(""))

	return b, nil
}
