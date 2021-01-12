package multibench

import "fmt"

// convert types take an int and return a string value.
type convert func(int) string

// ToBin implements convert, returning the binary representation of an int.
func ToBin(x int) string {
	return fmt.Sprintf("%b", x)
}

// value implements convert, returning x as string.
func value(x int) string {
	return fmt.Sprintf("%v", x)
}

// Quoted implements convert, returning x as quoted string.
func Quoted(x int) string {
	return fmt.Sprintf("%q", x)
}

// ListDigits returns a list of the digits from an integer
// in order of powers of 10 from greatest to least.
//
// e.g.
// ListDigits(1234) = {4, 3, 2, 1}
func ListDigits(number int) []int {
	if number == 0 {
		return []int{0}
	}
	var negFlag bool = false
	var neg int
	if number < 0 {
		negFlag = true
		number = -number
	}
	retval := make([]int, 0, 32)
	// remainder := 0

	for number != 0 {
		// remainder = number % 10
		retval = append(retval, number%10)
		number = number / 10
	}

	if negFlag {
		neg = 0 - retval[len(retval)-1]
		retval[len(retval)-1] = neg
	}
	return retval
}

func Reverse(lst []string) chan string {
	ret := make(chan string)
	n := len(lst) - 1
	go func() {
		for i := range lst {
			ret <- lst[n-i]
		}
		close(ret)
	}()
	return ret
}

func sumDigits(number int) int {
	remainder := 0
	sumResult := 0
	for number != 0 {
		remainder = number % 10
		sumResult += remainder
		number = number / 10
	}
	return sumResult
}

func recursiveSumDigits(number int) int {
	sumResult := 0
	if number == 0 {
		return 0
	}
	sumResult = number%10 + recursiveSumDigits(number/10)
	return sumResult
}

func convertTests() {
	var result string

	result = value(123)
	fmt.Println(result)
	// Output: 123

	result = ToBin(8190)
	fmt.Println(result)

	_ = convert(value) // confirm foo satisfies convert at runtime

	// fails due to argument type
	// _ = convert(func(x float64) string { return "" })

	// _ = convert(result)
}
