# _murex_ Shell Guide

## Command Reference: `let`

> Evaluate a mathmatical function and assign to variable

### Description

`let` evaluates a mathmatical function and then assigns it to a locally
scoped variable (like `set`)

### Usage

    let var_name=evaluation
    
    let var_name++
    
    let var_name--

### Examples

    » let: age=18
    » $age
    18
    
    » let: age++
    » $age
    19
    
    » let: under18=age<18
    » $under18
    false
    
    » let: under21 = age < 21
    » $under21
    true

### Detail

#### Variables

There are two ways you can use variables with the math functions. Either by
string interpolation like you would normally with any other function, or
directly by name.

String interpolation:

    » set abc=123
    » = $abc==123
    true
    
Directly by name:

    » set abc=123
    » = abc==123
    false
    
To understand the difference between the two, you must first understand how
string interpolation works; which is where the parser tokenised the parameters
like so

    command line: = $abc==123
    token 1: command (name: "=")
    token 2: parameter 1, string (content: "")
    token 3: parameter 1, variable (name: "abc")
    token 4: parameter 1, string (content: "==123")
    
Then when the command line gets executed, the parameters are compiled on demand
similarly to this crude pseudo-code

    command: "="
    parameters 1: concatenate("", GetValue(abc), "==123")
    output: "=" "123==123"
    
Thus the actual command getting run is literally `123==123` due to the variable
being replace **before** the command executes.

Whereas when you call the variable by name it's up to `=` or `let` to do the
variable substitution.

    command line: = abc==123
    token 1: command (name: "=")
    token 2: parameter 1, string (content: "abc==123")
    
    command: "="
    parameters 1: concatenate("abc==123")
    output: "=" "abc==123"
    
The main advantage (or disadvantage, depending on your perspective) of using
variables this way is that their data-type is preserved.

    » set str abc=123
    » = abc==123
    false
    
    » set int abc=123
    » = abc==123
    true
    
Unfortunately is one of the biggest areas in _murex_ where you'd need to be
careful. The simple addition or omission of the dollar prefix, `$`, can change
the behavior of `=` and `let`.

#### Strings

Because the usual _murex_ tools for encapsulating a string (`"`, `'` and `()`)
are interpreted by the shell language parser, it means we need a new token for
handling strings inside `=` and `let`. This is where backtick comes to our
rescue.

    » set str abc=123
    » = abc==`123`
    true
    
Please be mindful that if you use string interpolation then you will need to
instruct `=` and `let` that your field is a string

    » set str abc=123
    » = `$abc`==`123`
    true
    
#### Best practice recommendation

As you can see from the sections above, string interpolation offers us some
conveniences when comparing variables of differing data-types, such as a `str`
type with a number (eg `num` or `int`). However it makes for less readable code
when just comparing strings. Thus the recommendation is to avoid using string
interpolation except only where it really makes sense (ie use it sparingly).

#### Non-boolean logic

Thus far the examples given have been focused on comparisons however `=` and
`let` supports all the usual arithmetic operators:

    » = 10+10
    20
    
    » = 10/10
    1
    
    » = (4 * (3 + 2))
    20
    
    » = `foo`+`bar`
    foobar
    
#### Read more

_murex_ uses the govaluate package. More information can be found in it's manual:
https://github.com/Knetic/govaluate/blob/master/MANUAL.md

#### Scoping

Variables are only scoped inside the code block they're defined in (or any
children of that code block). For example `$foo` will return an empty string in
the following code because it's defined within a `try` block then being queried
outside of the `try` block:

    » try {
    »     set foo=bar
    » }
    » out "foo: $foo"
    foo:
    
However if we define `$foo` above the `try` block then it's value will be changed
even though it is being set inside the `try` block:

    » set foo
    » try {
    »     set foo=bar
    » }
    » out "foo: $foo"
    foo: bar
    
So unlike the previous example, this will return `bar`.

Where `global` differs from `set` is that the variables defined with `global`
will scoped at the global shell level (please note this is not the same as
environmental variables!) so will cascade down through all scoped code-blocks
including those running in other threads.

It's also worth remembering that any variable defined using `set` in the shell's
FID (ie in the interactive shell) is literally the same as using `global`

Exported variables (defined via `export`) are system environmental variables.
Inside _murex_ environmental variables behave much like `global` variables
however their real purpose is passing data to external processes. For example
`env` is an external process on Linux (eg `/usr/bin/env` on ArchLinux):

    » export foo=bar
    » env -> grep foo
    foo=bar
    
#### Function Names

As a security feature function names cannot include variables. This is done to
reduce the risk of code executing by mistake due to executables being hidden
behind variable names.

Instead _murex_ will assume you want the output of the variable printed:

    » out "Hello, world!" -> set hw
    » $hw
    Hello, world!
    
On the rare occasions you want to force variables to be expanded inside a
function name, then call that function via `exec`:

    » set cmd=grep
    » ls -> exec: $cmd main.go
    main.go
    
This only works for external executables. There is currently no way to call
aliases, functions nor builtins from a variable and even the above `exec` trick
is considered bad form because it reduces the readability of your shell scripts.

#### Usage Inside Quotation Marks

Like with Bash, Perl and PHP: _murex_ will expand the variable when it is used
inside a double quotes but will escape the variable name when used inside single
quotes:

    » out "$foo"
    bar
    
    » out '$foo'
    $foo
    
    » out ($foo)
    bar

### See Also

* [`(` (brace quote)](../commands/brace-quote.md):
  Write a string to the STDOUT without new line
* [`=` (artithmetic evaluation)](../commands/equ.md):
  Evaluate a mathmatical function
* [`=` (artithmetic evaluation)](../commands/equ.md):
  Evaluate a mathmatical function
* [`[` (index)](../commands/index.md):
  Outputs an element from an array, map or table
* [`export`](../commands/export.md):
  Define a local variable and set it's value
* [`global`](../commands/global.md):
  Define a global variable and set it's value
* [`global`](../commands/global.md):
  Define a global variable and set it's value
* [`if`](../commands/if.md):
  Conditional statement to execute different blocks of code depending on the result of the condition
* [`set`](../commands/set.md):
  Define a local variable and set it's value
* [element](../commands/element.md):
  