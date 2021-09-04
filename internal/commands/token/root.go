package token

import (
	"github.com/spf13/cobra"
)

// TokenCommand creates the "token" command
func TokenCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "user management",
	}
	cmd.AddCommand()
	return cmd
}
