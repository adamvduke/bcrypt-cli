package commands

import (
	"bytes"
	"fmt"
	"os"

	"github.com/alecthomas/kingpin/v2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/term"
)

type HashCommand struct {
	Cost int
}

func (command *HashCommand) run(context *kingpin.ParseContext) error {
	fmt.Print("Enter password:")
	password, err := readPassword()
	if err != nil {
		panic(err)
	}
	fmt.Print("Confirm password:")
	confirmation, err := readPassword()
	if err != nil {
		panic(err)
	}

	if !bytes.Equal(password, confirmation) {
		fmt.Println("password and confirmation don't match")
		os.Exit(1)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(password, command.Cost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))
	return nil
}

func ConfigureHashCommand(app *kingpin.Application) {
	command := &HashCommand{}
	hash := app.Command("hash", "Use bcrypt to hash a password").Action(command.run)
	hash.Flag("cost", "The hashing cost to use").Short('c').IntVar(&command.Cost)
}

func readPassword() ([]byte, error) {
	confirmation, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	return confirmation, err
}
