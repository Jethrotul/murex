# murex
(I'm not sold on that name either. However I am open to suggestions)

## Description

Murex is a cross-platform shell like Bash but with greater emphasis on
writing shell scripts and powerful one-liners while maintaining
readability.

To achieve this the language employs a relatively simple syntax modelled
loosely on functional and stack-based programming paradigms (albeit
without the LISP-style nested parentheses that scare a lot of developers.
For example, a program structure could look like the following:
```
command -> command -> if { then_command } -> else { else_command }.
```

The language supports multiple data types, with JSON (and later XML)
support as a native data type. Which makes passing data through the
pipeline easier when dealing with more complex arrangements of data than
a simple byte stream when compared to standard shells like Bash.

However despite the amount of features added to shell, I have  tried to
keep the amount of "magic" to a minimum and follow a pretty standard
structure so the language is predictable and guessable.
 
I will also be working on hardening the shell to make it more robust for
writing shell scripts. Bash, for all it's power, is littered with hidden
traps. I'm hoping to address as many of them as I can without taking
much flexibility nor power away from the command line.

## Dependencies
```
go get github.com/chzyer/readline
go get github.com/kr/pty
```

## Build
```
go build github.com/lmorg/murex
```

Test the binary (requires Bash):
```
test/regression_test.sh
```

## Language guides

Please read the following guides:

1. [GUIDE.syntax.md](./GUIDE.syntax.md) - this is recommended first as
it gives an overview if the shell scripting languages syntax and data
types.

2. [GUIDE.control-structures.md](./GUIDE.control-structures.md) - this
will list how to use if statements and iteration like for loops.

3. [GUIDE.builtin-functions.md](./GUIDE.builtin-functions.md) - lastly
this will list some of the builtin functions available for this shell.

## Known bugs / TODO

* _Currently no support for interactive commands._ This will need to be
addressed.

* _Interactive shell auto-completion is unreliable._ I have a nasty
feeling I may need to fork the readline package or even create my own
one.

* _Interactive shell does not support multiline scripts._ Related to
previous issue.

* _No support for mathematical operators._ This is going be the next
feature I include. Current plan is to build a `math` builtin function
rather than allowing users to inline mathematical operations. I'm also
considering using [Polish Notation](https://en.wikipedia.org/wiki/Polish_notation)
simply because it's easier to write a PN parser.

* _`foreach` only supports line splitting - not JSON objects._ This is a
TODO rather than bug.

* _No support for piping scripts to the shell executable._ This will be
supported via a `--stdin` flag. It's an easy thing to implement but
wasn't considered necessary for the MVP (minimum viable product).

* _No support for reading saved files._ Again this is an easy fix but
wasn't required for the MVP. However it will be needed if I want support
for sha-bang (`#!/bin/env`-prefixed scripts).

* _The lang/builins/encoders package exhibits weird behaviors as seen
from the failing regression tests._ This is a tougher problem to crack.
My current working theory is that there is a race condition with io.Copy
finishing before the encoder has finished writing to it's buffer. I've
spent some time investigating this issue but now sidelined it to focus
on developing other aspects of this project. My worry is that this is a
symptom of a much broader problem, however the project is modular enough
that even a substantial rewrite of core logic can be done safely in
isolation from other, working, code. If anyone wishes to contribute
their time to this project then that would be the one area I would most
appreciate your support.
