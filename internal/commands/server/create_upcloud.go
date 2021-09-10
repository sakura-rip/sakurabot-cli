package server

import (
	"github.com/UpCloudLtd/upcloud-go-api/upcloud"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud/client"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud/request"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud/service"
	"github.com/sakura-rip/sakurabot-cli/internal/database"
	"github.com/sakura-rip/sakurabot-cli/pkg/file"
	"github.com/sakura-rip/sakurabot-cli/pkg/logger"
	"os"
	"time"
)

// isUpcloudServerTagCreated check if tag is already created at upcloud
func (p *createParams) isUpcloudServerTagCreated(tag *upcloud.Tags, tagName string) bool {
	for _, t := range tag.Tags {
		if t.Name == tagName {
			return true
		}
	}
	return false
}

// tagUpcloudServer tag upcloud server
func (p *createParams) tagUpcloudServer(cl *service.Service, uuid string, tags []string) error {
	upcloudTags, err := cl.GetTags()
	if err != nil {
		return err
	}
	for _, tagName := range tags {
		if p.isUpcloudServerTagCreated(upcloudTags, tagName) {
			if _, err := cl.TagServer(&request.TagServerRequest{UUID: uuid, Tags: []string{tagName}}); err != nil {
				return err
			}
		} else {
			if _, err := cl.CreateTag(&request.CreateTagRequest{Tag: upcloud.Tag{Name: tagName, Servers: []string{uuid}}}); err != nil {
				return err
			}
		}
	}
	return nil
}

// createUpcloudServer create upcloud server for createParams value
func (p *createParams) createUpcloudServer() (*database.Server, error) {
	cl := service.New(client.New(os.Getenv("UPCLOUD_USER_NAME"), os.Getenv("UPCLOUD_PASSWORD")))
	detail, err := cl.CreateServer(&request.CreateServerRequest{
		Hostname: "sakura-bot",
		Networking: &request.CreateServerNetworking{
			Interfaces: createParam.asUpcloudIpAddressSlice(),
		},
		LoginUser: &request.LoginUser{
			CreatePassword: "no",
			SSHKeys:        []string{string(file.ReadAll(createParam.getSSHPublicKeyPath()))},
		},
		Plan:     "1xCPU-1GB",
		Title:    createParam.serverName,
		Zone:     "pl-waw1",
		Metadata: upcloud.True,
	})
	if err != nil {
		return nil, err
	}
	logger.Info().Msgf("waiting for server to start")

	if _, err = cl.WaitForServerState(&request.WaitForServerStateRequest{
		UUID:         detail.UUID,
		DesiredState: upcloud.ServerStateStarted,
		Timeout:      time.Minute * 7,
	}); err != nil {
		logger.Fatal().Err(err).Msgf("failed to wait for server start")
	}
	logger.Info().Msgf("server started")
	if err := p.tagUpcloudServer(cl, detail.UUID, createParam.tags); err != nil {
		logger.Fatal().Err(err).Msgf("failed to tag server")
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
