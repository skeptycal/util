# tree

## Description:
>Tree is a recursive directory listing command that produces a depth indented listing of files, which is colorized ala dircolors if color tty output is available.

>The idea was based on the functionality of the GNU freeware program `tree` and is similarly licensed (GPL v2)

>Tree has been tested under the following operating systems: Linux, macOS, Windows.

>The following likely work because of Go's wide compiler coverage, but no testing has been done: FreeBSD, Solaris, HP/UX, Cygwin, HP Nonstop and OS/2.

## Example:
Here is an example of the output of tree:

![terminal output from tree command](original/tree.jpg)

```sh
    ➜ tree -C
    .
    ├── CHANGES
    ├── INSTALL_RECEIPT.json
    ├── LICENSE
    ├── README
    ├── TODO
    ├── bin
    │   └── tree
    └── share
        └── man
            └── man1
                └── tree.1

    4 directories, 7 files
```

asdf
