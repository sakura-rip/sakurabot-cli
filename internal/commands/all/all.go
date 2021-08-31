package all

import (
	"github.com/spf13/cobra"
)

func BuildCommands() (*cobra.Command, error) {
	rootCmd := &cobra.Command{
		Use:   "sbcli",
		Short: "A manager for sakura bot",
	}
	return rootCmd, nil
}
