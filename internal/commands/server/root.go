package server

import "github.com/spf13/cobra"

// NewCommand creates the "server" command
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "server management",
	}
	cmd.AddCommand()
	return cmd
}
