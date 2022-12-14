package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"m3o.dev/apps/chat/api"
	"m3o.dev/apps/chat/api/seen/handler"
	pb "m3o.dev/apps/chat/api/seen/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("seen"),
		service.Version("latest"),
	)

	// Connect to the database
	db, err := api.NewDB("seen")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(&handler.SeenInstance{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterSeenHandler(srv.Server(), &handler.Seen{DB: db.Debug()})

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
