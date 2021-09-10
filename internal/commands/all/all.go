package all

import (
	"github.com/sakura-rip/sakurabot-cli/internal/commands/product"
	"github.com/sakura-rip/sakurabot-cli/internal/commands/proxy"
	"github.com/sakura-rip/sakurabot-cli/internal/commands/server"
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
		user.NewCommand(),
		token.NewCommand(),
		proxy.NewCommand(),
		server.NewCommand(),
		product.NewCommand(),
	)

	return rootCmd, nil
}
