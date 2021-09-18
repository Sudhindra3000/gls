package util

import (
	"strings"
)

func EndingWith(arg string) (string, bool) {
	if strings.HasPrefix(arg, "*") {
		return arg[1:], true
	}
	return "", false
}

func StartingWith(arg string) (string, bool) {
	if strings.HasSuffix(arg, "*") {
		return arg[0 : len(arg)-1], true
	}
	return "", false
}
