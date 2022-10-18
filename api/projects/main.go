package main

import (
	"m3o.dev/api/projects/handler"
	pb "m3o.dev/api/projects/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("projects"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterProjectsHandler(srv.Server(), handler.New(srv))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
