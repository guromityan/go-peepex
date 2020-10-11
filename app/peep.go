package app

import (
	"fmt"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Peep provides a function to get the value
// of the specified Excel file, sheet, cell.
func Peep(filename, sheets, cells string) ([]*PeepData, error) {
	f, err := excelize.OpenFile(filename)
	if err != nil {
		return nil, fmt.Errorf("Cloud not open file %v: %v", filename, err)
	}

	peepDatas := []*PeepData{}
	for _, sheet := range strings.Split(sheets, ",") {
		sheet = strings.TrimSpace(sheet)
		p := NewPeepData(filename, sheet)
		GetSheetCellValue(p, f, sheet, cells)
		peepDatas = append(peepDatas, p)
	}
	return peepDatas, nil
}

// GetSheetCellValue 指定したシート、セルの値を取得
// シートが存在しない場合はシート名を "None: sheet" とし、セルの数だけ空データを生成
func GetSheetCellValue(p *PeepData, f *excelize.File, sheet, cells string) {
	cellList := []string{}
	for _, cell := range strings.Split(cells, ",") {
		cell = strings.TrimSpace(cell)
		cellList = append(cellList, cell)
	}

	i := f.GetSheetIndex(sheet)
	if i == -1 {
		p.Sheet = fmt.Sprintf("None:%v", sheet)
		for _, cell := range cellList {
			v := NewPeepValue(cell, "")
			p.AddPeepValue(v)
		}
		return
	}

	for _, cell := range cellList {
		cellVal, err := f.GetCellValue(sheet, cell)
		if err != nil {
			cellVal = ""
		}
		v := NewPeepValue(cell, cellVal)
		p.AddPeepValue(v)
	}
}
