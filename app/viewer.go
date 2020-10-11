package app

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func TableView(p []*PeepData) {
	// header := []string{"Name", "Sheet", "A1", "B2"}
	// data := [][]string{
	// 	{"tes1.xlsx", "Sheet1", "compute", "fukuoka"},
	// 	{"tes2.xlsx", "Sheet1", "computer", "fukuoka"},
	// 	{"tes3.xlsx", "Sheet1", "compute", "tokyo"},
	// }

	header := p[0].GetTableHeader()
	data := [][]string{}
	for _, d := range p {
		data = append(data, d.GetTableSrc())
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetAutoWrapText(false)
	table.SetAutoFormatHeaders(true)
	table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetCenterSeparator("")
	table.SetColumnSeparator("")
	table.SetRowSeparator("")
	table.SetHeaderLine(false)
	table.SetBorder(false)
	table.SetTablePadding("\t") // pad with tabs
	table.SetNoWhiteSpace(true)
	table.AppendBulk(data) // Add Bulk Data
	table.Render()
}
