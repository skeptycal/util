package main

import (
	"path/filepath"
	"runtime"
)

func main() {

	// on Windows, using
	//      GOOS=windows go build
	// the getEmptyPath function returned :         .\
	// filepath.Clean("") returned :                        .

	println("getEmptyPath result: ", getEmptyPath())
	println("filepath.Clean() result: ", filepath.Clean(""))

}

func getEmptyPath() string {
	if runtime.GOOS == "windows" || runtime.GOOS == "plan9" {
		return ".\\"
	}
	return "."
}
