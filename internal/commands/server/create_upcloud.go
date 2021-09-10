package server

import (
	"github.com/UpCloudLtd/upcloud-go-api/upcloud"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud/request"
	"github.com/sakura-rip/sakurabot-cli/internal/database"
	"github.com/sakura-rip/sakurabot-cli/internal/utils"
)

// createUpcloudServer create upcloud server for createParams value
func (p *createParams) createUpcloudServer() (*database.Server, error) {
	cl := utils.NewUpcloudClient()
	//TODO: handle create server
	detail, err := cl.CreateServer(&request.CreateServerRequest{
		Hostname: "sakura-bot",
		Networking: &request.CreateServerNetworking{
			Interfaces: createParam.asUpcloudIpAddressSlice(),
		},
		LoginUser: &request.LoginUser{
			CreatePassword: "no",
			SSHKeys:        []string{string(utils.ReadAll(createParam.getSSHPublicKeyPath()))},
		},
		Plan:     "1xCPU-1GB",
		Title:    createParam.serverName,
		Zone:     "pl-waw1",
		Metadata: upcloud.True,
	})
	if err != nil {
		return nil, err
	}
	return &database.Server{
		IP:         detail.IPAddresses[0].Address,
		ServerType: database.ServerTypeUPCLOUD,
		UserName:   "root",
		SSHKeyPath: createParam.getSSHPrivateKeyPath(),
		Tags:       database.StringsToDBTags(p.tags),
		UUID:       detail.UUID,
	}, nil
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
