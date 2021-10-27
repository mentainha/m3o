const { FunctionService } = require("m3o/function");

// List all the deployed functions
async function listFunctions() {
  let functionService = new FunctionService(process.env.M3O_API_TOKEN);
  let rsp = await functionService.list({});
  console.log(rsp);
}

listFunctions();
