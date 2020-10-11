package app

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func TableView(p [][]*PeepData, fileNum int) {
	header := p[0][0].GetTableHeader()
	dataFlat := [][]string{}
	for _, peepData := range p {
		for _, d := range peepData {
			dataFlat = append(dataFlat, d.GetTableSrc())
		}
	}

	data := [][]string{}
	k := 0
	for i := 0; i < (len(dataFlat) / fileNum); i++ {
		k = i
		for j := 0; j < fileNum; j++ {
			if j != 0 {
				k += len(dataFlat) / fileNum
			}
			data = append(data, dataFlat[k])
		}
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
