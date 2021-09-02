package user

import (
	"github.com/sakura-rip/sakurabot-cli/internal/actor"
	"github.com/sakura-rip/sakurabot-cli/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

var getParam = new(getParams)

// GetCommand base command for "user get"
func GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get user",
		Run:   runGetCommand,
	}
	cmd.Flags().AddFlagSet(getParam.getFlagSet())
	return cmd
}

// getParams add commands parameter
type getParams struct {
	tags   []string
	groups []string
	name   string
	mids   []string
	email  string
}

// getFlagSet returns the flagSet for getParams
func (p *getParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)
	fs.StringVarP(&getParam.name, "name", "n", "", "user name")
	fs.StringVarP(&getParam.email, "email", "e", "", "user email")
	fs.StringArrayVarP(&getParam.tags, "tags", "t", []string{}, "user tags")
	fs.StringArrayVarP(&getParam.mids, "mids", "m", []string{}, "user mids")
	fs.StringArrayVarP(&getParam.groups, "groups", "g", []string{}, "user groups")
	return fs
}

// validate validate parameters
func (p *getParams) validate() error {
	return validator.New().Struct(p)
}

// processParams process parameters variable
func (p *getParams) processParams(args []string) {
	if err := p.validate(); err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
}

// processInteract process interact parameter initializer
func (p *getParams) processInteract(args []string) {
	name, err := actor.Actor.Prompt(actor.Input("user name"))
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	p.name = name

	email, err := actor.Actor.Prompt(actor.Input("user email"))
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	p.email = email

	tags, err := actor.Actor.Prompt(actor.Input("user tags"))
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	if tags != "" {
		p.tags = strings.Split(tags, ",")
	}

	mids, err := actor.Actor.Prompt(actor.Input("user mids"))
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	if mids != "" {
		p.mids = strings.Split(mids, ",")
	}

	groups, err := actor.Actor.Prompt(actor.Input("user groups"))
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	if groups != "" {
		p.groups = strings.Split(groups, ",")
	}
}

// runGetCommand execute "user get" command
func runGetCommand(cmd *cobra.Command, args []string) {
	if pflag.NFlag() == 0 {
		getParam.processInteract(args)
	}
	getParam.processParams(args)

}
