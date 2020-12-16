package replace

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
)

var LOG = log.New(os.Stdout, "DEBUG mutex", log.Lmicroseconds)

func logTrace(msg string) {
	var pc = make([]uintptr, 10)
	no := runtime.Callers(4, pc)
	frames := runtime.CallersFrames(pc)

	var callers []string

	for i := 0; i < no; i++ {
		frame, more := frames.Next()
		callers = append(callers, fmt.Sprintf("%s:%d %s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}

	LOG.Printf(msg + "\n\t" + strings.Join(callers, "\n\t"))
}

type Mutex interface {
	Lock()
	Unlock()
}

type RWMutex interface {
	Mutex
	RLock()
	RUnlock()
}

func NewMutex(name string, debug bool) Mutex {
	if !debug {
		return &sync.Mutex{}
	}
	return &mutex{Name: name}
}

func NewRWMutex(name string, debug bool) RWMutex {
	if !debug {
		return &sync.RWMutex{}
	}
	return &rwmutex{Name: name}
}

type mutex struct {
	Name string
	mx   sync.Mutex
}

func (m *mutex) logTracing(fn string) {
	logTrace(fmt.Sprintf("Mutex %q %s called", m.Name, fn))
}

func (m *mutex) Lock() {
	m.logTracing("Lock")
	m.mx.Lock()
}

func (m *mutex) Unlock() {
	m.logTracing("Unlock")
	m.mx.Unlock()
}

func (m *rwmutex) logTracing(fn string) {
	logTrace(fmt.Sprintf("RWMutex %q %s called", m.Name, fn))
}

type rwmutex struct {
	Name string
	mx   sync.RWMutex
}

func (m *rwmutex) RLock() {
	m.logTracing("RLock")
	m.mx.RLock()
}

func (m *rwmutex) Lock() {
	m.logTracing("Lock")
	m.mx.Lock()
}

func (m *rwmutex) Unlock() {
	m.logTracing("Unlock")
	m.mx.Unlock()
}

func (m *rwmutex) RUnlock() {
	m.logTracing("RUnlock")
	m.mx.RUnlock()
}
