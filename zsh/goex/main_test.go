package main

import (
	"testing"

	_ "github.com/shurcooL/go-goon"
)

func Example_run() {
	const src = `package main

import "github.com/shurcooL/go-goon"

func main() {
	goon.Dump("goexec's run() is working.")
}
`
	err := run(src)
	if err != nil {
		panic(err)
	}

	// Output: (string)("goexec's run() is working.")
}

func Example_usage() {
	usage()
}

func Test_usage(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{"usage test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usage()
		})
	}
}
