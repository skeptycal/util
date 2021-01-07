package main

import (
	"fmt"
	"math"

	"github.com/yourbasic/bit"
)

func main() {

	// Sieve of Eratosthenes
	const n = 50
	sieve := bit.New().AddRange(2, n)
	sqrtN := int(math.Sqrt(n))
	for p := 2; p <= sqrtN; p = sieve.Next(p) {
		for k := p * p; k < n; k += p {
			sieve.Delete(k)
		}
	}
	fmt.Println(sieve)
}
