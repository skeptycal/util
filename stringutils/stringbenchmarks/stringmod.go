package stringbenchmarks

import "strings"

func ToLower(s string) string {
	var sb strings.Builder
	for _, c := range []byte(s) {
		sb.WriteByte(ToLowerByte(c))
	}
	return sb.String()
}

func ToUpper(s string) string {
	var sb strings.Builder
	for _, c := range []byte(s) {
		sb.WriteByte(ToUpperByte(c))
	}
	return sb.String()
}

func toUpperNoDep(s string) string {
	var sb strings.Builder
	for _, c := range []byte(s) {
		if IsASCIIAlpha(c) {
			sb.WriteByte(c | upperMask)
			continue
		}
		sb.WriteByte(c)
	}
	return sb.String()
}

func ToLowerByte(c byte) byte {
	if IsASCIIAlpha(c) {
		return c & lowerMask
	}
	return c
}

func ToUpperByte(c byte) byte {
	if IsASCIIAlpha(c) {
		return c | upperMask
	}
	return c
}
