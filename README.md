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

## License

```
The MIT License (MIT)

Copyright (c) 2018 Adam Duke

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
