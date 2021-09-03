package tag

import "github.com/spf13/cobra"

// TagCommand creates the "user tag" command
func TagCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "tag",
		Short: "tag management",
	}
	cmd.AddCommand(
		AddCommand(),
		RemoveCommand(),
	)
	return cmd
}
