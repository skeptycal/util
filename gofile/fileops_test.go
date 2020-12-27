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

func TestGetFileInfo(t *testing.T) {
	fiSample, _ := os.Stat("fileops_test.go")
	fiTestJson, _ := os.Stat("testdata/test.json")
	fiDev, _ := os.Stat("/dev")

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
		{"fileops_test.go", args{"fileops_test.go"}, fiSample, false},
		{"fakefile", args{"fakefile"}, nil, true},
		{"directory", args{"~/"}, nil, true},
		{"json file", args{"testdata/test.json"}, fiTestJson, false},
		// filepath.EvalSymlinks() in GetFileInfo() parses the path to return the target of any symlinks before proceeding
		{"symlink", args{"testdata/test_ln.json"}, fiTestJson, false},
		{"directory permission error", args{"/dev"}, fiDev, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFileInfo(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFileInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
