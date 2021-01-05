package ini

import (
	"github.com/skeptycal/util/gofile"
)

// RemoveComments returns the contents of a text file with
// comments removed.
//
// Comments are defined as lines of text that begin with a
// comment string. Leading whitespace characters are  ignored,
// and optionally removed. The default implementation uses
// the following list of comment strings:
//
//  "#", "//", ";"
//
// Other options include:
//  removeWhiteSpace - strip all leading and trailing whitespace
//  remove trailing comments - strip comments at end of lines
//
func RemoveComments(file string) string {
	gofile.ReadFile(file)

}
