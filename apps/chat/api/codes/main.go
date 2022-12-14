package main

import (
	"time"

	"m3o.dev/apps/chat/api"
	"m3o.dev/apps/chat/api/codes/handler"
	pb "m3o.dev/apps/chat/api/codes/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("codes"),
		service.Version("latest"),
	)

	// Connect to the database
	db, err := api.NewDB("codes")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(&handler.Code{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterCodesHandler(srv.Server(), &handler.Codes{DB: db, Time: time.Now})

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
