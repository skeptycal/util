package main

import (
	"path/filepath"
)

func main() {
	buildtools.BuildList()
	println("filepath.Clean() result: ", filepath.Clean(""))

}
