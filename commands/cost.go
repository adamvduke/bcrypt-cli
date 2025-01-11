package commands

import (
	"fmt"
	"io"

	"github.com/alecthomas/kingpin/v2"

	"golang.org/x/crypto/bcrypt"
)

type CostCommand struct {
	In  io.Reader
	Out io.Writer
}

func ConfigureCostCommand(app *kingpin.Application, inputReader io.Reader, outputWriter io.Writer) {
	command := &CostCommand{In: inputReader, Out: outputWriter}
	app.Command("cost", "Print the hashing cost used to create the given hash").Action(command.Run)
}

func (command *CostCommand) Run(context *kingpin.ParseContext) error {
	fmt.Fprintln(command.Out, "Enter previously hashed password:")
	hash, err := readInput(command.In)
	if err != nil {
		return err
	}
	cost, err := bcrypt.Cost(hash)
	if err != nil {
		return err
	}
	fmt.Fprintln(command.Out, "Cost:", cost)

	return nil
}
