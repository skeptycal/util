package options

import (
	"io"
	"os"
)

var CommandLine = NewFlagSet(os.Args[0], ExitOnError)

// NewOptionSet returns a new, empty option set with the specified
// name and error handling property. The default name is the name
// of the command line program that was invoked. The default error
// handling is ContinueOnError.
func NewFlagSet(name string, errorHandling ErrorHandling) *OptionSet {
	f := &OptionSet{
		name:          name,
		errorHandling: errorHandling,
	}
	f.Usage = f.defaultUsage
	return f
}

// An OptionSet represents a set of defined options. The zero value of an OptionSet
// has no name and has ContinueOnError error handling.
//
// Flag names must be unique within a FlagSet. An attempt to define a flag whose
// name is already in use will cause a panic.
type OptionSet struct {
	// Usage is the function called when an error occurs while parsing flags.
	// The field is a function (not a method) that may be changed to point to
	// a custom error handler. What happens after Usage is called depends
	// on the ErrorHandling setting; for the command line, this defaults
	// to ExitOnError, which exits the program after calling Usage.
	Usage func()

	name          string
	parsed        bool
	option        map[string]*Option
	formal        map[string]*Option
	args          []string // arguments after flags
	errorHandling ErrorHandling
	output        io.Writer // nil means stderr; use Output() accessor
}

func (o OptionSet) Get(key string) interface{} {
	opt, ok := o[key]
	if !ok {
		return nil
	}
	return opt.Value
}

// NewFlagSet returns a new, empty flag set with the specified name and
// error handling property. If the name is not empty, it will be printed
// in the default usage message and in error messages.
func NewOptionSet(name string, errorHandling ErrorHandling) *OptionSet {
	o := &OptionSet{
		name:          name,
		errorHandling: errorHandling,
	}
	o.Usage = o.defaultUsage
	return o
}

// Output returns the destination for usage and error messages. os.Stderr is returned if
// output was not set or was set to nil.
func (f *OptionSet) Output() io.Writer {
	if f.output == nil {
		return os.Stderr
	}
	return f.output
}

// Name returns the name of the flag set.
func (f *OptionSet) Name() string {
	return f.name
}

// ErrorHandling returns the error handling behavior of the flag set.
func (f *OptionSet) ErrorHandling() ErrorHandling {
	return f.errorHandling
}

// SetOutput sets the destination for usage and error messages.
// If output is nil, os.Stderr is used.
func (f *OptionSet) SetOutput(output io.Writer) {
	f.output = output
}

// VisitAll visits the flags in lexicographical order, calling fn for each.
// It visits all flags, even those not set.
func (f *OptionSet) VisitAll(fn func(*Option)) {
	for _, flag := range sortFlags(f.formal) {
		fn(flag)
	}
}
