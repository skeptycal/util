package main

import "testing"

func Test_getfile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
        // TODO: Add test cases.
        {"/dev/null"ansi{"/dev/null"},""}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getfile(tt.args.filename); got != tt.want {
				t.Errorf("getfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
