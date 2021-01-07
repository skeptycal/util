package gofile

import (
	"testing"
)

func BenchmarkSplit2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split2(PWD())
	}
}

func BenchmarkSplit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split(PWD())
	}
}

func TestSplit2(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name     string
		args     args
		wantDir  string
		wantFile string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDir, gotFile := Split2(tt.args.path)
			if gotDir != tt.wantDir {
				t.Errorf("Split2() gotDir = %v, want %v", gotDir, tt.wantDir)
			}
			if gotFile != tt.wantFile {
				t.Errorf("Split2() gotFile = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}
