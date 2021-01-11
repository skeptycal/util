// sort2ext takes a directory of files and sorts
// all of them into new folders according to the
// file extension.
package main

import (
	"fmt"
	"log"
	"os"
	"path"
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
			} else {
				usage()
				log.Fatal(err)
			}
		}
	}

	fpln("base: %s", basename)
	fpln("argpath: %s", argpath)

}

func usage()                          { fmt.Fprintf(os.Stderr, "Usage: %s  [path]\n", basename) }
func fpln(f string, v ...interface{}) { fmt.Println(fmt.Sprintf(f, v...)) }
