package main

import (
	"m3o.dev/apps/chat/api"
	"m3o.dev/apps/chat/api/groups/handler"
	pb "m3o.dev/apps/chat/api/groups/proto"

	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("groups"),
		service.Version("latest"),
	)

	// Connect to the database
	db, err := api.NewDB("groups")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(&handler.Group{}, &handler.Membership{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterGroupsHandler(srv.Server(), &handler.Groups{DB: db.Debug()})

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
