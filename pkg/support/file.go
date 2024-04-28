package support

import (
	"log"
	"path/filepath"
	"runtime"
)

func GetFilePath(path string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Failed to get the current file's path")
	}
	dir := filepath.Dir(filename)
	return dir + "/../../" + path
}
