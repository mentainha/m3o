package example

import (
	"fmt"
	"os"

	"github.com/micro/services/clients/go/currency"
)

// Rates returns the currency rates for a given code e.g USD
func GetRatesForUsd() {
	currencyService := currency.NewCurrencyService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := currencyService.Rates(&currency.RatesRequest{
		Code: "USD",
	})
	fmt.Println(rsp, err)
}
