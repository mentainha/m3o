package main

import (
	"time"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/model/sql"
	"m3o.dev/api/chats/handler"
	pb "m3o.dev/api/chats/proto"
)

func main() {
	srv := service.New(
		service.Name("chats"),
	)

	// Connect to the database
	db, err := sql.New("chats")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}

	if err := db.AutoMigrate(&handler.Chat{}, &handler.Message{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterChatsHandler(srv.Server(), &handler.Chats{DB: db, Time: time.Now})

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
