const { CryptoService } = require("m3o/crypto");

// Get news related to a currency
async function getCryptocurrencyNews() {
  let cryptoService = new CryptoService(process.env.M3O_API_TOKEN);
  let rsp = await cryptoService.news({
    symbol: "BTCUSD",
  });
  console.log(rsp);
}

getCryptocurrencyNews();
