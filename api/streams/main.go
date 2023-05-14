package main

import (
	"time"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/events"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/model/sql"
	"m3o.dev/api/streams/handler"
	pb "m3o.dev/api/streams/proto"
)

func main() {
	srv := service.New(
		service.Name("streams"),
	)

	// Connect to the database
	db, err := sql.NewDB("streams")
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

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
