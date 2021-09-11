package user

import (
	"github.com/sakura-rip/sakurabot-cli/internal/database"
	actor "github.com/sakura-rip/sakurabot-cli/pkg/actor"
	"github.com/sakura-rip/sakurabot-cli/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
)

var chargeParam = new(chargeParams)

// ChargeCommand base command for "user charge"
func ChargeCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "charge",
		Short: "charge user",
		Run:   runChargeCommand,
	}
	cmd.Flags().AddFlagSet(chargeParam.getFlagSet())
	return cmd
}

// chargeParams add commands parameter
type chargeParams struct {
	userId     int
	amount     int
	chargeType string
}

// getFlagSet returns the flagSet for chargeParams
func (p *chargeParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)
	fs.IntVarP(&p.userId, "userid", "u", 0, "user id to charge")
	fs.IntVarP(&p.amount, "amount", "a", 0, "amount to charge")
	fs.StringVarP(&p.chargeType, "type", "t", "", "charge type")
	return fs
}

// validate validate parameters
func (p *chargeParams) validate() error {
	return validator.New().Struct(p)
}

// processParams process parameters variable
func (p *chargeParams) processParams(args []string) {
	if err := p.validate(); err != nil {
		logger.Fatal().Err(err).Msg("")
	}
}

// processInteract process interact parameter initializer
func (p *chargeParams) processInteract(args []string) {
	uid, err := actor.PromptAndRetry(actor.Input("user id "), actor.CheckIsAPositiveNumber, func(s string) error {
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
	u, _ := strconv.Atoi(uid)
	p.userId = u

	amount, err := actor.PromptAndRetry(actor.Input("amount"), func(s string) error {
		_, err := strconv.Atoi(s)
		return err
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	n, _ := strconv.Atoi(amount)
	p.amount = n

	type_, err := actor.PromptOptional(actor.Input("type"), "amazon")
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	p.chargeType = type_

}

// runChargeCommand execute "user charge" command
func runChargeCommand(cmd *cobra.Command, args []string) {
	if cmd.Flags().NFlag() == 0 {
		chargeParam.processInteract(args)
	}
	chargeParam.processParams(args)

	charge := &database.Charge{
		Amount: chargeParam.amount,
		Type:   database.ChargeType(chargeParam.chargeType),
		UserId: chargeParam.userId,
	}

	user, err := database.GetUser(chargeParam.userId)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	user.Balance += chargeParam.amount
	database.Save(user)

	err = database.Model(user).Association("Charges").Append(charge)
	if err != nil {
		logger.Error().Err(err).Msg("")
	}
	logger.Info().Msgf("DONE: update amount to user:[%v]  %v JPY ", user.Name, chargeParam.amount)
}
