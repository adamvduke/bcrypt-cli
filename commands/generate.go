package commands

import (
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/howeyc/gopass"
	"golang.org/x/crypto/bcrypt"
)

type GenerateCommand struct {
	Cost int
}

func (command *GenerateCommand) run(context *kingpin.ParseContext) error {
	fmt.Print("Enter password: ")
	password, err := gopass.GetPasswdMasked()
	if err != nil {
		panic(err)
	}
	fmt.Print("Confirm password: ")
	confirmation, err := gopass.GetPasswdMasked()
	if err != nil {
		panic(err)
	}

	if string(password) != string(confirmation) {
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

func ConfigureGenerateCommand(app *kingpin.Application) {
	command := &GenerateCommand{}
	generate := app.Command("generate", "Use bcrypt to hash a password").Action(command.run)
	generate.Flag("cost", "The hashing cost to use").Short('c').IntVar(&command.Cost)
}
