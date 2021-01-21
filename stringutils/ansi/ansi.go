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
// There is an enabled flag with interface methods Enable and Disable.
//
// There is a map for named settings of any type.
//
type Config struct {
    enable bool
    ansi AnsiSet
    enabled bool
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
