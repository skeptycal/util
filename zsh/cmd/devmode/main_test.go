package main

import (
	"os"
	"strings"
	"testing"
)

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
		{"/dev/null first word", args{"/dev/null"}, ""},
		{"main.go first word", args{"main.go"}, "package"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strings.Split(getfile(tt.args.filename), " ")[0]
			if got != tt.want {
				t.Errorf("getfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_myfile(t *testing.T) {

        home, err := os.UserHomeDir()
    if err != nil {
        t.Errorf("cannot locate user home directory: %v", err)
    }
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
        // TODO: Add test cases.
        {"myfile",args{"myfile"},home +"myfile"},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myfile(tt.args.filename); got != tt.want {
				t.Errorf("myfile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getFile(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
        // TODO: Add test cases.
        {"/dev/null first word", args{"/dev/null"}, ""},
		{"main.go first word", args{"main.go"}, "package"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := strings.Split(getFile(tt.args.filename), " ")[0]
			if got != tt.want {
				t.Errorf("getFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
