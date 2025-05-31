package bcryptio

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

const (
	// DefaultLength is the default length of the a password.
	DefaultLength = 20

	// Bcrypt does not support lengths longer than this.
	maxLength = 72

	// The printable characters that can be used for a password with symbols.
	extendedAlphabet = "!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~"

	// The number of printable characters in the extended alphabet.
	extendedAlphabetLen = 94
)

// Generate creates a random password of the specified length, hashes it using
// bcrypt with the specified cost, and writes both the plaintext password and
// the hashed password to the output writer.
//
// If symbols is true, the password will include symbols; otherwise, it will only
// include alphanumeric characters.
//
// If the length is less than DefaultLength, it will be set to DefaultLength.
//
// If the length exceeds maxLength, it will be set to maxLength.
//
// The cost parameter determines the computational cost of hashing the password.
//
// If an error occurs during password generation or hashing, it will be returned.
func Generate(outw io.Writer, symbols bool, length, cost int) error {
	if length < DefaultLength {
		length = DefaultLength
	}
	if length > maxLength {
		length = maxLength
	}
	var data []byte
	if symbols {
		data = randomBytesWithSymbols(length)
	} else {
		data = randomBytes(length)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword(data, cost)
	if err != nil {
		return err
	}
	fmt.Fprintln(outw, "Plaintext:", string(data))
	fmt.Fprintln(outw, "Bcrypt:", string(hashedPassword))

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
