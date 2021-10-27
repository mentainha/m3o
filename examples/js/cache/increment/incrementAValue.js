const { CacheService } = require("m3o/cache");

// Increment a value (if it's a number)
async function incrementAvalue() {
  let cacheService = new CacheService(process.env.M3O_API_TOKEN);
  let rsp = await cacheService.increment({
    key: "counter",
    value: 2,
  });
  console.log(rsp);
}

incrementAvalue();
