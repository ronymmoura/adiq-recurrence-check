package cmd

import (
	"log"

	"github.com/charmbracelet/huh"
	"github.com/ronymmoura/adiq-recurrence-check/internal/adiq"
	"github.com/ronymmoura/adiq-recurrence-check/internal/sql"
	"github.com/ronymmoura/adiq-recurrence-check/internal/util"
	"github.com/ronymmoura/adiq-recurrence-check/internal/xlsx"
)

var (
	filter      string
	filterValue string
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

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Deseja filtrar os resultados?").
				Options(
					huh.NewOption("Não", "nao"),
					huh.NewOption("Por CPF", "cpf"),
					huh.NewOption("Por Plano", "plano"),
					huh.NewOption("Por Assinatura", "assinatura"),
					huh.NewOption("Por Pagamento", "pagamento"),
				).
				Value(&filter),
		),
	)

	err = form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if filter != "nao" {
		form2 := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("Filtro").
					Value(&filterValue),
			),
		)

		err = form2.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

	billings, err := adiq.GetBilling(accessToken, filter, filterValue)
	if err != nil {
		log.Fatal("Cannot get billings:", err)
	}

	db, err := sql.CreateConnection(config)
	if err != nil {
		log.Fatal("Error opening connection:", err)
	}

	assinaturas, err := db.GetAssinaturas(filter, filterValue)
	if err != nil {
		log.Fatal("Error getting subscriptions:", err)
	}

	wb := xlsx.CreateFile()
	wb.AddAdiqBillings(billings)
	wb.AddAssinaturas(assinaturas)
	wb.Cross(billings, assinaturas)
	wb.SaveFile("adiq.xlsx")
}
