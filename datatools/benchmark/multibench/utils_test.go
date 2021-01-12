package multibench

import (
	"reflect"
	"testing"
)

func TestListDigits(t *testing.T) {
	type args struct {
		number int
	}
	tests := []struct {
		name       string
		args       args
		wantRetval []int
	}{
		// TODO: Add test cases.
		{"1234", args{1234}, []int{4, 3, 2, 1}},
		{"101010101010101043", args{101010101010101043}, []int{3, 4, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1}},
		{"59843", args{59843}, []int{3, 4, 8, 9, 5}},
		{"-473", args{-473}, []int{3, 7, -4}},
		{"0", args{0}, []int{0}},
		{"-1", args{-1}, []int{-1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRetval := ListDigits(tt.args.number); !reflect.DeepEqual(gotRetval, tt.wantRetval) {
				t.Errorf("ListDigits() = %v, want %v", gotRetval, tt.wantRetval)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	type args struct {
		lst []string
	}
	tests := []struct {
		name string
		args args
		want chan string
	}{
		// TODO: Add test cases.
		{"hello", args{[]string{"Test", "o", "l", "l", "e", "h"}}, ch.},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reverse(tt.args.lst); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
