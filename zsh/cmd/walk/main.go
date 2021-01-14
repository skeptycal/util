// This example demonstrates how to use go to walk a directory
// and filter the contents by using filepath.Walk with
// variables in the external scope to provide the filter and
// to capture the file names.
//
// To build:
//     go build (in parent directory 'walk')
// Usage:
//     walk <pattern> <files>
//
// License:
//     MIT Open Source
//     Copyright (c) 2016 by Joe Linoff
package main

import (
	"fmt"
	"os"

	"github.com/skeptycal/util/zsh/fileset"
)

func usage() {
	fmt.Println("Usage: Walk [pattern] [files]")
}

func main() {

	searchpath := "."
	pattern := "*"

	if len(os.Args) > 2 {
		searchpath = os.Args[2]
	}

	if len(os.Args) > 1 {
		pattern = os.Args[1]
	}

	fmt.Printf("Walk %s %s\n", pattern, searchpath)

	files := fileset.ReDir(pattern, searchpath)

	for _, f := range files {
		fmt.Printf(" %s\n", f)
	}

}
