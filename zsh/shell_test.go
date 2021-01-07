package zsh

import (
	"fmt"
	"io"
	"os/exec"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestBackground(t *testing.T) {
	c := DefaultContext
	if c == nil {
		t.Fatalf("Background returned nil")
	}
	select {
	case x := <-c.Done():
		t.Errorf("<-c.Done() == %v want nothing (it should block)", x)
	default:
	}
	if got, want := fmt.Sprint(c), "context.Background"; got != want {
		t.Errorf("Background().String() = %q want %q", got, want)
	}
}

func TestCombinedOutput(t *testing.T) {
	type args struct {
		command string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"macos specific", args{"uname -o"}, "Darwin"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CombinedOutput(tt.args.command); got != tt.want {
				t.Errorf("CombinedOutput() = %v, want %v", got, tt.want)
			}
		})
	}
	// additional tests
	s := CombinedOutput("fake_program_that_is_not_real")
	if s[:2] == `/x1b[` {
		t.Errorf("OutErr() = %v, want %v", s, "(error string beginning with '/x1b'")

	}
}

func ExampleRepl() {
	command := "printf '%s\n' 'Some string to be displayed'"
	if _, err := Repl(command); err != nil {
		log.Info(err)
	}

	// Output:
	// some io.Reader stream to be read
}

/*
func ExampleCopy() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}

	// Output:
	// some io.Reader stream to be read
}
*/

func ExampleCmd_StdinPipe() {
	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin are passed to cmd's standard input")
	}()

	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}
