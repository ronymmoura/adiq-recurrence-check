package xlsx

import (
	"github.com/ronymmoura/adiq-recurrence-check/internal/sql"
	"github.com/tealeg/xlsx/v3"
)

func (wb *Workbook) AddFichaContrib(ficha []sql.FichaFinanceira, correcao bool) error {
	var sheetName string
	if correcao {
		sheetName = "Ficha Financeira Correção"
	} else {
		sheetName = "Ficha Financeira"
	}

	sh, err := wb.AddSheet(sheetName)
	if err != nil {
		return err
	}

	headers := []ColHeader{
		{Name: "Status", Width: 15.0},
		{Name: "Nome", Width: 30.0},
		{Name: "CPF", Width: 15.0},
		{Name: "Tipo Cobranca", Width: 10.0},
		{Name: "Data Competência", Width: 15.0},
		{Name: "Data Referencia", Width: 15.0},
		{Name: "Data Aporte", Width: 15.0},
		{Name: "Valor", Width: 10.0},
	}

	row := sh.AddRow()

	headerStyles := xlsx.NewStyle()
	headerStyles.Font.Bold = true

	for idx, header := range headers {
		cell := row.AddCell()
		cell.SetStyle(headerStyles)
		cell.SetString(header.Name)

		sh.SetColWidth(idx+1, idx+1, header.Width)
	}

	duplicatedStyles := xlsx.NewStyle()
	duplicatedStyles.Font.Bold = true
	duplicatedStyles.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")
	duplicatedStyles.Fill.FgColor = "00ff0400"
	duplicatedStyles.Fill.PatternType = "solid"
	duplicatedStyles.ApplyFill = true

	okStyles := xlsx.NewStyle()
	okStyles.Font.Bold = true
	okStyles.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")
	okStyles.Fill.FgColor = "0030c130"
	okStyles.Fill.PatternType = "solid"
	okStyles.ApplyFill = true

	for _, item := range ficha {
		row = sh.AddRow()

		var style *xlsx.Style

		cell := row.AddCell()

		if item.Exists(ficha) {
			style = duplicatedStyles
			cell.SetString("Duplicado")
		} else {
			style = okStyles
			cell.SetString("OK")
		}
		cell.SetStyle(style)

		cell = row.AddCell()
		cell.SetStyle(style)
		cell.SetString(item.NoPessoa)

		cell = row.AddCell()
		cell.SetStyle(style)
		cell.SetString(item.CPF)

		cell = row.AddCell()
		cell.SetStyle(style)
		cell.SetInt(item.SqTipoCobranca)

		cell = row.AddCell()
		cell.SetStyle(style)
		cell.SetDateTime(item.DtCompetencia)

		cell = row.AddCell()
		cell.SetStyle(style)
		cell.SetDateTime(item.DtReferencia)

		cell = row.AddCell()
		cell.SetStyle(style)
		cell.SetDateTime(item.DtAporte)

		cell = row.AddCell()
		cell.SetStyle(style)
		cell.SetFloat(item.VlContribuicao)
	}

	return nil
}
