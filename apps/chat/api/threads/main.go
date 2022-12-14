package main

import (
	"time"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"m3o.dev/apps/chat/api"
	"m3o.dev/apps/chat/api/threads/handler"
	pb "m3o.dev/apps/chat/api/threads/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("threads"),
		service.Version("latest"),
	)

	// Connect to the database
	db, err := api.NewDB("threads")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(&handler.Conversation{}, &handler.Message{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterThreadsHandler(srv.Server(), &handler.Threads{DB: db, Time: time.Now})

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
