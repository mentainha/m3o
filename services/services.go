package services

import (
	"github.com/micro/micro/v3/service/client"
	"m3o.dev/services/address/proto"
	"m3o.dev/services/analytics/proto"
	"m3o.dev/services/answer/proto"
	"m3o.dev/services/app/proto"
	"m3o.dev/services/avatar/proto"
	"m3o.dev/services/bitcoin/proto"
	"m3o.dev/services/cache/proto"
	"m3o.dev/services/carbon/proto"
	"m3o.dev/services/chat/proto"
	"m3o.dev/services/comments/proto"
	"m3o.dev/services/contact/proto"
	"m3o.dev/services/crypto/proto"
	"m3o.dev/services/currency/proto"
	"m3o.dev/services/db/proto"
	"m3o.dev/services/dns/proto"
	"m3o.dev/services/email/proto"
	"m3o.dev/services/emoji/proto"
	"m3o.dev/services/evchargers/proto"
	"m3o.dev/services/event/proto"
	"m3o.dev/services/file/proto"
	"m3o.dev/services/forex/proto"
	"m3o.dev/services/function/proto"
	"m3o.dev/services/geocoding/proto"
	"m3o.dev/services/gifs/proto"
	"m3o.dev/services/github/proto"
	"m3o.dev/services/google/proto"
	"m3o.dev/services/helloworld/proto"
	"m3o.dev/services/holidays/proto"
	"m3o.dev/services/id/proto"
	"m3o.dev/services/image/proto"
	"m3o.dev/services/ip/proto"
	"m3o.dev/services/joke/proto"
	"m3o.dev/services/lists/proto"
	"m3o.dev/services/location/proto"
	"m3o.dev/services/memegen/proto"
	"m3o.dev/services/minecraft/proto"
	"m3o.dev/services/movie/proto"
	"m3o.dev/services/mq/proto"
	"m3o.dev/services/news/proto"
	"m3o.dev/services/nft/proto"
	"m3o.dev/services/notes/proto"
	"m3o.dev/services/otp/proto"
	"m3o.dev/services/password/proto"
	"m3o.dev/services/ping/proto"
	"m3o.dev/services/place/proto"
	"m3o.dev/services/postcode/proto"
	"m3o.dev/services/prayer/proto"
	"m3o.dev/services/price/proto"
	"m3o.dev/services/qr/proto"
	"m3o.dev/services/quran/proto"
	"m3o.dev/services/routing/proto"
	"m3o.dev/services/rss/proto"
	"m3o.dev/services/search/proto"
	"m3o.dev/services/sentiment/proto"
	"m3o.dev/services/sms/proto"
	"m3o.dev/services/space/proto"
	"m3o.dev/services/spam/proto"
	"m3o.dev/services/stock/proto"
	"m3o.dev/services/stream/proto"
	"m3o.dev/services/sunnah/proto"
	"m3o.dev/services/thumbnail/proto"
	"m3o.dev/services/time/proto"
	"m3o.dev/services/translate/proto"
	"m3o.dev/services/tunnel/proto"
	"m3o.dev/services/twitter/proto"
	"m3o.dev/services/url/proto"
	"m3o.dev/services/user/proto"
	"m3o.dev/services/vehicle/proto"
	"m3o.dev/services/wallet/proto"
	"m3o.dev/services/weather/proto"
	"m3o.dev/services/wordle/proto"
	"m3o.dev/services/youtube/proto"
)

type Client struct {
	Address    address.AddressService
	Analytics  analytics.AnalyticsService
	Answer     answer.AnswerService
	App        app.AppService
	Avatar     avatar.AvatarService
	Bitcoin    bitcoin.BitcoinService
	Cache      cache.CacheService
	Carbon     carbon.CarbonService
	Chat       chat.ChatService
	Comments   comments.CommentsService
	Contact    contact.ContactService
	Crypto     crypto.CryptoService
	Currency   currency.CurrencyService
	Db         db.DbService
	Dns        dns.DnsService
	Email      email.EmailService
	Emoji      emoji.EmojiService
	Evchargers evchargers.EvchargersService
	Event      event.EventService
	File       file.FileService
	Forex      forex.ForexService
	Function   function.FunctionService
	Geocoding  geocoding.GeocodingService
	Gifs       gifs.GifsService
	Github     github.GithubService
	Google     google.GoogleService
	Helloworld helloworld.HelloworldService
	Holidays   holidays.HolidaysService
	Id         id.IdService
	Image      image.ImageService
	Ip         ip.IpService
	Joke       joke.JokeService
	Lists      lists.ListsService
	Location   location.LocationService
	Memegen    memegen.MemegenService
	Minecraft  minecraft.MinecraftService
	Movie      movie.MovieService
	Mq         mq.MqService
	News       news.NewsService
	Nft        nft.NftService
	Notes      notes.NotesService
	Otp        otp.OtpService
	Password   password.PasswordService
	Ping       ping.PingService
	Place      place.PlaceService
	Postcode   postcode.PostcodeService
	Prayer     prayer.PrayerService
	Price      price.PriceService
	Qr         qr.QrService
	Quran      quran.QuranService
	Routing    routing.RoutingService
	Rss        rss.RssService
	Search     search.SearchService
	Sentiment  sentiment.SentimentService
	Sms        sms.SmsService
	Space      space.SpaceService
	Spam       spam.SpamService
	Stock      stock.StockService
	Stream     stream.StreamService
	Sunnah     sunnah.SunnahService
	Thumbnail  thumbnail.ThumbnailService
	Time       time.TimeService
	Translate  translate.TranslateService
	Tunnel     tunnel.TunnelService
	Twitter    twitter.TwitterService
	Url        url.UrlService
	User       user.UserService
	Vehicle    vehicle.VehicleService
	Wallet     wallet.WalletService
	Weather    weather.WeatherService
	Wordle     wordle.WordleService
	Youtube    youtube.YoutubeService
}

