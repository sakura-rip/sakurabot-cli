package utils

import (
	"github.com/UpCloudLtd/upcloud-go-api/upcloud/client"
	"github.com/UpCloudLtd/upcloud-go-api/upcloud/service"
	"os"
)

// NewUpcloudClient returns the client of UpCloud.inc api
func NewUpcloudClient() *service.Service {
	return service.New(client.New(os.Getenv("UPCLOUD_USER_NAME"), os.Getenv("UPCLOUD_PASSWORD")))
}
