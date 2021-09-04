package database

import (
	"gorm.io/gorm"
)

// String database table to save string array object
type String struct {
	*gorm.Model

	Referer int
	Value   string
}

// User database table for user
type User struct {
	*gorm.Model

	Name    string    `gorm:"unique"`
	Tags    []*Tag    `gorm:"many2many:user_tag"`
	Mids    []*String `gorm:"foreignKey:Referer"`
	Charges []*Charge `gorm:"foreignKey:UserId"`
	Email   string
	Balance int
	Group   string
}
