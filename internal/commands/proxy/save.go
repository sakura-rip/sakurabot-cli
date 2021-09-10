package proxy

import (
	"fmt"
	"github.com/sakura-rip/sakurabot-cli/internal/database"
	actor "github.com/sakura-rip/sakurabot-cli/pkg/actor"
	"github.com/sakura-rip/sakurabot-cli/pkg/file"
	"github.com/sakura-rip/sakurabot-cli/pkg/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

var saveParam = new(saveParams)

// SaveCommand base command for "token save"
func SaveCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save",
		Short: "save token",
		Run:   runSaveCommand,
	}
	cmd.Flags().AddFlagSet(saveParam.getFlagSet())
	return cmd
}

// saveParams add commands parameter
type saveParams struct {
	formatType string
	textPath   string
}

// getFlagSet returns the flagSet for saveParams
func (p *saveParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)
	fs.StringVarP(&saveParam.formatType, "format", "f", "brightdata-datacenter", "specific proxy type(brightdata-datacenter)")
	fs.StringVarP(&saveParam.textPath, "path", "p", "", "specific proxy data path")
	return fs
}

// validate validate parameters
func (p *saveParams) validate() error {
	return validator.New().Struct(p)
}

// processParams process parameters variable
func (p *saveParams) processParams(args []string) {
	if err := p.validate(); err != nil {
		logger.Fatal().Err(err).Msg("")
	}
}

// processInteract process interact parameter initializer
func (p *saveParams) processInteract(args []string) {
	formatType, err := actor.PromptOptional(actor.Input("proxy formatType"), "brightdata-datacenter")
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	p.formatType = formatType

	textPath, err := actor.Prompt(actor.Input("proxy textPath"), actor.CheckNotEmpty)
	if err != nil {
		logger.Fatal().Err(err).Msg("")
	}
	p.textPath = textPath
}

func parseBrightDataDataCenterFormat(path string) []*database.Proxy {
	var proxies []*database.Proxy
	for _, line := range file.ReadFileLines(path) {
		proxyStr := strings.Split(line, ":")
		host := proxyStr[0]
		port := proxyStr[1]
		userId := proxyStr[2]
		password := proxyStr[3]
		proxyUrl := fmt.Sprintf("https://%v:%v@%v:%v", userId, password, host, port)
		proxies = append(proxies, &database.Proxy{
			URL:      proxyUrl,
			UserId:   userId,
			Password: password,
			Host:     host,
			Port:     port,
			IP:       userId[3:],
		})
	}
	return proxies
}

// runSaveCommand execute "token save" command
func runSaveCommand(cmd *cobra.Command, args []string) {
	if cmd.Flags().NFlag() == 0 {
		saveParam.processInteract(args)
	}
	saveParam.processParams(args)

	var proxies []*database.Proxy
	switch saveParam.formatType {
	case "brightdata-datacenter":
		proxies = parseBrightDataDataCenterFormat(saveParam.textPath)
	}
	for _, p := range proxies {
		result := database.Create(&p)
		if result.Error != nil {
			logger.Error().Err(result.Error).Msg("")
		}
	}
}
