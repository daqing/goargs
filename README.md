Goargs
======

A Utility tool to replace xargs, written in Go.

How it works
============

This program was designed with the idea of positional arguments.

It will replace positional arguments with input from stdin.

For example:

$ `echo foo bar | goargs echo :2 :1`
> output: bar foo

So, in this example, goargs read `foo bar` from stdin, and make `foo` as `:1`, `bar` as `:2`.

That is, `foo` is the first positional argument, `bar` is the second, etc.

Then, goargs will replace positional arguments before executing the command,

So `goargs echo :2 :1` will become `goargs echo bar foo`, then `echo bar foo` was executed and you got the final output.

Usage
=====

`goargs` works like `xargs`:

1. `find . -name '*.go' | goargs wc -l`
2. `find . -name '*.go' | goargs mv :1 :1.bak`
3. `find . -name '*.go' | awk -F. '{print $1, $2, $3}' | goargs echo :3.:2.:1`
