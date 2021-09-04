package database

import "gorm.io/gorm"

// Tag database table for tag
type Tag struct {
	*gorm.Model
	Name        string `gorm:"unique"`
	Description string

	Tokens []*Token `gorm:"many2many:tokens_tag;"`
	Users  []*User  `gorm:"many2many:user_tag"`
}
