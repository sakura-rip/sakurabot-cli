package user

import (
	"github.com/sakura-rip/sakurabot-cli/internal/commands/user/tag"
	"github.com/spf13/cobra"
)

// NewCommand creates the "user" command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "user management",
	}
	cmd.AddCommand(
		AddCommand(),
		ListCommand(),
		GetCommand(),
		ChargeCommand(),
		tag.NewCommand(),
	)
	return cmd
}
