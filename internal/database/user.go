package database

import (
	"gorm.io/gorm"
)

type String struct {
	*gorm.Model

	Referer int
	Value   string
}

type User struct {
	*gorm.Model

	Name    string    `gorm:"unique"`
	Tags    []*Tag    `gorm:"many2many:user_tag"`
	Mids    []*String `gorm:"foreignKey:Referer"`
	Email   string
	Balance int
	Group   string
}
