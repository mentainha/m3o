const { CacheService } = require("m3o/cache");

// Set an item in the cache. Overwrites any existing value already set.
async function setAvalue() {
  let cacheService = new CacheService(process.env.M3O_API_TOKEN);
  let rsp = await cacheService.set({
    key: "foo",
    value: "bar",
  });
  console.log(rsp);
}

setAvalue();
