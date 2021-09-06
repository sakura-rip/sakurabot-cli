package server

import (
	"github.com/sakura-rip/sakurabot-cli/internal/actor"
	"github.com/sakura-rip/sakurabot-cli/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
	"strings"
)

var createParam = new(createParams)

// CreateCommand base command for "server create"
func CreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create server",
		Run:   runCreateCommand,
	}
	cmd.Flags().AddFlagSet(createParam.getFlagSet())
	return cmd
}

// createParams add commands parameter
type createParams struct {
	serverType string `validate:"oneof=upcloud vultr"`
	sshKeyPath string
	ipCount    int
	tags       []string
}

// getFlagSet returns the flagSet for createParams
func (p *createParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)

	return fs
}

// validate validate parameters
func (p *createParams) validate() error {
	return validator.New().Struct(p)
}

// processParams process parameters variable
func (p *createParams) processParams(args []string) {
	if err := p.validate(); err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
}

// processInteract process interact parameter initializer
func (p *createParams) processInteract(args []string) {
	serverType, err := actor.PromptOptional("server type", "upcloud")
	if err != nil {
		utils.Logger.Fatal().Err(err).Msgf("")
	}
	p.serverType = serverType

	ipCount, err := actor.PromptAndRetry("ip count", actor.CheckIsAPositiveNumber)
	if err != nil {
		utils.Logger.Fatal().Err(err).Msgf("")
	}
	n, _ := strconv.Atoi(ipCount)
	p.ipCount = n

	tags, err := actor.Prompt(actor.Input("user tags"))
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	if tags != "" {
		p.tags = strings.Split(tags, ",")
	}
}

// runCreateCommand execute "server create" command
func runCreateCommand(cmd *cobra.Command, args []string) {
	if cmd.Flags().NFlag() == 0 {
		createParam.processInteract(args)
	}
	createParam.processParams(args)

}
