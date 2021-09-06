package database

import "gorm.io/gorm"

type ServerType string

const (
	ServerTypeUPCLOUD ServerType = "upcloud"
	ServerTypeVULTR   ServerType = "vultr"
)

type Server struct {
	*gorm.Model

	IP         string
	ServerType ServerType
	UserName   string
	Password   string
	SSHKeyPath string
	Tags       []*Tag `gorm:"many2many:server_tag;"`
	UUID       string
}
