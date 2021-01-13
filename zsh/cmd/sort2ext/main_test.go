// sort2ext takes a directory of files and sorts
// all of them into new folders according to the
// file extension.
package main

import (
	"os"
	"testing"
)

type thing string

func (thing thing) String() string {
	return string(thing)
}

func ExampleEchoSB() {

	var s thing = "stringer"

	EchoSB("hello")
	EchoSB([]byte{65, 66, 67, 68})
	EchoSB(23)
	EchoSB(3.14)
	// EchoSB(nil)
	// EchoSB(os.Stdin)
	EchoSB(s)

	//Output:
	// hello
	// ABCD
	// 23
	// 3.140000

}

func TestEcho(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		// {"no args", args{}},
		// {"empty string", args{[]interface{}{""}}},
		// {"one string", args{[]interface{}{"hello"}}},
		{"two strings; without format string", args{[]interface{}{"hello, \n", "Mike"}}},
		// {"with format string", args{[]interface{}{"hello, %s\n", "James"}}},
		// {"with format string; multiple strings", args{[]interface{}{"hello, %6s (%T)\n", "Sarah", "Mary"}}},
		// {"with format string; multiple various args", args{[]interface{}{"int: %v (%T)\n", 1, 1}}},
		// {"with format string; %%v int", args{[]interface{}{"formatted int: %4d (%T)\n", 1, 1}}},
		// {"with format string; %%v float32", args{[]interface{}{"float: %v (%T)\n", float32(3.14), float32(3.14)}}},
		// {"with format string; multiple strings", args{[]interface{}{"hello, %s (%T)\n", "Mike", "Mike"}}},
		// {"just ints", args{[]interface{}{2, 3, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Echo(tt.args.v...)
		})
	}
}

// func ExampleEcho() {
// 	// Echo()
// 	Echo("", "hello")
// 	Echo(1, " - one")
// 	Echo("\n")
// 	Echo("hello, %s", "Mike")
// 	Echo("I like Pie: %.2f", math.Pi)
// 	Echo("I like Floating Pie: %T", math.Pi)
// 	Echo("", "No formatting ...")
// 	Echo("type: %T", "stuff")
// 	Echo()
// 	Echo("")

// 	// output:
// 	// hello
// 	// 1 - one
// 	//
// 	// hello, Mike
// 	// I like Pie: 3.14
// 	// I like Floating Pie: float64
// 	// No formatting ...

// }

func TestEchoSB(t *testing.T) {
	type args struct {
		v []interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
		{"hello", args{[]interface{}{"hello"}}, 6, false},
		{"hello", args{[]interface{}{[]byte{65, 66, 67, 68}}}, 5, false},
		{"hello", args{[]interface{}{23}}, 3, false},
		{"hello", args{[]interface{}{3.14}}, 9, false},
		{"hello", args{[]interface{}{nil}}, 1, false},
		{"hello", args{[]interface{}{os.Stdin}}, 1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := EchoSB(tt.args.v...)
			if (err != nil) != tt.wantErr {
				t.Errorf("EchoSB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("EchoSB() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
