colorize
========

Colorize terminal outputs of an arbitrary command.

Install
-------

1. Install [golang][]
2. Run the following command

```
go get -u github.com/tarao/colorize
```

Usage
-----

### Colorized outputs

For example, the following command will make your command outputs red.
The color applies to both stdout and stderr.

```
colorize --fg=red -- your command and arguments
```

### Different colors for stdout and stderr

Each command line option can have a prefix `out-` or `err-`, which
corresponds to stdout and stderr respectively.  The following command
make the stdout outputs blue and stderr outputs red.

```
colorize --out-fg=blue --err-fg=red -- your command and arguments
```

### Prefix

`--prefix` option insert a prefix text at the beginning of each output
line.  The following command make the stdout outputs blue with "out: "
prefix and stderr outputs red with "err: " prefix.

```
colorize --out-fg=blue --out-prefix='out: ' --err-fg=red --err-prefix='err: ' -- your command and arguments
```

### Patterns

You can restrict the colorization to be applied to matching parts of
lines.  The following command colorize lines start with `Error:` as
red, and lines start with `Warn:` as yellow.

```
colorize --pattern='^Error:.*' --fg=red -- colorize --force --pattern='^Warn:.*' --fg=yellow -- your command and arguments
```

Notice that `colorize` command can be nested; it can colorize another
outputs of `colorize` command.  The important thing is `--force`
option of the second `colorize` command.  By default, `colorize`
command ignores color options when the output is not a terminal.  You
need `--force` to stop ignoring color options for nested `colorize`
command, whose output is piped, not a terminal.

Patterns also applies to the prefix part if `--prefix` option is
specified together.

### Other options

Other than `--fg` or `--bg`, you can use style options such as
`--italic` and `--underline`.  See `colorize --help` for the full list
of options.

Available colors
----------------

You can use 16 ANSI colors listed as below.

- black
- red
- green
- yellow
- blue
- magenta
- cyan
- white
- hiblack
- hired
- higreen
- hiyellow
- hiblue
- himagenta
- hicyan
- hiwhite

License
-------

- Copyright (C) INA Lintaro
- MIT License

[golang]: https://golang.org/
