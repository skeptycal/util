package stringutils

import "math"

var mathPow = math.Pow

func pow(x, y int) int {
	if x == 0 {
		return 0
	}
	if y == 0 || x == 1 {
		return 1
	}
	if y == 1 {
		return x
	}
	retval := x
	for i := 1; i < y; i++ {
		retval *= x
	}
	return retval
}

func TenToThe(y int) int {
	return pow(10, y)
}

func TwoToThe(y int) int {
	return pow(2, y)
}
