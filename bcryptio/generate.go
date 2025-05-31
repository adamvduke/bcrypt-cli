package bcryptio

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/alecthomas/kingpin/v2"

	"golang.org/x/crypto/bcrypt"
)

const (
	defaultLength = 20

	// Bcrypt does not support lengths longer than this.
	maxLength = 72

	// The printable characters that can be used for a password with symbols.
	extendedAlphabet = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"

	// The number of printable characters in the extended alphabet.
	extendedAlphabetLen = 94
)

type GenerateCommand struct {
	Cost           int
	Length         int
	IncludeSymbols bool
	Out            io.Writer
}

func ConfigureGenerateCommand(app *kingpin.Application, outputWriter io.Writer) {
	command := &GenerateCommand{Out: outputWriter}
	generate := app.Command("generate", "Output a random password and it's bcrypt hash").Action(command.Run)
	generate.Flag("cost", "The hashing cost to use").Short('c').IntVar(&command.Cost)
	generate.Flag("symbols", "If the generated password should contain symbols").Short('s').BoolVar(&command.IncludeSymbols)
	generate.Flag("length", "The length of the random password to generate").Short('l').IntVar(&command.Length)
}

func (command *GenerateCommand) Run(context *kingpin.ParseContext) error {
	if command.Length < defaultLength {
		command.Length = defaultLength
	}
	if command.Length > maxLength {
		command.Length = maxLength
	}
	var data []byte
	if command.IncludeSymbols {
		data = randomBytesWithSymbols(command.Length)
	} else {
		data = randomBytes(command.Length)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(data, command.Cost)
	if err != nil {
		return err
	}
	fmt.Fprintln(command.Out, "Plaintext:", string(data))
	fmt.Fprintln(command.Out, "Bcrypt:", string(hashedPassword))

	return nil
}

// This is based on the implementation of crypto/rand.Text but with an extended
// alphabet.
func randomBytesWithSymbols(length int) []byte {
	src := make([]byte, length)

	// The crypto/rand package is documented as always filling the buffer and
	// never returning an error.
	rand.Read(src)
	for i := range src {
		src[i] = extendedAlphabet[src[i]%extendedAlphabetLen]
	}
	return src
}

func randomBytes(length int) []byte {
	buf := bytes.Buffer{}
	for buf.Len() < length {
		buf.WriteString(rand.Text())
	}
	return buf.Bytes()[:length]
}
