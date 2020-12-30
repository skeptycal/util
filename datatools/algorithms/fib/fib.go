package fib

// Fib returns the n-th number in the Fibonacci sequence.
// This version uses recursion and will be limited in
// effectiveness. It is a baseline implementation used for
// comparing other benchmarked methods.
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
