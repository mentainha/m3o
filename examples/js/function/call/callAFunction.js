const { FunctionService } = require("m3o/function");

// Call a function by name
async function callAfunction() {
  let functionService = new FunctionService(process.env.M3O_API_TOKEN);
  let rsp = await functionService.call({
    name: "my-first-func",
    request: {},
  });
  console.log(rsp);
}

callAfunction();
