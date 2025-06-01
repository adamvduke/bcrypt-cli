package cmd

import (
	"github.com/spf13/cobra"

	"github.com/adamvduke/bcrypt-cli/bcryptio"
)

const (
	hashCmd   = "hash"
	hashUsage = "Use bcrypt to hash a password"
)

func NewHashCommand() *cobra.Command {
	var cost int
	hsh := &cobra.Command{
		Use:   hashCmd,
		Short: hashUsage,
		Args:  cobra.NoArgs,
	}
	hsh.Flags().IntVarP(&cost, costFlag, "c", costFlagDefault, costFlagUsage)
	hsh.RunE = func(cmd *cobra.Command, args []string) error {
		return bcryptio.Hash(hsh.InOrStdin(), hsh.OutOrStdout(), cost)
	}

	return hsh
}
