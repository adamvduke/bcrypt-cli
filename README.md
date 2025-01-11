# bcrypt-cli

Wraps golang.org/x/crypto/bcrypt in a cli

## Features

## Install

- Make sure you have [Go](https://golang.org/doc/install) installed and have set your [GOPATH](https://golang.org/doc/code.html#GOPATH).
- `go get github.com/adamvduke/bcrypt-cli`

## Usage

```
$ bcrypt-cli --help
usage: bcrypt-cli [<flags>] <command> [<args> ...]

Wraps golang.org/x/crypto/bcrypt in a cli

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.

  compare
    Compare a hashed password to a plain text password

  cost
    Print the hashing cost used to create the given hash

  hash [<flags>]
    Use bcrypt to hash a password

  generate [<flags>]
    Output a random password and it's bcrypt hash
```
