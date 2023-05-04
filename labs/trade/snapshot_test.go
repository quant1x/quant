package internal

import (
	tableView "github.com/olekukonko/tablewriter"
	"os"
	"testing"
)

func TestQuoteSnapshot_Headers(t *testing.T) {
	tbl := tableView.NewWriter(os.Stdout)
	var quote QuoteSnapshot
	tbl.SetHeader(quote.Headers())
	tbl.Render()
}
