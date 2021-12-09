package main

import (
	"github.com/m3o/m3o/backend/mixpanel/handler"
	"github.com/m3o/m3o/backend/pkg/tracing"

	pb "github.com/m3o/m3o/backend/mixpanel/proto"

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
