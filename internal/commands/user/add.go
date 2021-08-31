package user

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var addParam = new(addParams)

// AddCommand base command for "user add"
func AddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add user",
		Run:   runAddCommand,
	}
	cmd.Flags().AddFlagSet(addParam.getFlagSet())
	return cmd
}

// addParams add commands parameter
type addParams struct {
}

// getFlagSet returns the flagSet for addParams
func (a *addParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)

	return fs
}

// processParams process parameters variable
func (a *addParams) processParams(args []string) {

}

// runAddCommand execute "use add" command
func runAddCommand(cmd *cobra.Command, args []string) {
	addParam.processParams(args)

}
