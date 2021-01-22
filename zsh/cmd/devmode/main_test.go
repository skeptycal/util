package main

import (
	"os"
	"path"
	"strings"
	"testing"
)

/* Benchmark results

using exec is 232 times slower ...
and uses 14.5 times more allocations

goos: darwin
goarch: amd64
pkg: github.com/skeptycal/util/zsh/cmd/devmode
BenchmarkGetFileUsingExec-8   	     453	   2729513 ns/op	   85448 B/op	      58 allocs/op
BenchmarkGetFile-8            	       102792	        11759 ns/op	          632 B/op	         4 allocs/op
*/

func BenchmarkGetFileUsingExec(b *testing.B) {
    for i := 0; i < b.N; i++ {
        getFileUsingExec("/dev/null")
    }
}

func BenchmarkGetFile(b *testing.B) {
    for i := 0; i < b.N; i++ {
        getFile("/dev/null")
    }
}

func Test_getFileUsingExec(t *testing.T) {
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
			got := strings.Split(string(getFileUsingExec(tt.args.filename)), " ")[0]
			if got != tt.want {
				t.Errorf("getFileUsingExec() = %v, want %v", got, tt.want)
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
        contents, err := getFile(tt.args.filename)
        if err != nil {
            t.Errorf("error opening config file: %v", err)
        }
		t.Run(tt.name, func(t *testing.T) {
			got := strings.Split(string(contents), " ")[0]
			if got != tt.want {
				t.Errorf("getFile() = %v, want %v", got, tt.want)
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
        {"myfile",args{"myfile"}, path.Join(home, "myfile")},
        {"configFile",args{configFile}, path.Join(home, configFile)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myfile(tt.args.filename); got != tt.want {
				t.Errorf("myfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
