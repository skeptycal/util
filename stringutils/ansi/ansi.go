// Copyright (c) 2020 Michael Treanor
// MIT License

// Package ansi provides fast ansi escape sequence processing based on strings.Builder.
// The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"
package ansi

import (
	"fmt"
	"os"
	"strings"
)


const (
    DefaultAll     string = "\033[39;49;0m"
	DefaultText    string = "\033[22;39m" // Normal text color and intensity
	Reset          string = "\033[0m"     // Turn off all attributes
)


const (
	// Character used for HR function
    HrChar string = "="
    // Mask to return only final nibble
    BasicMask byte = 0xF
    // Mask to return all except final nibble
    MSNibbleMask byte = 0xF0
)

type Config struct {
    enable bool
    ansi AnsiSet
    enabled bool
    configLock ioMutex
}


type OnOff struct {
    config Config
    locker *ioMutex
}

func (o *OnOff) Disable() {
    o.locker.Lock()
    defer o.locker.Unlock()
    o.enabled = false
}

func (o *OnOff) Enable() {
    o.locker.Lock()
    defer o.locker.Unlock()
    o.enabled = true
}

func (o *OnOff) Toggle() {
    o.locker.RLocker().Lock()
    defer o.locker.Unlock()
    o.enabled = !o.enabled
}



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
func (c ConfigMap) Get(key string) Any {
    if v, ok := c[key]; ok {
        return v
    }
    return nil
}
func (c ConfigMap) Set(key string, v Any) {
    c[key] = v
}
func (c ConfigMap) String() string {
    if len(c) < 1 {
        return "Empty Config."
    }

    sb := strings.Builder{}
    defer sb.Reset()

    // TODO - limit number of lines returned? or use Less format?
    for k, v := range c {
        sb.WriteString(fmt.Sprintf(" %s: %v\n",k,v))
    }
    sb.WriteString("\n")
    return sb.String()
}
