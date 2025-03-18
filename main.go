package main

import (
	"os"

	"github.com/alecthomas/kingpin/v2"

	"github.com/adamvduke/bcrypt-cli/commands"
)

const (
	version = "0.0.1"
)

func main() {
	app := kingpin.New("bcrypt-cli", "Wraps golang.org/x/crypto/bcrypt in a cli")
	app.Version(version).VersionFlag.Short('v')
	app.HelpFlag.Short('h')
	commands.ConfigureCompareCommand(app, os.Stdin, os.Stdout)
	commands.ConfigureCostCommand(app, os.Stdin, os.Stdout)
	commands.ConfigureHashCommand(app, os.Stdin, os.Stdout)
	commands.ConfigureGenerateCommand(app, os.Stdout)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
