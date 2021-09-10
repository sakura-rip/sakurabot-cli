package database

import "gorm.io/gorm"

type Product struct {
	*gorm.Model

	Name        string `gorm:"unique"`
	Description string
	Price       int
	Tags        []*Tag `gorm:"many2many:product_tag"`
}
