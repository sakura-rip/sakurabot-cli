package product

import (
	"github.com/sakura-rip/sakurabot-cli/internal/database"
	"github.com/sakura-rip/sakurabot-cli/pkg/actor"
	"github.com/sakura-rip/sakurabot-cli/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
	"strings"
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
	name        string `validate:"required"`
	description string `validate:"required"`
	price       int    `validate:"gte=1"`
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
	name, err := actor.PromptAndRetry(actor.Input("product name"), actor.CheckNotEmpty)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	p.name = name

	description, err := actor.Prompt(actor.Input("product description"))
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	p.description = description

	price, err := actor.PromptAndRetry(actor.Input("amount"), func(s string) error {
		_, err := strconv.Atoi(s)
		return err
	})
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	n, _ := strconv.Atoi(price)
	p.price = n

	tags, err := actor.Prompt(actor.Input("user tags"))
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	if tags != "" {
		p.tags = strings.Split(tags, ",")
	}
}

// runAddCommand execute "product add" command
func runAddCommand(cmd *cobra.Command, args []string) {
	if cmd.Flags().NFlag() == 0 {
		addParam.processInteract(args)
	}
	addParam.processParams(args)

	product := &database.Product{
		Model:       nil,
		Name:        addParam.name,
		Description: addParam.description,
		Price:       addParam.price,
		Tags:        database.StringsToDBTags(addParam.tags),
	}
	result := database.Create(product)
	if result.Error != nil {
		logger.Fatal().Err(result.Error).Msg("")
	}
	product.Print()
}
