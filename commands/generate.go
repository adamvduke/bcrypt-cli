package commands

import (
	"fmt"
	"io"
	"math/rand"

	"github.com/alecthomas/kingpin/v2"

	"golang.org/x/crypto/bcrypt"
)

const (
	defaultLength = 20
	maxLength     = 72
	printableMax  = 126
	printableMin  = 33
)

type GenerateCommand struct {
	Cost   int
	Length int
	Out    io.Writer
}

func ConfigureGenerateCommand(app *kingpin.Application, outputWriter io.Writer) {
	command := &GenerateCommand{Out: outputWriter}
	generate := app.Command("generate", "Output a random password and it's bcrypt hash").Action(command.Run)
	generate.Flag("cost", "The hashing cost to use").Short('c').IntVar(&command.Cost)
	generate.Flag("length", "The length of the random password to generate").Short('l').IntVar(&command.Length)
}

func (command *GenerateCommand) Run(context *kingpin.ParseContext) error {
	if command.Length < defaultLength {
		command.Length = defaultLength
	}
	if command.Length > maxLength {
		command.Length = maxLength
	}
	randomBytes := randomBytes(command.Length)
	hashedPassword, err := bcrypt.GenerateFromPassword(randomBytes, command.Cost)
	if err != nil {
		return err
	}
	fmt.Fprintln(command.Out, "Plaintext:", string(randomBytes))
	fmt.Fprintln(command.Out, "Bcrypt:", string(hashedPassword))

	return nil
}

func randomBytes(length int) []byte {
	b := make([]byte, length)
	for i := range b {
		j := rand.Intn((printableMax - printableMin + 1)) + printableMin
		b[i] = byte(j)
	}
	return b
}
