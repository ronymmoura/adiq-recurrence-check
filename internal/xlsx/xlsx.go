package xlsx

import (
	"github.com/tealeg/xlsx/v3"
)

type Workbook struct {
	*xlsx.File
}

func CreateFile() *Workbook {
	wb := xlsx.NewFile()

	book := &Workbook{wb}

	return book
}

func (wb *Workbook) SaveFile(path string) error {
	err := wb.Save(path)

	return err
}
