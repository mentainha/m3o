const { WeatherService } = require("m3o/weather");

// Get the current weather report for a location by postcode, city, zip code, ip address
async function getCurrentWeather() {
  let weatherService = new WeatherService(process.env.M3O_API_TOKEN);
  let rsp = await weatherService.now({
    location: "london",
  });
  console.log(rsp);
}

getCurrentWeather();
