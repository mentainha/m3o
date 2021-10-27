const { StockService } = require("m3o/stock");

// Get the last price for a given stock ticker
async function getAstockPrice() {
  let stockService = new StockService(process.env.M3O_API_TOKEN);
  let rsp = await stockService.price({
    symbol: "AAPL",
  });
  console.log(rsp);
}

getAstockPrice();
