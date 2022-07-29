package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"m3o.dev/services/analytics/handler"
	pb "m3o.dev/services/analytics/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("analytics"),
		service.Version("latest"),
	)

	h := handler.New()

	// Register handler
	pb.RegisterAnalyticsHandler(srv.Server(), h)

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
