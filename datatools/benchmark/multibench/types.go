// Package multibench provides benchmark comparisons between different
// algorithms used to solve the same problem. ** Experimental **
package multibench

import (
	"fmt"
	"os"
)

var (
	i = make(chan int)       // by default the capacity is 0
	s = make(chan string, 3) // non-zero capacity

	r = make(<-chan bool)          // can only read from
	w = make(chan<- []os.FileInfo) // can only write to
)

// Func is a function header showing arguments and returns that the test functions must match
// i.e. it is the test interface
type Func func(x, y int) int

type Functioning interface {
	Got() int
	Want() int
	Name() string
}

type funcTest struct {
	x int
	y int
	f Func
}

func NewFunc(x, y int, f Func) Functioning {
	return &funcTest{x, y, f}
}

func (f funcTest) Name() string {
	return fmt.Sprintf("%v", f.f)
}

// This is the result of the commonly accepted algorithm
func (f funcTest) Want() int {
	return AcceptedAnswer(f.x, f.y)
}

// This is the algorithm being tested
func (f funcTest) Got() int {
	return f.f(f.x, f.y)
}
