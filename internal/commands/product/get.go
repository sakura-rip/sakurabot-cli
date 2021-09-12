package product

import (
	"github.com/sakura-rip/sakurabot-cli/pkg/actor"
	"github.com/sakura-rip/sakurabot-cli/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
	"strings"
)

var getParam = new(getParams)

// GetCommand base command for "product get"
func GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get product",
		Run:   runGetCommand,
	}
	cmd.Flags().AddFlagSet(getParam.getFlagSet())
	return cmd
}

// getParams add commands parameter
type getParams struct {
	name        string
	description string
	price       int
	tags        []string
}

// getFlagSet returns the flagSet for getParams
func (p *getParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)
	fs.StringVarP(&p.name, "name", "n", "", "name")
	fs.StringVarP(&p.description, "description", "d", "", "descripton")
	fs.IntVarP(&p.price, "price", "p", 0, "price")
	fs.StringArrayVarP(&p.tags, "tags", "t", []string{}, "tags")
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
	name, err := actor.Prompt("product name")
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	p.name = name

	description, err := actor.Prompt("product description")
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	p.description = description

	tags, err := actor.Prompt("product tags")
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	if tags != "" {
		p.tags = strings.Split(tags, ",")
	}

	price, err := actor.Prompt("product price", func(s string) error {
		_, err := strconv.Atoi(s)
		return err
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	n, _ := strconv.Atoi(price)
	p.price = n
}

// runGetCommand execute "product get" command
func runGetCommand(cmd *cobra.Command, args []string) {
	if cmd.Flags().NFlag() == 0 {
		getParam.processInteract(args)
	}
	getParam.processParams(args)

}
