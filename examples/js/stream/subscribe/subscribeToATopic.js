const { StreamService } = require("m3o/stream");

// Subscribe to messages for a given topic.
async function subscribeToAtopic() {
  let streamService = new StreamService(process.env.M3O_API_TOKEN);
  let rsp = await streamService.subscribe({
    topic: "events",
  });
  console.log(rsp);
}

subscribeToAtopic();
