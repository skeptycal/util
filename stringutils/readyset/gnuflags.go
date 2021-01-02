package readyset

import (
	"time"

	"golang.org/x/text/language"
)

// flags that follow GNU CLI specification

// POSIX Program Argument Syntax Conventions
/* const posixflags = `POSIX recommends these conventions for command line arguments. getopt (see Getopt) and argp_parse (see Argp) make it easy to implement them.

Arguments are options if they begin with a hyphen delimiter (‘-’).
Multiple options may follow a hyphen delimiter in a single token if the options do not take arguments. Thus, ‘-abc’ is equivalent to ‘-a -b -c’.
Option names are single alphanumeric characters (as for isalnum; see Classification of Characters).
Certain options require an argument. For example, the ‘-o’ command of the ld command requires an argument—an output file name.
An option and its argument may or may not appear as separate tokens. (In other words, the whitespace separating them is optional.) Thus, ‘-o foo’ and ‘-ofoo’ are equivalent.
Options typically precede other non-option arguments.
The implementations of getopt and argp_parse in the GNU C Library normally make it appear as if all the option arguments were specified before all the non-option arguments for the purposes of parsing, even if the user of your program intermixed option and non-option arguments. They do this by reordering the elements of the argv array. This behavior is nonstandard; if you want to suppress it, define the _POSIX_OPTION_ORDER environment variable. See Standard Environment.

The argument ‘--’ terminates all options; any following arguments are treated as non-option arguments, even if they begin with a hyphen.
A token consisting of a single hyphen character is interpreted as an ordinary non-option argument. By convention, it is used to specify input from or output to the standard input and output streams.
Options may be supplied in any order, or appear multiple times. The interpretation is left up to the particular application program.
GNU adds long options to these conventions. Long options consist of ‘--’ followed by a name made of alphanumeric characters and dashes. Option names are typically one to three words long, with hyphens to separate words. Users can abbreviate the option names as long as the abbreviations are unique.

To specify an argument for a long option, write ‘--name=value’. This syntax enables a long option to accept an argument that is itself optional.

Eventually, GNU systems will provide completion for long option names in the shell.`
*/

// POSIX Standard Environment Variables
type ENV struct {
	// These environment variables have standard meanings. This doesn’t mean that they are always present in the environment; but if these variables are present, they have these meanings. You shouldn’t try to use these environment variable names for some other purpose.

	HOME string
	// This is a string representing the user’s home directory, or initial default working directory.

	// The user can set HOME to any value. If you need to make sure to obtain the proper home directory for a particular user, you should not use HOME; instead, look up the user’s name in the user database (see User Database).

	// For most purposes, it is better to use HOME, precisely because this lets the user specify the value.

	LOGNAME string
	// This is the name that the user used to log in. Since the value in the environment can be tweaked arbitrarily, this is not a reliable way to identify the user who is running a program; a function like getlogin (see Who Logged In) is better for that purpose.

	// For most purposes, it is better to use LOGNAME, precisely because this lets the user specify the value.

	PATH string
	// A path is a sequence of directory names which is used for searching for a file. The variable PATH holds a path used for searching for programs to be run.

	// The execlp and execvp functions (see Executing a File) use this environment variable, as do many shells and other utilities which are implemented in terms of those functions.

	// The syntax of a path is a sequence of directory names separated by colons. An empty string instead of a directory name stands for the current directory (see Working Directory).

	// A typical value for this environment variable might be a string like:

	// :/bin:/etc:/usr/bin:/usr/new/X11:/usr/new:/usr/local/bin
	// This means that if the user tries to execute a program named foo, the system will look for files named foo, /bin/foo, /etc/foo, and so on. The first of these files that exists is the one that is executed.

	TERM string
	// This specifies the kind of terminal that is receiving program output. Some programs can make use of this information to take advantage of special escape sequences or terminal modes supported by particular kinds of terminals. Many programs which use the termcap library (see Find in The Termcap Library Manual) use the TERM environment variable, for example.

	TZ time.Location
	// This specifies the time zone. See TZ Variable, for information about the format of this string and how it is used.

	LANG string
	// This specifies the default locale to use for attribute categories where neither LC_ALL nor the specific environment variable for that category is set. See Locales, for more information about locales.

	LC_ALL string
	// If this environment variable is set it overrides the selection for all the locales done using the other LC_* environment variables. The value of the other LC_* environment variables is simply ignored in this case.

	LC_COLLATE string
	// This specifies what locale to use for string sorting.

	LC_CTYPE language.CanonType
	// This specifies what locale to use for character sets and character classification.

	// LC_MESSAGES
	// This specifies what locale to use for printing messages and to parse responses.

	// LC_MONETARY
	// This specifies what locale to use for formatting monetary values.

	// LC_NUMERIC
	// This specifies what locale to use for formatting numbers.

	// LC_TIME
	// This specifies what locale to use for formatting date/time values.

	// NLSPATH
	// This specifies the directories in which the catopen function looks for message translation catalogs.

	// _POSIX_OPTION_ORDER
	// If this environment variable is defined, it suppresses the usual reordering of command line arguments by getopt and argp_parse. See Argument Syntax.
}
