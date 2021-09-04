package database

import (
	"github.com/line-org/line-account-generator/generator"
	"gorm.io/gorm"
)

// Token database table for token
type Token struct {
	*gorm.Model
	*generator.Account

	IsUsed bool
	Group  string
	Tags   []*Tag `gorm:"many2many:tokens_tag;"`
}
