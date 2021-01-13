// sort2ext takes a directory of files and sorts
// all of them into new folders according to the
// file extension.
package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
)

type Stringer interface {
	String() string
}

const (
	fmtEchoString string = "%v "
)

var (
	basename   string = path.Base(os.Args[0])
	pwd, _            = os.Getwd()
	echoWriter        = os.Stdout
)

func main() {
	argpath := pwd
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "-h", "--help":
			usage()
			os.Exit(0)
		default:
			arg, err := os.Stat(os.Args[1])
			if err == nil {
				argpath = arg.Name()
				Echo("argpath: %s", argpath)
			} else {
				usage()
				log.Fatal(err)
			}
		}
	}

	Echo("base: %s", basename)
	Echo("argpath: %s", argpath)

}

func usage() { fmt.Fprintf(os.Stderr, "Usage: %s  [path]\n", basename) }

// Echo checks for format strings, io.Writers, and ANSI tags before printing results.
func Echo(v ...interface{}) {

	// single string argument  ... newline assumed
	// (most common desireable outcome)
	// to print a single string without newline, use fmt.Print() instead.
	if len(v) == 1 {
		fmt.Fprintln(echoWriter, v[0])
		return
	}
	// no arguments ... interpreted as newline only
	if len(v) == 0 {
		fmt.Fprintf(echoWriter, "%v\n", "")
		return
	}

	switch s := v[0].(type) {
	case string:
		if strings.ContainsAny(s, "%") {
			// fmtString = v[0].(string)
			// args = args[1:]
			fmt.Fprintf(echoWriter, s, v[1:]...)
			return
		}
	default:
		fmt.Fprintln(echoWriter, v[1:])
		return
	}

	for _, a := range v[:len(v)-1] {
		fmt.Fprintf(echoWriter, fmtEchoString, a)
	}
	fmt.Fprintf(echoWriter, "%v\n", v[len(v)-1])

}

func EchoSB(v ...interface{}) (n int, err error) {
	sb := strings.Builder{}

	for _, a := range v {

		switch t := a.(type) {
		case string:
			sb.WriteString(t)
		case []byte:
			sb.Write(t)
		case int:
			sb.WriteString(fmt.Sprintf("%d", t))
		case float32, float64:
			sb.WriteString(fmt.Sprintf("%f", t))
		case nil:

		default:
			if s, ok := t.(Stringer); ok {
				sb.WriteString(s.String())
			}

		}
	}

	return fmt.Println(sb.String())
}
