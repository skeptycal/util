package main

import (
	"testing"
)

func BenchmarkDirShellBM(b *testing.B) {

}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestDirShellBM(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"pwd", args{"."}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirShellBM(tt.args.path); got != tt.want {
				t.Errorf("DirShellBM() = %v, want %v", got, tt.want)
			}
		})
	}
}
