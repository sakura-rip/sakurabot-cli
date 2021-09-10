package product

import (
	"github.com/sakura-rip/sakurabot-cli/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/go-playground/validator.v9"
)

var addParam = new(addParams)

// AddCommand base command for "product add"
func AddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "add product",
		Run:   runAddCommand,
	}
	cmd.Flags().AddFlagSet(addParam.getFlagSet())
	return cmd
}

// addParams add commands parameter
type addParams struct {
	name        string
	description string
	price       int
	tags        []string
}

// getFlagSet returns the flagSet for addParams
func (p *addParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)
	fs.StringVarP(&p.name, "name", "n", "", "product name")
	fs.StringVarP(&p.description, "description", "d", "", "product description")
	fs.IntVarP(&p.price, "price", "p", 0, "product price")
	fs.StringArrayVarP(&p.tags, "tags", "t", []string{}, "product tags")
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

}

// runAddCommand execute "product add" command
func runAddCommand(cmd *cobra.Command, args []string) {
	if cmd.Flags().NFlag() == 0 {
		addParam.processInteract(args)
	}
	addParam.processParams(args)

}
