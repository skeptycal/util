// Package format contains functions that format numeric values.
package format

import "testing"

func TestNumSpace(t *testing.T) {
	type args struct {
		n float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"12345.54321", args{12345.54321}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumSpace(tt.args.n); got != tt.want {
				t.Errorf("NumSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
