package main

import (
	"fmt"

	"github.com/skeptycal/util/datatools/math/primes"
)

type trials []trial

type trial struct {
	name string
	f    func(n uint8) uint8
}

func WhiteSpace(c rune) bool {
	switch c {
	case ' ', '\t', '\n', '\f', '\r':
		return true
	}
	return false
}

func main() {

	var T = &trials{
		{"IsEven", primes.IsEven},
		{"IsOdd", primes.IsOdd},
		{"And", primes.And},
		{"Xor", primes.Xor},
		{"Or", primes.Or},
		{"ShR", primes.ShR},
		{"ShL", primes.ShL},
	}

	for _, t := range *T {
		fmt.Printf("%v\n", t.name)

		for i := 0; i < 10; i++ {
			x := uint8(i)
			fmt.Printf("  %10s(%3d) = %10v", t.name, x, t.f(x))
			fmt.Printf("    %5b  :  %3b = %5b\n", i, 1, t.f(x))
			// fmt.Printf("IsEvenAnd(%d) = %v\n", i, primes.IsEvenAnd(i))
			// fmt.Printf("IsEvenAnd(%d) = %v\n", i, primes.IsEvenAnd(i))

		}
	}
}
