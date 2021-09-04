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

// GetTags return the string array of user tag
func (u *Token) GetTags() []string {
	var tags []string
	for _, t := range u.Tags {
		tags = append(tags, t.Name)
	}
	return tags
}
