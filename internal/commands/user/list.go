package user

import (
	"github.com/sakura-rip/sakurabot-cli/internal/database"
	"github.com/sakura-rip/sakurabot-cli/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm/clause"
)

var listParam = new(listParams)

// ListCommand base command for "user list"
func ListCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list user",
		Run:   runListCommand,
	}
	cmd.Flags().AddFlagSet(listParam.getFlagSet())
	return cmd
}

// listParams add commands parameter
type listParams struct {
}

// getFlagSet returns the flagSet for listParams
func (p *listParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)

	return fs
}

// validate validate parameters
func (p *listParams) validate() error {
	return validator.New().Struct(p)
}

// processParams process parameters variable
func (p *listParams) processParams(args []string) {
	if err := p.validate(); err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
}

// processInteract process interact parameter initializer
func (p *listParams) processInteract(args []string) {

}

// runListCommand execute "user list" command
func runListCommand(cmd *cobra.Command, args []string) {
	if cmd.Flags().NFlag() == 0 {
		listParam.processInteract(args)
	}
	listParam.processParams(args)
	var users []*database.User
	database.Preload(clause.Associations).Find(&users)
	database.PrintUsers(users)
}
