const { TwitterService } = require("m3o/twitter");

// Get the current global trending topics
async function getTheCurrentGlobalTrendingTopics() {
  let twitterService = new TwitterService(process.env.M3O_API_TOKEN);
  let rsp = await twitterService.trends({});
  console.log(rsp);
}

getTheCurrentGlobalTrendingTopics();
