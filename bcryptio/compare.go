// Package bcryptio provides minimal wrappers around the bcrypt package
// to allow for command-line operations such as comparing passwords, checking
// costs, generating hashes, and generating random passwords with their hashes.
package bcryptio

import (
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

// Compare reads a previously hashed password from the input reader and a plain
// text password, then compares them using bcrypt.
//
// If the plain text password matches the hashed password, it prints a success
// message; otherwise, it returns an error.
func Compare(inr io.Reader, outw io.Writer) error {
	fmt.Fprintln(outw, "Enter previously hashed password:")
	hashed, err := readInput(inr)
	if err != nil {
		return err
	}

	fmt.Fprintln(outw, "Enter password:")
	plain, err := readSensitive(inr)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword(hashed, plain); err != nil {
		return err
	}
	fmt.Fprintln(outw, "Password is correct")

	return nil
}
