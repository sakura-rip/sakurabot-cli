package tag

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

var removeParam = new(removeParams)

// RemoveCommand base command for "user tag remove"
func RemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "remove user tag",
		Run:   runRemoveCommand,
	}
	cmd.Flags().AddFlagSet(removeParam.getFlagSet())
	return cmd
}

// removeParams add commands parameter
type removeParams struct {
	userId int
	tags   []string
}

// getFlagSet returns the flagSet for removeParams
func (p *removeParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)
	fs.IntVarP(&removeParam.userId, "userid", "u", 0, "user id")
	fs.StringArrayVarP(&removeParam.tags, "tags", "t", []string{}, "tags")
	return fs
}

// validate validate parameters
func (p *removeParams) validate() error {
	return validator.New().Struct(p)
}

// processParams process parameters variable
func (p *removeParams) processParams(args []string) {
	if err := p.validate(); err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
}

// processInteract process interact parameter initializer
func (p *removeParams) processInteract(args []string) {
	userId, err := actor.PromptAndRetry(actor.Input("user id"), actor.CheckIsAPositiveNumber, func(s string) error {
		user, err := database.GetUser(s)
		if err != nil {
			return err
		}
		utils.Logger.Info().Msgf("user name: %v", user.Name)
		return nil
	})
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	n, _ := strconv.Atoi(userId)
	p.userId = n
	tags, err := actor.PromptAndRetry(actor.Input("user tags"), actor.CheckNotEmpty)
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	if tags != "" {
		p.tags = strings.Split(tags, ",")
	}
}

// getMatchedTags return the array of tags that matched to tagNames
func (p *removeParams) getMatchedTags(tags []*database.Tag) []*database.Tag {
	var result []*database.Tag
	for _, t := range tags {
		for _, tn := range p.tags {
			if t.Name == tn {
				result = append(result, t)
			}
		}
	}
	return result
}

// runRemoveCommand execute "user tag remove" command
func runRemoveCommand(cmd *cobra.Command, args []string) {
	if cmd.Flags().NFlag() == 0 {
		removeParam.processInteract(args)
	}
	removeParam.processParams(args)

	user, err := database.GetUser(removeParam.userId)
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}

	err = database.Model(user).Association("Tags").Delete(removeParam.getMatchedTags(user.Tags))
	if err != nil {
		utils.Logger.Error().Err(err).Msg("")
	}
	utils.Logger.Info().Msgf("DONE: remove %v tags from user: [%v]", len(removeParam.tags), user.Name)
}
