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

func GetAllSheetList(filenames []string) []string {
	sheetList := []string{}

	for _, filename := range filenames {
		f, err := excelize.OpenFile(filename)
		if err != nil {
			continue
		}
		sheetList = MergeSlices(sheetList, f.GetSheetList())
	}

	return sheetList
}

func MergeSlices(s1, s2 []string) []string {
	merged := []string{}
	a := []string{}
	b := []string{}

	if len(s1) >= len(s2) {
		a = s1
		b = s2
	} else {
		a = s2
		b = s1
	}

	merged = a
	for i, e := range a {
		if !(i < len(b)) {
			continue
		}
		if e != b[i] {
			merged = append(merged[:i+1], merged[i:]...)
			merged[i+1] = b[i]
		}
	}
	return merged
}
