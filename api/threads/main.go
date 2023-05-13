package main

import (
	"time"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"github.com/micro/micro/v3/service/model/sql"
	"m3o.dev/api/threads/handler"
	pb "m3o.dev/api/threads/proto"
)

func main() {
	srv := service.New(
		service.Name("threads"),
	)

	// Connect to the database
	db, err := sql.New("threads")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(&handler.Conversation{}, &handler.Message{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterThreadsHandler(srv.Server(), &handler.Threads{DB: db, Time: time.Now})

	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
