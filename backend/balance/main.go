package main

import (
	"github.com/m3o/m3o/backend/balance/handler"
	pb "github.com/m3o/m3o/backend/balance/proto"
	"github.com/m3o/m3o/backend/pkg/tracing"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("balance"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterBalanceHandler(srv.Server(), handler.NewHandler(srv))

	traceCloser := tracing.SetupOpentracing("balance")
	defer traceCloser.Close()

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
