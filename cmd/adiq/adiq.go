package cmd

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/ronymmoura/adiq-recurrence-check/internal/adiq"
	"github.com/ronymmoura/adiq-recurrence-check/internal/sql"
	"github.com/ronymmoura/adiq-recurrence-check/internal/util"
	"github.com/ronymmoura/adiq-recurrence-check/internal/xlsx"
)

var (
	op          string = "nao"
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
				Title("Selecione uma operação:").
				Options(
					huh.NewOption("Cruzar dados", "cruzar"),
					huh.NewOption("Corrigir CPF Assinaturas", "corrCpfAss"),
				).
				Value(&op),
		),
	)

	err = form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if op == "corrCpfAss" {
		filter = "nao"
	}

	if op == "cruzar" {
		form = huh.NewForm(
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
			form = huh.NewForm(
				huh.NewGroup(
					huh.NewInput().
						Title("Filtro").
						Value(&filterValue),
				),
			)

			err = form.Run()
			if err != nil {
				log.Fatal(err)
			}
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

	if op == "cruzar" {
		wb := xlsx.CreateFile()
		wb.AddAdiqBillings(billings)
		wb.AddAssinaturas(assinaturas)
		wb.Cross(billings, assinaturas)
		wb.SaveFile("adiq.xlsx")
	}

	if op == "corrCpfAss" {
		count := 0
		fixed := 0

		for _, assinatura := range assinaturas {
			if assinatura.Status == "AGU" {

				plan, err := adiq.GetPlan(accessToken, assinatura.IdPlano)
				if err != nil {
					log.Fatal("Error getting plan:", err)
				}

				nomePlano := plan.Name
				cpf := nomePlano[len(nomePlano)-11:]

				if cpf != assinatura.CPF {
					fmt.Printf("CPF API: %s\nCPF DB : %s\n", cpf, assinatura.CPF)
					fmt.Printf("Plano API: %s\nPlano DB : %s\n\n", plan.Id, assinatura.IdPlano)

					count++

					rows, err := db.UpdateAssinatura(plan.Id, cpf)
					if err != nil {
						log.Fatal("Error updating plan:", err)
					}

					fixed += int(rows)
				}
			}
		}

		fmt.Printf("%d inconsistências encontradas\n", count)
		fmt.Printf("%d inconsistências corrigidas\n", fixed)
	}
}
