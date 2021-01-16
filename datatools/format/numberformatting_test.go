// Package format contains functions that format numeric values.
package format

import "testing"

func TestNumSpace(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"12345.54321e42", args{"12345.54321e42"}, "12345.54321e42"},
		{"1", args{"1"}, "1"},
		{"-1", args{"-1"}, "-1"},
		{"0.123", args{"0.123"}, "0.123"},
		{"1", args{"1"}, "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumSpace(tt.args.n); got != tt.want {
				t.Errorf("NumSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
