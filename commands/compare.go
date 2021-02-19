package commands

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/alecthomas/kingpin"
	"github.com/howeyc/gopass"
	"golang.org/x/crypto/bcrypt"
)

type CompareCommand struct {
}

func ConfigureCompareCommand(app *kingpin.Application) {
	command := &CompareCommand{}
	app.Command("compare", "Compare a hashed password to a plain text password").Action(command.run)
}

func (command *CompareCommand) run(context *kingpin.ParseContext) error {
	fmt.Print("Enter hashed password: ")
	reader := bufio.NewReader(os.Stdin)
	hashedPassword, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	fmt.Print("Enter password: ")
	password, err := gopass.GetPasswdMasked()
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), password)
	if err != nil {
		return errors.New("password is not correct")
	}
	fmt.Println("Password is correct")

	return nil
}
