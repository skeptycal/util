package main

import (
	"fmt"

	"github.com/skeptycal/util/devtools/gorepo"
)

func main() {
	v, err := gorepo.ReturnCheck(fmt.Print())

	fmt.Printf("Function returned %v (%v)", v, err)
}
