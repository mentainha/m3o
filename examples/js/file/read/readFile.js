const { FileService } = require("m3o/file");

// Read a file by path
async function readFile() {
  let fileService = new FileService(process.env.M3O_API_TOKEN);
  let rsp = await fileService.read({
    path: "/document/text-files/file.txt",
    project: "examples",
  });
  console.log(rsp);
}

readFile();
