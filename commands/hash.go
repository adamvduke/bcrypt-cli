package commands

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/alecthomas/kingpin/v2"

	"golang.org/x/crypto/bcrypt"
)

var (
	errConfirmationMismatch = errors.New("password and confirmation do not match")
)

type HashCommand struct {
	Cost int
	In   io.Reader
	Out  io.Writer
}

func ConfigureHashCommand(app *kingpin.Application, inputReader io.Reader, outputWriter io.Writer) {
	command := &HashCommand{In: inputReader, Out: outputWriter}
	hash := app.Command("hash", "Use bcrypt to hash a password").Action(command.Run)
	hash.Flag("cost", "The hashing cost to use").Short('c').IntVar(&command.Cost)
}

func (command *HashCommand) Run(context *kingpin.ParseContext) error {
	fmt.Fprintln(command.Out, "Enter password:")
	password, err := readSensitive(command.In)
	if err != nil {
		return err
	}

	fmt.Fprintln(command.Out, "Confirm password:")
	confirmation, err := readSensitive(command.In)
	if err != nil {
		return err
	}

	if !bytes.Equal(password, confirmation) {
		return errConfirmationMismatch
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(password, command.Cost)
	if err != nil {
		return err
	}
	fmt.Fprintln(command.Out, "Bcrypt:", string(hashedPassword))

	return nil
}
