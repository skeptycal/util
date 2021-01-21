package ansi

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

// ioMutex is a reader/writer mutual exclusion lock.
// The lock can be held by an arbitrary number of readers OR a single writer.
// The zero value for a RWMutex is an unlocked mutex.
//
// A RWMutex must not be copied after first use.
//
// If a goroutine holds a RWMutex for reading and another goroutine might
// call Lock, no goroutine should expect to be able to acquire a read lock
// until the initial read lock is released. In particular, this prohibits
// recursive read locking. This is to ensure that the lock eventually becomes
// available; a blocked Lock call excludes new readers from acquiring the
// lock.
//
// Reference: /go/src/sync/rwmutex.go (go standard library)
type ioMutex = sync.RWMutex

var defaultconfig ConfigMap = ConfigMap{
    "name": "ansi",
    "enabled": true,
    "defaultWriter": os.Stdout,
}
var (
    DefaultAnsiSet  = NewAnsiSet(StyleNormal, White, Black,Normal)
)

type Any = interface{}
type ConfigMap map[string]Any

// Config represents a configuration structure that can be used to configure
// a variety of objects.
//
// There is an 'enabled' flag with interface methods Enable and Disable.
// By default, there is an ANSI cli interface AnsiWriter that can be used
// to easily write console and terminal output.
//
// There is a built in RWmutex to allow concurrent reading or writing of
// settings. and corresponding interface methods Lock() and Unlock().
//
// There is a map for named settings of any type as well as the standard
// Get() and Set() interface methods.
//
// The default String() interface method pretty prints the configuration.
//
type Config struct {
    enabled bool
    ansi AnsiSet
    settings ConfigMap
    locker *ioMutex
}

func (o *Config) Disable() {
    o.Lock(); defer o.Unlock()
    o.enabled = false
}
func (o *Config) Enable() {
    o.Lock(); defer o.Unlock()
    o.enabled = true
}
func (o *Config) Lock() {
    o.locker.Lock()
}
func (o *Config) Unlock() {
    o.locker.Unlock()
}
func (o *Config) Get(key string) Any {
    o.Lock(); defer o.Unlock()
    if v, ok := o.settings[key]; ok {
        return v
    }
    return nil
}
func (o *Config) Set(key string, v Any) {
    o.Lock(); defer o.Unlock()
    o.settings[key] = v
}
func (o *Config) String() string {
    o.Lock(); defer o.Unlock()

    if len(o.settings) < 1 {
        return "Empty Config."
    }

    sb := strings.Builder{}
    defer sb.Reset()

    // TODO - limit number of lines returned? or use Less format?
    for k, v := range o.settings {
        sb.WriteString(fmt.Sprintf(" %s: %v\n",k,v))
    }
    sb.WriteString("\n")
    return sb.String()
}
