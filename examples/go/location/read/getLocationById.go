package example

import (
	"fmt"
	"os"

	"github.com/micro/services/clients/go/location"
)

// Read an entity by its ID
func GetLocationById() {
	locationService := location.NewLocationService(os.Getenv("M3O_API_TOKEN"))
	rsp, err := locationService.Read(&location.ReadRequest{
		Id: "1",
	})
	fmt.Println(rsp, err)
}
