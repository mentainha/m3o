package main

import (
	"github.com/m3o/m3o/backend/pkg/tracing"
	"github.com/m3o/m3o/backend/publicapi/handler"
	pb "github.com/m3o/m3o/backend/publicapi/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("publicapi"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterPublicapiHandler(srv.Server(), handler.NewPublicAPIHandler(srv))
	pb.RegisterExploreHandler(srv.Server(), handler.NewExploreAPIHandler(srv))
	traceCloser := tracing.SetupOpentracing("publicapi")
	defer traceCloser.Close()

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
