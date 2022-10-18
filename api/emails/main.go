package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"m3o.dev/api/emails/handler"
	pb "m3o.dev/api/emails/proto"
	"m3o.dev/api/pkg/tracing"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("emails"),
		service.Version("latest"),
	)

	// Register handler
	pb.RegisterEmailsHandler(srv.Server(), handler.NewEmailsHandler())
	traceCloser := tracing.SetupOpentracing("emails")
	defer traceCloser.Close()

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
