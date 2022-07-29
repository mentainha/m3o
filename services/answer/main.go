package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"m3o.dev/services/answer/handler"
	pb "m3o.dev/services/answer/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("answer"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterAnswerHandler(srv.Server(), new(handler.Answer))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
