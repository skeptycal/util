package main

import (
	"fmt"
	"strings"
	"time"
)

const Mu = "µ"

func main() {
	n := 1000000
	original := "cat and dog"
	// Create Replacer (excluded from benchmark time).
	r := strings.NewReplacer("cat", "car",
		"and", "or",
		"dog", "truck")

	t0 := time.Now()

	// Version 1: use Replacer.
	for i := 0; i < n; i++ {
		result := r.Replace(original)
		if result != "car or truck" {
			fmt.Println(0)
		}
	}

	t1 := time.Now()

	// Version 2: use Replace calls.
	for i := 0; i < n; i++ {
		temp := strings.Replace(original, "cat", "car", -1)
		temp = strings.Replace(temp, "and", "or", -1)
		result := strings.Replace(temp, "dog", "truck", -1)
		if result != "car or truck" {
			fmt.Println(0)
		}
	}

	t2 := time.Now()

	// Results.

	v1 := t1.Sub(t0).Microseconds()
	v2 := t2.Sub(t1).Microseconds()

	fmt.Printf("Using Replacer: %5d %s\n", v1, Mu)
	fmt.Printf("Using Replace:  %5d %s\n", v2, Mu)
}
