package cmd

import (
	"github.com/spf13/cobra"

	"github.com/adamvduke/bcrypt-cli/bcryptio"
)

const (
	generateCmd   = "generate"
	generateUsage = "Output a random password and its bcrypt hash"

	costFlag        = "cost"
	costFlagDefault = 10
	costFlagUsage   = "The cost to use when generating the password hash"

	lengthFlag      = "length"
	lengthFlagUsage = "The length of the random password to generate"

	symbolsFlag      = "symbols"
	symbolsFlagUsage = "If the generated password should contain symbols"
)

func NewGenerateCommand() *cobra.Command {
	var includeSymbols bool
	var length int
	var cost int

	gen := &cobra.Command{
		Use:   generateCmd,
		Short: generateUsage,
		Args:  cobra.NoArgs,
	}
	gen.Flags().BoolVarP(&includeSymbols, symbolsFlag, "s", true, symbolsFlagUsage)
	gen.Flags().IntVarP(&length, lengthFlag, "l", bcryptio.DefaultLength, lengthFlagUsage)
	gen.Flags().IntVarP(&cost, costFlag, "c", costFlagDefault, costFlagUsage)
	gen.RunE = func(cmd *cobra.Command, args []string) error {
		return bcryptio.Generate(gen.OutOrStdout(), includeSymbols, length, cost)
	}

	return gen
}
