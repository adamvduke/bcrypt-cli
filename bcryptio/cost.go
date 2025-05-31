package bcryptio

import (
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

// Cost reads a previously hashed password from the input reader and prints the
// cost used to create that hash.
//
// If the hash is invalid, an error is returned.
func Cost(inr io.Reader, outw io.Writer) error {
	fmt.Fprintln(outw, "Enter previously hashed password:")
	hash, err := readInput(inr)
	if err != nil {
		return err
	}
	cost, err := bcrypt.Cost(hash)
	if err != nil {
		return err
	}
	fmt.Fprintln(outw, "Cost:", cost)

	return nil
}
