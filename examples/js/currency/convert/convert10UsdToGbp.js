const { CurrencyService } = require("m3o/currency");

// Convert returns the currency conversion rate between two pairs e.g USD/GBP
async function convert10usdToGbp() {
  let currencyService = new CurrencyService(process.env.M3O_API_TOKEN);
  let rsp = await currencyService.convert({
    amount: 10,
    from: "USD",
    to: "GBP",
  });
  console.log(rsp);
}

convert10usdToGbp();
