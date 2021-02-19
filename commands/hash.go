package commands

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/alecthomas/kingpin"
	"github.com/howeyc/gopass"
	"golang.org/x/crypto/bcrypt"
)

type HashCommand struct {
	Cost int
}

func (command *HashCommand) run(context *kingpin.ParseContext) error {
	fmt.Print("Enter password: ")
	password, err := gopass.GetPasswdMasked()
	if err != nil {
		return err
	}
	fmt.Print("Confirm password: ")
	confirmation, err := gopass.GetPasswdMasked()
	if err != nil {
		return err
	}

	if !bytes.Equal(password, confirmation) {
		return errors.New("password and confirmation don't match")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(password, command.Cost)
	if err != nil {
		return err
	}
	fmt.Println(string(hashedPassword))
	return nil
}

func ConfigureHashCommand(app *kingpin.Application) {
	command := &HashCommand{}
	hash := app.Command("hash", "Use bcrypt to hash a password").Action(command.run)
	hash.Flag("cost", "The hashing cost to use").Short('c').IntVar(&command.Cost)
}
