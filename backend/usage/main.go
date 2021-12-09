package main

import (
	"github.com/m3o/m3o/backend/pkg/tracing"
	"github.com/m3o/m3o/backend/usage/handler"
	pb "github.com/m3o/m3o/backend/usage/proto"
	"github.com/robfig/cron/v3"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	dbproto "github.com/m3o/m3o/services/db/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("usage"),
		service.Version("latest"),
	)

	p := handler.NewHandler(srv, dbproto.NewDbService("db", srv.Client()))
	// Register handler
	pb.RegisterUsageHandler(srv.Server(), p)

	c := cron.New()
	c.AddFunc("1 0 * * *", p.UsageCron)
	c.AddFunc("1 1 * * *", p.RankingCron)
	c.Start()
	traceCloser := tracing.SetupOpentracing("usage")
	defer traceCloser.Close()

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
