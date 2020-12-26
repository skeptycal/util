package main

import (
	"fmt"
	"time"

	"github.com/skeptycal/util/devtools/cgo/rand"
)

var (
	maxNum   int = 10
	timeSeed int = time.Now().Nanosecond()

	r, c int
)

func main() {
	c = 0
	rand.Seed(timeSeed)

	data := map[int]map[int]int{}

	for i := 0; i < maxNum; i++ {
		jmap := map[int]int{}
		for j := 0; j < maxNum; j++ {
			r = rand.Random()
			// fmt.Println("Data: ", data)
			jmap[j] = r
			fmt.Printf("Data [%d][%d] = [%d]\n", i, j, r)
		}
		data[i] = jmap
	}
	fmt.Println(data)

}
