package database

import (
	"gorm.io/gorm"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
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

func (u *User) GetTags() []string {
	var tags []string
	for _, t := range u.Tags {
		tags = append(tags, t.Name)
	}
	return tags
}

func (u *User) GetMids() []string {
	var mids []string
	for _, m := range u.Mids {
		mids = append(mids, m.Value)
	}
	return mids
}

func (u *User) Print() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"id", "name", "tags", "mids", "email", "balance", "group"})
	t.AppendRows([]table.Row{{u.ID, u.Name, u.GetTags(), u.GetMids(), u.Email, u.Balance, u.Group}})
	t.AppendSeparator()
	t.Render()
}
