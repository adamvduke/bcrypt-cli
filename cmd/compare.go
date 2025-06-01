// Package cmd provides command wrappers around the bcryptio package
// to allow for command-line operations such as comparing passwords, checking
// costs, generating hashes, and generating random passwords with their hashes.
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/adamvduke/bcrypt-cli/bcryptio"
)

const (
	cmpCmd   = "compare"
	cmpUsage = "Compare a previously hashed password to a plain text password"
)

func NewCompareCommand() *cobra.Command {
	return &cobra.Command{
		Use:   cmpCmd,
		Short: cmpUsage,
		Args:  cobra.NoArgs,
		RunE: func(cmp *cobra.Command, args []string) error {
			return bcryptio.Compare(cmp.InOrStdin(), cmp.OutOrStdout())
		},
	}
}
