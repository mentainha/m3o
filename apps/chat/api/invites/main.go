package main

import (
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	"m3o.dev/apps/chat/api"
	"m3o.dev/apps/chat/api/invites/handler"
	pb "m3o.dev/apps/chat/api/invites/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("invites"),
		service.Version("latest"),
	)

	// Connect to the database
	db, err := api.NewDB("invites")
	if err != nil {
		logger.Fatalf("Error connecting to database: %v", err)
	}
	if err := db.AutoMigrate(&handler.Invite{}); err != nil {
		logger.Fatalf("Error migrating database: %v", err)
	}

	// Register handler
	pb.RegisterInvitesHandler(srv.Server(), &handler.Invites{DB: db})

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
