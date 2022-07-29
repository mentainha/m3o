package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"m3o.dev/services/wordle/handler"
	pb "m3o.dev/services/wordle/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("wordle"),
	)

	// Register handler
	pb.RegisterWordleHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
