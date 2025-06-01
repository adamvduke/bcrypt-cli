// Package main implements a command-line interface for bcrypt operations.
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/adamvduke/bcrypt-cli/cmd"
)

//go:generate go tool -modfile=tools.mod golangci-lint run --fix

const (
	version = "0.0.10"
)

func main() {
	ctx := context.Background()
	if err := cmd.ExecuteContext(ctx, version); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err.Error())
		os.Exit(1)
	}
}
