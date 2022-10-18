package main

import (
	"m3o.dev/api/oauth/handler"

	"github.com/micro/micro/v3/service"
	mauth "github.com/micro/micro/v3/service/auth/client"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("oauth"),
		service.Version("latest"),
	)

	// passing in auth because the DefaultAuth is the one used to set up the service
	auth := mauth.NewAuth()

	// Register Handler
	srv.Handle(handler.NewOauth(srv, auth))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
