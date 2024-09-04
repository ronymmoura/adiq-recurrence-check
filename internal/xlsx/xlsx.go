package xlsx

import (
	"github.com/tealeg/xlsx/v3"
)

type Sheet struct {
	*xlsx.File
}

func CreateFile() *Sheet {
	wb := xlsx.NewFile()

	sheet := &Sheet{wb}

	return sheet
}

func (wb *Sheet) SaveFile(path string) error {
	err := wb.Save(path)

	return err
}
