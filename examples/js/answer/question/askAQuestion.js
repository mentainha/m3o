const { AnswerService } = require("m3o/answer");

// Ask a question and receive an instant answer
async function askAquestion() {
  let answerService = new AnswerService(process.env.M3O_API_TOKEN);
  let rsp = await answerService.question({
    query: "microsoft",
  });
  console.log(rsp);
}

askAquestion();
