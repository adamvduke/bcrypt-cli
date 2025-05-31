package main

import (
	"os"

	"github.com/alecthomas/kingpin/v2"

	"github.com/adamvduke/bcrypt-cli/bcryptio"
)

//go:generate go tool -modfile=tools.mod golangci-lint run

const (
	version = "0.0.8"
)

func main() {
	app := kingpin.New("bcrypt-cli", "Wraps golang.org/x/crypto/bcrypt in a cli")
	app.Version(version).VersionFlag.Short('v')
	app.HelpFlag.Short('h')
	bcryptio.ConfigureCompareCommand(app, os.Stdin, os.Stdout)
	bcryptio.ConfigureCostCommand(app, os.Stdin, os.Stdout)
	bcryptio.ConfigureHashCommand(app, os.Stdin, os.Stdout)
	bcryptio.ConfigureGenerateCommand(app, os.Stdout)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
