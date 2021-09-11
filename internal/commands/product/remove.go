package product

import (
	"github.com/sakura-rip/sakurabot-cli/internal/database"
	"github.com/sakura-rip/sakurabot-cli/pkg/actor"
	"github.com/sakura-rip/sakurabot-cli/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
)

var removeParam = new(removeParams)

// RemoveCommand base command for "product remove"
func RemoveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove",
		Short: "remove product",
		Run:   runRemoveCommand,
	}
	cmd.Flags().AddFlagSet(removeParam.getFlagSet())
	return cmd
}

// removeParams add commands parameter
type removeParams struct {
	id int `validate:"required"`
}

// getFlagSet returns the flagSet for removeParams
func (p *removeParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)
	fs.IntVarP(&p.id, "id", "i", 0, "product id")
	return fs
}

// validate validate parameters
func (p *removeParams) validate() error {
	return validator.New().Struct(p)
}

// processParams process parameters variable
func (p *removeParams) processParams(args []string) {
	if err := p.validate(); err != nil {
		logger.Fatal().Err(err).Msg("")
	}
}

// processInteract process interact parameter initializer
func (p *removeParams) processInteract(args []string) {
	price, err := actor.PromptAndRetry(actor.Input("product id"), func(s string) error {
		_, err := strconv.Atoi(s)
		return err
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	n, _ := strconv.Atoi(price)
	p.id = n
}

// runRemoveCommand execute "product remove" command
func runRemoveCommand(cmd *cobra.Command, args []string) {
	if cmd.Flags().NFlag() == 0 {
		removeParam.processInteract(args)
	}
	removeParam.processParams(args)

	var product *database.Product
	if result := database.First(product, removeParam.id); result.Error != nil {
		logger.Fatal().Err(result.Error).Msgf("failed to find product")
	}

	product.Print()
	if result := database.Delete(&database.Product{}, removeParam.id); result.Error != nil {
		logger.Fatal().Err(result.Error).Msgf("failed to delete product")
	}
	logger.Info().Msgf("delete product done")
}
