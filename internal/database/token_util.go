package database

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"strings"
)

// PrintTokens output tokens information given to args to console
func PrintTokens(tokens []*Token) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"id", "name", "mid", "tags", "group"})
	var data []table.Row
	for _, token := range tokens {
		data = append(data, table.Row{
			token.ID, token.Name, token.Mid, strings.Join(token.GetTags(), ","), token.Group,
		})
	}
	t.AppendRows(data)
	t.AppendSeparator()
	t.Render()
}
