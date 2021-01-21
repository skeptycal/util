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
	// Character used for HR function
    HrChar string = "="
    // Mask to return only final nibble
    BasicMask byte = 0xF
    // Mask to return all except final nibble
    MSNibbleMask byte = 0xF0
)

var config Config = Config{
    "name": "anansi",
    "enabled": true,
    "defaultWriter": os.Stdout,
}

type Any = interface{}

type Config map[string]Any

func (c Config) Get(key string) Any {
    if v, ok := c[key]; ok {
        return v
    }
    return nil
}

func (c Config) Set(key string, v Any) {
    c[key] = v
}

func (c Config) String() string {
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
