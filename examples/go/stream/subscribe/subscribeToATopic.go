package example

import (
	"fmt"
	"os"

	"github.com/micro/services/clients/go/stream"
)

// Subscribe to messages for a given topic.
func SubscribeToAtopic() {
	streamService := stream.NewStreamService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := streamService.Subscribe(&stream.SubscribeRequest{
		Topic: "events",
	})
	fmt.Println(rsp, err)
}
