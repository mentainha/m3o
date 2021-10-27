const { CryptoService } = require("m3o/crypto");

// Get the last price for a given crypto ticker
async function getCryptocurrencyPrice() {
  let cryptoService = new CryptoService(process.env.M3O_API_TOKEN);
  let rsp = await cryptoService.price({
    symbol: "BTCUSD",
  });
  console.log(rsp);
}

getCryptocurrencyPrice();
