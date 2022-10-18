package main

import (
	"m3o.dev/api/mixpanel/handler"
	"m3o.dev/api/pkg/tracing"

	pb "m3o.dev/api/mixpanel/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("mixpanel"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterMixpanelHandler(srv.Server(), handler.NewHandler(srv))
	traceCloser := tracing.SetupOpentracing("mixpanel")
	defer traceCloser.Close()

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
