package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aquasecurity/table"
	"github.com/liamg/tml"
	"github.com/ronymmoura/adiq-recurrence-check/internal/adiq"
	"github.com/ronymmoura/adiq-recurrence-check/internal/util"
)

func Run() {

	config, err := util.LoadConfig(".env")
	if err != nil {
		log.Fatal("Error loading config file:", err)
	}

	accessToken, err := adiq.Auth(config)
	if err != nil {
		log.Fatal("Cannot autenticate:", err)
	}

	billings, err := adiq.GetBilling(accessToken)
	if err != nil {
		log.Fatal("Cannot get billings:", err)
	}

	t := table.New(os.Stdout)
	t.SetAlignment(table.AlignLeft, table.AlignLeft, table.AlignRight)
	t.SetHeaderStyle(table.StyleBold)
	t.SetLineStyle(table.StyleBlue)
	t.SetDividers(table.UnicodeRoundedDividers)

	t.SetHeaders("Status", "CPF", "Plano", "Status Plano", "Data Criação Plano", "Assinatura", "Status Assinatura", "Data Criação Assinatura", "Valor", "Data Transação", "Data Modificação", "Data Expiração")

	for _, billing := range billings.Items {
		var paidString string

		switch billing.Status {
		case "Paid":
			paidString = "<bold><green>Paid</green></bold>"
		case "Opened":
			paidString = "<bold><magenta>Opened<magenta></bold>"
		case "Denied":
			paidString = "<bold><red>Denied</red></bold>"
		}

		nomePlano := billing.Subscription.Plan.Name
		cpf := nomePlano[len(nomePlano)-11:]

		t.AddRow(tml.Sprintf(paidString), cpf, billing.Subscription.Plan.Id, billing.Subscription.Plan.Status, billing.Subscription.Plan.CreatedDate.Format(time.RFC822), billing.Subscription.Id, billing.Subscription.Status, billing.Subscription.CreatedDate.Format(time.RFC822), fmt.Sprintf("%.2f", billing.Amount), billing.CreatedDate.Format(time.RFC822), billing.ModifiedDate.Format(time.RFC822), billing.ExpireAt.Format(time.RFC822))
	}

	t.Render()
}
