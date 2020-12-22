

Reference: http://www.maizure.org/projects/decoded-gnu-coreutils/

## Helpful background for code reading

The GNU coreutils has its foibles. Many of these utilities are approaching 30 years old and include revisions by many people over the years. Here are some things to keep in mind when reading the code:
- **Tiny programs** - These utilities are small, (mostly) single-source file programs designed to do one thing and do it well. They are not designed for long life or to scale beyond their role. Consequently, we see designs often considered 'bad practice' such - as:
  - Many globals
  - Liberal use of macros
  - goto statements
  - Long functions with nested switchs/loops

- **Know POSIX**- Start with the [Utility Syntax Guidelines][guidelines]. In general, POSIX supports interoperability by defining appropriate inputs and outputs, but leaves the 'work' to the implementation. While the GNU coreutils [may not strictly conform][maybe-not] to POSIX, many ideas are entrenched: permission bits, uids/gids, environment variables, - exit status, and about [3718 pages][3718] of more trivia.

- **Outside help** - Portability is a complex problem and coreutils relies on extra help from a related project: [gnulib][gnulib]. Almost every utility includes functions from gnulib which are specially designed for common problems used in many places across various - systems - No need to reinvent the wheel.

- **Launched from a shell** - The Core utilities expect support from a shell such as bash, zsh, ksh, and others. The shell forks/clones in to the utility, passes the arguments, sets up the environment, - redirects I/O via pipes, and retains exit values.

- **Three families** - GNU coreutils were originally three distinct packages for shell, text, and file utilities. Utilities within the - same type share many of the same design patterns.

## Basic design
Most CLI utilities look something close to this:

The key ideas:

A setup phase for flags, options, localization, etc
An argument parsing phase thats reads input to set execution parameters
A processing/execution phase that prepares input for one or more syscalls
Many opportunities to check constraints and fail out of execution
Distinct EXIT status hint about problem location
EXIT_FAILURE is general and commonly used
Providing feedback after failed execution
This is the framework I'll use to organize the decoding of each utility. We'll see that each has a unique variant of this idea which range from a few lines to thousands of lines. I'd categorize the variants in three groups: trivial, wrappers, and full utilities

Trivial utilities
Trivial utilities have a unique set up phase which defines a macro in a couple lines. Then it 'includes' the source of another utility in which the macro forces a specific flow control. Examples include: arch, dir, and vdir

Wrapper utilities
Wrappers perform setup and parse command line options which are passed directly as arguments to a syscall. The result of the syscall is the result of the utility. These utilities do little processing on their own. Examples include: link, whoami, hostid, logname, and more

Full utilities
The diagram above shows a design for full utilities. A setup phase, an option/argument parsing phase, and execution. Execution means processing input data and may invoke many syscalls along the way to handle more data until complete. Most utilities fall in to this category.

[maybe-not]: (https://www.gnu.org/software/coreutils/manual/html_node/Standards-conformance.html)
[guidelines]: (http://pubs.opengroup.org/onlinepubs/9699919799/basedefs/V1_chap12.html#tag_12_02)
[3718]: (http://www.open-std.org/jtc1/sc22/open/n4217.pdf)
[gnulib]: (https://www.gnu.org/s/gnulib/)
[getopts]: (https://pubs.opengroup.org/onlinepubs/9699919799/functions/getopt.html)
