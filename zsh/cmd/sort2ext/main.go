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
	"github.com/skeptycal/util/zsh"
)

const (
	fmtEchoString string = "%v "
)

var (
	basename string = path.Base(os.Args[0])
	pwd, _          = os.Getwd()
	echo            = os.Stdout
)

// ---------------------------------------------
// -------------------------------------- set

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
	usage()

	zsh.Dir(pwd)
}

func usage() { Echo("Usage: %s [path]\n", basename) }

// Echo checks for format strings, io.Writers, and ANSI tags before printing results.
// Output is sent to echo, a buffered io.Writer implemented by EchoWriter.
func Echo(v ...interface{}) {

	// single string argument  ... newline assumed
	// (most common desireable outcome)
	// to print a single string without newline, use fmt.Print() instead.
	if len(v) == 1 {
		fmt.Fprintln(echo, v[0])
		return
	}
	// no arguments ... interpreted as newline only
	if len(v) == 0 {
		fmt.Fprintf(echo, "%v\n", "")
		return
	}

	switch s := v[0].(type) {
	case string:
		if strings.ContainsAny(s, "%") {
			// fmtString = v[0].(string)
			// args = args[1:]
			fmt.Fprintf(echo, s, v[1:]...)
			fmt.Println()
			return
		}
	default:
		fmt.Fprintln(echo, v[1:]...)
		return
	}

	for _, a := range v[:len(v)-1] {
		fmt.Fprintf(echo, fmtEchoString, a)
	}
	fmt.Fprintf(echo, "%v\n", v[len(v)-1])

}
