package main

import (
	"github.com/m3o/m3o/backend/emails/handler"
	pb "github.com/m3o/m3o/backend/emails/proto"
	"github.com/m3o/m3o/backend/pkg/tracing"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
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
