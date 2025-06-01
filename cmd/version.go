package cmd

import (
	"text/template"

	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of bcrypt-cli",
		Args:  cobra.NoArgs,
		RunE: func(ver *cobra.Command, args []string) error {
			root := ver.Root()
			templateString := root.VersionTemplate()
			templ := template.Must(template.New("version").Parse(templateString))
			return templ.Execute(ver.OutOrStdout(), root)
		},
	}
}
