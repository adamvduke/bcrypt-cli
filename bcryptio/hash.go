package bcryptio

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

var (
	errConfirmationMismatch = errors.New("password and confirmation do not match")
)

// Hash reads a password from the input reader, confirms it by reading it again,
// and then hashes it using bcrypt with the specified cost. The hashed password
// is written to the output writer.
//
// If the passwords do not match, an error is returned.
func Hash(inr io.Reader, out io.Writer, cost int) error {
	fmt.Fprintln(out, "Enter password:")
	password, err := readSensitive(inr)
	if err != nil {
		return err
	}

	fmt.Fprintln(out, "Confirm password:")
	confirmation, err := readSensitive(inr)
	if err != nil {
		return err
	}

	if !bytes.Equal(password, confirmation) {
		return errConfirmationMismatch
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(password, cost)
	if err != nil {
		return err
	}
	fmt.Fprintln(out, "Bcrypt:", string(hashedPassword))

	return nil
}
