package xlsx

import (
	"fmt"

	"github.com/ronymmoura/adiq-recurrence-check/internal/sql"
	"github.com/tealeg/xlsx/v3"
)

func (wb *Workbook) AddAssinaturas(assinaturas []sql.Assinatura) error {
	sheetName := "Assinaturas BD"
	sh, err := wb.AddSheet(sheetName)
	if err != nil {
		return err
	}

	headers := []ColHeader{
		{Name: "CPF", Width: 15.0},
		{Name: "Plano Prev", Width: 15.0},
		{Name: "Plano", Width: 50.0},
		{Name: "Assinatura", Width: 50.0},
		{Name: "Status", Width: 15.0},
		{Name: "Data Pagamento", Width: 20.0},
		{Name: "Valor", Width: 10.0},
		{Name: "ID Pagamento", Width: 50.0},
		{Name: "Codigo Autorizacao", Width: 20.0},
		{Name: "Lançado", Width: 15.0},
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

	aguStyles := xlsx.NewStyle()
	aguStyles.Font.Bold = true
	aguStyles.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")
	aguStyles.Fill.FgColor = "0060adfe"
	aguStyles.Fill.PatternType = "solid"
	aguStyles.ApplyFill = true
	aguStyles.ApplyBorder = true

	atiStyles := xlsx.NewStyle()
	atiStyles.Font.Bold = true
	atiStyles.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")
	atiStyles.Fill.FgColor = "0030c130"
	atiStyles.Fill.PatternType = "solid"
	atiStyles.ApplyFill = true

	canStyles := xlsx.NewStyle()
	canStyles.Font.Bold = true
	canStyles.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")
	canStyles.Fill.FgColor = "00ff0400"
	canStyles.Fill.PatternType = "solid"
	canStyles.ApplyFill = true

	payStyles := xlsx.NewStyle()
	payStyles.Border = *xlsx.NewBorder("thin", "thin", "thin", "thin")
	payStyles.Fill.FgColor = "00F1F1F1"
	payStyles.Fill.PatternType = "solid"
	payStyles.ApplyFill = true

	for _, assinatura := range assinaturas {
		row = sh.AddRow()

		var style *xlsx.Style

		if assinatura.Status == "AGU" {
			style = aguStyles
		} else if assinatura.Status == "ATI" {
			style = atiStyles
		} else if assinatura.Status == "CAN" {
			style = canStyles
		}

		cell := row.AddCell()
		cell.SetStyle(style)
		cell.SetString(assinatura.CPF)

		cell = row.AddCell()
		cell.SetStyle(style)
		cell.SetString(assinatura.SqPlanoPrevidencial)

		cell = row.AddCell()
		cell.SetStyle(style)
		cell.SetString(assinatura.IdPlano)

		cell = row.AddCell()
		cell.SetStyle(style)
		if assinatura.IdAssinat != nil {
			cell.SetString(assinatura.IdAssinat.String)
		} else {
			cell.SetString("-")
		}

		cell = row.AddCell()
		cell.SetStyle(style)
		cell.SetString(assinatura.Status)

		cell = row.AddCell()
		cell.SetStyle(style)
		cell.SetString("")

		cell = row.AddCell()
		cell.SetStyle(style)
		cell.SetString("")

		cell = row.AddCell()
		cell.SetStyle(style)
		cell.SetString("")

		cell = row.AddCell()
		cell.SetStyle(style)
		cell.SetString("")

		cell = row.AddCell()
		cell.SetStyle(style)
		cell.SetString("")

		for _, pagamento := range assinatura.Pagamentos {
			row = sh.AddRow()

			cell := row.AddCell()
			cell.SetStyle(payStyles)
			cell.SetString("")

			cell = row.AddCell()
			cell.SetStyle(payStyles)
			cell.SetString("")

			cell = row.AddCell()
			cell.SetStyle(payStyles)
			cell.SetString("")

			cell = row.AddCell()
			cell.SetStyle(payStyles)
			cell.SetString("")

			cell = row.AddCell()
			cell.SetStyle(payStyles)
			cell.SetString("")

			cell = row.AddCell()
			cell.SetStyle(payStyles)
			cell.SetDateTime(pagamento.DataPagamento)

			cell = row.AddCell()
			cell.SetStyle(payStyles)
			cell.SetNumeric(fmt.Sprintf("%.2f", pagamento.Valor))

			cell = row.AddCell()
			cell.SetStyle(payStyles)
			cell.SetString(pagamento.IdPagamento)

			cell = row.AddCell()
			cell.SetStyle(payStyles)
			cell.SetString(pagamento.CodigoAutorizacao)

			cell = row.AddCell()
			cell.SetStyle(payStyles)
			cell.SetString(pagamento.Lancado)
		}
	}

	return nil
}
