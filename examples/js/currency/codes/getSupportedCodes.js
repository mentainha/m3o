const { CurrencyService } = require("m3o/currency");

// Codes returns the supported currency codes for the API
async function getSupportedCodes() {
  let currencyService = new CurrencyService(process.env.M3O_API_TOKEN);
  let rsp = await currencyService.codes({});
  console.log(rsp);
}

getSupportedCodes();
