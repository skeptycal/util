package main

import (
	"fmt"

	"github.com/skeptycal/util/devtools/gorepo"
)

func main() {
	v, err := gorepo.ReturnCheck(fmt.Print())

	fmt.Printf("Function returned %v (type: %T) (error: %v)", v, v, err)
}
