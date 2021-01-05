package benchmarks

func mod(n, r int) int {
	return n % r
}

// if r is a power of 2 ...
func modAnd(n, r int) int {
	return n & (r - 1)
}

/*
// if r is a power of 2 ...
func modAnd(n, r int) int {
	if r&1 == 0 {
		return n & (r - 1)
	}
	return n % r
}
*/
