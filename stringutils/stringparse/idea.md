## stringparse
>ref: https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/

In this example, we’ll create a command-line tool called stringparse, that will count the characters, words, or lines in a string. Here's our inteded usage to start (we'll add in some subcommands and more flags later):

 -metric string     Metric {chars|words|lines}. (default "chars")  -text string     Text to parse. (required)  -unique     Measure unique values of a metric.

## Parsing Arguments

### Arguments, Options, Flags -- What’s the difference?

These terms are often used interchangeably. In general, they refer to the list of strings following a CLI command. Here are the traditional definitions:

- Arguments are all strings that follow a CLI command.
- Options are arguments with dashes (single or double) that are followed by user input and modify the operation of the command.
- Flags are boolean options that do not take user input.

Go does not stick to these definitions. In the context of its flag package, all strings beginning with dashes (single or double) are considered flags, and all strings without dashes that are not user input are considered arguments. All strings after the first argument are considered arguments.
