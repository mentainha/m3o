package main

import (
	"time"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/events"
	"github.com/micro/micro/v3/service/logger"
	"m3o.dev/apps/chat/api"
	"m3o.dev/apps/chat/api/streams/handler"
	pb "m3o.dev/apps/chat/api/streams/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("streams"),
		service.Version("latest"),
	)

	// Connect to the database
	db, err := api.NewDB("streams")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(&handler.Token{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterStreamsHandler(srv.Server(), &handler.Streams{
		DB:     db,
		Events: events.DefaultStream,
		Time:   time.Now,
	})

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
