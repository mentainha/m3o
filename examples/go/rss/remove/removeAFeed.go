package example

import (
	"fmt"
	"os"

	"github.com/micro/services/clients/go/rss"
)

// Remove an RSS feed by name
func RemoveAfeed() {
	rssService := rss.NewRssService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := rssService.Remove(&rss.RemoveRequest{
		Name: "bbc",
	})
	fmt.Println(rsp, err)
}
