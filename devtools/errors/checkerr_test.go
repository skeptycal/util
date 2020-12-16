package errors

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"testing"
	"text/template"
)

var (
	stderr       = os.Stderr
	errFakeError = errors.New("this is a fake error: foo")
)

func BenchmarkTemplateParallel(b *testing.B) {
	// example of parallel setup ...
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			err := templ.Execute(&buf, "World")
			if err != nil {
				b.Fail()
			}
		}
	})
}

// ************************************ Errf Tests
func TestErrf(t *testing.T) {
	if Errf(errFakeError, "fake error: %v") == true {
		t.Fail()
	}
	if Errf(nil, "") == false {
		t.Fail()
	}
}
func ExampleErrf() {
	Errf(errFakeError, "fake error: %v")
	// Output:
	// fake error: this is a fake error: foo
}

// ************************************ TryExceptPass Tests
func TestTryExceptPassNoOutput(t *testing.T) {
	// should run without errors and return nothing
	// produces no output
	TryExceptPass(errFakeError, false)
}
func TestTryExceptPassWithOutput(t *testing.T) {
	// should run without errors and return nothing
	// produces stdout output
	TryExceptPass(errFakeError, true)
}
func ExampleTryExceptPass() {
	TryExceptPass(errFakeError, false)
	// Output:
	//
	TryExceptPass(errFakeError, true)
	// Output:
	// An error occurred: this is a fake error: foo
}
func BenchmarkTryExceptPassNoOutput(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TryExceptPass(errFakeError, false)
	}
}
func BenchmarkTryExceptPassWithOutput(b *testing.B) {
	b.Log("running BenchmarkTryExceptPassWithOutput ...")
	// too slow ...
	// for i := 0; i < b.N; i++ {
	// 	TryExceptPass(errFakeError, true)
	// }
}

func Test_checkPanic(t *testing.T) {

	tests := []struct {
		name      string
		e         error
		want      error
		wantErr   bool
		wantPanic bool
	}{
		//TODO: write test cases
		{"test 1", nil, nil, false, false},
		{"test 2", nil, nil, true, false},
		{"test 3", nil, nil, false, true},
		// {"Panic", errFakeError, errFakeError, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				r := recover()
				if (r != nil) && !tt.wantPanic {
					t.Errorf("SequenceInt() recover = %v, wantPanic = %v", r, tt.wantPanic)
				}
			}()

			err := checkPanic(tt.e)
			if err != nil {
				if !tt.wantErr {
					fmt.Fprintf(stderr, "unexpected  error %v, wantErr %v", err, tt.wantErr)
				}
			}
			if (err != nil) && !tt.wantErr {
				t.Fail(fmt.Errorf("YourFunc() error = %v, wantErr %v", err, tt.wantErr))
				return
			}
			if err != tt.want {
				t.Errorf("YourFunc() = %v, want %v", err, tt.want)
			}
		})
	}
}
