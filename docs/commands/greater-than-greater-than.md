# _murex_ Shell Docs

## Command Reference: `>>` (append file)

> Writes STDIN to disk - appending contents if file already exists

## Description

Redirects output to file.

If a file already exists, the contents will be appended to existing contents.
Otherwise a new file is created.

## Usage

    <stdin> -> >> filename

## Examples

    g * -> >> files.txt

## Synonyms

* `>>`
* `fappend`


## See Also

* [commands/`>` (truncate file)](../commands/greater-than.md):
  Writes STDIN to disk - overwriting contents if file already exists
* [commands/`g`](../commands/g.md):
  Glob pattern matching for file system objects (eg *.txt)
* [commands/`pipe`](../commands/pipe.md):
  Manage _murex_ named pipes
* [commands/`tmp`](../commands/tmp.md):
  Create a temporary file and write to it