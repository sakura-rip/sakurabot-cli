package database

import "gorm.io/gorm"

// Proxy database table for proxy
type Proxy struct {
	*gorm.Model
	URL    string
	IsUsed bool

	UserId   string
	Password string
	Host     string
	Port     string
	IP       string `gorm:"unique"`
	UnUsable bool
}
