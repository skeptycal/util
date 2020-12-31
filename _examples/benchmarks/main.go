package main

import (
	"fmt"
	"time"

	"github.com/skeptycal/util/datatools/algorithms/fib"
)

type node struct {
	n  int
    dt time.Duration
}

type DataSet struct {
	name  string
	f     func()
	args  []interface{}
	n     int
	avg   time.Duration
	nodes []*node
}

func (ds *DataSet) Add(n *node) {
	ds.nodes = append(ds.nodes, n)
}

func NewDataSet(name string, f func, args []interface{}, n int) *DataSet {
	return &DataSet{name, f, args, n, 0, []*node{}}
}

func test(f fn, n int) (*node, interface {}, error) {
    err error := nil
    d0 := time.Now()
    y := f(n)
    d1 := time.Now()
    dt := t1.Sub(t0)
    return &node{n,dt}, y, err
}

func Fib(n int) int {
	return fib.Fib(n)
}

func main() {

	var data *DataSet = NewDataSet("Baseline Recursive Fib", Fib, []interface{}, 10)

	var maxNum int = 100
	var t0, t1 time.Time
	var dt time.Duration

	for i := 0; i < maxNum; i++ {
		t0 := time.Now()
		fmt.Printf("%5d  - %8d\n", i, Fib(i))
		t1 := time.Now()
		dt := t1.Sub(t0)

	}
}
