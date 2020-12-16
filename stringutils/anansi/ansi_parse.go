package anansi

import "io"

type AnsiWriter interface {
	io.WriterTo
	io.StringWriter
}
