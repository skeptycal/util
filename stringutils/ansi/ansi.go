// Copyright (c) 2020 Michael Treanor
// MIT License

// Package ansi provides fast ansi escape sequence processing based on strings.Builder.
// The standard is defined by the ECMA-48 standard "Control Functions for Coded Character Sets - Fifth Edition"
package ansi

const (
	// Character used for HR function
    HrChar string = "="
    // Mask to return only final nibble
    BasicMask byte = 0xF
    // Mask to return all except final nibble
    MSNibbleMask byte = 0xF0
)
