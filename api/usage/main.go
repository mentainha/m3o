package main

import (
	"github.com/robfig/cron/v3"
	"m3o.dev/api/pkg/tracing"
	"m3o.dev/api/usage/handler"
	pb "m3o.dev/api/usage/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	dbproto "github.com/micro/services/db/proto"
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
