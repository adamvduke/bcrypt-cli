package main

import (
	"os"

	"github.com/adamvduke/bcrypt-cli/commands"
	"github.com/alecthomas/kingpin/v2"
)

func main() {
	app := kingpin.New("bcrypt-cli", "Wraps golang.org/x/crypto/bcrypt in a cli")
	commands.ConfigureCompareCommand(app)
	commands.ConfigureCostCommand(app)
	commands.ConfigureHashCommand(app)
	commands.ConfigureGenerateCommand(app)
	kingpin.MustParse(app.Parse(os.Args[1:]))
}
