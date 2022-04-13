package utils

import (
	"os"
	"strings"
)

func ReadLinesOfFile(path string) []string {
	var data, err = os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var content = strings.TrimSpace(string(data))
	return strings.Split(content, "\n")
}
