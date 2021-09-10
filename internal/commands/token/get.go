package token

import (
	"github.com/sakura-rip/sakurabot-cli/internal/actor"
	"github.com/sakura-rip/sakurabot-cli/internal/database"
	"github.com/sakura-rip/sakurabot-cli/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

var getParam = new(getParams)

// GetCommand base command for "token get"
func GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get token",
		Run:   runGetCommand,
	}
	cmd.Flags().AddFlagSet(getParam.getFlagSet())
	return cmd
}

// getParams add commands parameter
type getParams struct {
	count int
	tags  []string
	group string
}

// getFlagSet returns the flagSet for getParams
func (p *getParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)
	fs.IntVarP(&getParam.count, "count", "c", 1, "token count")
	fs.StringVarP(&getParam.group, "group", "g", "", "group")
	fs.StringArrayVarP(&getParam.tags, "tags", "t", []string{}, "tags")
	return fs
}

// validate validate parameters
func (p *getParams) validate() error {
	return validator.New().Struct(p)
}

// processParams process parameters variable
func (p *getParams) processParams(args []string) {
	if err := p.validate(); err != nil {
		logger.Fatal().Err(err).Msg("")
	}
}

// processInteract process interact parameter initializer
func (p *getParams) processInteract(args []string) {
	count, err := actor.PromptAndRetry(actor.Input("count"), actor.CheckIsAPositiveNumber)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	n, _ := strconv.Atoi(count)
	p.count = n

	tags, err := actor.Prompt(actor.Input("token tags"))
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	if tags != "" {
		p.tags = strings.Split(tags, ",")
	}

	group, err := actor.Prompt(actor.Input("token group"))
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	p.group = group
}

// buildDatabaseQuery return the query for database search
func (p *getParams) buildDatabaseQuery() *gorm.DB {
	query := database.Limit(getParam.count)
	if getParam.group != "" {
		query = query.Where("group LIKE ?", "%"+getParam.group+"%")
	}
	if len(getParam.tags) != 0 {
		query = query.Preload("Tags", "name IN ? ", getParam.tags)
	} else {
		query = query.Preload("Tags")
	}
	return query
}

// runGetCommand execute "token get" command
func runGetCommand(cmd *cobra.Command, args []string) {
	if cmd.Flags().NFlag() == 0 {
		getParam.processInteract(args)
	}
	getParam.processParams(args)
	var tokens []*database.Token
	if getParam.buildDatabaseQuery().Find(&tokens).RowsAffected == 0 {
		logger.Fatal().Msg("no tokens found")
	}
	//validate tags
	var resultTokens []*database.Token
	for _, token := range tokens {
		if len(getParam.tags) != 0 && len(token.Tags) == 0 {
			continue
		}
		resultTokens = append(resultTokens, token)
	}
	if len(resultTokens) == 0 {
		logger.Fatal().Msgf("no tokens found")
	}
	database.PrintTokens(resultTokens)
}
