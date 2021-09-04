package proxy

import "github.com/spf13/cobra"

// ProxyCommand creates the "proxy" command
func ProxyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proxy",
		Short: "proxy management",
	}
	cmd.AddCommand(
		SaveCommand(),
	)
	return cmd
}
