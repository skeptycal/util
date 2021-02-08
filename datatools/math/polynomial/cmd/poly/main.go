package main

import (
	"fmt"

	"github.com/skeptycal/util/datatools/math/polynomial"
)

func main() {
    l1 := polynomial.New(77777)

    fmt.Println(l1.String())
    fmt.Println(l1)

    fmt.Println("---")
    fmt.Println(polynomial.StringDigits(12345))

}
