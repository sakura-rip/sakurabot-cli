package user

import (
	"github.com/sakura-rip/sakurabot-cli/internal/actor"
	"github.com/sakura-rip/sakurabot-cli/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"strconv"
	"strings"
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
	name    string
	tags    []string
	mids    []string
	email   string
	balance int
	group   string
}

// getFlagSet returns the flagSet for addParams
func (a *addParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)
	fs.StringVarP(&addParam.name, "name", "n", "", "user name")
	fs.StringArrayVarP(&addParam.tags, "tags", "t", []string{}, "tag names")
	fs.StringArrayVarP(&addParam.mids, "mids", "m", []string{}, "mids")
	fs.StringVarP(&addParam.email, "email", "e", "", "email")
	fs.IntVarP(&addParam.balance, "balance", "b", 0, "balance")
	fs.StringVarP(&addParam.group, "group", "g", utils.GenUUID(), "specific group")
	return fs
}

// processParams process parameters variable
func (a *addParams) processParams(args []string) {

}

// processInteract process interact parameter initializer
func (a *addParams) processInteract(args []string) {
	name, err := actor.Actor.PromptAndRetry(actor.Input("user name"), actor.CheckNotEmpty)
	if err != nil {
		utils.Logger.Error().Err(err).Msg("")
	}
	a.name = name

	tags, err := actor.Actor.Prompt(actor.Input("user tags"))
	if err != nil {
		utils.Logger.Error().Err(err).Msg("")
	}
	a.tags = strings.Split(tags, ",")

	mids, err := actor.Actor.Prompt(actor.Input("user mids"))
	if err != nil {
		utils.Logger.Error().Err(err).Msg("")
	}
	a.mids = strings.Split(mids, ",")

	email, err := actor.Actor.Prompt(actor.Input("user email"))
	if err != nil {
		utils.Logger.Error().Err(err).Msg("")
	}

	a.email = email

	balance, err := actor.Actor.PromptAndRetry(actor.Input("user balance"), actor.CheckIsAPositiveNumber)
	if err != nil {
		utils.Logger.Error().Err(err).Msg("")
	}
	n, _ := strconv.Atoi(balance)
	a.balance = n

	group, err := actor.Actor.PromptAndRetry(actor.Input("user group"), actor.CheckNotEmpty)
	if err != nil {
		utils.Logger.Error().Err(err).Msg("")
	}
	a.group = group
}

// runAddCommand execute "user add" command
func runAddCommand(cmd *cobra.Command, args []string) {
	if pflag.NFlag() == 0 {
		addParam.processInteract(args)
	}
	addParam.processParams(args)

}
