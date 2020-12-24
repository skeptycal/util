//
/*
    Package options extends the standard library package flag to include automated setup and POSIX compliance.

	Usage

	Define flags using flag.String(), Bool(), Int(), etc.

    from package flag:
	This declares an integer flag, -n, stored in the pointer nFlag, with type *int:
		import "flag"
		var nFlag = flag.Int("n", 1234, "help message for flag n")

	After all flags are defined, call
		flag.Parse()
	to parse the command line into the defined flags.

	Flags may then be used directly. If you're using the flags themselves,
	they are all pointers; if you bind to variables, they're values.
		fmt.Println("ip has value ", *ip)
		fmt.Println("flagvar has value ", flagvar)

	After parsing, the arguments following the flags are available as the
	slice flag.Args() or individually as flag.Arg(i).
	The arguments are indexed from 0 through flag.NArg()-1.

	Command line flag syntax

	The following forms are permitted:

		-flag
		-flag=x
		-flag x  // non-boolean flags only
	One or two minus signs may be used; they are equivalent.
	The last form is not permitted for boolean flags because the
	meaning of the command
		cmd -x *
	where * is a Unix shell wildcard, will change if there is a file
	called 0, false, etc. You must use the -flag=false form to turn
	off a boolean flag.

	Flag parsing stops just before the first non-flag argument
	("-" is a non-flag argument) or after the terminator "--".

	Integer flags accept 1234, 0664, 0x1234 and may be negative.
	Boolean flags may be:
		1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False
	Duration flags accept any input valid for time.ParseDuration.

	The default set of command-line flags is controlled by
	top-level functions.  The FlagSet type allows one to define
	independent sets of flags, such as to implement subcommands
	in a command-line interface. The methods of FlagSet are
	analogous to the top-level functions for the command-line
	flag set.
*/
package options

import (
	"fmt"
	"reflect"
)

type Option interface {
	String() string
	Default() string
	TypeOf() string
}

type value struct {
	name  string
	value interface{}
}

// Value is the interface to the dynamic value stored in a flag.
// (The default value is represented as a string.)
//
// If a Value has an IsBoolFlag() bool method returning true,
// the command-line parser makes -name equivalent to -name=true
// rather than using the next command-line argument.
//
// Set is called once, in command line order, for each flag present.
// The flag package may call the String method with a zero-valued receiver,
// such as a nil pointer.
type Value interface {
	String() string
	Set(string) error
}

// Getter is an interface that allows the contents of a Value to be retrieved.
// It wraps the Value interface, rather than being part of it, because it
// appeared after Go 1 and its compatibility rules. All Value types provided
// by this package satisfy the Getter interface.
type Getter interface {
	Value
	Get() interface{}
}

func (v *value) String() string {
	return fmt.Sprintf("%v", v.value)
}

func (v *value) Set(value string) error {
	v.value = value
	return nil
}

func (v *value) Get() interface{} {
	return v.value
}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

// unknownOption represents an unknown unknownOption that
// will be parsed for usage information using the standard GNU
// option parser and other lists.
//
// They are intended to be used to create a set of options that
// conforms to the POSIX 12.1 Utility Argument Syntax[1].
//
// To create a set of options from scratch, use the NewOption
//command. This will use the more efficient cliOption struct.
//
// This structure is used to create option sets from the help
// screen of a similar utility. These options can be modified,
// automated, and put into service much more quickly than
// option sets created by hand.
//
// Once options have been loaded into an OptionSet, the default
// values and descriptions can be modified to match standard CLI
// option patterns.
//
// The zero value of most options is set using the guidelines.
// The getopt() function [2] in the System Interfaces volume of
// POSIX.1-2017 that assists utilities in handling options and
// operands that conform to these guidelines.
//
// References:
//      1. https://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap12.html#tag_12_02
//      2. https://pubs.opengroup.org/onlinepubs/9699919799/functions/getopt.html
type cliOption struct {
	Name      string
	Usage     string
	Short     string
	Long      string
	Value     Value
	valueType reflect.Type
	ZeroValue interface{}
}

func (o *cliOption) TypeOf() reflect.Type {
	if o.valueType == nil {
		// o.valueType = typeof(o.Value)
		o.valueType = reflect.TypeOf(o.Value)
	}
	return o.valueType
}

func (o *cliOption) String() string {
	return fmt.Sprintf("Command Line Option %s (%v): %s", o.Name, o.TypeOf(), o.Usage)
}

// isZeroValue determines whether the string represents the zero
// value for a flag.
func (o *cliOption) isZeroValue() bool {
	// Build a zero value of the flag's Value type, and see if the
	// result of calling its String method equals the value passed in.
	// This works unless the Value type is itself an interface type.
	typ := reflect.TypeOf(o.Value)
	var z reflect.Value
	if typ.Kind() == reflect.Ptr {
		z = reflect.New(typ.Elem())
	} else {
		z = reflect.Zero(typ)
	}
	return o.Value == z.Interface().(Value)
}

func (o *cliOption) setZeroValue() {
	if o.TypeOf().Kind() == reflect.Ptr {
		o.ZeroValue = reflect.Zero(o.TypeOf().Elem())
	}
	o.ZeroValue = reflect.Zero(o.TypeOf())
}

func (o *cliOption) getZeroValue() reflect.Value {
	if o.TypeOf().Kind() == reflect.Ptr {
		return reflect.Zero(o.TypeOf().Elem())
	}
	return reflect.Zero(o.TypeOf())
}
