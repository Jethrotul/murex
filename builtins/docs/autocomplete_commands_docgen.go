package docs

func init() {

	Definition["autocomplete"] = "# _murex_ Shell Docs\n\n## Command Reference: `autocomplete`\n\n> Set definitions for tab-completion in the command line\n\n## Description\n\n`autocomplete` digests a JSON schema and uses that to define the tab-\ncompletion rules for suggestions in the interactive command line.\n\n## Usage\n\n    autocomplete get command -> <stdout>\n    \n    autocomplete set command { mxjson }\n\n## Flags\n\n* `get`\n    output all autocompletion schemas\n* `set`\n    define a new autocompletion schema\n\n## Detail\n\n### Undefining autocomplete\n\nCurrently there is no support for undefining an autocompletion rule however\nyou can overwrite existing rules.\n\n## Directives\n\nThe directives are listed below. Headings are formatted as follows:\n\n    \"DirectiveName\": json data-type (default value)\n    \nWhere \"default value\" is what will be auto-populated at run time if you\ndon't define an autocomplete schema manually.\n\n### \"Alias\": string (\"\")\n\n_description to follow_\n\n### \"AllowMultiple\": boolean (false)\n\nSet to `true` to enable multiple parameters following the same rules as\ndefined in this index. For example the following will suggest directories\non each tab for multiple parameters:\n\n    autocomplete set example { [{\n        \"IncDirs\": true,\n        \"AllowMultiple\": true\n    }] }\n    \n### \"AnyValue\": boolean (false)\n\nThe way autocompletion works in _murex_ is the suggestion engine looks for\nmatches and if it fines one, it then moves onto the next index in the JSON\nschema. This means unexpected values typed in the interactive terminal will\nbreak the suggestion engine's ability to predict what the next expected\nparameter should be. Setting **AnyValue** to `true` tells the suggestion\nengine to accept any value as the next parameter thus allowing it to then\npredict the next parameter afterwards.\n\nThis directive isn't usually nessisary because such fields are often the\nlast parameter or most parameters can be detectable with a reasonable\namount of effort. However **AnyValue** is often required for more complex\ncommand line tools.\n\n### \"AutoBranch\": boolean (false)\n\nUse this in conjunction with **Dynamic**. If the return is an array of paths,\nfor example `[ \"/home/foo\", \"/home/bar\" ]` then **AutoBranch** will return\nthe following patterns in the command line:\n\n    » example [tab]\n    # suggests \"/home/\"\n    \n    » example /home/[tab]\n    # suggests \"/home/foo\" and \"/home/bar\"\n    \nPlease note that **AutoBranch**'s behavior is also dependant on a \"shell\"\n`config` setting, recursive-enabled\":\n\n    » config get shell recursive-enabled\n    true\n    \n### \"Dynamic\": string (\"\")\n\nThis is a _murex_ block which returns an array of suggestions.\n\nCode inside that block are executed like a function and the parameters will\nmirror the same as those parameters entered in the interactive terminal.\n\n### \"DynamicDesc\": string (\"\")\n\nThis is very similar to **Dynamic** except your function should return a\nmap instead of an array. Where each key is the suggestion and the value is\na description.\n\nThe description will appear either in the hint text or alongside the\nsuggestion - depending on which suggestion \"popup\" you define (see\n**ListView**).\n\n### \"ExecCmdline\": boolean (false)\n\nSometimes you'd want your autocomplete suggestions to aware of the output\nreturned from the commands that preceded it. For example the suggestions\nfor `[` (index) will depend entirely on what data is piped into it.\n\n**ExecCmdline** tells _murex_ to run the commandline up until the command\nwhich your cursor is editing and pipe that output to the STDIN of that\ncommands **Dynamic** or **DynamicDesc** code block.\n\n> This is a dangerous feature to enable so **ExecCmdline** is only honoured\n> if the commandline is considered \"safe\". **Dynamic** / **DynamicDesc**\n> will still be executed however if the commandline is \"unsafe\" then your\n> dynamic autocompletion blocks will have no STDIN.\n\nBecause this is a dangerous feature, your partial commandline will only\nexecute if the following conditions are met:\n\n* the commandline must be one pipeline (eg `;` tokens are not allowed)\n* the commandline must not have any new line characters\n* there must not be any redirection, including named pipes\n  (eg `cmd <namedpipe>`) and the STDOUT/STDERR switch token (`?`)\n* the commandline doesn't inline any variables (`$strings`, `@arrays`) or\n  functions (`${subshell}`, `$[index]`)\n* lastly all commands are whitelisted in \"safe-commands\"\n  (`config get shell safe-commands`)\n\nIf these criteria are met, the commandline is considered \"safe\"; if any of\nthose conditions fail then the commandline is considered \"unsafe\".\n\n_murex_ will come with a number of sane commands already included in its\n`safe-commands` whitelist however you can add or remove them using `config`\n\n    » function: foobar { -> match foobar }\n    » config: eval shell safe-commands { -> append foobar }\n    \nRemember that **ExecCmdline** is designed to be included with either\n**Dynamic** or **DynamicDesc** and those code blocks would need to read\nfrom STDIN:\n\n    autocomplete set \"[\" { [{\n        \"AnyValue\": true,\n        \"AllowMultiple\": true,\n        \"ExecCmdline\": true,\n        \"Dynamic\": ({\n            switch ${ get-type: stdin } {\n                case * {\n                    <stdin> -> [ 0: ] -> format json -> [ 0 ]\n                }\n                \n                catch {\n                    <stdin> -> formap k v { out $k } -> cast str -> append \"]\"\n                }\n            }\n        })\n    }] }\n    \n### \"FlagValues\": map of arrays (null)\n\nThis is a map of the flags with the values being the same array of directive\nas the top level.\n\nThis allows you to nest operations by flags. eg when a flag might accept\nmultiple parameters.\n\n### \"Flags\": array of strings (auto-populated from man pages)\n\nSetting **Flags** is the fastest and easiest way to populate suggestions\nbecause it is just an array of strings. eg\n\n    autocomplete set example { [{\n        \"Flags\": [ \"foo\", \"bar\" ]\n    }] }\n    \nIf a command doesn't **Flags** already defined when you request a completion\nsuggestion but that command does have a man page, then **Flags** will be\nautomatically populated with any flags identified from an a quick parse of\nthe man page. However because man pages are written to be human readable\nrather than machine parsable, there may not be a 100% success rate with the\nautomatic man page parsing.\n    \n### \"FlagsDesc\": map of strings (null)\n\nThis is the same concept as **Flags** except it is a map with the suggestion\nas a key and description as a value. This distinction is the same as the\ndifference between **Dynamic** and **DynamicDesc**.\n\nPlease note that currently man page parsing cannot provide a description so\nonly **Flags** get auto-populated.\n\n### \"IncDirs\": boolean (false)\n\nEnable to include directories.\n\nNot needed if **IncFiles** is set to `true`.\n\nBehavior of this directive can be altered with `config set shell\nrecursive-enabled`\n\n### \"IncExePath\": boolean (false)\n\nEnable this to any executables in `$PATH`. Suggestions will not include\naliases, functions nor privates.\n\n### \"IncFiles\": boolean (true)\n\nInclude files and directories. This is enabled by default for any commands\nthat don't have autocomplete defined but you will need to manually enable\nit in any `autocomplete` schemas you create and want files as part of the\nsuggestions.\n\n### \"ListView\": boolean (false)\n\nThis alters the appearance of the autocompletion suggestions \"popup\". Rather\nthan suggestions being in a grid layout (with descriptions overwriting the\nhint text) the suggestions are in a list view with the descriptions next to\nthem on the same row (similar to how an IDE might display it's suggestions).\n\n### \"NestedCommand\": boolean (false)\n\nOnly enable this if the command you are autocompleting is a nested parameter\nof the parent command you have types. For example with `sudo`, once you've\ntyped the command name you wish to elivate, then you would want suggestions\nfor that command rather than for `sudo` itself.\n\n### \"Optional\": boolean (false)\n\nSpecifies if a match is required for the index in this schema. ie optional\nflags.\n\n## See Also\n\n* [commands/`<stdin>` ](../commands/stdin.md):\n  Read the STDIN belonging to the parent code block\n* [commands/`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [commands/`alias`](../commands/alias.md):\n  Create an alias for a command\n* [commands/`config`](../commands/config.md):\n  Query or define _murex_ runtime settings\n* [commands/`function`](../commands/function.md):\n  Define a function block\n* [commands/`get-type`](../commands/get-type.md):\n  Returns the data-type of a variable or pipe\n* [commands/`private`](../commands/private.md):\n  Define a private function block\n* [commands/`summary` ](../commands/summary.md):\n  Defines a summary help text for a command\n* [commands/`switch`](../commands/switch.md):\n  Blocks of cascading conditionals\n* [types/mxjson](../types/mxjson.md):\n  Murex-flavoured JSON (primitive)"

}
