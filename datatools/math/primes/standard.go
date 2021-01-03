package primes

import (
	"fmt"

	"github.com/skeptycal/util/stringutils/ansi"
)

const zeroRune = rune('0')

var (
	Reset = ansi.Ansi(0) // ANSI reset
	Red   = ansi.Ansi(ansi.Red)
)

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
//  5 for "ends in 5"
//  3 for "digits sum to 3"
//  -1 for unknown / untested
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
		if sumDigits(n)%3 == 0 {
			return 3
		}

		return -1
	}
}

func sumDigits(number int) int {
	rem := 0
	retval := 0
	for number != 0 {
		rem = number % 10
		retval += rem
		number = number / 10
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
