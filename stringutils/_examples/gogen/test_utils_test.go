package gogen

import (
	"testing"
)

func TestGotWant(t *testing.T) {
	type args struct {
		name    string
		subname string
		got     string
		want    string
		t       *testing.T
	}
	tests := []struct {
		name       string
		args       args
		wantRetval bool
	}{
		// TODO: Add test cases.
		{"fakeGood", args{"name", "subname", "result", "result", t}, false},
		{"fakeBad", args{"name", "subname", "stuff", "nostuff", t}, true},
		// not possible if things are correct:
		// (non matching but not expecting error??)
		{"impossible", args{"match", "args", "stuff", "nostuff", t}, false},
		// not possible, but not tested ... if they match; doesn't matter
		// (matching but expecting error??)
		{"unrealistic", args{"match", "args", "stuff", "stuff", t}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = GotWant(tt.args.name, tt.args.subname, tt.args.got, tt.args.want, tt.wantRetval, tt.args.t)
		})
	}
}

func TestErrTest(t *testing.T) {
	var a Any = nil
	tests := []struct {
		name    string
		test    interface{}
		want    bool
		wantErr bool
	}{
		{"int 0", 0, true, true},
		{"int 42", 42, true, true},
		{"blank", a, false, false},
		{"nil", nil, false, false},
		{"empty string", "", true, true},
		{"bool true", true, true, false},
		{"bool false", false, false, false},
		{"float64 3.14", float64(3.14), true, true},
		{"float32 3.14", float32(3.14), true, true},
		{"nil interface", interface{}(nil), false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ErrTest(tt.name, tt.test, tt.wantErr, t); got != tt.want {
				t.Errorf("ErrTest %v(%v) = %v, want %v", tt.name, tt.test, got, tt.want)
			}
		})
	}
}
