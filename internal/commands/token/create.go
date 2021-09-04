package token

import (
	"github.com/line-org/line-account-generator/generator"
	"github.com/line-org/lineall/lineapp/service/line"
	"github.com/sakura-rip/sakurabot-cli/internal/actor"
	"github.com/sakura-rip/sakurabot-cli/internal/database"
	"github.com/sakura-rip/sakurabot-cli/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
	"strings"
	"sync"
	"time"
)

var createParam = new(createParams)

// CreateCommand base command for "token create"
func CreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create token",
		Run:   runCreateCommand,
	}
	cmd.Flags().AddFlagSet(createParam.getFlagSet())
	return cmd
}

// createParams add commands parameter
type createParams struct {
	count    int
	manually bool
	debug    bool
	appType  string
	proxy    string
	tags     []string
	group    string
}

// getFlagSet returns the flagSet for createParams
func (p *createParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)
	fs.IntVarP(&createParam.count, "count", "c", 1, "tokens count")
	fs.BoolVarP(&createParam.manually, "manually", "m", false, "create manually?")
	fs.BoolVarP(&createParam.debug, "debug", "d", false, "debug")
	fs.StringVarP(&createParam.appType, "apptype", "a", "android", "application type")
	fs.StringVarP(&createParam.proxy, "proxy", "p", "", "proxy url")
	fs.StringArrayVarP(&createParam.tags, "tags", "t", []string{}, "tags")
	fs.StringVarP(&createParam.group, "group", "g", "", "group")
	return fs
}

// validate validate parameters
func (p *createParams) validate() error {
	return validator.New().Struct(p)
}

// processParams process parameters variable
func (p *createParams) processParams(args []string) {
	if err := p.validate(); err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
}

// processInteract process interact parameter initializer
func (p *createParams) processInteract(args []string) {
	count, err := actor.Actor.PromptAndRetry(actor.Input("count"), actor.CheckIsAPositiveNumber)
	if err != nil {
		utils.Logger.Fatal().Err(err)
	}
	n, _ := strconv.Atoi(count)
	p.count = n

	appType, err := actor.Actor.PromptOptional(actor.Input("appType"), "android")
	if err != nil {
		utils.Logger.Fatal().Err(err)
	}
	p.appType = appType

	tags, err := actor.Actor.Prompt(actor.Input("tags"))
	if err != nil {
		utils.Logger.Fatal().Err(err)
	}
	if tags != "" {
		p.tags = strings.Split(tags, ",")
	}

	group, err := actor.Actor.Prompt(actor.Input("group"))
	if err != nil {
		utils.Logger.Fatal().Err(err)
	}
	p.group = group
}

func (p *createParams) getProxy() *database.Proxy {
	if p.proxy == "" {
		return database.GetRandomFreeProxy(1)[0]
	}
	return database.ParseProxyUrl(p.proxy)
}

func (p *createParams) getAppType() line.ApplicationType {
	switch p.appType {
	case "android":
		return line.ApplicationType_ANDROID
	case "ios":
		return line.ApplicationType_IOS
	case "lite":
		return line.ApplicationType_ANDROIDLITE
	}
	utils.Logger.Fatal().Msg("wrong app type")
	return line.ApplicationType(-1)
}

// runCreateCommand execute "token create" command
func runCreateCommand(cmd *cobra.Command, args []string) {
	if cmd.Flags().NFlag() == 0 {
		createParam.processInteract(args)
	}
	createParam.processParams(args)
	if createParam.count <= 0 {
		utils.Logger.Fatal().Msgf("count must be at least 1")
	}
	var successCount int
	wg := &sync.WaitGroup{}
	for i := 0; i < createParam.count; i++ {
		wg.Add(1)
		proxy := createParam.getProxy()
		cl := generator.NewClient()
		cl.Setting.Debug = createParam.debug
		cl.Setting.AppType = createParam.getAppType()
		cl.Setting.Proxy = proxy.URL
		go func() {
			defer wg.Done()
			err := cl.Start()
			if err != nil {
				utils.Logger.Error().Err(err).Int("idx", i).Msg("failed to generate account")
				return
			}
			result, err := cl.GetResult()
			if err != nil {
				utils.Logger.Error().Err(err).Int("idx", i).Msg("failed to get generation result")
				return
			}
			successCount++
			utils.Logger.Info().Int("idx", i).Msgf("successfully create account %v/%v", successCount, createParam.count)
			database.Create(&database.Token{
				Account: result,
				Group:   createParam.group,
				Tags:    database.StringsToDBTags(createParam.tags),
			})
			proxy.IsUsed = true
			database.Save(proxy)
		}()
		time.Sleep(time.Second * 5)
	}
	wg.Wait()
	utils.Logger.Info().Msgf("create %d accounts done, Group: %v", successCount, createParam.group)
}
