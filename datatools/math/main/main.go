package main

func WhiteSpace(c rune) bool {
	switch c {
	case ' ', '\t', '\n', '\f', '\r':
		return true
	}
	return false
}

func main() {
	pi := float.Float(3.14)
}
