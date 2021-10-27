const { QuranService } = require("m3o/quran");

// Get a summary for a given chapter (surah)
async function getChapterSummary() {
  let quranService = new QuranService(process.env.M3O_API_TOKEN);
  let rsp = await quranService.summary({
    chapter: 1,
  });
  console.log(rsp);
}

getChapterSummary();
