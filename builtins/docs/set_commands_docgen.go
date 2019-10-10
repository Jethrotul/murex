package docs

func init() {

	Definition["set"] = "# _murex_ Shell Docs\n\n## Command Reference: `set`\n\n> Define a local variable and set it's value\n\n### Description\n\nDefines, updates or deallocates a local variable.\n\n### Usage\n\n    <stdin> -> set var_name\n    \n    # Assume value from STDIN, define the data type manually\n    <stdin> -> set datatype var_name\n    \n    # Define value manually (data type defaults to string; `str`)\n    set var_name=data\n    \n    # Define value and data type manually\n    set datatype var_name=data\n    \n    # Define a variable but don't set any value\n    set var_name\n    set datatype var_name\n\n### Examples\n\nAs a method:\n\n    » out \"Hello, world!\" -> set hw\n    » out \"$hw\"\n    Hello, World!\n    \nAs a function:\n\n    » set hw=\"Hello, world!\"\n    » out \"$hw\"\n    Hello, World!\n\n### Detail\n\n#### Deallocation\n\nYou can unset variable names with the bang prefix:\n\n    !set var_name\n    \n#### Scoping\n\nVariables are only scoped inside the code block they're defined in (or any\nchildren of that code block). For example `$foo` will return an empty string in\nthe following code because it's defined within a `try` block then being queried\noutside of the `try` block:\n\n    » try {\n    »     set foo=bar\n    » }\n    » out \"foo: $foo\"\n    foo:\n    \nHowever if we define `$foo` above the `try` block then it's value will be changed\neven though it is being set inside the `try` block:\n\n    » set foo\n    » try {\n    »     set foo=bar\n    » }\n    » out \"foo: $foo\"\n    foo: bar\n    \nSo unlike the previous example, this will return `bar`.\n\nWhere `global` differs from `set` is that the variables defined with `global`\nwill scoped at the global shell level (please note this is not the same as\nenvironmental variables!) so will cascade down through all scoped code-blocks\nincluding those running in other threads.\n\nIt's also worth remembering that any variable defined using `set` in the shell's\nFID (ie in the interactive shell) is literally the same as using `global`\n\nExported variables (defined via `export`) are system environmental variables.\nInside _murex_ environmental variables behave much like `global` variables\nhowever their real purpose is passing data to external processes. For example\n`env` is an external process on Linux (eg `/usr/bin/env` on ArchLinux):\n\n    » export foo=bar\n    » env -> grep foo\n    foo=bar\n    \n#### Function Names\n\nAs a security feature function names cannot include variables. This is done to\nreduce the risk of code executing by mistake due to executables being hidden\nbehind variable names.\n\nInstead _murex_ will assume you want the output of the variable printed:\n\n    » out \"Hello, world!\" -> set hw\n    » $hw\n    Hello, world!\n    \nOn the rare occasions you want to force variables to be expanded inside a\nfunction name, then call that function via `exec`:\n\n    » set cmd=grep\n    » ls -> exec: $cmd main.go\n    main.go\n    \nThis only works for external executables. There is currently no way to call\naliases, functions nor builtins from a variable and even the above `exec` trick\nis considered bad form because it reduces the readability of your shell scripts.\n\n#### Usage Inside Quotation Marks\n\nLike with Bash, Perl and PHP: _murex_ will expand the variable when it is used\ninside a double quotes but will escape the variable name when used inside single\nquotes:\n\n    » out \"$foo\"\n    bar\n    \n    » out '$foo'\n    $foo\n    \n    » out ($foo)\n    bar\n    \n#### Declaration Without Values\n\nYou can declare a variable without a value. This is largely only of use when\nyou want to overide the scoping of a variable inside a nested code-block.\n(see the text above about variable scoping).\n\n### Synonyms\n\n* `set`\n* `!set`\n\n\n### See Also\n\n* [commands/`(` (brace quote)](../commands/brace-quote.md):\n  Write a string to the STDOUT without new line\n* [commands/`=` (arithmetic evaluation)](../commands/equ.md):\n  Evaluate a mathematical function\n* [commands/`[[` (element)](../commands/element.md):\n  Outputs an element from a nested structure\n* [commands/`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [commands/`exec`](../commands/exec.md):\n  Runs an executable\n* [commands/`export`](../commands/export.md):\n  Define a local variable and set it's value\n* [commands/`global`](../commands/global.md):\n  Define a global variable and set it's value\n* [commands/`let`](../commands/let.md):\n  Evaluate a mathematical function and assign to variable"

}
