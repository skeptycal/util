// Package arrayperf performs benchmark tests on variaous array based structures.
package arrayperf

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	pool              string = `_:$%&/()`
	defaultStringSize int    = 32
	defaultMinIntSize int    = 1
	defaultInt16      int    = 1<<16 - 1
	defaultInt32      int    = 1<<32 - 1
	defaultInt8       int    = 1<<8 - 1
)

type array struct {
	Len  int
	Size int
	S    []string
	I    []int
}

func (a array) String() string {
	sb := strings.Builder{}
	for i, s := range a.S {
		sb.WriteString(fmt.Sprintf("%d - %v  ...  %32v : %-32v\n", i, s, a.S[i], a.I[i]))
	}
	return sb.String()
}

func (a array) Display() string {
	sb := strings.Builder{}
	for i, _ := range a.S {
		sb.WriteString(fmt.Sprintf("  %v = %v\n", a.S[i], a.I[i]))
	}
	return sb.String()
}

func MakeArray(count, size int) *array {
	if size == 0 {
		size = defaultStringSize
	}
	if count == 0 {
		count = defaultInt8
	}
	ss := make([]string, count)
	ii := make([]int, count)

	// a := array{
	//     s: [n]string "",
	//     i: [n]0,
	// }
	for i := 0; i < count; i++ {
		ss[i] = randUpper(size)
		ii[i] = randomInt(defaultMinIntSize, defaultInt32)
	}
	return &array{
		Len:  count,
		Size: size,
		S:    ss,
		I:    ii,
	}
}

// Returns an int >= min, < max
func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func randUpper(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(65, 90))
	}
	return string(bytes)
}

// Generate a random string of A-Z chars with len = l
func randomLower(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		bytes[i] = byte(randomInt(97, 122))
	}
	return string(bytes)
}

// Generate a random string of A-Z chars with len = l
func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = pool[rand.Intn(len(pool))]
	}
	return string(bytes)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
