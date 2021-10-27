const { HelloworldService } = require("m3o/helloworld");

// Stream returns a stream of "Hello $name" responses
async function streamsAreCurrentlyTemporarilyNotSupportedInClients() {
  let helloworldService = new HelloworldService(process.env.M3O_API_TOKEN);
  let rsp = await helloworldService.stream({
    name: "not supported",
  });
  console.log(rsp);
}

streamsAreCurrentlyTemporarilyNotSupportedInClients();
