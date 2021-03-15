// package main provides a simplified re-implementation of the head command,
// which displays the first several lines of a given file.
//
// This is predominately an example of CLI flag usage.
//
// Reference: https://www.digitalocean.com/community/tutorials/how-to-use-the-flag-package-in-go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

const (
	colorUsage string = `colorize the output; WHEN can be 'always' (default if omitted), 'auto', or 'never'; more info below`
)

var (
	debug = kingpin.Flag("debug", "enable debug mode").Default("false").Bool()
	// serverIP = kingpin.Flag("server", "server address").Default("127.0.0.1").IP()

	count  = *kingpin.Flag("count", "number of lines to read from the file").Default("10").Int()
	color  = *kingpin.Flag("color", colorUsage).Default("always").String()
	output = *kingpin.Flag("output", "send output to fil (default ").Default("").String()
)

func main() {

	kingpin.Usage()

	fmt.Println("")
	fmt.Printf("debug: %v\n", debug)
	fmt.Printf("count: %v\n", count)
	fmt.Printf("color: %q\n", color)
	fmt.Printf("output: %q\n", output)
	fmt.Println("")

	var in io.Reader
	if filename := flag.Arg(0); filename != "" {
		f, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file: err:", err)
			os.Exit(1)
		}
		defer f.Close()

		in = f
	} else {
		in = os.Stdin
	}

	buf := bufio.NewScanner(in)

	for i := 0; i < count; i++ {
		if !buf.Scan() {
			break
		}
		fmt.Println(buf.Text())
	}

	if err := buf.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading: err:", err)
	}
}
