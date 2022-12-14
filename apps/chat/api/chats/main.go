package main

import (
	"time"

	"m3o.dev/apps/chat/api"
	"m3o.dev/apps/chat/api/chats/handler"
	pb "m3o.dev/apps/chat/api/chats/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("chats"),
		service.Version("latest"),
	)

	// Connect to the database
	db, err := api.NewDB("chats")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}

	if err := db.AutoMigrate(&handler.Chat{}, &handler.Message{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterChatsHandler(srv.Server(), &handler.Chats{DB: db, Time: time.Now})

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
