package docs

func init() {

	Definition["source"] = "# _murex_ Shell Docs\n\n## Command Reference: `source` \n\n> Import _murex_ code from another file of code block\n\n## Description\n\n`source` imports code from another file or code block. It can be used as either\nan \"import\" / \"include\" directive (eg Python, Go, C, etc) or an \"eval\" (eg\nPython, Perl, etc).\n\n## Usage\n\nExecute source from STDIN\n\n    <stdin> -> source\n    \nExecute source from a file\n\n    source: filename.mx\n    \nExecute a code block from parameter\n\n    source: { code-block }\n\n## Examples\n\nExecute source from stdin:\n\n    » tout: block { out: \"Hello, world!\" } -> source\n    Hello, world!\n    \nExecute source from file:\n\n    » tout: block { out: \"Hello, world!\" } |> example.mx\n    » source: example.mx\n    Hello, world!\n    \nExecute a code block from parameter\n\n    » source { out: \"Hello, world!\" }\n    Hello, world!\n\n## Synonyms\n\n* `source`\n* `.`\n\n\n## See Also\n\n* [commands/`autocomplete`](../commands/autocomplete.md):\n  Set definitions for tab-completion in the command line\n* [commands/`config`](../commands/config.md):\n  Query or define _murex_ runtime settings\n* [commands/`exec`](../commands/exec.md):\n  Runs an executable\n* [commands/`fexec` ](../commands/fexec.md):\n  Execute a command or function, bypassing the usual order of precedence.\n* [commands/`function`](../commands/function.md):\n  Define a function block\n* [commands/`murex-parser` ](../commands/murex-parser.md):\n  Runs the _murex_ parser against a block of code \n* [commands/`private`](../commands/private.md):\n  Define a private function block\n* [commands/`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of _murex_\n* [commands/`version` ](../commands/version.md):\n  Get _murex_ version\n* [commands/args](../commands/args.md):\n  \n* [commands/params](../commands/params.md):\n  "

}