func NewClient(c client.Client) *Client {
	return &Client{
		Address:    address.NewAddressService("address", c),
		Analytics:  analytics.NewAnalyticsService("analytics", c),
		Answer:     answer.NewAnswerService("answer", c),
		App:        app.NewAppService("app", c),
		Avatar:     avatar.NewAvatarService("avatar", c),
		Bitcoin:    bitcoin.NewBitcoinService("bitcoin", c),
		Cache:      cache.NewCacheService("cache", c),
		Carbon:     carbon.NewCarbonService("carbon", c),
		Chat:       chat.NewChatService("chat", c),
		Comments:   comments.NewCommentsService("comments", c),
		Contact:    contact.NewContactService("contact", c),
		Crypto:     crypto.NewCryptoService("crypto", c),
		Currency:   currency.NewCurrencyService("currency", c),
		Db:         db.NewDbService("db", c),
		Dns:        dns.NewDnsService("dns", c),
		Email:      email.NewEmailService("email", c),
		Emoji:      emoji.NewEmojiService("emoji", c),
		Evchargers: evchargers.NewEvchargersService("evchargers", c),
		Event:      event.NewEventService("event", c),
		File:       file.NewFileService("file", c),
		Forex:      forex.NewForexService("forex", c),
		Function:   function.NewFunctionService("function", c),
		Geocoding:  geocoding.NewGeocodingService("geocoding", c),
		Gifs:       gifs.NewGifsService("gifs", c),
		Github:     github.NewGithubService("github", c),
		Google:     google.NewGoogleService("google", c),
		Helloworld: helloworld.NewHelloworldService("helloworld", c),
		Holidays:   holidays.NewHolidaysService("holidays", c),
		Id:         id.NewIdService("id", c),
		Image:      image.NewImageService("image", c),
		Ip:         ip.NewIpService("ip", c),
		Joke:       joke.NewJokeService("joke", c),
		Lists:      lists.NewListsService("lists", c),
		Location:   location.NewLocationService("location", c),
		Memegen:    memegen.NewMemegenService("memegen", c),
		Minecraft:  minecraft.NewMinecraftService("minecraft", c),
		Movie:      movie.NewMovieService("movie", c),
		Mq:         mq.NewMqService("mq", c),
		News:       news.NewNewsService("news", c),
		Nft:        nft.NewNftService("nft", c),
		Notes:      notes.NewNotesService("notes", c),
		Otp:        otp.NewOtpService("otp", c),
		Password:   password.NewPasswordService("password", c),
		Ping:       ping.NewPingService("ping", c),
		Place:      place.NewPlaceService("place", c),
		Postcode:   postcode.NewPostcodeService("postcode", c),
		Prayer:     prayer.NewPrayerService("prayer", c),
		Price:      price.NewPriceService("price", c),
		Qr:         qr.NewQrService("qr", c),
		Quran:      quran.NewQuranService("quran", c),
		Routing:    routing.NewRoutingService("routing", c),
		Rss:        rss.NewRssService("rss", c),
		Search:     search.NewSearchService("search", c),
		Sentiment:  sentiment.NewSentimentService("sentiment", c),
		Sms:        sms.NewSmsService("sms", c),
		Space:      space.NewSpaceService("space", c),
		Spam:       spam.NewSpamService("spam", c),
		Stock:      stock.NewStockService("stock", c),
		Stream:     stream.NewStreamService("stream", c),
		Sunnah:     sunnah.NewSunnahService("sunnah", c),
		Thumbnail:  thumbnail.NewThumbnailService("thumbnail", c),
		Time:       time.NewTimeService("time", c),
		Translate:  translate.NewTranslateService("translate", c),
		Tunnel:     tunnel.NewTunnelService("tunnel", c),
		Twitter:    twitter.NewTwitterService("twitter", c),
		Url:        url.NewUrlService("url", c),
		User:       user.NewUserService("user", c),
		Vehicle:    vehicle.NewVehicleService("vehicle", c),
		Wallet:     wallet.NewWalletService("wallet", c),
		Weather:    weather.NewWeatherService("weather", c),
		Wordle:     wordle.NewWordleService("wordle", c),
		Youtube:    youtube.NewYoutubeService("youtube", c),
	}
}
