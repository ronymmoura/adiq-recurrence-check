package xlsx

import (
	"fmt"

	"github.com/ronymmoura/adiq-recurrence-check/internal/adiq"
	"github.com/ronymmoura/adiq-recurrence-check/internal/sql"
	"github.com/tealeg/xlsx/v3"
)

func (wb *Workbook) Cross(billings []adiq.Billing, assinaturas []sql.Assinatura) error {
	sheetName := "Cruzamento"
	sh, err := wb.AddSheet(sheetName)
	if err != nil {
		return err
	}

	headers := []ColHeader{
		{Name: "CPF", Width: 15.0},
		{Name: "Plano", Width: 40.0},
		{Name: "Status Plano API", Width: 25.0},
		{Name: "Status Plano DB", Width: 25.0},
		{Name: "Assinatura", Width: 40.0},
		{Name: "Status Assinatura API", Width: 25.0},
		{Name: "Status Assinatura DB", Width: 25.0},
		{Name: "ID Pagamento", Width: 40.0},
		{Name: "Status Pagamento API", Width: 25.0},
		{Name: "Status Pagamento DB", Width: 25.0},
		{Name: "Data Criação Plano API", Width: 20.0},
		{Name: "Data Criação Assinatura API", Width: 20.0},
		{Name: "Data Criação Plano DB", Width: 20.0},
		{Name: "Data Modificação", Width: 20.0},
		{Name: "Data Expiração", Width: 20.0},
		{Name: "Data Pagamento API", Width: 20.0},
		{Name: "Data Pagamento DB", Width: 20.0},
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

	notFoundBillings := []adiq.Billing{}

	for _, billing := range billings {
		for idx, assinatura := range assinaturas {
			if assinatura.IdAssinat != nil && billing.Subscription.Id == assinatura.IdAssinat.String {

				var pagamento *sql.Pagamento
				for _, pag := range assinatura.Pagamentos {
					if pag.IdPagamento == billing.Id {
						pagamento = &pag
					}
				}

				row = sh.AddRow()

				nomePlano := billing.Subscription.Plan.Name
				cpf := nomePlano[len(nomePlano)-11:]

				cell := row.AddCell()
				cell.SetString(cpf)

				cell = row.AddCell()
				cell.SetString(billing.Subscription.Plan.Id)

				cell = row.AddCell()
				cell.SetString(billing.Subscription.Plan.Status)

				cell = row.AddCell()
				cell.SetString(assinatura.Status)

				cell = row.AddCell()
				cell.SetString(billing.Subscription.Id)

				cell = row.AddCell()
				cell.SetString(billing.Subscription.Status)

				cell = row.AddCell()
				if assinatura.IdAssinat != nil {
					cell.SetString("Existente")
				} else {
					cell.SetString("Inexistente")
				}

				cell = row.AddCell()
				cell.SetString(billing.Id)

				cell = row.AddCell()
				cell.SetString(billing.Status)

				cell = row.AddCell()
				if pagamento != nil {
					if pagamento.Lancado == "SIM" {
						cell.SetString("Pago e Lançado")
					} else if pagamento.Lancado == "NAO" {
						cell.SetString("Pago")
					}
				} else {
					cell.SetString("Inexistente")
				}

				cell = row.AddCell()
				cell.SetDateTime(billing.Subscription.Plan.CreatedDate.Time)

				cell = row.AddCell()
				cell.SetDateTime(billing.Subscription.CreatedDate.Time)

				cell = row.AddCell()
				cell.SetDateTime(assinatura.DataCriacao)

				cell = row.AddCell()
				cell.SetDateTime(billing.ModifiedDate.Time)

				cell = row.AddCell()
				cell.SetDateTime(billing.ExpireAt.Time)

				cell = row.AddCell()
				cell.SetDateTime(billing.CreatedDate.Time)

				cell = row.AddCell()
				if pagamento != nil {
					cell.SetDateTime(pagamento.DataPagamento)
				} else {
					cell.SetString("Inexistente")
				}
			}

			if idx == len(assinaturas)-1 {
				notFoundBillings = append(notFoundBillings, billing)
			}
		}
	}

	fmt.Printf("%d not found billings", len(notFoundBillings))

	return nil
}
