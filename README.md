# bcrypt-cli

Wraps golang.org/x/crypto/bcrypt in a cli

[![Go Reference](https://pkg.go.dev/badge/github.com/adamvduke/bcrypt-cli.svg)](https://pkg.go.dev/github.com/adamvduke/bcrypt-cli)

## Install
### Download
- Download an archive that matches your platform from the [releases page](https://github.com/adamvduke/bcrypt-cli/releases).
- Extract the archive.
- On Mac OS, remove the quarantine attribute to avoid Gatekeeper preventing the application from running.
  - `xattr -dr com.apple.quarantine bcrypt-cli`
- Move the `bcrypt-cli` file to a directory on your `$PATH`.
 
### Build From Source
- Make sure you have [Go](https://golang.org/doc/install) installed.
- `go install github.com/adamvduke/bcrypt-cli@v0.0.9`

## Usage

```
$ bcrypt-cli --help
Wraps golang.org/x/crypto/bcrypt in a cli

A simple CLI for hashing and comparing passwords using bcrypt.

Usage:
  bcrypt-cli [command]

Available Commands:
  compare     Compare a previously hashed password to a plain text password
  cost        Use bcrypt to hash a password
  generate    Output a random password and its bcrypt hash
  hash        Use bcrypt to hash a password
  help        Help about any command
  version     Print the version number of bcrypt-cli

Flags:
  -h, --help      help for bcrypt-cli
  -v, --version   version for bcrypt-cli

Use "bcrypt-cli [command] --help" for more information about a command.
```
