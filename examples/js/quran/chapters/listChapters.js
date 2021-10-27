const { QuranService } = require("m3o/quran");

// List the Chapters (surahs) of the Quran
async function listChapters() {
  let quranService = new QuranService(process.env.M3O_API_TOKEN);
  let rsp = await quranService.chapters({
    language: "en",
  });
  console.log(rsp);
}

listChapters();
