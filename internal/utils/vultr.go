package utils

import (
	"context"
	"os"

	"github.com/vultr/govultr/v2"
	"golang.org/x/oauth2"
)

// NewVultrClient returns the client of UpCloud.com api
func NewVultrClient() *govultr.Client {
	config := &oauth2.Config{}
	ctx := context.Background()
	ts := config.TokenSource(ctx, &oauth2.Token{AccessToken: os.Getenv("VULTR_API_KEY")})
	return govultr.NewClient(oauth2.NewClient(ctx, ts))
}
