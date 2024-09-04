package cmd

import (
	"log"

	"github.com/ronymmoura/adiq-recurrence-check/internal/adiq"
	"github.com/ronymmoura/adiq-recurrence-check/internal/util"
	"github.com/ronymmoura/adiq-recurrence-check/internal/xlsx"
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

	wb := xlsx.CreateFile()
	xlsx.AddAdiqBillings(wb, billings)
	xlsx.SaveFile("adiq.xlsx", wb)
}
