package commands

import (
	"bufio"
	"fmt"
	"os"

	"github.com/alecthomas/kingpin/v2"
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
		panic(err)
	}
	fmt.Print("Enter password:")
	password, err := readPassword()
	if err != nil {
		panic(err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), password)
	if err != nil {
		fmt.Println("Password is not correct")
		os.Exit(1)
	}
	fmt.Println("Password is correct")

	return nil
}
