package sieve

import (
	"bufio"
	"io"
	"strings"

	gofile "github.com/skeptycal/util/gofile/bak"
)

// OutputFormat describes the output format
type OutputFormat int

// OutputFormat constants define the choices for file format.
const (
	text OutputFormat = iota
	csv
	json
	md
)

// AllowedFormats is a list of output formats mapped to
// their respective io.Writers.
var AllowedFormats map[string]io.Writer = map[string]io.Writer{
	"text": bufio.NewWriter(nil),
	"csv":  gofile.NewCSVWriter(io.Writer, ','),
	"json": io.Writer,
	"md":   io.Writer,
}

// Set sets the output format and returns flag value.
// Choices are defined by OutputFormat constants.
func (o OutputFormat) Set(s string) (retval OutputFormat) {
	switch strings.ToLower(s) {
	case "json":
		o = json
	case "csv":
		o = csv
	default:
		o = text
	}
	return o
}

type CLIoptions map[string]interface{}
