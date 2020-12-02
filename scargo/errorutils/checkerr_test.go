package errorutils

import (
	"bytes"
	"errors"
	"testing"
	"text/template"
)

var errFakeError = errors.New("this is a fake error: foo")

func BenchmarkTemplateParallel(b *testing.B) {
	// example of parallel setup ...
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	b.RunParallel(func(pb *testing.PB) {
		var buf bytes.Buffer
		for pb.Next() {
			buf.Reset()
			templ.Execute(&buf, "World")
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
	return
	// too slow ...
	// for i := 0; i < b.N; i++ {
	// 	TryExceptPass(errFakeError, true)
	// }
}
