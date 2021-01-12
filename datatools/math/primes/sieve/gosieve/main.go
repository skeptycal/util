// A concurrent prime sieve

package main

import "fmt"

var last int = 0

// Generate - Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Filter - Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	// maxnum is the number of primes to calculate.
	maxnum := 1000
	// 	percentInterval is the percentage of primes that are listed.
	percentInterval := 5.0
	seriesStart := 1
	seriesStop := maxnum

	statusInterval := int(float64(maxnum) * percentInterval / 100.0)

	maxnumLengthFmt := fmt.Sprintf("%v", 5)
	maxnumLengthFmtLen := len(maxnumLengthFmt)

	// for maxnum ==5, generates " %5d : %-d\n"
	maxLenStrFmt := fmt.Sprintf(" %%%dd : %%-d\n", maxnumLengthFmtLen)

	ch := make(chan int) // Create a new channel.
	go Generate(ch)      // Launch Generate goroutine.
	for i := seriesStart; i < seriesStop; i++ {
		prime := <-ch
		if (i+1)%statusInterval == 0 {
			fmt.Printf(maxLenStrFmt, i, prime)
		}
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
		if i > seriesStop {
			last = prime
			fmt.Printf(maxLenStrFmt, i, prime)
		}
	}
	fmt.Printf("That was %.0f%% of the first %d prime numbers.\n", percentInterval, maxnum)
}
