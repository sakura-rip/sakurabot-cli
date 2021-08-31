package all

import (
	"github.com/sakura-rip/sakurabot-cli/internal/commands/user"
	"github.com/spf13/cobra"
)

func BuildCommands() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   "sbcli",
		Short: "A manager for sakura bot",
	}

	rootCmd.AddCommand(
		user.UserCommand(),
	)

	return rootCmd, nil
}
