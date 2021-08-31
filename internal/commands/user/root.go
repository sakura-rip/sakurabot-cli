package user

import "github.com/spf13/cobra"

// UserCommand creates the "user" command
func UserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "user management",
	}
	cmd.AddCommand(
		AddCommand(),
	)
	return cmd
}
