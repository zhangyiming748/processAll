package util

import (
	"path"
	"runtime"
)

var Root string

func SetRoot() {
	_, filename, _, _ := runtime.Caller(0)
	Root = path.Dir(filename)
}

func GetRoot() string {
	return Root
}
