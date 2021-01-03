package main

import (
    "github.com/skeptycal/util/datatools/math/floatingints"
)

func WhiteSpace(c rune) bool {
	switch c {
	case ' ', '\t', '\n', '\f', '\r':
		return true
	}
	return false
}

func main() {
	pi := floatingints.float(3.14)
}
