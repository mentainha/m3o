package example

import (
	"fmt"
	"os"

	"github.com/micro/services/clients/go/stock"
)

// Get the last price for a given stock ticker
func GetAstockPrice() {
	stockService := stock.NewStockService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := stockService.Price(&stock.PriceRequest{
		Symbol: "AAPL",
	})
	fmt.Println(rsp, err)
}
