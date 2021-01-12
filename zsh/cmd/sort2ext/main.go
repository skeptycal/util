// sort2ext takes a directory of files and sorts
// all of them into new folders according to the
// file extension.
package main

import (
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	basename string = path.Base(os.Args[0])
	pwd, _          = os.Getwd()
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
    if len(v) < 2 {
        fmt.Println(v...)
    }
	var w io.Writer = os.Stdout
	var fmtString string = "%v\n"
	var args []interface{} = v[0:]

	g := v[0]

	// gt := reflect.TypeOf(g)

	fmt.Printf("userType: %T\n", g)
	t := fmt.Sprintf("type: %T\n", g)
	if strings.Contains(t, "os.File") {
		fmt.Println("yes!")
	}

	// if reflect.Kind(v[0]) == reflect.Ptr {

	// }

	if strings.Contains(t, "os.File") {
		w = args[0].(io.Writer)
		args = args[1:]
		args = append(args, " ")
	}

	// if strings.ContainsAny(args[0].(string), `%`) && len(args) > 1 {
	// 	fmtString = v[0].(string)
	// 	args = args[1:]
	// }

	// fmt.Printf(fmtString, args...)
	fmt.Fprintf(w, fmtString, args...)
}
