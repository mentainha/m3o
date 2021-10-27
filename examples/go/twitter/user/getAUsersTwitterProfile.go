package example

import (
	"fmt"
	"os"

	"github.com/micro/services/clients/go/twitter"
)

// Get a user's twitter profile
func GetAusersTwitterProfile() {
	twitterService := twitter.NewTwitterService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := twitterService.User(&twitter.UserRequest{
		Username: "crufter",
	})
	fmt.Println(rsp, err)
}
