package user

import (
	"github.com/sakura-rip/sakurabot-cli/internal/actor"
	"github.com/sakura-rip/sakurabot-cli/internal/database"
	"github.com/sakura-rip/sakurabot-cli/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/go-playground/validator.v9"
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
	name    string `validate:"required,gt=0,lt=20"`
	tags    []string
	mids    []string
	email   string
	balance int
	group   string `validate:"required,gt=0,lt=34"`
}

// getFlagSet returns the flagSet for addParams
func (p *addParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)
	fs.StringVarP(&addParam.name, "name", "n", "", "user name")
	fs.StringArrayVarP(&addParam.tags, "tags", "t", []string{}, "tag names")
	fs.StringArrayVarP(&addParam.mids, "mids", "m", []string{}, "mids")
	fs.StringVarP(&addParam.email, "email", "e", "", "email")
	fs.IntVarP(&addParam.balance, "balance", "b", 0, "balance")
	fs.StringVarP(&addParam.group, "group", "g", "", "specific group")
	return fs
}

func (p *addParams) validate() error {
	return validator.New().Struct(p)
}

// processParams process parameters variable
func (p *addParams) processParams(args []string) {
	if err := p.validate(); err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
}

// processInteract process interact parameter initializer
func (p *addParams) processInteract(args []string) {
	name, err := actor.Actor.PromptAndRetry(actor.Input("user name"), actor.CheckNotEmpty)
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	p.name = name

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

	email, err := actor.Actor.Prompt(actor.Input("user email"))
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	p.email = email

	balance, err := actor.Actor.PromptAndRetry(actor.Input("user balance"), actor.CheckIsAPositiveNumber)
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	n, _ := strconv.Atoi(balance)
	p.balance = n

	group, err := actor.Actor.Prompt(actor.Input("user group"))
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	p.group = group
}

func (p *addParams) DBTags() []*database.Tag {
	var tags []*database.Tag
	for _, tagName := range p.tags {
		var tag *database.Tag
		result := database.Client.Where(&database.Tag{Name: tagName}).First(tag)
		if result.RowsAffected == 0 {
			tag = &database.Tag{Name: tagName}
			database.Client.Create(tag)
		}
		tags = append(tags, tag)
	}
	return tags
}

func (p *addParams) DBMids() []*database.String {
	var mids []*database.String
	for _, m := range p.mids {
		mids = append(mids, &database.String{Value: m})
	}
	return mids
}

// runAddCommand execute "user add" command
func runAddCommand(cmd *cobra.Command, args []string) {
	if cmd.Flags().NFlag() == 0 {
		addParam.processInteract(args)
	}
	addParam.processParams(args)

	user := &database.User{
		Name:    addParam.name,
		Tags:    addParam.DBTags(),
		Mids:    addParam.DBMids(),
		Email:   addParam.email,
		Balance: addParam.balance,
		Group:   addParam.group,
	}
	result := database.Client.Create(user)
	if result.Error != nil {
		utils.Logger.Fatal().Err(result.Error).Msg("")
	}
	user.Print()
}
