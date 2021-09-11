package product

import "github.com/spf13/cobra"

// NewCommand creates the "product" command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "product",
		Short: "product management",
	}
	cmd.AddCommand(
		AddCommand(),
		RemoveCommand(),
	)
	return cmd
}
