goex
======

[![Build Status](https://travis-ci.org/skeptycal/goex.svg?branch=master)](https://travis-ci.org/skeptycal/goex) [![GoDoc](https://godoc.org/github.com/skeptycal/goex?status.svg)](https://godoc.org/github.com/skeptycal/goex)

goex is a command line tool to execute Go code. Output is printed as goons to stdout.
It is heavily modeled after [goexec](https://github.com/shurcooL/goexec) and is
my interpretation of that project.

Installation
------------

```sh

# download and install a local copy to the go src directory
go get -u github.com/skeptycal/goex

# shell reload
exec ${SHELL} -l
```

Usage
-----

```
Usage: goex [flags] [packages] [package.]function(parameters)
       echo parameters | goex -stdin [flags] [packages] [package.]function
  -compiler string
    	Compiler to use, one of: "gc", "gopherjs". (default "gc")
  -n	Print the generated source but do not run it.
  -quiet
    	Do not dump the return values as a goon.
  -stdin
    	Read func parameters from stdin instead.
```

Examples
--------

```sh
$ goex 'strings.Repeat("Go! ", 5)'
(string)("Go! Go! Go! Go! Go! ")

$ goex strings 'Replace("Calling Go functions from the terminal is hard.", "hard", "easy", -1)'
(string)("Calling Go functions from the terminal is easy.")

# Note that parser.ParseExpr returns 2 values (ast.Expr, error).
$ goex 'parser.ParseExpr("5 + 7")'
(*ast.BinaryExpr)(&ast.BinaryExpr{
	X: (*ast.BasicLit)(&ast.BasicLit{
		ValuePos: (token.Pos)(1),
		Kind:     (token.Token)(5),
		Value:    (string)("5"),
	}),
	OpPos: (token.Pos)(3),
	Op:    (token.Token)(12),
	Y: (*ast.BasicLit)(&ast.BasicLit{
		ValuePos: (token.Pos)(5),
		Kind:     (token.Token)(5),
		Value:    (string)("7"),
	}),
})
(interface{})(nil)

# Run function RepoRootForImportPath from package "golang.org/x/tools/go/vcs".
$ goex 'vcs.RepoRootForImportPath("rsc.io/pdf", false)'
(*vcs.RepoRoot)(...)
(interface{})(nil)

$ goex -quiet 'fmt.Println("Use -quiet to disable output of goon; useful if you want to print to stdout.")'
Use -quiet to disable output of goon; useful if you want to print to stdout.

$ echo '"fmt"' | goex -stdin 'gist4727543.GetForcedUse'
(string)("var _ = fmt.Errorf")
```

Alternatives
------------

-   [goexec](https://github.com/shurcooL/goexec) - the inspiration for this program
-	[gommand](https://github.com/sno6/gommand) - Go one liner program, similar to python -c.
-	[gorram](https://github.com/natefinch/gorram) - Like go run for any Go function.
-	[goeval](https://github.com/dolmen-go/goeval) - Run Go snippets instantly from the command-line.

License
-------

-	[MIT License](https://opensource.org/licenses/mit-license.php)
