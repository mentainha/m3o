const { CurrencyService } = require("m3o/currency");

// Rates returns the currency rates for a given code e.g USD
async function getRatesForUsd() {
  let currencyService = new CurrencyService(process.env.M3O_API_TOKEN);
  let rsp = await currencyService.rates({
    code: "USD",
  });
  console.log(rsp);
}

getRatesForUsd();
