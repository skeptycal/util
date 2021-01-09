// Package arrayperf performs benchmark tests on variaous array based structures.
package arrayperf

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	MaxInt16               int    = 1<<16 - 1
	MaxInt32               int    = 1<<32 - 1
	MaxInt8                int    = 1<<8 - 1
	Pool                   string = `_:$%&/()`
	DefaultIntFieldSizeMin int    = 0
	DefaultIntFieldSizeMax int    = MaxInt16
	DefaultStringFieldSize int    = 8
	DefaultArrayFieldCount int    = MaxInt8
)

type arrayParallel struct {
	Count int
	Size  int
	S     []string
	I     []int
}

func (a arrayParallel) String() string {
	sb := strings.Builder{}
	for i, s := range a.S {
		sb.WriteString(fmt.Sprintf("%d - %v  ...  %32v : %-32v\n", i, s, a.S[i], a.I[i]))
	}
	return sb.String()
}

func (a arrayParallel) Display() string {
	sb := strings.Builder{}
	for i, _ := range a.S {
		sb.WriteString(fmt.Sprintf("  %v = %v\n", a.S[i], a.I[i]))
	}
	return sb.String()
}

func MakeParallelArray(fieldCount, fieldSize int) *arrayParallel {
	if fieldSize <= 0 {
		fieldSize = DefaultStringFieldSize
	}
	if fieldCount <= 0 {
		fieldCount = DefaultArrayFieldCount
	}

	a := arrayParallel{
		Count: fieldCount,
		Size:  fieldSize,
		S:     make([]string, fieldCount),
		I:     make([]int, fieldCount),
	}

	for i := 0; i < fieldCount; i++ {
		a.S[i] = randUpper(fieldSize)
		a.I[i] = randomInt(DefaultIntFieldSizeMin, DefaultIntFieldSizeMax)
	}

	return &a
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
