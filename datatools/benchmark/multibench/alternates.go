package multibench

import "strings"

// AcceptedAnswer returns the "expected" answer for this problem.
// All of the other functions are pursuing alternate methods
// but should give the same answer.
//
// This represents the commonly accepted algorithm.
func AcceptedAnswer(x, y int) int {
	return x * y
}

func Multi2(x, y int) int {
	retval := 0
	for i := 0; i < y; i++ {
		retval += x
	}
	return retval
}
func Multi3(x, y int) int {
	retval := 0
	i := 0
	for i < y {
		i--

	}
	return retval
}

func Multi4() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func Multi5(x, y int) int {
	a := strings.Repeat(".", x)
	b := strings.Repeat(a, y)
	return len(b)
}

func Multi6(x, y int) int {
	a := "asdf"
	b := a
	// i := make([x]int, y, y)
	return len(b)
}

func Multi7(x, y int) int {

	// i := make([x]int, y, y)
	return 0
}

// func Multi8(x, y int) int {
// 	i := 0
// 	chint := make(chan int, 0)
// 	for i < y {
// 		chint <- x
// 	}
// 	return chint
// }

// StandardMulti performs math like a gradeschool kid is taught.
/*e.g.

     3498
   x 1842
  -------

       8 x 2      ------>        16
      90 x 2      ------>       180
     400 x 2      ------>       800
    3000 x 2      ------>      6000
                       |-------->         6996
       8 x 40     ------>       320
      90 x 40     ------>      3600
     400 x 40     ------>     16000
    3000 x 40     ------>    120000
                                        139920
       [...]                  [...]
                3498 x 800 ------>     2798400
       [...]                  [...]
               3498 x 1000 ------>     3498000    =======>  6443316
*/
// func StandardMulti(x, y int) int {

// 	if x > 9 {
// 		x = numSplit(x)
// 	}
// 	for number != 0 {
// 		remainder = number % 10
// 		sumResult += remainder
// 		number = number / 10
// 	}
// 	for y % 10 {

// 	}
// }
