package test

import (
	"path"
	"runtime"
)

func GetTestDataDir(testFileName string) string {
	_, filename, _, _ := runtime.Caller(0)

	dir := path.Join(path.Dir(filename), "testdata", testFileName)
	return dir
}
