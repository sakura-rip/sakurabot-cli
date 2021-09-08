package server

import (
	"github.com/UpCloudLtd/upcloud-go-api/upcloud"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud/request"
	"github.com/sakura-rip/sakurabot-cli/internal/actor"
	"github.com/sakura-rip/sakurabot-cli/internal/database"
	"github.com/sakura-rip/sakurabot-cli/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/go-playground/validator.v9"
	"strconv"
	"strings"
)

var createParam = new(createParams)

// CreateCommand base command for "server create"
func CreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create server",
		Run:   runCreateCommand,
	}
	cmd.Flags().AddFlagSet(createParam.getFlagSet())
	return cmd
}

// createParams add commands parameter
type createParams struct {
	serverType  string `validate:"oneof=upcloud vultr"`
	pubKeyPath  string
	privKeyPath string
	serverName  string
	ipCount     int `validate:"gte=1,lt=6"`
	tags        []string
}

// getFlagSet returns the flagSet for createParams
func (p *createParams) getFlagSet() *pflag.FlagSet {
	fs := new(pflag.FlagSet)
	fs.StringVarP(&createParam.serverType, "type", "t", "", "server type (upcloud | vultr)")
	fs.StringVar(&createParam.pubKeyPath, "pubkey", utils.GetHomeDir()+"/.ssh/id_rsa.pub", "ssh public key")
	fs.StringVar(&createParam.privKeyPath, "privkey", utils.GetHomeDir()+"/.ssh/id_rsa", "ssh private key")

	fs.StringVarP(&createParam.serverName, "name", "n", "", "server name")
	fs.IntVarP(&createParam.ipCount, "ipcount", "c", 1, "server ipv4 address count")
	fs.StringArrayVarP(&createParam.tags, "tags", "t", []string{}, "server tags")
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
	if p.serverName == "" {
		uid := "pro-bot-" + utils.GenUUID()[:5]
		utils.Logger.Info().Msgf("server name not given. use: %v", uid)
		p.serverName = uid
	}
}

// processInteract process interact parameter initializer
func (p *createParams) processInteract(args []string) {
	serverType, err := actor.PromptOptional("server type (upcloud | vultr)", "upcloud")
	if err != nil {
		utils.Logger.Fatal().Err(err).Msgf("")
	}
	p.serverType = serverType

	ipCount, err := actor.PromptAndRetry("ip count", actor.CheckIsAPositiveNumber)
	if err != nil {
		utils.Logger.Fatal().Err(err).Msgf("")
	}
	n, _ := strconv.Atoi(ipCount)
	p.ipCount = n

	tags, err := actor.Prompt("server tags")
	if err != nil {
		utils.Logger.Fatal().Err(err).Msg("")
	}
	if tags != "" {
		p.tags = strings.Split(tags, ",")
	}
}

// createVultrServer create vultr server for createParams value
func (p *createParams) createVultrServer() (*database.Server, error) {
	return nil, nil
}

// asUpcloudIpAddressSlice returns the array of ip address for upcloud server
func (p *createParams) asUpcloudIpAddressSlice() request.CreateServerInterfaceSlice {
	var ips request.CreateServerInterfaceSlice
	for i := 0; i < p.ipCount; i++ {
		ips = append(ips, request.CreateServerInterface{
			IPAddresses: []request.CreateServerIPAddress{{
				Family: upcloud.IPAddressFamilyIPv4,
			}},
			Type: upcloud.IPAddressAccessPublic,
		})
	}
	return ips
}

func (p *createParams) getSSHPublicKey() string {
	if p.pubKeyPath == "" {
		//TODO: use github.com/lxn/walk
		return ""
	}
	return p.pubKeyPath
}

func (p *createParams) getSSHPrivateKey() string {
	if p.privKeyPath == "" {
		//TODO: use github.com/lxn/walk
		return ""
	}
	return p.privKeyPath
}

// createUpcloudServer create upcloud server for createParams value
func (p *createParams) createUpcloudServer() (*database.Server, error) {
	cl := utils.NewUpcloudClient()
	//TODO: handle create server
	_, err := cl.CreateServer(&request.CreateServerRequest{
		Hostname: "sakura-bot",
		Networking: &request.CreateServerNetworking{
			Interfaces: createParam.asUpcloudIpAddressSlice(),
		},
		LoginUser: &request.LoginUser{
			CreatePassword: "no",
			SSHKeys:        []string{createParam.getSSHKey()},
		},
		Plan:     "1xCPU-1GB",
		Title:    createParam.serverName,
		Zone:     "pl-waw1",
		Metadata: upcloud.True,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// runCreateCommand execute "server create" command
func runCreateCommand(cmd *cobra.Command, args []string) {
	if cmd.Flags().NFlag() == 0 {
		createParam.processInteract(args)
	}
	createParam.processParams(args)

	var server *database.Server
	var createErr error
	switch createParam.serverType {
	case "vultr":
		server, createErr = createParam.createVultrServer()
	case "upcloud":
		server, createErr = createParam.createUpcloudServer()
	default:
		utils.Logger.Fatal().Msgf("invalid server type")
	}
	if createErr != nil {
		utils.Logger.Fatal().Err(createErr).Msgf("failed to create server")
	}

	database.Create(&server)
	utils.Logger.Info().Msgf("create server done")
}
