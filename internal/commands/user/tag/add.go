package tag

import (
	"github.com/sakura-rip/sakurabot-cli/internal/database"
	actor "github.com/sakura-rip/sakurabot-cli/pkg/actor"
	"github.com/sakura-rip/sakurabot-cli/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
	"strings"
)

var addParam = new(addParams)

// AddCommand base command for "user tag add"
func AddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add user tag",
		Run:   runAddCommand,
	}
	cmd.Flags().AddFlagSet(addParam.getFlagSet())
	return cmd
}

// addParams add commands parameter
type addParams struct {
	userId int
	tags   []string
}

// getFlagSet returns the flagSet for addParams
func (p *addParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)
	fs.IntVarP(&p.userId, "userid", "u", 0, "user id")
	fs.StringArrayVarP(&p.tags, "tags", "t", []string{}, "tags")
	return fs
}

// validate validate parameters
func (p *addParams) validate() error {
	return validator.New().Struct(p)
}

// processParams process parameters variable
func (p *addParams) processParams(args []string) {
	if err := p.validate(); err != nil {
		logger.Fatal().Err(err).Msg("")
	}
}

// processInteract process interact parameter initializer
func (p *addParams) processInteract(args []string) {
	userId, err := actor.PromptAndRetry("user id", actor.CheckIsAPositiveNumber, func(s string) error {
		user, err := database.GetUser(s)
		if err != nil {
			return err
		}
		logger.Info().Msgf("user name: %v", user.Name)
		return nil
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	n, _ := strconv.Atoi(userId)
	p.userId = n

	tags, err := actor.Prompt("user tags", actor.CheckNotEmpty)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	if tags != "" {
		p.tags = strings.Split(tags, ",")
	}
}

func (p *addParams) DBTags() []*database.Tag {
	var tags []*database.Tag
	for _, tagName := range p.tags {
		var tag *database.Tag
		result := database.Where(&database.Tag{Name: tagName}).First(tag)
		if result.RowsAffected == 0 {
			tag = &database.Tag{Name: tagName}
			database.Create(tag)
		}
		tags = append(tags, tag)
	}
	return tags
}

// runAddCommand execute "user tag add" command
func runAddCommand(cmd *cobra.Command, args []string) {
	if cmd.Flags().NFlag() == 0 {
		addParam.processInteract(args)
	}
	addParam.processParams(args)
	user, err := database.GetUser(addParam.userId)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	err = database.Model(user).Association("Tags").Append(addParam.DBTags())
	if err != nil {
		logger.Error().Err(err).Msg("")
	}
	logger.Info().Msgf("DONE: add %v tags to user: [%v]", len(addParam.tags), user.Name)
}
