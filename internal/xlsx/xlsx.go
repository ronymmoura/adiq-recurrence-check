package xlsx

import (
	"fmt"

	"github.com/ronymmoura/adiq-recurrence-check/internal/adiq"
	"github.com/tealeg/xlsx/v3"
)

func CreateFile() *xlsx.File {
	wb := xlsx.NewFile()

	return wb
}

func SaveFile(path string, wb *xlsx.File) error {
	err := wb.Save(path)

	return err
}

type Header struct {
	Name  string
	Width float64
}

func AutoWidth(s string) float64 {
	return 20.0
}

func AddAdiqBillings(wb *xlsx.File, billings []adiq.Billing) (*xlsx.File, error) {
	sheetName := "API Billings"
	sh, err := wb.AddSheet(sheetName)
	if err != nil {
		return nil, err
	}

	headers := []Header{
		{Name: "Status", Width: 20.0},
		{Name: "CPF", Width: 15.0},
		{Name: "Plano", Width: 40.0},
		{Name: "Status Plano", Width: 15.0},
		{Name: "Data Criação Plano", Width: 20.0},
		{Name: "Assinatura", Width: 40.0},
		{Name: "Status Assinatura", Width: 15.0},
		{Name: "Data Criação Assinatura", Width: 20.0},
		{Name: "Valor", Width: 10.0},
		{Name: "Data Transação", Width: 20.0},
		{Name: "Data Modificação", Width: 20.0},
		{Name: "Data Expiração", Width: 20.0},
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

	for _, billing := range billings {
		row = sh.AddRow()

		nomePlano := billing.Subscription.Plan.Name
		cpf := nomePlano[len(nomePlano)-11:]

		// Cell Status
		cell := row.AddCell()
		statusStyle := xlsx.NewStyle()
		statusStyle.Font.Bold = true

		switch billing.Status {
		case "Paid":
			statusStyle.Font.Color = "FF30c130"
		case "Opened":
			statusStyle.Font.Color = "FF60adfe"
		case "Denied":
			statusStyle.Font.Color = "00ff0400"
		case "PaymentInvalid":
			statusStyle.Font.Color = "00ff0400"
		default:
			statusStyle.Font.Color = "00000000"
		}

		cell.SetStyle(statusStyle)
		cell.SetString(billing.Status)

		// Cell CPF
		cell = row.AddCell()
		cell.SetString(cpf)

		// Cell Plan ID
		cell = row.AddCell()
		cell.SetString(billing.Subscription.Plan.Id)

		cell = row.AddCell()
		cell.SetString(billing.Subscription.Plan.Status)

		cell = row.AddCell()
		cell.SetDateTime(billing.Subscription.Plan.CreatedDate.Time)

		cell = row.AddCell()
		cell.SetString(billing.Subscription.Id)

		cell = row.AddCell()
		cell.SetString(billing.Subscription.Status)

		cell = row.AddCell()
		cell.SetDateTime(billing.Subscription.CreatedDate.Time)

		cell = row.AddCell()
		cell.SetString(fmt.Sprintf("%.2f", billing.Amount))

		cell = row.AddCell()
		cell.SetDateTime(billing.CreatedDate.Time)

		cell = row.AddCell()
		cell.SetDateTime(billing.ModifiedDate.Time)

		cell = row.AddCell()
		cell.SetDateTime(billing.ExpireAt.Time)
	}

	return wb, nil
}
