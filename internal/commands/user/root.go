package user

import "github.com/spf13/cobra"

func UserCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "user",
		Short: "user management",
	}
	cmd.AddCommand()
	return cmd
}
