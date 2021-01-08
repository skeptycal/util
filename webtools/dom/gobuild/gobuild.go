// Package gobuild is a small program to build and initialize
// js WebAssembly projects.
//
// Reference: https://github.com/golang/go/wiki/WebAssembly
package gobuild

import (
	"github.com/skeptycal/util/gofile"
	"github.com/skeptycal/util/zsh"
)

const (
	wasmBuildCommandString = "GOOS=js GOARCH=wasm go build -o main.wasm"
)

func BuildWASM(path string) error {
	if path == "" || path == "." {
		path = gofile.PWD()
	}

	path = gofile.Abs(path)
	err := zsh.Status(wasmBuildCommandString)
	if err != nil {
		return err
	}
	return nil
}
