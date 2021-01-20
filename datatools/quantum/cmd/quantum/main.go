package main

import (
	"fmt"
	"os"

	"github.com/skeptycal/util/gofile"
)

func main() {
    fmt.Println(gofile.Here())
    fmt.Println(os.Args[0])
}
