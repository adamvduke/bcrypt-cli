# bcrypt-cli

Wraps golang.org/x/crypto/bcrypt in a cli

[![Go Reference](https://pkg.go.dev/badge/github.com/adamvduke/bcrypt-cli.svg)](https://pkg.go.dev/github.com/adamvduke/bcrypt-cli)

## Install

- Make sure you have [Go](https://golang.org/doc/install) installed.
- `go install github.com/adamvduke/bcrypt-cli@v0.0.5`

## Usage

```
$ bcrypt-cli --help
usage: bcrypt-cli [<flags>] <command> [<args> ...]

Wraps golang.org/x/crypto/bcrypt in a cli


Flags:
  -h, --[no-]help     Show context-sensitive help (also try --help-long and --help-man).
  -v, --[no-]version  Show application version.

Commands:
help [<command>...]
    Show help.

compare
    Compare a previously hashed password to a plain text password

cost
    Print the hashing cost used to create the given hash

hash [<flags>]
    Use bcrypt to hash a password

generate [<flags>]
    Output a random password and it's bcrypt hash
```
