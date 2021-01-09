package gofile

import (
	"os"
	"reflect"
	"testing"
)

func TestChunkMultiple(t *testing.T) {
	tt := []struct {
		name     string
		size     int64
		expected int64
	}{
		{"size 1.x chunk", 550, 1024},
		{"size 2.x chunk", 1200, 1536},
		{"chunk size 16", 100, 512},
		{"1.x size", 5234, 5632},
		{"42kb file", 42000, 42496},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			result := chunkMultiple(tc.size)
			if result != tc.expected {
				t.Errorf("expected value <%v> does not match result: %v", tc.expected, result)
			}
		})
	}
}

func TestInitialCapacity(t *testing.T) {
	type args struct {
		capacity int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"below default (4096)", args{16}, 4096},
		{"1.x capacity (chunk 512)", args{5333}, 5632},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitialCapacity(tt.args.capacity); got != tt.want {
				t.Errorf("InitialCapacity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRegularFileInfo(t *testing.T) {
	fiSample, _ := os.Stat("fileops_test.go")
	fiTestJSON, _ := os.Stat("testdata/test.json")
	// fiDev, _ := os.Stat("/dev")

	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    os.FileInfo
		wantErr bool
	}{
		// TODO: Add test cases.
		{"this file: fileops_test.go", args{"fileops_test.go"}, fiSample, false},
		{"filetest testdata/test.json", args{"testdata/test.json"}, fiTestJSON, false},
		// filepath.EvalSymlinks() in GetRegularFileInfo() parses the path to return the target of any symlinks before proceeding
		{"symlink testdata/test_ln.json", args{"testdata/test_ln.json"}, fiTestJSON, false},
		{"not exist error", args{"fakefile"}, nil, true},
		{"PWD() directory error", args{PWD()}, nil, true},
		{"not regular file error", args{"/dev/null"}, nil, true},
		{"directory permission error", args{"/dev/rdisk0"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRegularFileInfo(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRegularFileInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRegularFileInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
