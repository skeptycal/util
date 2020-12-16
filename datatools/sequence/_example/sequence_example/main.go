package main

import (
	"crypto/rand"
	"math/big"

	"github.com/skeptycal/util/data/sequence"
)

func main() {
	size := 100
	sequence := make(sequence.Sequence, 0, size)

	println(sequence.Len())

	for i := 0; i < size; i++ {
		r = rand.Int(rand., max*big.Int)
		sequence = append(sequence, i)
	}

	println(sequence.Len())
	println(sequence.String())
}
