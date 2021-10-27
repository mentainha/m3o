const { QuranService } = require("m3o/quran");

// Lookup the verses (ayahs) for a chapter including
// translation, interpretation and breakdown by individual
// words.
async function getVersesOfAchapter() {
  let quranService = new QuranService(process.env.M3O_API_TOKEN);
  let rsp = await quranService.verses({
    chapter: 1,
  });
  console.log(rsp);
}

getVersesOfAchapter();
