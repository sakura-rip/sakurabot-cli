package token

import (
	"github.com/spf13/cobra"
)

// TokenCommand creates the "token" command
func TokenCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "token",
		Short: "token management",
	}
	cmd.AddCommand(
		CreateCommand(),
	)
	return cmd
}
