package server

import (
	"context"
	"fmt"
	"github.com/sakura-rip/sakurabot-cli/internal/database"
	"github.com/sakura-rip/sakurabot-cli/pkg/file"
	"github.com/sakura-rip/sakurabot-cli/pkg/logger"
	"github.com/vultr/govultr/v2"
	"golang.org/x/oauth2"
	"os"
	"time"
)

// newVultrClient
func (p *createParams) newVultrClient() *govultr.Client {
	config := &oauth2.Config{}
	ctx := context.Background()
	ts := config.TokenSource(ctx, &oauth2.Token{AccessToken: os.Getenv("VULTR_API_KEY")})
	return govultr.NewClient(oauth2.NewClient(ctx, ts))
}

func (p *createParams) waitForVultrServerToStart(cl *govultr.Client, id string) error {
	attempts := 0
	timeout := time.Minute * 5
	for {
		attempts++
		server, err := cl.Instance.Get(context.Background(), id)
		if err != nil {
			return err
		}
		if server.Status == "active" {
			return nil
		}
		sleep := time.Second * 10
		time.Sleep(sleep)
		if time.Duration(attempts)*sleep > timeout {
			return fmt.Errorf("time out: %v", timeout)
		}
	}
}

// createVultrServer create vultr server for createParams value
func (p *createParams) createVultrServer() (*database.Server, error) {
	cl := p.newVultrClient()
	detail, err := cl.Instance.Create(context.Background(), &govultr.InstanceCreateReq{
		Region:     "nrt", //Japan
		Plan:       "vc2-1c-1gb",
		OsID:       413, // Ubuntu
		Hostname:   "sakura-bot",
		EnableIPv6: govultr.BoolToBoolPtr(true),
		Label:      createParam.serverName,
		SSHKeys:    []string{string(file.ReadAll(createParam.getSSHPublicKeyPath()))},
	})
	if err != nil {
		return nil, err
	}

	logger.Info().Msgf("waiting for server to start")

	if err := p.waitForVultrServerToStart(cl, detail.ID); err != nil {
		logger.Fatal().Err(err).Msgf("failed to wait for server start")
	}

	logger.Info().Msgf("server started")

	for i := 0; i < createParam.ipCount; i++ {
		if _, err := cl.Instance.CreateIPv4(context.Background(), detail.ID, govultr.BoolToBoolPtr(false)); err != nil {
			logger.Fatal().Err(err).Msgf("failed to create ipv4")
		}
	}

	if err := cl.Instance.Reboot(context.Background(), detail.ID); err != nil {
		logger.Fatal().Err(err).Msgf("failed to reboot server")
	}

	if err := p.waitForVultrServerToStart(cl, detail.ID); err != nil {
		logger.Fatal().Err(err).Msgf("failed to wait for server start")
	}

	return &database.Server{
		Model:      nil,
		IP:         detail.MainIP,
		ServerType: database.ServerTypeVULTR,
		UserName:   "root",
		SSHKeyPath: createParam.getSSHPrivateKeyPath(),
		Tags:       database.StringsToDBTags(createParam.tags),
		UUID:       detail.ID,
	}, nil
}
