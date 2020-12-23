package main

import (
	"fmt"
	"os"

	"github.com/skeptycal/util/gofile"
)

const ansi256Fmt = "\033[%d;"

type Ansi int8

func (a Ansi) String() string {
	return fmt.Sprintf(ansi256Fmt, a)
}

func Usage() {
	programName := gofile.Base(os.Args[0])
	fmt.Println(os.Args)
	fmt.Printf("Usage:%s %s %s-h -v\n", Ansi(34), programName, Ansi(33))
}

func main() {
	if len(os.Args) < 2 {
		Usage()
		os.Exit(1)
	}

	filename := os.Args[1]
	fmt.Println("File info:")
	fmt.Println(gofile.Parents(filename))
}
