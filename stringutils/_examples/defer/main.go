// This an example of the defer statement in action.
/*
The arguments to the deferred function (which include the receiver if the function is a method) are evaluated when the defer executes, not when the call executes. Besides avoiding worries about variables changing values as the function executes, this means that a single deferred call site can defer multiple function executions. Here's a silly example.

    for i := 0; i < 5; i++ {
        defer fmt.Printf("%d ", i)
    }

Deferred functions are executed in LIFO order, so this code will cause 4 3 2 1 0 to be printed when the function returns. A more plausible example is a simple way to trace function execution through the program. We could write a couple of simple tracing routines like this:

    func trace(s string)   { fmt.Println("entering:", s) }
    func untrace(s string) { fmt.Println("leaving:", s) }

    // Use them like this:
    func a() {
        trace("a")
        defer untrace("a")
        // do something....
    }

We can do better by exploiting the fact that arguments to deferred functions are evaluated when the defer executes. The tracing routine can set up the argument to the untracing routine. This example:

    func trace(s string) string {
        fmt.Println("entering:", s)
        return s
    }

    func un(s string) {
        fmt.Println("leaving:", s)
    }

    func a() {
        defer un(trace("a"))
        fmt.Println("in a")
    }

    func b() {
        defer un(trace("b"))
        fmt.Println("in b")
        a()
    }

    func main() {
        b()
    }

prints

    entering: b
    in b
    entering: a
    in a
    leaving: a
    leaving: b

For programmers accustomed to block-level resource management from other languages, defer may seem peculiar, but its most interesting and powerful applications come precisely from the fact that it's not block-based but function-based. In the section on panic and recover we'll see another example of its possibilities.

...
*/
package main

import (
	"fmt"

	"github.com/skeptycal/util/stringutils/ansi"
)

const (
	ansiColor = "\033[38;5;%dm" // set ANSI foreground color to code %d using printf
	ansiReset = "\033[39;49;0m" // reset ANSI terminal output to default foreground and background colors
)

func ansiString(i int) string {
	return fmt.Sprintf(ansiColor, i)
}

func Cprint(i int, args ...interface{}) {
	fmt.Print(ansiString(i))
	fmt.Print(args...)
	fmt.Print(ansiReset)
}

func Cprintln(i int, args ...interface{}) {
	fmt.Print(ansiString(i))
	fmt.Print(args...)
	fmt.Println(ansiReset)
}

func main() {
	fmt.Println("")
	ansi.Cprintln(83, "Example of 'defer' statement.")
	Cprintln(83, "-----------------------------")
	fmt.Println("")
	Cprintln(35, "This code contains a loop that counts *UP* from 0 to 500.")
	Cprintln(35, "- Within the loop, the loop counter is printed in the matching ANSI color using 'defer print ...'")
	Cprintln(35, "- shows the 'reversing' effect of the defer statement. ")
	Cprintln(35, "- ANSI color codes above 255 become x mod 255 ... only the LSB is used.")
	fmt.Println("")
	for i := 0; i < 500; i++ {
		defer Cprint(i, i, " ")
	}
	fmt.Println("")
}
