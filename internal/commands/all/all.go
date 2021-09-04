package all

import (
	"github.com/sakura-rip/sakurabot-cli/internal/commands/token"
	"github.com/sakura-rip/sakurabot-cli/internal/commands/user"
	"github.com/spf13/cobra"
)

// BuildAllCommands initialize all commands
func BuildAllCommands() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   "sbcli",
		Short: "A manager for sakura bot",
	}

	rootCmd.AddCommand(
		user.UserCommand(),
		token.TokenCommand(),
	)

	return rootCmd, nil
}
