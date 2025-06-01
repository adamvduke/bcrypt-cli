package cmd

import (
	"github.com/spf13/cobra"

	"github.com/adamvduke/bcrypt-cli/bcryptio"
)

const (
	costCmd      = "cost"
	costCmdUsage = "Print the hashing cost used to create the given hash"
)

func NewCostCommand() *cobra.Command {
	return &cobra.Command{
		Use:   costCmd,
		Short: costCmdUsage,
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return bcryptio.Cost(cmd.InOrStdin(), cmd.OutOrStdout())
		},
	}
}
