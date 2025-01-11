package main

import (
	"os"

	"github.com/alecthomas/kingpin/v2"

	"github.com/adamvduke/bcrypt-cli/commands"
)

func main() {
	app := kingpin.New("bcrypt-cli", "Wraps golang.org/x/crypto/bcrypt in a cli")
	commands.ConfigureCompareCommand(app, os.Stdin, os.Stdout)
	commands.ConfigureCostCommand(app, os.Stdin, os.Stdout)
	commands.ConfigureHashCommand(app, os.Stdin, os.Stdout)
	commands.ConfigureGenerateCommand(app, os.Stdout)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
