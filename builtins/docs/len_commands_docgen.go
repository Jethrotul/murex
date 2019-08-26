package docs

func init() {

	Definition["len"] = "# _murex_ Shell Guide\n\n## Command Reference: `len` \n\n> Outputs the length of an array\n\n### Description\n\nThis will read an array from STDIN and outputs the length for that array\n\n### Usage\n\n    <STDIN> -> len -> <stdout>\n\n### Examples\n\n    » tout: json ([\"a\", \"b\", \"c\"]) -> len \n    3\n\n### Detail\n\nPlease note that this returns the length of the _array_ rather than string.\nFor example `out \"foobar\" -> len` would return `1` because an array in the\n`str` data type would be new line separated (eg `out \"foo\\nbar\" -> len`\nwould return `2`). If you need to count characters in a string and are\nrunning POSIX (eg Linux / BSD / OSX) then it is recommended to use `wc`\ninstead. But be mindful that `wc` will also count new line characters\n\n    » out: \"foobar\" -> len\n    1\n    \n    » out: \"foo\\nbar\" -> len\n    2\n    \n    » out: \"foobar\" -> wc: -c\n    7\n    \n    » out: \"foo\\nbar\" -> wc: -c\n    8\n    \n    » printf: \"foobar\" -> wc: -c\n    6\n    # (printf does not print a trailing new line)\n\n### See Also\n\n* [`@[` (range) ](../commands/range.md):\n  Outputs a ranged subset of data from STDIN\n* [`[` (index)](../commands/index.md):\n  Outputs an element from an array, map or table\n* [`a` (mkarray)](../commands/a.md):\n  A sophisticated yet simple way to build an array or list\n* [`append`](../commands/append.md):\n  Add data to the end of an array\n* [`ja`](../commands/ja.md):\n  A sophisticated yet simply way to build a JSON array\n* [`jsplit` ](../commands/jsplit.md):\n  Splits STDIN into a JSON array based on a regex parameter\n* [`map` ](../commands/map.md):\n  Creates a map from two data sources\n* [`msort` ](../commands/msort.md):\n  Sorts an array - data type agnostic\n* [`mtac`](../commands/mtac.md):\n  Reverse the order of an array\n* [`prepend` ](../commands/prepend.md):\n  Add data to the start of an array\n* [element](../commands/element.md):\n  "

}
