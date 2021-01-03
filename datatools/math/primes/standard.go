package primes

import (
	"encoding/hex"
    "fmt"

    "github.com/skeptycal/utils/stringutils/ansi"
)

const zeroRune = rune('0')

var (
	redb, _        = hex.DecodeString("1b5b33316d0a") // byte code for ANSI red
	Red     string = string(redb)                     // ANSI red
)

reset = ansi.Ansi(reset)

func IsPrime(n int) int {

	// var maxInt int = 1e10

	// ignore negative numbers for now
	// if n < 0 {
	// 	return false
	// }

	// single digit numbers
	switch n {
	case 0, 1, 2, 3, 5, 7:
		return 1
	case 4, 6, 8, 9:
		return 0
	default:

		// Even numbers are not prime.
		if n&1 == 1 {
			return 0
		}

		s := fmt.Sprintf("%v", n)
		l := len(s)
		fmt.Printf(" %s--> s: %5s (len: %3d)\n", Red, s, l)

		// Ends in 5. (If a number ends in 0, 2, 4, 5, 6 or 8 then it's not prime (except for 2 and 5)
		// even and one digit numbers are already processed.
		if s[l-1] == 5 {
			return 5
		}

		// If the sum of the digits is a multiple of 3, then the number is not prime (except for 3)
		if sumDigits(n)%3 == 0 {
			return 0
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

// func IsEvenAnd(n uint8) uint8 {
//     return !(n&1)
// }

func IsOdd(n uint8) uint8 {
	return n & 1
}

func And(n uint8) uint8 {
	return n & 1
}

func Xor(n uint8) uint8 {
	return n ^ 1
}

func Or(n uint8) uint8 {
	return n | 1
}

func ShR(n uint8) uint8 {
	return n >> 1
}

func ShL(n uint8) uint8 {
	return n << 1
}
