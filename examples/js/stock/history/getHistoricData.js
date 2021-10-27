const { StockService } = require("m3o/stock");

// Get the historic open-close for a given day
async function getHistoricData() {
  let stockService = new StockService(process.env.M3O_API_TOKEN);
  let rsp = await stockService.history({
    date: "2020-10-01",
    stock: "AAPL",
  });
  console.log(rsp);
}

getHistoricData();
