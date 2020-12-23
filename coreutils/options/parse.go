package options

import (
	"fmt"
	"os"
)

// Parse parses the command-line flags from os.Args[1:]. Must be called
// after all flags are defined and before flags are accessed by the program.
func Parse() {
	// Ignore errors; CommandLine is set for ExitOnError.
	CommandLine.Parse(os.Args[1:])
}

func (o *OptionSet) Parse(args []string) error {
	o.parsed = true
	o.args = args
	for {
		seen, err := o.parseArg()
		if seen {
			continue
		}
		if err == nil {
			break
		}
		switch o.errorHandling {
		case ContinueOnError:
			return err
		case ExitOnError:
			if err == ErrHelp {
				os.Exit(0)
			}
			os.Exit(2)
		case PanicOnError:
			panic(err)
		}
	}

	return nil
}

// parseArg parses one option and reports whether an option was seen.
func (o *OptionSet) parseArg() (bool, error) {
	// if len(o.args) == 0 {
	// 	return false, nil
	// }
	// s := o.args[0]

	// if len(s) < 2 || s[0] != '-' {
	// 	return false, nil
	// }
	// numMinuses := 1
	// if s[1] == '-' {
	// 	numMinuses++
	// 	if len(s) == 2 { // "--" terminates the flags
	// 		o.args = o.args[1:]
	// 		return false, nil
	// 	}
	// }
	// name := s[numMinuses:]
	// if len(name) == 0 || name[0] == '-' || name[0] == '=' {
	// 	return false, o.failf("bad option syntax: %s", s)
	// }

	// // it's an option. does it have an argument?
	// o.args = o.args[1:]
	// hasValue := false
	// value := ""
	// for i := 1; i < len(name); i++ { // equals cannot be first
	// 	if name[i] == '=' {
	// 		value = name[i+1:]
	// 		hasValue = true
	// 		name = name[0:i]
	// 		break
	// 	}
	// }
	// m := o.formal
	// option, alreadythere := m[name] // BUG
	// if !alreadythere {
	// 	if name == "help" || name == "h" { // special case for nice help message.
	// 		o.Usage()
	// 		return false, ErrHelp
	// 	}
	// 	return false, o.failf("option provided but not defined: -%s", name)
	// }

	// if option.TypeOf().Kind() == reflect.Bool { // special case: doesn't need an arg
	// 	if hasValue {
	// 		if err := .Set(value); err != nil {
	// 			return false, o.failf("invalid boolean value %q for -%s: %v", value, name, err)
	// 		}
	// 	} else {
	// 		if err := fv.Set("true"); err != nil {
	// 			return false, o.failf("invalid boolean option %s: %v", name, err)
	// 		}
	// 	}
	// } else {
	// 	// It must have a value, which might be the next argument.
	// 	if !hasValue && len(o.args) > 0 {
	// 		// value is the next arg
	// 		hasValue = true
	// 		value, o.args = o.args[0], o.args[1:]
	// 	}
	// 	if !hasValue {
	// 		return false, o.failf("option needs an argument: -%s", name)
	// 	}
	// 	if err := option.Value.Set(value); err != nil {
	// 		return false, o.failf("invalid value %q for option -%s: %v", value, name, err)
	// 	}
	// }
	// if o.option == nil {
	// 	// 	o.actual = make(map[string]*Option)
	// 	// }
	// 	o.option[name] = option
	return true, nil
	// }
}

// failf prints to standard error a formatted error and usage message and
// returns the error.
func (o *OptionSet) failf(format string, args ...interface{}) error {
	err := fmt.Errorf(format, args...)
	fmt.Fprintln(o.Output(), err)
	o.Usage()
	return err
}
