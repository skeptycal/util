// sort2ext takes a directory of files and sorts
// all of them into new folders according to the
// file extension.
package main

import (
	"os"
	"testing"
)

func TestEcho(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"hello", args{[]interface{}{"hello"}}},
		{"hello, name", args{[]interface{}{"hello, %s", "Mike"}}},
		{"hello", args{[]interface{}{"hello"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Echo(tt.args.v...)
		})
	}
}
func ExampleEcho() {
	Echo(os.Stdout, "hello")
	// output:
	// false
}

// func ExampleEcho() {
// 	Echo(os.Stdout, "hello")
// Echo("hello, %s\n", "Mike")
// Echo("I like Pie: %.2f\n", math.Pi)
// Echo("I like Floating Pie: %T\n", math.Pi)
// Echo("No formatting ...")
// Echo("type: %T", os.Stdout, "stuff")

// output:
// hello
// hello, Mike
// I like Pie: 3.14
// I like Floating Pie: float64
// No formatting ...

// }
