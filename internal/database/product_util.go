package database

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"strings"
)

// GetTags return the string array of user tag
func (p *Product) GetTags() []string {
	var tags []string
	for _, t := range p.Tags {
		tags = append(tags, t.Name)
	}
	return tags
}

func (p *Product) Print() {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"id", "name", "description", "tags", "price"})
	t.AppendRows([]table.Row{{p.ID, p.Name, p.Description, strings.Join(p.GetTags(), ","), p.Price}})
	t.AppendSeparator()
	t.Render()
}

// PrintProducts output users information given to args to console
func PrintProducts(users []*Product) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"id", "name", "description", "tags", "price"})
	var data []table.Row
	for _, p := range users {
		data = append(data, table.Row{p.ID, p.Name, p.Description, strings.Join(p.GetTags(), ","), p.Price})
	}
	t.AppendRows(data)
	t.AppendSeparator()
	t.Render()
}
