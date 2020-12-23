package options

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
)

var CommandLine = NewFlagSet(os.Args[0], ExitOnError)

func defaultUsage() {
	fmt.Printf("%v\n", os.Args[0])
	flag.Parse()

}

// NewOptionSet returns a new, empty option set with the specified
// name and error handling property. The default name is the name
// of the command line program that was invoked. The default error
// handling is ContinueOnError.
func NewFlagSet(name string, errorHandling ErrorHandling) *OptionSet {
	f := &OptionSet{
		name:          name,
		errorHandling: errorHandling,
	}
	f.Usage = defaultUsage
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
	option        map[string]*cliOption
	formal        map[string]*cliOption
	args          []string // arguments after flags
	errorHandling ErrorHandling
	output        io.Writer // nil means stderr; use Output() accessor
}

func (o OptionSet) Get(key string) interface{} {
	opt, ok := o.option[key]
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
	o.Usage = defaultUsage
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
func (f *OptionSet) VisitAll(fn func(*cliOption)) {
	for _, flag := range sortFlags(f.formal) {
		fn(flag)
	}
}

// Visit visits the flags in lexicographical order, calling fn for each.
// It visits only those flags that have been set.
func (f *OptionSet) Visit(fn func(*cliOption)) {
	for _, flag := range sortFlags(f.option) {
		fn(flag)
	}
}

// sortFlags returns the flags as a slice in lexicographical sorted order.
func sortFlags(flags map[string]*cliOption) []*cliOption {
	result := make([]*cliOption, len(flags))
	i := 0
	for _, f := range flags {
		result[i] = f
		i++
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})
	return result
}
