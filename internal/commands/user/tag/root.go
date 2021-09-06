package tag

import "github.com/spf13/cobra"

// NewCommand creates the "user tag" command
func NewCommand() *cobra.Command {
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
