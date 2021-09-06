package database

import (
	"errors"
	"github.com/jedib0t/go-pretty/v6/table"
	"gorm.io/gorm/clause"
	"os"
	"strings"
)

// GetTags return the string array of user tag
func (u *User) GetTags() []string {
	var tags []string
	for _, t := range u.Tags {
		tags = append(tags, t.Name)
	}
	return tags
}

// GetMids return the string array of user mid
func (u *User) GetMids() []string {
	var mids []string
	for _, m := range u.Mids {
		mids = append(mids, m.Value)
	}
	return mids
}

// Print output user information to console
func (u *User) Print() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"id", "name", "tags", "email", "balance", "group"})
	t.AppendRows([]table.Row{{u.ID, u.Name, strings.Join(u.GetTags(), ","), u.Email, u.Balance, u.Group}})
	t.AppendSeparator()
	t.Render()
}

// PrintUsers output users information given to args to console
func PrintUsers(users []*User) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"id", "name", "tags", "email", "balance", "group"})
	var data []table.Row
	for _, u := range users {
		data = append(data, table.Row{
			u.ID, u.Name, strings.Join(u.GetTags(), ","), u.Email, u.Balance, u.Group,
		})
	}
	t.AppendRows(data)
	t.AppendSeparator()
	t.Render()
}

// GetUser return user which has same "id" given to param
func GetUser(id interface{}) (*User, error) {
	var user *User
	result := Preload(clause.Associations).First(&user, id)
	if result.RowsAffected == 0 {
		return nil, errors.New("no user found")
	}
	return user, nil
}
