package user

import (
	"github.com/sakura-rip/sakurabot-cli/internal/actor"
	"github.com/sakura-rip/sakurabot-cli/internal/database"
	"github.com/sakura-rip/sakurabot-cli/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
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
	tags  []string
	group string
	name  string
	mids  []string
	email string
}

// getFlagSet returns the flagSet for getParams
func (p *getParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)
	fs.StringVarP(&getParam.name, "name", "n", "", "user name")
	fs.StringVarP(&getParam.email, "email", "e", "", "user email")
	fs.StringArrayVarP(&getParam.tags, "tags", "t", []string{}, "user tags")
	fs.StringArrayVarP(&getParam.mids, "mids", "m", []string{}, "user mids")
	fs.StringVarP(&getParam.group, "group", "g", "", "user group")
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
	name, err := actor.Prompt(actor.Input("user name"))
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	p.name = name

	email, err := actor.Prompt(actor.Input("user email"))
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	p.email = email

	tags, err := actor.Prompt(actor.Input("user tags"))
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	if tags != "" {
		p.tags = strings.Split(tags, ",")
	}

	mids, err := actor.Prompt(actor.Input("user mids"))
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	if mids != "" {
		p.mids = strings.Split(mids, ",")
	}

	group, err := actor.Prompt(actor.Input("user group"))
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	p.group = group
}

// buildDatabaseQuery return the query for database search
func (p *getParams) buildDatabaseQuery() *gorm.DB {
	query := database.DefaultClient
	if getParam.name != "" {
		query = query.Where("name LIKE ?", "%"+getParam.name+"%")
	}
	if getParam.email != "" {
		query = query.Where("email LIKE ?", "%"+getParam.email+"%")
	}
	if getParam.group != "" {
		query = query.Where("group LIKE ?", "%"+getParam.group+"%")
	}
	if len(getParam.tags) != 0 {
		query = query.Preload("Tags", "name IN ? ", getParam.tags)
	} else {
		query = query.Preload("Tags")
	}
	if len(getParam.mids) != 0 {
		query = query.Preload("Mids", "value IN ? ", getParam.mids)
	} else {
		query = query.Preload("Mids")
	}
	return query
}

// runGetCommand execute "user get" command
func runGetCommand(cmd *cobra.Command, args []string) {
	if cmd.Flags().NFlag() == 0 {
		getParam.processInteract(args)
	}
	getParam.processParams(args)

	var users []*database.User
	if getParam.buildDatabaseQuery().Find(&users).RowsAffected == 0 {
		utils.Logger.Fatal().Msg("no users found")
	}
	//validate tags and mids
	var resultUsers []*database.User
	for _, user := range users {
		if len(getParam.tags) != 0 && len(user.Tags) == 0 {
			continue
		}
		if len(getParam.mids) != 0 && len(user.Mids) == 0 {
			continue
		}
		resultUsers = append(resultUsers, user)
	}
	if len(resultUsers) == 0 {
		utils.Logger.Fatal().Msg("no users found")
	}
	database.PrintUsers(resultUsers)
}
