package bench

import (
	"fmt"
	"io"
	"time"
)

const (
	defaultMaxDuration = time.Second * 1
	timeUnits          = time.Second
)

const ()

type scinot struct {
	mantissa float64
	exponent float64
}

func (s *scinot) String() string {
	return fmt.Sprintf("%1.6f E %.2f", s.mantissa, s.exponent)
}

const (
	S = iota
	ms
	us
	ns
)

type mark struct {
	name     string
	testFunc func()
}

type benchmarkSet struct {
	name        string
	maxDuration time.Duration
	tests       []*mark
	logfile     io.Writer
}

func NewMark(name string, f func()) *mark {
	return &mark{
		name:     name,
		testFunc: f,
	}
}

func NewBenchmarkSet(name string, dur time.Duration, logfile io.Writer) *benchmarkSet {
	if dur == 0 {
		dur == defaultMaxDuration
	}
	return &benchmarkSet{
		name:        name,
		maxDuration: dur,
		tests:       make([]*mark, 0, 8),
		logfile:     logfile,
	}
}
