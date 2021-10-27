const { FunctionService } = require("m3o/function");

// Delete a function by name
async function deleteAfunction() {
  let functionService = new FunctionService(process.env.M3O_API_TOKEN);
  let rsp = await functionService.delete({
    name: "my-first-func",
    project: "tests",
  });
  console.log(rsp);
}

deleteAfunction();
