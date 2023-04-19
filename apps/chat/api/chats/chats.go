package chats

import (
	"time"

	"m3o.dev/apps/chat/api/chats/handler"
	pb "m3o.dev/apps/chat/api/chats/proto"
	"m3o.dev/apps/chat/api/db"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func Register(srv *service.Service) {
	// Connect to the database
	db, err := db.New("chats")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}

	if err := db.AutoMigrate(&handler.Chat{}, &handler.Message{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterChatsHandler(srv.Server(), &handler.Chats{DB: db, Time: time.Now})
}
