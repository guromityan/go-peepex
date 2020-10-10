package app

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// Peep provides a function to get the value
// of the specified Excel file, sheet, cell.
func Peep(path, sheet, cell string) (string, error) {
	f, err := excelize.OpenFile(path)
	if err != nil {
		return "", fmt.Errorf("Cloud not open file %v: %v", path, err)
	}

	i := f.GetSheetIndex(sheet)
	if i == -1 {
		return "** None sheet. **", nil
	}

	v, err := f.GetCellValue(sheet, cell)
	if err != nil {
		return "", err
	}

	return v, err
}
