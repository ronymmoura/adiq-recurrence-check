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
		{Name: "Plano", Width: 40.0},
		{Name: "Assinatura", Width: 40.0},
		{Name: "Status", Width: 15.0},
		{Name: "Data Pagamento", Width: 15.0},
		{Name: "Valor", Width: 10.0},
		{Name: "ID Pagamento", Width: 40.0},
		{Name: "Codigo Autorizacao", Width: 40.0},
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

	for _, assinatura := range assinaturas {
		for _, pagamento := range assinatura.Pagamentos {
			row = sh.AddRow()

			cell := row.AddCell()
			cell.SetString(assinatura.CPF)

			cell = row.AddCell()
			cell.SetString(assinatura.SqPlanoPrevidencial)

			cell = row.AddCell()
			cell.SetString(assinatura.IdPlano)

			cell = row.AddCell()
			cell.SetString(assinatura.IdAssinat.String)

			cell = row.AddCell()
			cell.SetString(assinatura.Status)

			cell = row.AddCell()
			cell.SetDateTime(pagamento.DataPagamento)

			cell = row.AddCell()
			cell.SetNumeric(fmt.Sprintf("%.2f", pagamento.Valor))

			cell = row.AddCell()
			cell.SetString(pagamento.IdPagamento)

			cell = row.AddCell()
			cell.SetString(pagamento.CodigoAutorizacao)

			cell = row.AddCell()
			cell.SetString(pagamento.Lancado)
		}
	}

	return nil
}
