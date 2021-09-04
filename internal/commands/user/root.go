package user

import (
	"github.com/sakura-rip/sakurabot-cli/internal/commands/user/tag"
	"github.com/spf13/cobra"
)

// UserCommand creates the "user" command
func UserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "user management",
	}
	cmd.AddCommand(
		AddCommand(),
		ListCommand(),
		GetCommand(),
		ChargeCommand(),
		tag.TagCommand(),
	)
	return cmd
}
