package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

const (
	appName  = "bcrypt-cli"
	appUsage = "Wraps golang.org/x/crypto/bcrypt in a cli"
)

func ExecuteContext(ctx context.Context, version string) error {
	return NewRootCmd(version).ExecuteContext(ctx)
}

func NewRootCmd(version string) *cobra.Command {
	root := &cobra.Command{
		Use:   appName,
		Short: appUsage,
		Long:  appUsage + "\n\n" + "A simple CLI for hashing and comparing passwords using bcrypt.",
		Args:  cobra.NoArgs,
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd: true,
		},
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	root.Version = version

	// Add subcommands to the new command
	root.AddCommand(NewCompareCommand())
	root.AddCommand(NewGenerateCommand())
	root.AddCommand(NewHashCommand())
	root.AddCommand(NewVersionCommand())
	root.AddCommand(NewCostCommand())

	return root
}
