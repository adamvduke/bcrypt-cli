package bcryptio

import (
	"fmt"
	"io"

	"github.com/alecthomas/kingpin/v2"

	"golang.org/x/crypto/bcrypt"
)

type CompareCommand struct {
	In  io.Reader
	Out io.Writer
}

func ConfigureCompareCommand(app *kingpin.Application, inputReader io.Reader, outputWriter io.Writer) {
	command := &CompareCommand{In: inputReader, Out: outputWriter}
	app.Command("compare", "Compare a previously hashed password to a plain text password").Action(command.Run)
}

func (command *CompareCommand) Run(_ *kingpin.ParseContext) error {
	fmt.Fprintln(command.Out, "Enter previously hashed password:")
	hashed, err := readInput(command.In)
	if err != nil {
		return err
	}

	fmt.Fprintln(command.Out, "Enter password:")
	plain, err := readSensitive(command.In)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword(hashed, plain); err != nil {
		return err
	}
	fmt.Fprintln(command.Out, "Password is correct")

	return nil
}
