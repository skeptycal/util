package primes

import (
	"fmt"
	"unsafe"

	"github.com/skeptycal/util/stringutils/ansi"
)

const zeroRune = rune('0')

var (
	Reset = ansi.Ansi(0) // ANSI reset
	Red   = ansi.Ansi(ansi.Red)
)

// Float32bits returns the IEEE 754 binary representation of f,
// with the sign bit of f and the result in the same bit position.
// Float32bits(Float32frombits(x)) == x.
func Float32bits(f float32) uint32 { return *(*uint32)(unsafe.Pointer(&f)) }

// Float32frombits returns the floating-point number corresponding
// to the IEEE 754 binary representation b, with the sign bit of b
// and the result in the same bit position.
// Float32frombits(Float32bits(x)) == x.
func Float32frombits(b uint32) float32 { return *(*float32)(unsafe.Pointer(&b)) }

// Float64bits returns the IEEE 754 binary representation of f,
// with the sign bit of f and the result in the same bit position,
// and Float64bits(Float64frombits(x)) == x.
func Float64bits(f float64) uint64 { return *(*uint64)(unsafe.Pointer(&f)) }

// Float64frombits returns the floating-point number corresponding
// to the IEEE 754 binary representation b, with the sign bit of b
// and the result in the same bit position.
// Float64frombits(Float64bits(x)) == x.
func Float64frombits(b uint64) float64 { return *(*float64)(unsafe.Pointer(&b)) }

// Fib calculates the nth Fibonacci number
// BenchmarkFib10-8   	 2516619	       448 ns/op	       0 B/op	       0 allocs/op
// BenchmarkFib20-8   	   19195	     54657 ns/op	       0 B/op	       0 allocs/op
//
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// IsPrime returns:
//  0 for true - number is prime
//  1 for false
//  2 for even number
//  3 for "digits sum to 3"
//  5 for "ends in 5"
//  -1 for unknown / untested / not yet implemented
//
func IsPrime(n int) int {

	// var maxInt int = 1e10

	// ignore negative numbers for now
	// if n < 0 {
	// 	return false
	// }

	// single digit numbers
	switch n {
	case 0, 1, 2, 3, 5, 7:
		return 0
	case 4, 6, 8, 9:
		return 1
	default:

		// Even numbers are not prime.
		if n&1 == 0 {
			return 2
		}

		s := fmt.Sprintf("%v", n)
		l := len(s)

		// Ends in 5. (If a number ends in 0, 2, 4, 5, 6 or 8 then it's not prime (except for 2 and 5)
		// even and one digit numbers are already processed.
		if s[l-1] == '5' {
			return 5
		}

		// If the sum of the digits is a multiple of 3, then the number is not prime (except for 3)
		if sumDigits(uint(n))%3 == 0 {
			return 3
		}

		return -1
	}
}

func sumDigits(n uint) uint {
	var rem uint = 0
	var retval uint = 0

	for n != 0 {
		rem = n & 1
		retval += rem
		n = n / 10
	}
	return retval
}

func sumDigitsStr(s string) int {
	retval := 0
	for _, n := range s {
		retval += int(n - zeroRune)
	}
	return retval
}

func IsEven(n uint8) uint8 {
	for n%2 == 0 {
		return 1
	}
	return 0
}

func IsEvenAnd(n uint8) uint8 {
	if n&1 == 0 {
		return 0
	}
	return 1
}

func IsOdd(n uint8) uint8 {
	return n & 1
}

func And(n, m uint8) uint8 {
	return n & m
}

func Xor(n, m uint8) uint8 {
	return n ^ m
}

func Or(n, m uint8) uint8 {
	return n | m
}

func ShR(n, m uint8) uint8 {
	return n >> m
}

func ShL(n, m uint8) uint8 {
	return n << m
}
