package commands

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/alecthomas/kingpin"
	"golang.org/x/crypto/bcrypt"
)

const DEFAULT_LENGTH = 20
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

type GenerateCommand struct {
	Cost   int
	Length int
}

func (command *GenerateCommand) run(context *kingpin.ParseContext) error {
	if command.Length == 0 {
		command.Length = DEFAULT_LENGTH
	}
	randomBytes := randomBytes(command.Length)
	hashedPassword, err := bcrypt.GenerateFromPassword(randomBytes, command.Cost)
	if err != nil {
		panic(err)
	}
	fmt.Println("Plaintext:", string(randomBytes))
	fmt.Println("BCrypt:", string(hashedPassword))
	return nil
}

func randomBytes(length int) []byte {
	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return b
}

func ConfigureGenerateCommand(app *kingpin.Application) {
	command := &GenerateCommand{}
	generate := app.Command("generate", "Output a random password and it's bcrypt hash").Action(command.run)
	generate.Flag("cost", "The hashing cost to use").Short('c').IntVar(&command.Cost)
	generate.Flag("length", "The length of the random password to generate").Short('l').IntVar(&command.Length)
}
