package user

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var addParam = new(addParams)

func AddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add user",
		Run:   runAddCommand,
	}
	cmd.Flags().AddFlagSet(addParam.getFlagSet())
	return cmd
}

type addParams struct {
}

func (a *addParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)

	return fs
}

func (a *addParams) processParams(args []string) {

}

func runAddCommand(cmd *cobra.Command, args []string) {
	addParam.processParams(args)

}
