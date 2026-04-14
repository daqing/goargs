Goargs
======

A utility tool to replace `xargs`, written in Go.

How it works
============

This program was designed with the idea of positional arguments.

It will replace positional arguments with input from stdin.

For example:

```bash
$ echo "foo bar" | goargs echo :1 :0
bar foo
```

So, in this example, goargs read `foo bar` from stdin, and make `foo` as `:0`, `bar` as `:1`.

That is, `foo` is the first positional argument, `bar` is the second, etc.

Then, goargs will replace positional arguments before executing the command.

So `goargs echo :1 :0` will become `goargs echo bar foo`, then `echo bar foo` was executed and you got the final output.

Placeholders
============

- `:0, :1, :2, ...` — Reference fields split by whitespace (0-indexed).
- `:@` — Reference the entire line (useful for filenames with spaces).

Options
=======

- `-F sep` — Use `sep` as the field separator instead of whitespace.

Usage
=====

### Basic usage

```bash
find . -name '*.go' | goargs wc -l
find . -name '*.go' | goargs mv :0 :0.bak
```

### Whole line placeholder (`:@`)

For filenames with spaces, use `:@` to reference the entire line:

```bash
echo 'Frame 123.svg' | goargs mv :@ :1.svg
# renames "Frame 123.svg" to "123.svg"
```

### Custom field separator (`-F`)

Use `-F` to split fields by a custom separator (like `awk -F`):

```bash
echo 'a.txt' | goargs -F. mv :@ :0.md
# renames "a.txt" to "a.md"
```

Another example combining `awk` with positional fields:

```bash
find . -name '*.go' | awk -F/ '{print $1, $2}' | goargs echo :0/:1
```
