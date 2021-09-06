package proxy

import "github.com/spf13/cobra"

// NewCommand creates the "proxy" command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proxy",
		Short: "proxy management",
	}
	cmd.AddCommand(
		SaveCommand(),
	)
	return cmd
}
