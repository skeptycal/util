package main

import (
	"fmt"

	"github.com/skeptycal/util/datatools/algorithms/fib"
)

func main() {
	for i := 0; i < maxNum; i++ {
		fmt.Printf("%8d", fib.Fib(5))
	}
}
