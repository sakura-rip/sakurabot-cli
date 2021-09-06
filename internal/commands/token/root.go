package token

import (
	"github.com/spf13/cobra"
)

// NewCommand creates the "token" command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token",
		Short: "token management",
	}
	cmd.AddCommand(
		CreateCommand(),
		GetCommand(),
	)
	return cmd
}
