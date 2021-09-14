package product

import (
	"github.com/sakura-rip/sakurabot-cli/internal/database"
	"github.com/sakura-rip/sakurabot-cli/pkg/actor"
	"github.com/sakura-rip/sakurabot-cli/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
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
	if err == nil {
		n, _ := strconv.Atoi(price)
		p.price = n
	}
}

// buildDatabaseQuery return the query for database search
func (p *getParams) buildDatabaseQuery() *gorm.DB {
	query := database.DefaultClient
	if getParam.name != "" {
		query = query.Where("name LIKE ?", "%"+getParam.name+"%")
	}

	if getParam.description != "" {
		query = query.Where("description LIKE ?", "%"+getParam.description+"%")
	}

	if len(getParam.tags) != 0 {
		query = query.Preload("Tags", "name IN ? ", getParam.tags)
	} else {
		query = query.Preload("Tags")
	}

	if getParam.price != 0 {
		query = query.Where("price = ?", getParam.price)
	}
	return query
}

// runGetCommand execute "product get" command
func runGetCommand(cmd *cobra.Command, args []string) {
	if cmd.Flags().NFlag() == 0 {
		getParam.processInteract(args)
	}
	getParam.processParams(args)

	var products []*database.Product
	if getParam.buildDatabaseQuery().Find(&products).RowsAffected == 0 {
		logger.Fatal().Msg("no products found")
	}

	//validate tags
	var resultProducts []*database.Product
	for _, product := range products {
		if len(getParam.tags) != 0 && len(product.Tags) == 0 {
			continue
		}
		resultProducts = append(resultProducts, product)
	}
	if len(resultProducts) == 0 {
		logger.Fatal().Msg("no products found")
	}
	database.PrintProducts(resultProducts)
}
