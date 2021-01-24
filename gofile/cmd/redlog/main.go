package main

import (
	"os"

	"github.com/skeptycal/util/gofile/redlogger"
)

var r = redlogger.New(os.Stderr)

func main() {
    defer r.Flush()
    r.WriteString("Hello World!")
}
