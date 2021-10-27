const { ForexService } = require("m3o/forex");

// Returns the data for the previous close
async function getPreviousClose() {
  let forexService = new ForexService(process.env.M3O_API_TOKEN);
  let rsp = await forexService.history({
    symbol: "GBPUSD",
  });
  console.log(rsp);
}

getPreviousClose();
